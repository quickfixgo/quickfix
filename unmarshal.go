package quickfix

import (
	"reflect"
	"time"
)

type decoder struct {
	FieldMap
}

func (d decoder) decodeValue(fixTag Tag, t reflect.Type, v reflect.Value) MessageRejectError {
	switch v.Kind() {
	case reflect.Ptr:
		v.Set(reflect.New(t.Elem()))
		return d.decodeValue(fixTag, t.Elem(), v.Elem())
	case reflect.Struct:
		switch v.Interface().(type) {
		case time.Time:
			fieldValue := new(FIXUTCTimestamp)
			if err := d.FieldMap.GetField(Tag(fixTag), fieldValue); err != nil {
				return err
			}
			v.Set(reflect.ValueOf(fieldValue.Time))
		}
	case reflect.String:
		var fieldValue FIXString
		if err := d.FieldMap.GetField(Tag(fixTag), &fieldValue); err != nil {
			return err
		}
		v.SetString(string(fieldValue))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		var fieldValue FIXInt
		if err := d.FieldMap.GetField(Tag(fixTag), &fieldValue); err != nil {
			return err
		}
		v.SetInt(int64(fieldValue))
	case reflect.Bool:
		var fieldValue FIXBoolean
		if err := d.FieldMap.GetField(Tag(fixTag), &fieldValue); err != nil {
			return err
		}
		v.SetBool(bool(fieldValue))
	case reflect.Float32, reflect.Float64:
		var fieldValue FIXFloat
		if err := d.FieldMap.GetField(Tag(fixTag), &fieldValue); err != nil {
			return err
		}
		v.SetFloat(float64(fieldValue))
	case reflect.Slice:
		elem := v.Type().Elem()
		if elem.Kind() != reflect.Struct {
			panic("repeating group must be a slice of type struct")
		}

		repeatingGroup := NewRepeatingGroup(fixTag, buildGroupTemplate(elem))
		if err := d.FieldMap.GetGroup(repeatingGroup); err != nil {
			return err
		}

		v.Set(reflect.MakeSlice(v.Type(), repeatingGroup.Len(), repeatingGroup.Len()))

		for i := 0; i < repeatingGroup.Len(); i++ {
			groupDecoder := decoder{FieldMap: repeatingGroup.Get(i).FieldMap}

			for j := 0; j < v.Type().Elem().NumField(); j++ {
				sf := v.Type().Elem().Field(j)
				sv := v.Index(i).Field(j)

				if err := groupDecoder.decodeField(sf, sf.Type, sv); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (d decoder) decodeField(sf reflect.StructField, t reflect.Type, v reflect.Value) MessageRejectError {
	if sf.Tag.Get("fix") != "" {
		fixTag, _, _ := parseStructTag(sf.Tag.Get("fix"))
		if !d.FieldMap.Has(fixTag) {
			return nil
		}

		return d.decodeValue(Tag(fixTag), t, v)
	}

	switch t.Kind() {
	case reflect.Ptr:
		v.Set(reflect.New(t.Elem()))
		return d.decodeField(sf, t.Elem(), v.Elem())
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			if err := d.decodeField(t.Field(i), t.Field(i).Type, v.Field(i)); err != nil {
				return err
			}
		}
	}

	return nil
}

// Unmarshal populates v from a Message
func Unmarshal(m Message, v interface{}) MessageRejectError {
	reflectValue := reflect.ValueOf(v)

	switch reflectValue.Kind() {
	case reflect.Ptr:
		reflectValue = reflectValue.Elem()
	default:
		panic("v not a pointer")
	}

	reflectType := reflectValue.Type()
	d := decoder{}

	for i := 0; i < reflectType.NumField(); i++ {
		sf := reflectType.Field(i)
		sv := reflectValue.Field(i)

		switch sf.Name {
		case "FIXMsgType":
			continue
		case "Header":
			d.FieldMap = m.Header.FieldMap
		case "Trailer":
			d.FieldMap = m.Trailer.FieldMap
		default:
			d.FieldMap = m.Body.FieldMap
		}

		if err := d.decodeField(sf, sf.Type, sv); err != nil {
			return err
		}
	}

	return nil
}
