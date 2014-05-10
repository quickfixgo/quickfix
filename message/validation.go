package message

import (
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
)

//Validate tests the message against the provided data dictionary.
func Validate(d *datadictionary.DataDictionary, msg Message) errors.MessageRejectError {
	msgType := new(StringField)
	if err := msg.Header.GetField(tag.MsgType, msgType); err != nil {
		if err.RejectReason() == errors.RejectReasonConditionallyRequiredFieldMissing {
			return errors.RequiredTagMissing(tag.MsgType)
		}

		return err
	}

	if _, validMsgType := d.Messages[msgType.Value]; validMsgType == false {
		return errors.InvalidMessageType()
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

func ValidateFIXTApp(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msg Message) errors.MessageRejectError {
	msgType := new(StringField)
	if err := msg.Header.GetField(tag.MsgType, msgType); err != nil {
		if err.RejectReason() == errors.RejectReasonConditionallyRequiredFieldMissing {
			return errors.RequiredTagMissing(tag.MsgType)
		}

		return err
	}

	if _, validMsgType := appDD.Messages[msgType.Value]; validMsgType == false {
		return errors.InvalidMessageType()
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

func validateWalk(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msg Message) errors.MessageRejectError {
	remainingFields := msg.fields
	var err errors.MessageRejectError

	if remainingFields, err = validateWalkComponent(transportDD.Header, remainingFields); err != nil {
		return err
	}

	msgType := new(StringField)
	msg.Header.GetField(tag.MsgType, msgType)

	if remainingFields, err = validateWalkComponent(appDD.Messages[msgType.Value], remainingFields); err != nil {
		return err
	}

	if remainingFields, err = validateWalkComponent(transportDD.Trailer, remainingFields); err != nil {
		return err
	}

	if len(remainingFields) != 0 {
		return errors.TagNotDefinedForThisMessageType(remainingFields[0].Tag)
	}

	return nil
}

func validateWalkComponent(messageDef *datadictionary.MessageDef, fields []*fieldBytes) ([]*fieldBytes, errors.MessageRejectError) {
	var fieldDef *datadictionary.FieldDef
	var ok bool
	var err errors.MessageRejectError
	iteratedTags := make(datadictionary.TagSet)

	for len(fields) > 0 {
		field := fields[0]
		//field not defined for this component
		if fieldDef, ok = messageDef.Fields[field.Tag]; !ok {
			break
		}

		if _, duplicate := iteratedTags[field.Tag]; duplicate {
			return nil, errors.TagAppearsMoreThanOnce(field.Tag)
		}
		iteratedTags.Add(field.Tag)

		if fields, err = validateVisitField(fieldDef, fields); err != nil {
			return nil, err
		}
	}

	return fields, nil
}

func validateVisitField(fieldDef *datadictionary.FieldDef, fields []*fieldBytes) ([]*fieldBytes, errors.MessageRejectError) {
	var err errors.MessageRejectError

	if fieldDef.IsGroup() {
		if fields, err = validateVisitGroupField(fieldDef, fields); err != nil {
			return nil, err
		}
	}

	return fields[1:], nil
}

func validateVisitGroupField(fieldDef *datadictionary.FieldDef, fieldStack []*fieldBytes) ([]*fieldBytes, errors.MessageRejectError) {
	numInGroupTag := fieldStack[0].Tag
	numInGroup := new(IntValue)

	if err := numInGroup.Read(fieldStack[0].Value); err != nil {
		return nil, errors.IncorrectDataFormatForValue(numInGroupTag)
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
			var err errors.MessageRejectError
			if fieldStack, err = validateVisitField(childDefs[0], fieldStack); err != nil {
				return fieldStack, err
			}
		} else {
			if childDefs[0].Required {
				return fieldStack, errors.RequiredTagMissing(childDefs[0].Tag)
			}
		}

		childDefs = childDefs[1:]
	}

	if groupCount != numInGroup.Value {
		return fieldStack, errors.IncorrectNumInGroupCountForRepeatingGroup(numInGroupTag)
	}

	return fieldStack, nil
}

func validateOrder(msg Message) errors.MessageRejectError {

	inHeader := true
	inTrailer := false
	for _, field := range msg.fields {
		t := field.Tag
		switch {
		case inHeader && tag.IsHeader(t):
		case inHeader && !tag.IsHeader(t):
			inHeader = false
		case !inHeader && tag.IsHeader(t):
			return errors.TagSpecifiedOutOfRequiredOrder(t)
		case tag.IsTrailer(t):
			inTrailer = true
		case inTrailer && !tag.IsTrailer(t):
			return errors.TagSpecifiedOutOfRequiredOrder(t)
		}
	}

	return nil
}

func validateRequired(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msgType string, message Message) errors.MessageRejectError {
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

func validateRequiredFieldMap(msg Message, requiredTags map[fix.Tag]struct{}, fieldMap FieldMap) errors.MessageRejectError {
	for required := range requiredTags {
		field := new(StringValue)
		if err := fieldMap.GetField(required, field); err != nil {
			//FIXME: add "has..." method?
			if err.RejectReason() == errors.RejectReasonConditionallyRequiredFieldMissing {
				return errors.RequiredTagMissing(required)
			}
			return err
		}
	}

	return nil
}

func validateFields(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msgType string, message Message) errors.MessageRejectError {
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

func validateFieldMapFields(d *datadictionary.DataDictionary, message Message, validFields datadictionary.TagSet, fieldMap FieldMap) errors.MessageRejectError {
	for tag, fieldValue := range fieldMap.fieldLookup {
		if len(fieldValue.Value) == 0 {
			return errors.TagSpecifiedWithoutAValue(tag)
		}
	}

	for tag := range fieldMap.fieldLookup {
		if _, valid := d.FieldTypeByTag[tag]; !valid {
			return errors.InvalidTagNumber(tag)
		}
	}

	for tag, fieldValue := range fieldMap.fieldLookup {
		allowedValues := d.FieldTypeByTag[tag].Enums
		if len(allowedValues) != 0 {
			if _, validValue := allowedValues[string(fieldValue.Value)]; !validValue {
				return errors.ValueIsIncorrect(tag)
			}
		}
	}

	for tag := range fieldMap.fieldLookup {
		fieldType := d.FieldTypeByTag[tag]
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

		if err := fieldMap.GetField(tag, prototype); err != nil {
			return err
		}
	}

	return nil
}
