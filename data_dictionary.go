package quickfixgo

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/spec"
	"github.com/cbusbey/quickfixgo/tag"
)

type tagSet map[tag.Tag]struct{}

func (t tagSet) add(tag tag.Tag) {
	t[tag] = struct{}{}
}

type enumSet map[string]struct{}

func (e enumSet) add(enum string) {
	e[enum] = struct{}{}
}

type fieldDefinition struct {
	enums     enumSet
	prototype FieldValue
}

func newFieldDefinition(fieldType spec.FieldType) (fieldDef fieldDefinition, err error) {
	fieldDef.enums = make(enumSet)

	switch fieldType.Type {
	case "STRING":
		fieldDef.prototype = new(StringValue)
	case "MULTIPLESTRINGVALUE", "MULTIPLEVALUESTRING":
		fieldDef.prototype = new(MultipleStringValue)
	case "MULTIPLECHARVALUE":
		fieldDef.prototype = new(MultipleCharValue)
	case "CHAR":
		fieldDef.prototype = new(CharValue)
	case "CURRENCY":
		fieldDef.prototype = new(CurrencyValue)
	case "DATA":
		fieldDef.prototype = new(DataValue)
	case "MONTHYEAR":
		fieldDef.prototype = new(MonthYearValue)
	case "LOCALMKTDATE", "DATE":
		fieldDef.prototype = new(LocalMktDateValue)
	case "EXCHANGE":
		fieldDef.prototype = new(ExchangeValue)
	case "LANGUAGE":
		fieldDef.prototype = new(LanguageValue)
	case "XMLDATA":
		fieldDef.prototype = new(XMLDataValue)
	case "COUNTRY":
		fieldDef.prototype = new(CountryValue)
	case "UTCTIMEONLY":
		fieldDef.prototype = new(UTCTimeOnlyValue)
	case "UTCDATEONLY", "UTCDATE":
		fieldDef.prototype = new(UTCDateOnlyValue)
	case "TZTIMEONLY":
		fieldDef.prototype = new(TZTimeOnlyValue)
	case "TZTIMESTAMP":
		fieldDef.prototype = new(TZTimestampValue)
	case "BOOLEAN":
		fieldDef.prototype = new(BooleanValue)
	case "INT":
		fieldDef.prototype = new(IntValue)
	case "LENGTH":
		fieldDef.prototype = new(LengthValue)
	case "DAYOFMONTH":
		fieldDef.prototype = new(DayOfMonthValue)
	case "NUMINGROUP":
		fieldDef.prototype = new(NumInGroupValue)
	case "SEQNUM":
		fieldDef.prototype = new(SeqNumValue)
	case "UTCTIMESTAMP", "TIME":
		fieldDef.prototype = new(UTCTimestampValue)
	case "FLOAT":
		fieldDef.prototype = new(FloatValue)
	case "QTY", "QUANTITY":
		fieldDef.prototype = new(QtyValue)
	case "AMT":
		fieldDef.prototype = new(AmtValue)
	case "PRICE":
		fieldDef.prototype = new(PriceValue)
	case "PRICEOFFSET":
		fieldDef.prototype = new(PriceOffsetValue)
	case "PERCENTAGE":
		fieldDef.prototype = new(PercentageValue)
	default:
		err = fmt.Errorf("Unknown type '%v' for tag '%v'", fieldType.Type, fieldType.Number)
		return
	}

	for _, value := range fieldType.FieldValues {
		fieldDef.enums.add(value.Enum)
	}

	return
}

type DataDictionary struct {
	allFields map[tag.Tag]fieldDefinition

	messageTags  map[string]tagSet
	requiredTags map[string]tagSet

	headerTags         tagSet
	requiredHeaderTags tagSet

	trailerTags         tagSet
	requiredTrailerTags tagSet
}

func NewDataDictionary(path string) (*DataDictionary, error) {
	fixSpec, err := spec.ParseFixSpec(path)
	if err != nil {
		return nil, err
	}

	d := new(DataDictionary)

	d.allFields = make(map[tag.Tag]fieldDefinition)

	for _, fieldType := range fixSpec.FieldTypeMap {
		if fieldDef, err := newFieldDefinition(fieldType); err != nil {
			return nil, err
		} else {
			d.allFields[tag.Tag(fieldType.Number)] = fieldDef
		}
	}

	d.messageTags = make(map[string]tagSet)
	d.requiredTags = make(map[string]tagSet)

	collectTags := func(fieldSpecs []spec.Field, components []spec.Component) (allTags tagSet, requiredTags tagSet, err error) {
		allTags = make(tagSet)
		requiredTags = make(tagSet)

		for _, component := range components {
			if componentType, ok := fixSpec.ComponentTypeMap[component.Name]; ok {
				fieldSpecs = append(fieldSpecs, componentType.Fields...)
			} else {
				err = fmt.Errorf("Component %q not given in components section", component.Name)
				return
			}
		}

		for _, field := range fieldSpecs {
			if fieldType, ok := fixSpec.FieldTypeMap[field.Name]; ok {
				tag := tag.Tag(fieldType.Number)
				allTags.add(tag)
				if field.Required == "Y" {
					requiredTags.add(tag)
				}
			} else {
				err = fmt.Errorf("Field %q not given in fields section", field.Name)
				return
			}
		}

		return
	}

	for _, msg := range fixSpec.Messages {
		if d.messageTags[msg.MsgType], d.requiredTags[msg.MsgType], err = collectTags(msg.Fields, msg.Components); err != nil {
			return nil, err
		}
	}

	if d.headerTags, d.requiredHeaderTags, err = collectTags(fixSpec.Header, []spec.Component{}); err != nil {
		return nil, err
	}

	if d.trailerTags, d.requiredTrailerTags, err = collectTags(fixSpec.Trailer, []spec.Component{}); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *DataDictionary) validate(message Message) (reject MessageReject) {
	msgType := new(StringField)
	if err := message.Header.GetField(tag.MsgType, msgType); err != nil {
		switch err.(type) {
		case FieldNotFoundError:
			return NewRequiredTagMissing(message, tag.MsgType)
		default:
			panic(fmt.Sprintf("Unhandled error: %v", err))
		}
	}

	if _, validMsgType := d.messageTags[msgType.Value]; validMsgType == false {
		return NewInvalidMessageType(message)
	}

	if reject = d.checkRequired(msgType.Value, message); reject != nil {
		return
	}

	if reject = d.iterate(msgType.Value, message); reject != nil {
		return
	}

	return
}

func (d *DataDictionary) checkRequired(msgType string, message Message) (reject MessageReject) {
	if reject = d.checkRequiredFieldMap(message, d.requiredHeaderTags, message.Header); reject != nil {
		return
	}

	if reject = d.checkRequiredFieldMap(message, d.requiredTags[msgType], message.Body); reject != nil {
		return
	}

	if reject = d.checkRequiredFieldMap(message, d.requiredTrailerTags, message.Trailer); reject != nil {
		return
	}

	return
}

func (d *DataDictionary) checkRequiredFieldMap(msg Message, requiredTags map[tag.Tag]struct{}, fieldMap FieldMap) (reject MessageReject) {
	for required := range requiredTags {
		//FIXME ugly...
		field := new(StringField)
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

func (d *DataDictionary) iterate(msgType string, message Message) (reject MessageReject) {
	if reject = d.iterateFieldMap(message, d.headerTags, message.Header); reject != nil {
		return
	}
	if reject = d.iterateFieldMap(message, d.messageTags[msgType], message.Body); reject != nil {
		return
	}
	if reject = d.iterateFieldMap(message, d.trailerTags, message.Trailer); reject != nil {
		return
	}

	return
}

func (d *DataDictionary) iterateFieldMap(message Message, validFields tagSet, fieldMap FieldMap) MessageReject {
	for tag, fieldValue := range fieldMap.fields {
		if len(fieldValue.Value) == 0 {
			return NewTagSpecifiedWithoutAValue(message, tag)
		}
	}

	for tag := range fieldMap.fields {
		if _, valid := d.allFields[tag]; !valid {
			return NewInvalidTagNumber(message, tag)
		}
	}

	for tag, fieldValue := range fieldMap.fields {
		allowedValues := d.allFields[tag].enums
		if len(allowedValues) != 0 {
			if _, validValue := allowedValues[string(fieldValue.Value)]; !validValue {
				return NewValueIsIncorrect(message, &tag)
			}
		}
	}

	for tag, _ := range fieldMap.fields {
		prototype := d.allFields[tag].prototype
		if err := fieldMap.GetField(tag, prototype); err != nil {
			return NewIncorrectDataFormatForValue(message, tag)
		}
	}

	for tag := range fieldMap.fields {
		if _, valid := validFields[tag]; !valid {
			return NewTagNotDefinedForThisMessageType(message, tag)
		}
	}

	return nil
}
