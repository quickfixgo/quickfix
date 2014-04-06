package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/tag"
)

//validate tests the message against the provided data dictionary.
func validate(d *datadictionary.DataDictionary, message Message) (reject MessageReject) {
	msgType := new(field.StringField)
	if err := message.Header.GetField(tag.MsgType, msgType); err != nil {
		switch err.(type) {
		case FieldNotFoundError:
			return newRequiredTagMissing(message, tag.MsgType)
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

func validateFIXTApp(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, message Message) (reject MessageReject) {
	msgType := new(field.StringField)
	if err := message.Header.GetField(tag.MsgType, msgType); err != nil {
		switch err.(type) {
		case FieldNotFoundError:
			return newRequiredTagMissing(message, tag.MsgType)
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

	msgType := new(field.StringField)
	message.Header.GetField(tag.MsgType, msgType)

	if remainingFields, err = validateWalkComponent(appDD.Messages[msgType.Value], remainingFields, message); err != nil {
		return err
	}

	if remainingFields, err = validateWalkComponent(transportDD.Trailer, remainingFields, message); err != nil {
		return err
	}

	if len(remainingFields) != 0 {
		return newTagNotDefinedForThisMessageType(message, remainingFields[0].Tag)
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
			return nil, newTagAppearsMoreThanOnce(msg, field.Tag)
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
	numInGroup := new(field.IntValue)

	if err := numInGroup.Read(fieldStack[0].Value); err != nil {
		return nil, newValueIsIncorrect(message, &numInGroupTag)
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
				return fieldStack, newRequiredTagMissing(message, childDefs[0].Tag)
			}
		}

		childDefs = childDefs[1:]
	}

	if groupCount != numInGroup.Value {
		return fieldStack, newIncorrectNumInGroupCountForRepeatingGroup(message, numInGroupTag)
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
			return newTagSpecifiedOutOfRequiredOrder(message, tag)
		case tag.IsTrailer():
			inTrailer = true
		case inTrailer && !tag.IsTrailer():
			return newTagSpecifiedOutOfRequiredOrder(message, tag)
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
		field := new(field.StringValue)
		if err := fieldMap.GetField(required, field); err != nil {
			switch e := err.(type) {
			case FieldNotFoundError:
				return newRequiredTagMissing(msg, e.Tag)
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
			return newTagSpecifiedWithoutAValue(message, tag)
		}
	}

	for tag := range fieldMap.lookup {
		if _, valid := d.FieldTypeByTag[tag]; !valid {
			return newInvalidTagNumber(message, tag)
		}
	}

	for tag, fieldValue := range fieldMap.lookup {
		allowedValues := d.FieldTypeByTag[tag].Enums
		if len(allowedValues) != 0 {
			if _, validValue := allowedValues[string(fieldValue.Value)]; !validValue {
				return newValueIsIncorrect(message, &tag)
			}
		}
	}

	for tag := range fieldMap.lookup {
		fieldType := d.FieldTypeByTag[tag]
		var prototype FieldValue
		switch fieldType.Type {
		case "STRING":
			prototype = new(field.StringValue)
		case "MULTIPLESTRINGVALUE", "MULTIPLEVALUESTRING":
			prototype = new(field.MultipleStringValue)
		case "MULTIPLECHARVALUE":
			prototype = new(field.MultipleCharValue)
		case "CHAR":
			prototype = new(field.CharValue)
		case "CURRENCY":
			prototype = new(field.CurrencyValue)
		case "DATA":
			prototype = new(field.DataValue)
		case "MONTHYEAR":
			prototype = new(field.MonthYearValue)
		case "LOCALMKTDATE", "DATE":
			prototype = new(field.LocalMktDateValue)
		case "EXCHANGE":
			prototype = new(field.ExchangeValue)
		case "LANGUAGE":
			prototype = new(field.LanguageValue)
		case "XMLDATA":
			prototype = new(field.XMLDataValue)
		case "COUNTRY":
			prototype = new(field.CountryValue)
		case "UTCTIMEONLY":
			prototype = new(field.UTCTimeOnlyValue)
		case "UTCDATEONLY", "UTCDATE":
			prototype = new(field.UTCDateOnlyValue)
		case "TZTIMEONLY":
			prototype = new(field.TZTimeOnlyValue)
		case "TZTIMESTAMP":
			prototype = new(field.TZTimestampValue)
		case "BOOLEAN":
			prototype = new(field.BooleanValue)
		case "INT":
			prototype = new(field.IntValue)
		case "LENGTH":
			prototype = new(field.LengthValue)
		case "DAYOFMONTH":
			prototype = new(field.DayOfMonthValue)
		case "NUMINGROUP":
			prototype = new(field.NumInGroupValue)
		case "SEQNUM":
			prototype = new(field.SeqNumValue)
		case "UTCTIMESTAMP", "TIME":
			prototype = new(field.UTCTimestampValue)
		case "FLOAT":
			prototype = new(field.FloatValue)
		case "QTY", "QUANTITY":
			prototype = new(field.QtyValue)
		case "AMT":
			prototype = new(field.AmtValue)
		case "PRICE":
			prototype = new(field.PriceValue)
		case "PRICEOFFSET":
			prototype = new(field.PriceOffsetValue)
		case "PERCENTAGE":
			prototype = new(field.PercentageValue)
		}

		if err := fieldMap.GetField(tag, prototype); err != nil {
			return newIncorrectDataFormatForValue(message, tag)
		}
	}

	return nil
}
