package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/tag"
)

//Validate tests the message against the provided data dictionary.
func Validate(d *datadictionary.DataDictionary, message Message) (reject MessageReject) {
	msgType := new(StringField)
	if err := message.Header.GetField(tag.MsgType, msgType); err != nil {
		switch err.(type) {
		case FieldNotFoundError:
			return NewRequiredTagMissing(message, tag.MsgType)
		default:
			panic(fmt.Sprintf("Unhandled error: %v", err))
		}
	}

	if _, validMsgType := d.Messages[msgType.Value]; validMsgType == false {
		return NewInvalidMessageType(message)
	}

	if reject = validateRequired(d, d, msgType.Value, message); reject != nil {
		return
	}

	if reject = validateOrder(message); reject != nil {
		return
	}

	if reject = validateFields(d, d, msgType.Value, message); reject != nil {
		return
	}

	if reject = validateWalk(d, d, message); reject != nil {
		return
	}

	return nil
}

//Validate tests the message against the provided data dictionary.
func ValidateFIXTApp(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, message Message) (reject MessageReject) {
	msgType := new(StringField)
	if err := message.Header.GetField(tag.MsgType, msgType); err != nil {
		switch err.(type) {
		case FieldNotFoundError:
			return NewRequiredTagMissing(message, tag.MsgType)
		default:
			panic(fmt.Sprintf("Unhandled error: %v", err))
		}
	}

	if _, validMsgType := appDD.Messages[msgType.Value]; validMsgType == false {
		return NewInvalidMessageType(message)
	}

	if reject = validateRequired(transportDD, appDD, msgType.Value, message); reject != nil {
		return
	}

	if reject = validateOrder(message); reject != nil {
		return
	}

	if reject = validateWalk(transportDD, appDD, message); reject != nil {
		return
	}

	if reject = validateFields(transportDD, appDD, msgType.Value, message); reject != nil {
		return
	}

	return nil
}

func validateWalk(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, message Message) (reject MessageReject) {
	remainingFields := message.fields
	var err MessageReject

	if remainingFields, err = validateWalkComponent(transportDD.Header, remainingFields, message); err != nil {
		return err
	}

	msgType := new(StringField)
	message.Header.GetField(tag.MsgType, msgType)

	if remainingFields, err = validateWalkComponent(appDD.Messages[msgType.Value], remainingFields, message); err != nil {
		return err
	}

	if remainingFields, err = validateWalkComponent(transportDD.Trailer, remainingFields, message); err != nil {
		return err
	}

	if len(remainingFields) != 0 {
		return NewTagNotDefinedForThisMessageType(message, remainingFields[0].Tag)
	}

	return nil
}

func validateWalkComponent(messageDef *datadictionary.MessageDef, fields []*fieldBytes, msg Message) ([]*fieldBytes, MessageReject) {
	var fieldDef *datadictionary.FieldDef
	var ok bool
	var err MessageReject
	iteratedTags := make(datadictionary.TagSet)

	for len(fields) > 0 {
		field := fields[0]
		//field not defined for this component
		if fieldDef, ok = messageDef.Fields[field.Tag]; !ok {
			break
		}

		if _, duplicate := iteratedTags[field.Tag]; duplicate {
			return nil, NewTagAppearsMoreThanOnce(msg, field.Tag)
		}
		iteratedTags.Add(field.Tag)

		if fields, err = validateVisitField(fieldDef, fields, msg); err != nil {
			return nil, err
		}
	}

	return fields, nil
}

func validateVisitField(fieldDef *datadictionary.FieldDef, fields []*fieldBytes, message Message) ([]*fieldBytes, MessageReject) {
	var err MessageReject

	if fieldDef.IsGroup() {
		if fields, err = validateVisitGroupField(fieldDef, fields, message); err != nil {
			return nil, err
		}
	}

	return fields[1:], nil
}

func validateVisitGroupField(fieldDef *datadictionary.FieldDef, fieldStack []*fieldBytes, message Message) ([]*fieldBytes, MessageReject) {
	numInGroupTag := fieldStack[0].Tag
	numInGroup := new(IntValue)

	if err := numInGroup.Read(fieldStack[0].Value); err != nil {
		return nil, NewValueIsIncorrect(message, &numInGroupTag)
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
			var reject MessageReject
			if fieldStack, reject = validateVisitField(childDefs[0], fieldStack, message); reject != nil {
				return fieldStack, reject
			}
		} else {
			if childDefs[0].Required {
				return fieldStack, NewRequiredTagMissing(message, childDefs[0].Tag)
			}
		}

		childDefs = childDefs[1:]
	}

	if groupCount != numInGroup.Value {
		return fieldStack, NewIncorrectNumInGroupCountForRepeatingGroup(message, numInGroupTag)
	}

	return fieldStack, nil
}

func validateOrder(message Message) (reject MessageReject) {

	inHeader := true
	inTrailer := false
	for _, field := range message.fields {
		tag := field.Tag
		switch {
		case inHeader && tag.IsHeader():
		case inHeader && !tag.IsHeader():
			inHeader = false
		case !inHeader && tag.IsHeader():
			return NewTagSpecifiedOutOfRequiredOrder(message, tag)
		case tag.IsTrailer():
			inTrailer = true
		case inTrailer && !tag.IsTrailer():
			return NewTagSpecifiedOutOfRequiredOrder(message, tag)
		}
	}

	return nil
}

func validateRequired(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msgType string, message Message) (reject MessageReject) {
	if reject = validateRequiredFieldMap(message, transportDD.Header.RequiredTags, message.Header); reject != nil {
		return
	}

	if reject = validateRequiredFieldMap(message, appDD.Messages[msgType].RequiredTags, message.Body); reject != nil {
		return
	}

	if reject = validateRequiredFieldMap(message, transportDD.Trailer.RequiredTags, message.Trailer); reject != nil {
		return
	}

	return
}

func validateRequiredFieldMap(msg Message, requiredTags map[tag.Tag]struct{}, fieldMap FieldMap) (reject MessageReject) {
	for required := range requiredTags {
		field := new(StringValue)
		if err := fieldMap.GetField(required, field); err != nil {
			switch e := err.(type) {
			case FieldNotFoundError:
				return NewRequiredTagMissing(msg, e.Tag)
			default:
				panic(fmt.Sprintf("Unhandled error: %v", err))
			}
		}
	}

	return
}

func validateFields(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msgType string, message Message) (reject MessageReject) {
	if reject = validateFieldMapFields(transportDD, message, transportDD.Header.Tags, message.Header); reject != nil {
		return
	}
	if reject = validateFieldMapFields(appDD, message, appDD.Messages[msgType].Tags, message.Body); reject != nil {
		return
	}
	if reject = validateFieldMapFields(transportDD, message, transportDD.Trailer.Tags, message.Trailer); reject != nil {
		return
	}

	return
}

func validateFieldMapFields(d *datadictionary.DataDictionary, message Message, validFields datadictionary.TagSet, fieldMap FieldMap) MessageReject {
	for tag, fieldValue := range fieldMap.lookup {
		if len(fieldValue.Value) == 0 {
			return NewTagSpecifiedWithoutAValue(message, tag)
		}
	}

	for tag := range fieldMap.lookup {
		if _, valid := d.FieldTypeByTag[tag]; !valid {
			return NewInvalidTagNumber(message, tag)
		}
	}

	for tag, fieldValue := range fieldMap.lookup {
		allowedValues := d.FieldTypeByTag[tag].Enums
		if len(allowedValues) != 0 {
			if _, validValue := allowedValues[string(fieldValue.Value)]; !validValue {
				return NewValueIsIncorrect(message, &tag)
			}
		}
	}

	for tag := range fieldMap.lookup {
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
			return NewIncorrectDataFormatForValue(message, tag)
		}
	}

	return nil
}
