package basic

import(
  "sort"
  "bytes"
  "quickfixgo/message"
)

type fieldSort struct {
  tags []message.Tag
  compare fieldOrder
}
// fieldOrder true if tag i should occur before tag j
type fieldOrder func(i,j message.Tag) bool

// Normal FieldOrder (ascending tags)
func normalFieldOrder(i,j message.Tag) bool {return i<j}

func (t fieldSort) Len() int {return len(t.tags)}
func (t fieldSort) Swap(i,j int) {t.tags[i],t.tags[j]=t.tags[j],t.tags[i]}
func (t fieldSort) Less(i,j int) bool {return t.compare(t.tags[i],t.tags[j])}

//Basic implementation of message.FieldMap. Is mutable.
type FieldMap struct {
  fields map[message.Tag] message.Field
  fieldOrder
}

func (fieldMap * FieldMap) init(ordering fieldOrder) {
  fieldMap.fields=make(map[message.Tag] message.Field)
  fieldMap.fieldOrder = ordering
}

func (m *FieldMap) Tags() []message.Tag {
  tags := make([]message.Tag,0,len(m.fields))
  for tag :=range m.fields {
    tags = append(tags, tag)
  }

  return tags
}

func (m *FieldMap) Set(field message.Field) {
  m.fields[field.Tag()]=field
}

func (m *FieldMap) Get(tag message.Tag) (field message.Field, ok bool) {
  field,ok=m.fields[tag]
  return
}

func (m *FieldMap) Remove(tag message.Tag) {
  delete(m.fields, tag)
}


func (m *FieldMap) StringField(tag message.Tag) (message.StringField, error) {
  message_field,ok:=m.Get(tag)
  if !ok {
    return nil, message.FieldNotFoundError{tag}
  }

  switch typeField:=message_field.(type) {
    case message.StringField:
      return typeField, nil
  }

  return NewStringField(tag, message_field.Value()), nil
}

func (m *FieldMap) IntField(tag message.Tag) (message.IntField, error) {
  message_field,ok:=m.Get(tag)
  if !ok {
    return nil, message.FieldNotFoundError{tag} 
  }

  switch typeField:=message_field.(type) {
    case message.IntField:
      return typeField, nil
  }

  if intField,err:=ToIntField(message_field); err==nil {
    return intField, nil
  }

  return nil, message.FieldConvertError{Tag: tag, Value: message_field.Value()}
}

func (m *FieldMap) UTCTimestampField(tag message.Tag) (message.UTCTimestampField, error) {
  message_field,ok:=m.Get(tag)
  if !ok {
    return nil, message.FieldNotFoundError{tag}
  }

  switch typeField:=message_field.(type) {
    case message.UTCTimestampField:
      return typeField, nil
  }

  if utcTimestampField,err:=ToUTCTimestampField(message_field); err==nil {
    return utcTimestampField, nil
  }

  return nil,message.FieldConvertError{Tag: tag, Value: message_field.Value()}
}

func (m FieldMap) sortedTags() []message.Tag {
  tags := m.Tags()
  sort.Sort(fieldSort{tags, m.fieldOrder})
  return tags
}

func(m * FieldMap) Length() int {
  length:=0
  for _,field :=range m.fields {
    switch field.Tag() {
      case 8,9,10: //tags do not contribute to length
      default: length+=field.Length()
    }
  }

  return length
}

func(m *FieldMap) Total() int {
  total:=0
  for _,field :=range m.fields {
    switch field.Tag() {
      case 10: //tag does not contribute to total
      default: total+=field.Total()
    }
  }

  return total
}

func (m *FieldMap) Write(b *bytes.Buffer) {
  tags := m.sortedTags()

  for _,tag := range tags {
    if field,ok:=m.fields[tag]; ok {
      b.WriteString(field.String())
    }
  }
}
