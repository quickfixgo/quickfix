// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import (
	"github.com/quickfixgo/quickfix/datadictionary"
)

const (
	UserDefinedTagMin int = 5000
)

// Validator validates a FIX message.
type Validator interface {
	Validate(*Message) MessageRejectError
}

// ValidatorSettings describe validation behavior.
type ValidatorSettings struct {
	CheckFieldsOutOfOrder     bool
	RejectInvalidMessage      bool
	AllowUnknownMessageFields bool
	CheckUserDefinedFields    bool
}

// Default configuration for message validation.
// See http://www.quickfixengine.org/quickfix/doc/html/configuration.html.
var defaultValidatorSettings = ValidatorSettings{
	CheckFieldsOutOfOrder:     true,
	RejectInvalidMessage:      true,
	AllowUnknownMessageFields: false,
	CheckUserDefinedFields:    true,
}

type fixValidator struct {
	dataDictionary *datadictionary.DataDictionary
	settings       ValidatorSettings
}

type fixtValidator struct {
	transportDataDictionary *datadictionary.DataDictionary
	appDataDictionary       *datadictionary.DataDictionary
	settings                ValidatorSettings
}

// NewValidator creates a FIX message validator from the given data dictionaries.
func NewValidator(settings ValidatorSettings, appDataDictionary, transportDataDictionary *datadictionary.DataDictionary) Validator {
	if transportDataDictionary != nil {
		return &fixtValidator{
			transportDataDictionary: transportDataDictionary,
			appDataDictionary:       appDataDictionary,
			settings:                settings,
		}
	}
	return &fixValidator{
		dataDictionary: appDataDictionary,
		settings:       settings,
	}
}

// Validate tests the message against the provided data dictionary.
func (v *fixValidator) Validate(msg *Message) MessageRejectError {
	if !msg.Header.Has(tagMsgType) {
		return RequiredTagMissing(tagMsgType)
	}
	msgType, err := msg.Header.GetString(tagMsgType)
	if err != nil {
		return err
	}

	return validateFIX(v.dataDictionary, v.settings, msgType, msg)
}

// Validate tests the message against the provided transport and app data dictionaries.
// If the message is an admin message, it will be validated against the transport data dictionary.
func (v *fixtValidator) Validate(msg *Message) MessageRejectError {
	if !msg.Header.Has(tagMsgType) {
		return RequiredTagMissing(tagMsgType)
	}
	msgType, err := msg.Header.GetString(tagMsgType)
	if err != nil {
		return err
	}

	if isAdminMessageType([]byte(msgType)) {
		return validateFIX(v.transportDataDictionary, v.settings, msgType, msg)
	}
	return validateFIXT(v.transportDataDictionary, v.appDataDictionary, v.settings, msgType, msg)
}

func validateFIX(d *datadictionary.DataDictionary, settings ValidatorSettings, msgType string, msg *Message) MessageRejectError {
	if err := validateMsgType(d, msgType, msg); err != nil {
		return err
	}

	if err := validateRequired(d, d, msgType, msg); err != nil {
		return err
	}

	if settings.CheckFieldsOutOfOrder {
		if err := validateOrder(msg); err != nil {
			return err
		}
	}

	if settings.RejectInvalidMessage {
		if err := validateFields(d, d, settings, msgType, msg); err != nil {
			return err
		}

		if err := validateWalk(d, d, settings, msgType, msg); err != nil {
			return err
		}
	}

	return nil
}

func validateFIXT(transportDD, appDD *datadictionary.DataDictionary, settings ValidatorSettings, msgType string, msg *Message) MessageRejectError {
	if err := validateMsgType(appDD, msgType, msg); err != nil {
		return err
	}

	if err := validateRequired(transportDD, appDD, msgType, msg); err != nil {
		return err
	}

	if settings.CheckFieldsOutOfOrder {
		if err := validateOrder(msg); err != nil {
			return err
		}
	}

	if settings.RejectInvalidMessage {
		if err := validateFields(transportDD, appDD, settings, msgType, msg); err != nil {
			return err
		}

		if err := validateWalk(transportDD, appDD, settings, msgType, msg); err != nil {
			return err
		}
	}

	return nil
}

func validateMsgType(d *datadictionary.DataDictionary, msgType string, _ *Message) MessageRejectError {
	if _, validMsgType := d.Messages[msgType]; !validMsgType {
		return InvalidMessageType()
	}
	return nil
}

func validateWalk(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, settings ValidatorSettings, msgType string, msg *Message) MessageRejectError {
	remainingFields := msg.fields
	iteratedTags := make(datadictionary.TagSet)

	var messageDef *datadictionary.MessageDef
	var fieldDef *datadictionary.FieldDef
	var err MessageRejectError
	var ok bool

	for len(remainingFields) > 0 {
		field := remainingFields[0]
		tag := field.tag

		switch {
		case tag.IsHeader():
			messageDef = transportDD.Header
		case tag.IsTrailer():
			messageDef = transportDD.Trailer
		default: // is body
			messageDef = appDD.Messages[msgType]
		}

		if _, duplicate := iteratedTags[int(tag)]; duplicate {
			return tagAppearsMoreThanOnce(tag)
		}
		iteratedTags.Add(int(tag))

		if fieldDef, ok = messageDef.Fields[int(tag)]; !ok {
			if !checkFieldNotDefined(settings, tag) {
				return TagNotDefinedForThisMessageType(tag)
			}
			remainingFields = remainingFields[1:]
			continue
		}

		if remainingFields, err = validateVisitField(fieldDef, remainingFields); err != nil {
			return err
		}
	}

	if len(remainingFields) != 0 {
		return TagNotDefinedForThisMessageType(remainingFields[0].tag)
	}

	return nil
}

func validateVisitField(fieldDef *datadictionary.FieldDef, fields []TagValue) ([]TagValue, MessageRejectError) {
	if fieldDef.IsGroup() {
		var err MessageRejectError
		if fields, err = validateVisitGroupField(fieldDef, fields); err != nil {
			return nil, err
		}
		return fields, nil
	}

	return fields[1:], nil
}

func validateVisitGroupField(fieldDef *datadictionary.FieldDef, fieldStack []TagValue) ([]TagValue, MessageRejectError) {
	numInGroupTag := fieldStack[0].tag
	var numInGroup FIXInt

	if err := numInGroup.Read(fieldStack[0].value); err != nil {
		return nil, IncorrectDataFormatForValue(numInGroupTag)
	}

	fieldStack = fieldStack[1:]

	var childDefs []*datadictionary.FieldDef
	groupCount := 0

	for len(fieldStack) > 0 {

		// Start of repeating group.
		if int(fieldStack[0].tag) == fieldDef.Fields[0].Tag() {
			childDefs = fieldDef.Fields
			groupCount++
		}

		// Group complete.
		if len(childDefs) == 0 {
			break
		}

		if int(fieldStack[0].tag) == childDefs[0].Tag() {
			var err MessageRejectError
			if fieldStack, err = validateVisitField(childDefs[0], fieldStack); err != nil {
				return fieldStack, err
			}
		} else {
			if childDefs[0].Required() {
				return fieldStack, RequiredTagMissing(Tag(childDefs[0].Tag()))
			}
		}

		childDefs = childDefs[1:]
	}

	if groupCount != int(numInGroup) {
		return fieldStack, incorrectNumInGroupCountForRepeatingGroup(numInGroupTag)
	}

	return fieldStack, nil
}

func validateOrder(msg *Message) MessageRejectError {
	inHeader := true
	inTrailer := false
	for _, field := range msg.fields {
		t := field.tag
		switch {
		case inHeader && t.IsHeader():
		case inHeader && !t.IsHeader():
			inHeader = false
		case !inHeader && t.IsHeader():
			return tagSpecifiedOutOfRequiredOrder(t)
		case t.IsTrailer():
			inTrailer = true
		case inTrailer && !t.IsTrailer():
			return tagSpecifiedOutOfRequiredOrder(t)
		}
	}

	return nil
}

func validateRequired(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msgType string, message *Message) MessageRejectError {
	if err := validateRequiredFieldMap(message, transportDD.Header.RequiredTags, message.Header.FieldMap); err != nil {
		return err
	}

	if err := validateRequiredFieldMap(message, appDD.Messages[msgType].RequiredTags, message.Body.FieldMap); err != nil {
		return err
	}

	if err := validateRequiredFieldMap(message, transportDD.Trailer.RequiredTags, message.Trailer.FieldMap); err != nil {
		return err
	}

	return nil
}

func validateRequiredFieldMap(_ *Message, requiredTags map[int]struct{}, fieldMap FieldMap) MessageRejectError {
	for required := range requiredTags {
		requiredTag := Tag(required)
		if !fieldMap.Has(requiredTag) {
			return RequiredTagMissing(requiredTag)
		}
	}

	return nil
}

func validateFields(transportDD *datadictionary.DataDictionary,
	appDD *datadictionary.DataDictionary,
	settings ValidatorSettings,
	msgType string,
	message *Message,
) MessageRejectError {
	for _, field := range message.fields {
		switch {
		case field.tag.IsHeader():
			if err := validateField(transportDD, settings, transportDD.Header.Tags, field); err != nil {
				return err
			}
		case field.tag.IsTrailer():
			if err := validateField(transportDD, settings, transportDD.Trailer.Tags, field); err != nil {
				return err
			}
		default:
			if err := validateField(appDD, settings, appDD.Messages[msgType].Tags, field); err != nil {
				return err
			}
		}
	}

	return nil
}

func getFieldType(d *datadictionary.DataDictionary, field int) (*datadictionary.FieldType, bool) {
	fieldType, isMessageField := d.FieldTypeByTag[field]
	return fieldType, isMessageField
}

func checkFieldNotDefined(settings ValidatorSettings, field Tag) bool {
	fail := false
	if int(field) < UserDefinedTagMin {
		fail = !settings.AllowUnknownMessageFields
	} else {
		fail = settings.CheckUserDefinedFields
	}
	return !fail
}

func validateField(d *datadictionary.DataDictionary,
	settings ValidatorSettings,
	_ datadictionary.TagSet,
	field TagValue,
) MessageRejectError {
	if len(field.value) == 0 {
		return TagSpecifiedWithoutAValue(field.tag)
	}

	fieldType, isMessageField := getFieldType(d, int(field.tag))
	if !isMessageField && !checkFieldNotDefined(settings, field.tag) {
		return InvalidTagNumber(field.tag)
	}

	if !isMessageField {
		return nil
	}

	allowedValues := d.FieldTypeByTag[int(field.tag)].Enums
	if len(allowedValues) != 0 {
		if _, validValue := allowedValues[string(field.value)]; !validValue {
			return ValueIsIncorrect(field.tag)
		}
	}

	var prototype FieldValue
	switch fieldType.Type {
	case "MULTIPLESTRINGVALUE", "MULTIPLEVALUESTRING":
		fallthrough
	case "MULTIPLECHARVALUE":
		fallthrough
	case "CHAR":
		fallthrough
	case "CURRENCY":
		fallthrough
	case "DATA":
		fallthrough
	case "MONTHYEAR":
		fallthrough
	case "LOCALMKTDATE", "DATE":
		fallthrough
	case "EXCHANGE":
		fallthrough
	case "LANGUAGE":
		fallthrough
	case "XMLDATA":
		fallthrough
	case "COUNTRY":
		fallthrough
	case "UTCTIMEONLY":
		fallthrough
	case "UTCDATEONLY", "UTCDATE":
		fallthrough
	case "TZTIMEONLY":
		fallthrough
	case "TZTIMESTAMP":
		fallthrough
	case "STRING":
		prototype = new(FIXString)

	case "BOOLEAN":
		prototype = new(FIXBoolean)

	case "LENGTH":
		fallthrough
	case "DAYOFMONTH":
		fallthrough
	case "NUMINGROUP":
		fallthrough
	case "SEQNUM":
		fallthrough
	case "INT":
		prototype = new(FIXInt)

	case "UTCTIMESTAMP", "TIME":
		prototype = new(FIXUTCTimestamp)

	case "QTY", "QUANTITY":
		fallthrough
	case "AMT":
		fallthrough
	case "PRICE":
		fallthrough
	case "PRICEOFFSET":
		fallthrough
	case "PERCENTAGE":
		fallthrough
	case "FLOAT":
		prototype = new(FIXFloat)

	}

	if err := prototype.Read(field.value); err != nil {
		return IncorrectDataFormatForValue(field.tag)
	}

	return nil
}
