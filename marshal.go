package quickfix

import (
	"reflect"
	"strconv"
	"time"
)

type encoder struct {
	FieldMap
}

func (e encoder) encodeField(f reflect.StructField, t reflect.Type, v reflect.Value) {
	if f.Tag.Get("fix") != "" {
		tag, omitEmpty, defaultVal := parseStructTag(f.Tag.Get("fix"))
		e.encodeValue(tag, v, omitEmpty, defaultVal)
	}

	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return
		}
		e.encodeField(f, v.Elem().Type(), v.Elem())
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			e.encodeField(t.Field(i), t.Field(i).Type, v.Field(i))
		}
	}
}

func (e encoder) encodeValue(fixTag Tag, v reflect.Value, omitEmpty bool, defaultVal *string) {
	if defaultVal != nil {
		e.FieldMap.SetField(fixTag, FIXString(*defaultVal))
		return
	}

	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return
		}
		e.encodeValue(fixTag, v.Elem(), omitEmpty, defaultVal)
	case reflect.Struct:
		switch t := v.Interface().(type) {
		case time.Time:
			e.FieldMap.SetField(fixTag, FIXUTCTimestamp{Value: t})
		}
	case reflect.String:
		e.FieldMap.SetField(fixTag, FIXString(v.String()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		e.FieldMap.SetField(fixTag, FIXInt(v.Int()))
	case reflect.Bool:
		e.FieldMap.SetField(fixTag, FIXBoolean(v.Bool()))
	case reflect.Float32, reflect.Float64:
		e.FieldMap.SetField(fixTag, FIXFloat(v.Float()))
	case reflect.Slice:
		if v.Len() == 0 && omitEmpty {
			return
		}
		elem := v.Type().Elem()
		if elem.Kind() != reflect.Struct {
			panic("repeating group must be a slice of type struct")
		}

		template := make(GroupTemplate, elem.NumField())
		for i := 0; i < elem.NumField(); i++ {
			sf := elem.Field(i)
			fixTag, err := strconv.Atoi(sf.Tag.Get("fix"))

			if err != nil {
				panic(err)
			}

			template[i] = GroupElement(Tag(fixTag))
		}

		repeatingGroup := RepeatingGroup{Tag: fixTag, GroupTemplate: template}

		for i := 0; i < v.Len(); i++ {
			group := repeatingGroup.Add()
			groupEncoder := encoder{FieldMap: group.FieldMap}

			sliceV := v.Index(i)
			for i := 0; i < sliceV.NumField(); i++ {
				sf := sliceV.Type().Field(i)
				groupEncoder.encodeField(sf, sf.Type, sliceV.Field(i))
			}
		}

		e.FieldMap.SetGroup(repeatingGroup)
	}
}

// Marshal returns the Message encoding of v
func Marshal(v interface{}) Message {
	m := NewMessage()

	reflectValue := reflect.ValueOf(v)
	if !reflectValue.IsValid() {
		panic("Message value is invalid")
	}

	var msgType string

	reflectType := reflectValue.Type()
	e := encoder{}
	for i := 0; i < reflectType.NumField(); i++ {
		sf := reflectType.Field(i)

		switch sf.Name {
		case "FIXMsgType":
			msgType = sf.Tag.Get("fix")
			if msgType == "" {
				panic("FIXMsgType has no Tag Value")
			}
			continue
		case "Header":
			e.FieldMap = m.Header
		case "Trailer":
			e.FieldMap = m.Trailer
		default:
			e.FieldMap = m.Body
		}

		e.encodeField(sf, sf.Type, reflectValue.Field(i))
	}

	if msgType != "" {
		m.Header.SetField(tagMsgType, FIXString(msgType))
	}

	return m
}
