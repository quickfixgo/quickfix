package quickfixgo

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/datadictionary"
	"github.com/cbusbey/quickfixgo/tag"
)

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

	if reject = ValidateRequired(d, msgType.Value, message); reject != nil {
		return
	}

	if reject = ValidateOrder(message); reject != nil {
		return
	}

	if reject = validateStructure(d, msgType.Value, message); reject != nil {
		return
	}

	return nil
}

func ValidateOrder(message Message) (reject MessageReject) {

	inHeader := true
	inTrailer := false
	for _, tag := range message.tags {
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

func ValidateRequired(d *datadictionary.DataDictionary, msgType string, message Message) (reject MessageReject) {
	if reject = validateRequiredFieldMap(message, d.Header.RequiredTags, message.Header); reject != nil {
		return
	}

	if reject = validateRequiredFieldMap(message, d.Messages[msgType].RequiredTags, message.Body); reject != nil {
		return
	}

	if reject = validateRequiredFieldMap(message, d.Trailer.RequiredTags, message.Trailer); reject != nil {
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

func validateStructure(d *datadictionary.DataDictionary, msgType string, message Message) (reject MessageReject) {
	if reject = validateFieldMapStructure(d, message, d.Header.Tags, message.Header); reject != nil {
		return
	}
	if reject = validateFieldMapStructure(d, message, d.Messages[msgType].Tags, message.Body); reject != nil {
		return
	}
	if reject = validateFieldMapStructure(d, message, d.Trailer.Tags, message.Trailer); reject != nil {
		return
	}

	return
}

func validateFieldMapStructure(d *datadictionary.DataDictionary, message Message, validFields datadictionary.TagSet, fieldMap FieldMap) MessageReject {
	for tag, fieldValue := range fieldMap.lookup {
		if len(fieldValue.Value) == 0 {
			return NewTagSpecifiedWithoutAValue(message, tag)
		}
	}

	iteratedTags := make(datadictionary.TagSet)
	for _, tag := range fieldMap.tags {
		if _, duplicate := iteratedTags[tag]; duplicate {
			return NewTagAppearsMoreThanOnce(message, tag)
		}
		iteratedTags.Add(tag)
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

	for tag, _ := range fieldMap.lookup {
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
		default:
			//FIXME
			//err := fmt.Errorf("Unknown type '%v' for tag '%v'", fieldType.Type, fieldType.Tag)
			//			return
		}

		if err := fieldMap.GetField(tag, prototype); err != nil {
			return NewIncorrectDataFormatForValue(message, tag)
		}
	}

	for tag := range fieldMap.lookup {
		if _, valid := validFields[tag]; !valid {
			return NewTagNotDefinedForThisMessageType(message, tag)
		}
	}

	return nil
}
