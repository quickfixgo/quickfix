package quickfixgo

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/field"
	"github.com/cbusbey/quickfixgo/spec"
	"github.com/cbusbey/quickfixgo/tag"
)

type DataDictionary struct {
	allFields      map[tag.Tag]struct{}
	messageFields  map[string]map[tag.Tag]struct{}
	requiredFields map[string]map[tag.Tag]struct{}

	headerFields         map[tag.Tag]struct{}
	headerRequiredFields map[tag.Tag]struct{}

	trailerFields         map[tag.Tag]struct{}
	trailerRequiredFields map[tag.Tag]struct{}
}

func NewDataDictionary(path string) (*DataDictionary, error) {
	fixSpec, err := spec.ParseFixSpec(path)
	if err != nil {
		return nil, err
	}

	d := DataDictionary{}

	d.allFields = make(map[tag.Tag]struct{})
	fieldNameLookup := make(map[string]tag.Tag)
	for _, fieldType := range fixSpec.FieldTypeMap {
		d.allFields[tag.Tag(fieldType.Number)] = struct{}{}
		fieldNameLookup[fieldType.Name] = tag.Tag(fieldType.Number)
	}

	d.messageFields = make(map[string]map[tag.Tag]struct{})
	d.requiredFields = make(map[string]map[tag.Tag]struct{})

	assignFields := func(fieldSpecs []spec.Field, allFields map[tag.Tag]struct{}, requiredFields map[tag.Tag]struct{}) error {
		for _, field := range fieldSpecs {
			if tag, ok := fieldNameLookup[field.Name]; ok {
				allFields[tag] = struct{}{}
				if field.Required == "Y" {
					requiredFields[tag] = struct{}{}
				}
			} else {
				return fmt.Errorf("Field %q not given in fields section", field.Name)
			}
		}

		return nil
	}

	for _, msg := range fixSpec.Messages {
		d.messageFields[msg.MsgType] = make(map[tag.Tag]struct{})
		d.requiredFields[msg.MsgType] = make(map[tag.Tag]struct{})

		if err := assignFields(msg.Fields, d.messageFields[msg.MsgType], d.requiredFields[msg.MsgType]); err != nil {
			return nil, err
		}
	}

	d.headerFields = make(map[tag.Tag]struct{})
	d.headerRequiredFields = make(map[tag.Tag]struct{})
	if err := assignFields(fixSpec.Header, d.headerFields, d.headerRequiredFields); err != nil {
		return nil, err
	}

	d.trailerFields = make(map[tag.Tag]struct{})
	d.trailerRequiredFields = make(map[tag.Tag]struct{})
	if err := assignFields(fixSpec.Trailer, d.trailerFields, d.trailerRequiredFields); err != nil {
		return nil, err
	}

	return &d, nil
}

func (d *DataDictionary) validate(message Message) (reject MessageReject) {
	msgType := new(field.MsgType)
	if err := message.Header.Get(msgType); err != nil {
		switch err.(type) {
		case FieldNotFoundError:
			return NewRequiredTagMissing(message, tag.MsgType)
		default:
			panic(fmt.Sprintf("Unhandled error: %v", err))
		}
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
	if reject = d.checkRequiredFieldMap(message, d.headerRequiredFields, message.Header.FieldMap); reject != nil {
		return
	}

	if reject = d.checkRequiredFieldMap(message, d.requiredFields[msgType], message.Body); reject != nil {
		return
	}

	if reject = d.checkRequiredFieldMap(message, d.trailerRequiredFields, message.Trailer.FieldMap); reject != nil {
		return
	}

	return
}

func (d *DataDictionary) checkRequiredFieldMap(msg Message, requiredFields map[tag.Tag]struct{}, fieldMap FieldMap) (reject MessageReject) {
	for required := range requiredFields {
		//FIXME ugly...
		field := field.NewStringField(required, "")
		if err := fieldMap.Get(field); err != nil {
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
	if reject = d.iterateFieldMap(message, d.headerFields, message.Header.FieldMap); reject != nil {
		return
	}
	if reject = d.iterateFieldMap(message, d.messageFields[msgType], message.Body); reject != nil {
		return
	}
	if reject = d.iterateFieldMap(message, d.trailerFields, message.Trailer.FieldMap); reject != nil {
		return
	}

	return
}

func (d *DataDictionary) iterateFieldMap(message Message, validFields map[tag.Tag]struct{}, fieldMap FieldMap) MessageReject {
	for tag, fieldValue := range fieldMap.fields {
		if len(fieldValue.Value) == 0 {
			return NewTagSpecifiedWithoutAValue(message, tag)
		}
		if _, valid := d.allFields[tag]; !valid {
			return NewInvalidTagNumber(message, tag)
		}

		if _, valid := validFields[tag]; !valid {
			return NewTagNotDefinedForThisMessageType(message, tag)
		}
	}

	return nil
}
