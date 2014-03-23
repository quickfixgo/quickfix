package quickfixgo

import (
	"fmt"
	"github.com/cbusbey/quickfixgo/field"
	"github.com/cbusbey/quickfixgo/spec"
	"github.com/cbusbey/quickfixgo/tag"
)

type DataDictionary struct {
	fields         map[string]map[tag.Tag]struct{}
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

	fieldNameLookup := make(map[string]tag.Tag)
	for _, fieldType := range fixSpec.FieldTypeMap {
		fieldNameLookup[fieldType.Name] = tag.Tag(fieldType.Number)
	}

	d.fields = make(map[string]map[tag.Tag]struct{})
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
		d.fields[msg.MsgType] = make(map[tag.Tag]struct{})
		d.requiredFields[msg.MsgType] = make(map[tag.Tag]struct{})

		if err := assignFields(msg.Fields, d.fields[msg.MsgType], d.requiredFields[msg.MsgType]); err != nil {
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

func (d *DataDictionary) validate(message Message) error {
	msgType := new(field.MsgType)
	if err := message.Header.Get(msgType); err != nil {
		return err
	}

	if err := d.checkRequired(msgType.Value, message); err != nil {
		return err
	}

	if err := d.iterate(msgType.Value, message); err != nil {
		return err
	}

	return nil
}

func (d *DataDictionary) checkRequired(msgType string, message Message) (err error) {
	if err = d.checkRequiredFieldMap(d.headerRequiredFields, message.Header.FieldMap); err != nil {
		return
	}

	if err = d.checkRequiredFieldMap(d.requiredFields[msgType], message.Body); err != nil {
		return
	}

	if err = d.checkRequiredFieldMap(d.trailerRequiredFields, message.Trailer.FieldMap); err != nil {
		return
	}

	return
}

func (d *DataDictionary) checkRequiredFieldMap(requiredFields map[tag.Tag]struct{}, fieldMap FieldMap) (err error) {
	for required := range requiredFields {
		//FIXME ugly...
		field := field.NewStringField(required, "")
		if err = fieldMap.Get(field); err != nil {
			return err
		}
	}

	return
}

func (d *DataDictionary) iterate(msgType string, message Message) (err error) {
	if err = d.iterateFieldMap(d.headerFields, message.Header.FieldMap); err != nil {
		return
	}
	if err = d.iterateFieldMap(d.fields[msgType], message.Body); err != nil {
		return
	}
	if err = d.iterateFieldMap(d.trailerFields, message.Trailer.FieldMap); err != nil {
		return
	}

	return
}

func (d *DataDictionary) iterateFieldMap(validFields map[tag.Tag]struct{}, fieldMap FieldMap) error {
	for tag := range fieldMap.fields {
		if _, valid := validFields[tag]; !valid {
			return InvalidTagNumberError{tag}
		}
	}

	return nil
}
