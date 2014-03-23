package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/spec"
)

type DataDictionary struct {
	fields map[int]spec.FieldType
}

func NewDataDictionary(path string) (*DataDictionary, error) {
	fixSpec, err := spec.ParseFixSpec(path)
	if err != nil {
		return nil, err
	}

	d := DataDictionary{}
	d.fields = make(map[int]spec.FieldType)
	for _, fieldType := range fixSpec.FieldTypeMap {
		d.fields[fieldType.Number] = fieldType
	}

	return &d, nil
}

func (d *DataDictionary) validate(message Message) error {
	if err := d.iterate(message, message.Header.FieldMap); err != nil {
		return err
	}
	if err := d.iterate(message, message.Body); err != nil {
		return err
	}
	if err := d.iterate(message, message.Trailer.FieldMap); err != nil {
		return err
	}

	return nil
}

func (d *DataDictionary) iterate(msg Message, fieldMap FieldMap) error {
	for tag := range fieldMap.fields {
		if _, valid := d.fields[int(tag)]; !valid {
			return InvalidTagNumberError{tag}
		}
	}

	return nil
}
