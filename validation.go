package quickfix

import (
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
)

//validate tests the message against the provided data dictionary.
func validate(d *datadictionary.DataDictionary, msg Message) MessageRejectError {
	msgType := new(fix.StringField)
	if err := msg.Header.GetField(tag.MsgType, msgType); err != nil {
		if err.RejectReason() == rejectReasonConditionallyRequiredFieldMissing {
			return requiredTagMissing(tag.MsgType)
		}

		return err
	}

	if _, validMsgType := d.Messages[msgType.Value]; validMsgType == false {
		return invalidMessageType()
	}

	if err := validateRequired(d, d, msgType.Value, msg); err != nil {
		return err
	}

	if err := validateOrder(msg); err != nil {
		return err
	}

	if err := validateFields(d, d, msgType.Value, msg); err != nil {
		return err
	}

	if err := validateWalk(d, d, msg); err != nil {
		return err
	}

	return nil
}

func validateFIXTApp(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msg Message) MessageRejectError {
	msgType := new(fix.StringField)
	if err := msg.Header.GetField(tag.MsgType, msgType); err != nil {
		if err.RejectReason() == rejectReasonConditionallyRequiredFieldMissing {
			return requiredTagMissing(tag.MsgType)
		}

		return err
	}

	if _, validMsgType := appDD.Messages[msgType.Value]; validMsgType == false {
		return invalidMessageType()
	}

	if err := validateRequired(transportDD, appDD, msgType.Value, msg); err != nil {
		return err
	}

	if err := validateOrder(msg); err != nil {
		return err
	}

	if err := validateWalk(transportDD, appDD, msg); err != nil {
		return err
	}

	if err := validateFields(transportDD, appDD, msgType.Value, msg); err != nil {
		return err
	}

	return nil
}

func validateWalk(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msg Message) MessageRejectError {
	remainingFields := msg.fields
	var err MessageRejectError

	if remainingFields, err = validateWalkComponent(transportDD.Header, remainingFields); err != nil {
		return err
	}

	msgType := new(fix.StringField)
	msg.Header.GetField(tag.MsgType, msgType)

	if remainingFields, err = validateWalkComponent(appDD.Messages[msgType.Value], remainingFields); err != nil {
		return err
	}

	if remainingFields, err = validateWalkComponent(transportDD.Trailer, remainingFields); err != nil {
		return err
	}

	if len(remainingFields) != 0 {
		return tagNotDefinedForThisMessageType(remainingFields[0].Tag)
	}

	return nil
}

func validateWalkComponent(messageDef *datadictionary.MessageDef, fields []*fieldBytes) ([]*fieldBytes, MessageRejectError) {
	var fieldDef *datadictionary.FieldDef
	var ok bool
	var err MessageRejectError
	iteratedTags := make(datadictionary.TagSet)

	for len(fields) > 0 {
		field := fields[0]
		//field not defined for this component
		if fieldDef, ok = messageDef.Fields[field.Tag]; !ok {
			break
		}

		if _, duplicate := iteratedTags[field.Tag]; duplicate {
			return nil, tagAppearsMoreThanOnce(field.Tag)
		}
		iteratedTags.Add(field.Tag)

		if fields, err = validateVisitField(fieldDef, fields); err != nil {
			return nil, err
		}
	}

	return fields, nil
}

func validateVisitField(fieldDef *datadictionary.FieldDef, fields []*fieldBytes) ([]*fieldBytes, MessageRejectError) {
	var err MessageRejectError

	if fieldDef.IsGroup() {
		if fields, err = validateVisitGroupField(fieldDef, fields); err != nil {
			return nil, err
		}
	}

	return fields[1:], nil
}

func validateVisitGroupField(fieldDef *datadictionary.FieldDef, fieldStack []*fieldBytes) ([]*fieldBytes, MessageRejectError) {
	numInGroupTag := fieldStack[0].Tag
	numInGroup := new(fix.IntValue)

	if err := numInGroup.Read(fieldStack[0].Value); err != nil {
		return nil, incorrectDataFormatForValue(numInGroupTag)
	}

	fieldStack = fieldStack[1:]

	var childDefs []*datadictionary.FieldDef
	groupCount := 0

	for len(fieldStack) > 0 {

		//start of repeating group
		if fieldStack[0].Tag == fieldDef.ChildFields[0].Tag {
			childDefs = fieldDef.ChildFields
			groupCount++
		}

		//group complete
		if len(childDefs) == 0 {
			break
		}

		if fieldStack[0].Tag == childDefs[0].Tag {
			var err MessageRejectError
			if fieldStack, err = validateVisitField(childDefs[0], fieldStack); err != nil {
				return fieldStack, err
			}
		} else {
			if childDefs[0].Required {
				return fieldStack, requiredTagMissing(childDefs[0].Tag)
			}
		}

		childDefs = childDefs[1:]
	}

	if groupCount != numInGroup.Value {
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
		case inHeader && tag.IsHeader(t):
		case inHeader && !tag.IsHeader(t):
			inHeader = false
		case !inHeader && tag.IsHeader(t):
			return tagSpecifiedOutOfRequiredOrder(t)
		case tag.IsTrailer(t):
			inTrailer = true
		case inTrailer && !tag.IsTrailer(t):
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

func validateRequiredFieldMap(msg Message, requiredTags map[fix.Tag]struct{}, fieldMap FieldMap) MessageRejectError {
	for required := range requiredTags {
		field := new(fix.StringValue)
		if err := fieldMap.GetField(required, field); err != nil {
			//FIXME: add "has..." method?
			if err.RejectReason() == rejectReasonConditionallyRequiredFieldMissing {
				return requiredTagMissing(required)
			}
			return err
		}
	}

	return nil
}

func validateFields(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msgType string, message Message) MessageRejectError {
	if err := validateFieldMapFields(transportDD, message, transportDD.Header.Tags, message.Header); err != nil {
		return err
	}
	if err := validateFieldMapFields(appDD, message, appDD.Messages[msgType].Tags, message.Body); err != nil {
		return err
	}
	if err := validateFieldMapFields(transportDD, message, transportDD.Trailer.Tags, message.Trailer); err != nil {
		return err
	}

	return nil
}

func validateFieldMapFields(d *datadictionary.DataDictionary, message Message, validFields datadictionary.TagSet, fieldMap FieldMap) MessageRejectError {
	for tag, fieldValue := range fieldMap.fieldLookup {
		if len(fieldValue.Value) == 0 {
			return tagSpecifiedWithoutAValue(tag)
		}
	}

	for tag := range fieldMap.fieldLookup {
		if _, valid := d.FieldTypeByTag[tag]; !valid {
			return invalidTagNumber(tag)
		}
	}

	for tag, fieldValue := range fieldMap.fieldLookup {
		allowedValues := d.FieldTypeByTag[tag].Enums
		if len(allowedValues) != 0 {
			if _, validValue := allowedValues[string(fieldValue.Value)]; !validValue {
				return valueIsIncorrect(tag)
			}
		}
	}

	for tag := range fieldMap.fieldLookup {
		fieldType := d.FieldTypeByTag[tag]
		var prototype FieldValue
		switch fieldType.Type {
		case "STRING":
			prototype = new(fix.StringValue)
		case "MULTIPLESTRINGVALUE", "MULTIPLEVALUESTRING":
			prototype = new(fix.MultipleStringValue)
		case "MULTIPLECHARVALUE":
			prototype = new(fix.MultipleCharValue)
		case "CHAR":
			prototype = new(fix.CharValue)
		case "CURRENCY":
			prototype = new(fix.CurrencyValue)
		case "DATA":
			prototype = new(fix.DataValue)
		case "MONTHYEAR":
			prototype = new(fix.MonthYearValue)
		case "LOCALMKTDATE", "DATE":
			prototype = new(fix.LocalMktDateValue)
		case "EXCHANGE":
			prototype = new(fix.ExchangeValue)
		case "LANGUAGE":
			prototype = new(fix.LanguageValue)
		case "XMLDATA":
			prototype = new(fix.XMLDataValue)
		case "COUNTRY":
			prototype = new(fix.CountryValue)
		case "UTCTIMEONLY":
			prototype = new(fix.UTCTimeOnlyValue)
		case "UTCDATEONLY", "UTCDATE":
			prototype = new(fix.UTCDateOnlyValue)
		case "TZTIMEONLY":
			prototype = new(fix.TZTimeOnlyValue)
		case "TZTIMESTAMP":
			prototype = new(fix.TZTimestampValue)
		case "BOOLEAN":
			prototype = new(fix.BooleanValue)
		case "INT":
			prototype = new(fix.IntValue)
		case "LENGTH":
			prototype = new(fix.LengthValue)
		case "DAYOFMONTH":
			prototype = new(fix.DayOfMonthValue)
		case "NUMINGROUP":
			prototype = new(fix.NumInGroupValue)
		case "SEQNUM":
			prototype = new(fix.SeqNumValue)
		case "UTCTIMESTAMP", "TIME":
			prototype = new(fix.UTCTimestampValue)
		case "FLOAT":
			prototype = new(fix.FloatValue)
		case "QTY", "QUANTITY":
			prototype = new(fix.QtyValue)
		case "AMT":
			prototype = new(fix.AmtValue)
		case "PRICE":
			prototype = new(fix.PriceValue)
		case "PRICEOFFSET":
			prototype = new(fix.PriceOffsetValue)
		case "PERCENTAGE":
			prototype = new(fix.PercentageValue)
		}

		if err := fieldMap.GetField(tag, prototype); err != nil {
			return err
		}
	}

	return nil
}
