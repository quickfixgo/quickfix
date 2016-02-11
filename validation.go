package quickfix

import (
	"github.com/quickfixgo/quickfix/datadictionary"
)

//validate tests the message against the provided data dictionary.
func validate(d *datadictionary.DataDictionary, msg Message) MessageRejectError {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err != nil {
		if err.RejectReason() == rejectReasonConditionallyRequiredFieldMissing {
			return requiredTagMissing(tagMsgType)
		}

		return err
	}

	if _, validMsgType := d.Messages[string(msgType)]; validMsgType == false {
		return invalidMessageType()
	}

	if err := validateRequired(d, d, string(msgType), msg); err != nil {
		return err
	}

	if err := validateOrder(msg); err != nil {
		return err
	}

	if err := validateFields(d, d, string(msgType), msg); err != nil {
		return err
	}

	if err := validateWalk(d, d, msg); err != nil {
		return err
	}

	return nil
}

func validateFIXTApp(transportDD *datadictionary.DataDictionary, appDD *datadictionary.DataDictionary, msg Message) MessageRejectError {
	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err != nil {
		if err.RejectReason() == rejectReasonConditionallyRequiredFieldMissing {
			return requiredTagMissing(tagMsgType)
		}

		return err
	}

	if _, validMsgType := appDD.Messages[string(msgType)]; validMsgType == false {
		return invalidMessageType()
	}

	if err := validateRequired(transportDD, appDD, string(msgType), msg); err != nil {
		return err
	}

	if err := validateOrder(msg); err != nil {
		return err
	}

	if err := validateWalk(transportDD, appDD, msg); err != nil {
		return err
	}

	if err := validateFields(transportDD, appDD, string(msgType), msg); err != nil {
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

	var msgType FIXString
	msg.Header.GetField(tagMsgType, &msgType)

	if remainingFields, err = validateWalkComponent(appDD.Messages[string(msgType)], remainingFields); err != nil {
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

func validateWalkComponent(messageDef *datadictionary.MessageDef, fields []tagValue) ([]tagValue, MessageRejectError) {
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

func validateVisitField(fieldDef *datadictionary.FieldDef, fields []tagValue) ([]tagValue, MessageRejectError) {
	var err MessageRejectError

	if fieldDef.IsGroup() {
		if fields, err = validateVisitGroupField(fieldDef, fields); err != nil {
			return nil, err
		}
	}

	return fields[1:], nil
}

func validateVisitGroupField(fieldDef *datadictionary.FieldDef, fieldStack []tagValue) ([]tagValue, MessageRejectError) {
	numInGroupTag := fieldStack[0].Tag
	var numInGroup FIXInt

	if err := numInGroup.Read(fieldStack[0].Value); err != nil {
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

func validateField(d *datadictionary.DataDictionary, validFields datadictionary.TagSet, field tagValue) MessageRejectError {
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
		return incorrectDataFormatForValue(field.Tag)
	}

	return nil
}
