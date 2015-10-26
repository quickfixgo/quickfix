package quickfix

import (
	"github.com/quickfixgo/quickfix/datadictionary"
)

//validate tests the message against the provided data dictionary.
func validate(d *datadictionary.DataDictionary, msg Message) MessageRejectError {
	msgType := new(StringField)
	if err := msg.Header.GetField(tagMsgType, msgType); err != nil {
		if err.RejectReason() == rejectReasonConditionallyRequiredFieldMissing {
			return requiredTagMissing(tagMsgType)
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
	msgType := new(StringField)
	if err := msg.Header.GetField(tagMsgType, msgType); err != nil {
		if err.RejectReason() == rejectReasonConditionallyRequiredFieldMissing {
			return requiredTagMissing(tagMsgType)
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

	msgType := new(StringField)
	msg.Header.GetField(tagMsgType, msgType)

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

func validateWalkComponent(messageDef *datadictionary.MessageDef, fields []TagValue) ([]TagValue, MessageRejectError) {
	var fieldDef *datadictionary.FieldDef
	var ok bool
	var err MessageRejectError
	iteratedTags := make(datadictionary.TagSet)

	for len(fields) > 0 {
		field := fields[0]
		//field not defined for this component
		if fieldDef, ok = messageDef.Fields[int(field.Tag)]; !ok {
			break
		}

		if _, duplicate := iteratedTags[int(field.Tag)]; duplicate {
			return nil, tagAppearsMoreThanOnce(field.Tag)
		}
		iteratedTags.Add(int(field.Tag))

		if fields, err = validateVisitField(fieldDef, fields); err != nil {
			return nil, err
		}
	}

	return fields, nil
}

func validateVisitField(fieldDef *datadictionary.FieldDef, fields []TagValue) ([]TagValue, MessageRejectError) {
	var err MessageRejectError

	if fieldDef.IsGroup() {
		if fields, err = validateVisitGroupField(fieldDef, fields); err != nil {
			return nil, err
		}
	}

	return fields[1:], nil
}

func validateVisitGroupField(fieldDef *datadictionary.FieldDef, fieldStack []TagValue) ([]TagValue, MessageRejectError) {
	numInGroupTag := fieldStack[0].Tag
	numInGroup := new(IntValue)

	if err := numInGroup.Read(fieldStack); err != nil {
		return nil, incorrectDataFormatForValue(numInGroupTag)
	}

	fieldStack = fieldStack[1:]

	var childDefs []*datadictionary.FieldDef
	groupCount := 0

	for len(fieldStack) > 0 {

		//start of repeating group
		if int(fieldStack[0].Tag) == fieldDef.ChildFields[0].Tag {
			childDefs = fieldDef.ChildFields
			groupCount++
		}

		//group complete
		if len(childDefs) == 0 {
			break
		}

		if int(fieldStack[0].Tag) == childDefs[0].Tag {
			var err MessageRejectError
			if fieldStack, err = validateVisitField(childDefs[0], fieldStack); err != nil {
				return fieldStack, err
			}
		} else {
			if childDefs[0].Required {
				return fieldStack, requiredTagMissing(Tag(childDefs[0].Tag))
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
		field := new(StringValue)
		if err := fieldMap.GetField(Tag(required), field); err != nil {
			//FIXME: add "has..." method?
			if err.RejectReason() == rejectReasonConditionallyRequiredFieldMissing {
				return requiredTagMissing(Tag(required))
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

func validateField(d *datadictionary.DataDictionary, validFields datadictionary.TagSet, field TagValue) MessageRejectError {
	if len(field.Value) == 0 {
		return tagSpecifiedWithoutAValue(field.Tag)
	}

	if _, valid := d.FieldTypeByTag[int(field.Tag)]; !valid {
		return invalidTagNumber(field.Tag)
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
	case "STRING":
		prototype = new(StringValue)
	case "MULTIPLESTRINGVALUE", "MULTIPLEVALUESTRING":
		prototype = new(MultipleStringValue)
	case "MULTIPLECHARVALUE":
		prototype = new(MultipleCharValue)
	case "CHAR":
		prototype = new(CharValue)
	case "CURRENCY":
		prototype = new(CurrencyValue)
	case "DATA":
		prototype = new(DataValue)
	case "MONTHYEAR":
		prototype = new(MonthYearValue)
	case "LOCALMKTDATE", "DATE":
		prototype = new(LocalMktDateValue)
	case "EXCHANGE":
		prototype = new(ExchangeValue)
	case "LANGUAGE":
		prototype = new(LanguageValue)
	case "XMLDATA":
		prototype = new(XMLDataValue)
	case "COUNTRY":
		prototype = new(CountryValue)
	case "UTCTIMEONLY":
		prototype = new(UTCTimeOnlyValue)
	case "UTCDATEONLY", "UTCDATE":
		prototype = new(UTCDateOnlyValue)
	case "TZTIMEONLY":
		prototype = new(TZTimeOnlyValue)
	case "TZTIMESTAMP":
		prototype = new(TZTimestampValue)
	case "BOOLEAN":
		prototype = new(BooleanValue)
	case "INT":
		prototype = new(IntValue)
	case "LENGTH":
		prototype = new(LengthValue)
	case "DAYOFMONTH":
		prototype = new(DayOfMonthValue)
	case "NUMINGROUP":
		prototype = new(NumInGroupValue)
	case "SEQNUM":
		prototype = new(SeqNumValue)
	case "UTCTIMESTAMP", "TIME":
		prototype = new(UTCTimestampValue)
	case "FLOAT":
		prototype = new(FloatValue)
	case "QTY", "QUANTITY":
		prototype = new(QtyValue)
	case "AMT":
		prototype = new(AmtValue)
	case "PRICE":
		prototype = new(PriceValue)
	case "PRICEOFFSET":
		prototype = new(PriceOffsetValue)
	case "PERCENTAGE":
		prototype = new(PercentageValue)
	}

	if err := prototype.Read([]TagValue{field}); err != nil {
		return incorrectDataFormatForValue(field.Tag)
	}

	return nil
}
