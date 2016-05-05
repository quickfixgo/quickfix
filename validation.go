package quickfix

import (
	"github.com/quickfixgo/quickfix/datadictionary"
)

type validator interface {
	Validate(Message) MessageRejectError
}

type validatorSettings struct {
	CheckFieldsOutOfOrder bool
}

//Default configuration for message validation.
//See http://www.quickfixengine.org/quickfix/doc/html/configuration.html.
var defaultValidatorSettings = validatorSettings{
	CheckFieldsOutOfOrder: true,
}

type fixValidator struct {
	dataDictionary *datadictionary.DataDictionary
	settings       validatorSettings
}

type fixtValidator struct {
	transportDataDictionary *datadictionary.DataDictionary
	appDataDictionary       *datadictionary.DataDictionary
	settings                validatorSettings
}

//Validate tests the message against the provided data dictionary.
func (v *fixValidator) Validate(msg Message) MessageRejectError {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err != nil {
		if err.RejectReason() == rejectReasonConditionallyRequiredFieldMissing {
			return RequiredTagMissing(tagMsgType)
		} else {
			return err
		}
	} else {
		return validateFIX(v.dataDictionary, v.settings, string(msgType), msg)
	}
}

//Validate tests the message against the provided transport and app data dictionaries.
//If the message is an admin message, it will be validated against the transport data dictionary.
func (v *fixtValidator) Validate(msg Message) MessageRejectError {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err != nil {
		if err.RejectReason() == rejectReasonConditionallyRequiredFieldMissing {
			return RequiredTagMissing(tagMsgType)
		} else {
			return err
		}
	} else if isAdminMessageType(string(msgType)) {
		return validateFIX(v.transportDataDictionary, v.settings, string(msgType), msg)
	} else {
		return validateFIXT(v.transportDataDictionary, v.appDataDictionary, v.settings, string(msgType), msg)
	}
}

func validateFIX(d *datadictionary.DataDictionary, settings validatorSettings, msgType string, msg Message) MessageRejectError {
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

	if err := validateFields(d, d, msgType, msg); err != nil {
		return err
	}

	if err := validateWalk(d, d, msgType, msg); err != nil {
		return err
	}

	return nil
}

func validateFIXT(transportDD, appDD *datadictionary.DataDictionary, settings validatorSettings, msgType string, msg Message) MessageRejectError {
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

	if err := validateWalk(transportDD, appDD, msgType, msg); err != nil {
		return err
	}

	if err := validateFields(transportDD, appDD, msgType, msg); err != nil {
		return err
	}

	return nil
}

func validateMsgType(d *datadictionary.DataDictionary, msgType string, msg Message) MessageRejectError {
	if _, validMsgType := d.Messages[msgType]; validMsgType == false {
		return InvalidMessageType()
	}
	return nil
}

func validateWalk(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msgType string, msg Message) MessageRejectError {
	remainingFields := msg.fields
	iteratedTags := make(datadictionary.TagSet)

	var messageDef *datadictionary.MessageDef
	var fieldDef *datadictionary.FieldDef
	var err MessageRejectError
	var ok bool

	for len(remainingFields) > 0 {
		field := remainingFields[0]
		tag := field.Tag

		switch {
		case tag.IsHeader():
			messageDef = transportDD.Header
		case tag.IsTrailer():
			messageDef = transportDD.Trailer
		default: // is body
			messageDef = appDD.Messages[msgType]
		}

		if fieldDef, ok = messageDef.Fields[int(tag)]; !ok {
			return TagNotDefinedForThisMessageType(tag)
		}

		if _, duplicate := iteratedTags[int(tag)]; duplicate {
			return tagAppearsMoreThanOnce(tag)
		}
		iteratedTags.Add(int(tag))

		if remainingFields, err = validateVisitField(fieldDef, remainingFields); err != nil {
			return err
		}
	}

	if len(remainingFields) != 0 {
		return TagNotDefinedForThisMessageType(remainingFields[0].Tag)
	}

	return nil
}

func validateVisitField(fieldDef *datadictionary.FieldDef, fields []tagValue) ([]tagValue, MessageRejectError) {
	if fieldDef.IsGroup() {
		var err MessageRejectError
		if fields, err = validateVisitGroupField(fieldDef, fields); err != nil {
			return nil, err
		} else {
			return fields, nil
		}
	}

	return fields[1:], nil
}

func validateVisitGroupField(fieldDef *datadictionary.FieldDef, fieldStack []tagValue) ([]tagValue, MessageRejectError) {
	numInGroupTag := fieldStack[0].Tag
	var numInGroup FIXInt

	if err := numInGroup.Read(fieldStack[0].Value); err != nil {
		return nil, IncorrectDataFormatForValue(numInGroupTag)
	}

	fieldStack = fieldStack[1:]

	var childDefs []*datadictionary.FieldDef
	groupCount := 0

	for len(fieldStack) > 0 {

		//start of repeating group
		if int(fieldStack[0].Tag) == fieldDef.ChildFields[0].Tag() {
			childDefs = fieldDef.ChildFields
			groupCount++
		}

		//group complete
		if len(childDefs) == 0 {
			break
		}

		if int(fieldStack[0].Tag) == childDefs[0].Tag() {
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

func validateOrder(msg Message) MessageRejectError {
	inHeader := true
	inTrailer := false
	for _, field := range msg.fields {
		t := field.Tag
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

func validateRequired(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msgType string, message Message) MessageRejectError {
	if err := validateRequiredFieldMap(message, transportDD.Header.RequiredTags, message.Header); err != nil {
		return err
	}

	if err := validateRequiredFieldMap(message, appDD.Messages[msgType].RequiredTags, message.Body); err != nil {
		return err
	}

	if err := validateRequiredFieldMap(message, transportDD.Trailer.RequiredTags, message.Trailer); err != nil {
		return err
	}

	return nil
}

func validateRequiredFieldMap(msg Message, requiredTags map[int]struct{}, fieldMap FieldMap) MessageRejectError {
	for required := range requiredTags {
		field := new(FIXString)
		if err := fieldMap.GetField(Tag(required), field); err != nil {
			//FIXME: add "has..." method?
			if err.RejectReason() == rejectReasonConditionallyRequiredFieldMissing {
				return RequiredTagMissing(Tag(required))
			}
			return err
		}
	}

	return nil
}

func validateFields(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msgType string, message Message) MessageRejectError {
	for _, field := range message.fields {
		switch {
		case field.Tag.IsHeader():
			if err := validateField(transportDD, transportDD.Header.Tags, field); err != nil {
				return err
			}
		case field.Tag.IsTrailer():
			if err := validateField(transportDD, transportDD.Trailer.Tags, field); err != nil {
				return err
			}
		default:
			if err := validateField(appDD, appDD.Messages[msgType].Tags, field); err != nil {
				return err
			}
		}
	}

	return nil
}

func validateField(d *datadictionary.DataDictionary, validFields datadictionary.TagSet, field tagValue) MessageRejectError {
	if len(field.Value) == 0 {
		return TagSpecifiedWithoutAValue(field.Tag)
	}

	if _, valid := d.FieldTypeByTag[int(field.Tag)]; !valid {
		return InvalidTagNumber(field.Tag)
	}

	allowedValues := d.FieldTypeByTag[int(field.Tag)].Enums
	if len(allowedValues) != 0 {
		if _, validValue := allowedValues[string(field.Value)]; !validValue {
			return ValueIsIncorrect(field.Tag)
		}
	}

	fieldType := d.FieldTypeByTag[int(field.Tag)]
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

	if err := prototype.Read(field.Value); err != nil {
		return IncorrectDataFormatForValue(field.Tag)
	}

	return nil
}
