// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/quickfixgo/quickfix/datadictionary"
)

// ToJSON serializes the Message to the FIX JSON Encoding format.
// See: https://github.com/FIXTradingCommunity/fix-json-encoding-spec
//
// If dd is nil, numeric tag numbers are used as field names and repeating groups
// are serialized as a flat string rather than as JSON arrays.
func (m *Message) ToJSON(dd *datadictionary.DataDictionary) ([]byte, error) {
	msgType, _ := m.MsgType()

	var headerFields, bodyFields, trailerFields map[int]*datadictionary.FieldDef
	if dd != nil {
		if dd.Header != nil {
			headerFields = dd.Header.Fields
		}
		if dd.Trailer != nil {
			trailerFields = dd.Trailer.Fields
		}
		if msgType != "" {
			if msgDef, ok := dd.Messages[msgType]; ok {
				bodyFields = msgDef.Fields
			}
		}
	}

	var buf bytes.Buffer
	buf.WriteString(`{"Header":`)
	writeFieldMapJSON(&buf, &m.Header.FieldMap, headerFields, dd)
	buf.WriteString(`,"Body":`)
	writeFieldMapJSON(&buf, &m.Body.FieldMap, bodyFields, dd)
	buf.WriteString(`,"Trailer":`)
	writeFieldMapJSON(&buf, &m.Trailer.FieldMap, trailerFields, dd)
	buf.WriteByte('}')

	return buf.Bytes(), nil
}

// FromJSON populates the Message from a FIX JSON Encoding byte slice.
// See: https://github.com/FIXTradingCommunity/fix-json-encoding-spec
//
// If dd is nil, field names must be numeric tag numbers and repeating groups
// are not supported.
func (m *Message) FromJSON(data []byte, dd *datadictionary.DataDictionary) error {
	var raw struct {
		Header  json.RawMessage `json:"Header"`
		Body    json.RawMessage `json:"Body"`
		Trailer json.RawMessage `json:"Trailer"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("error parsing FIX JSON message: %w", err)
	}

	var headerFields map[int]*datadictionary.FieldDef
	if dd != nil && dd.Header != nil {
		headerFields = dd.Header.Fields
	}

	if err := populateFieldMapFromJSON(&m.Header.FieldMap, raw.Header, headerFields, dd); err != nil {
		return fmt.Errorf("error parsing FIX JSON Header: %w", err)
	}

	// Parse MsgType from the header to determine the body field context.
	msgType, _ := m.MsgType()

	var trailerFields, bodyFields map[int]*datadictionary.FieldDef
	if dd != nil {
		if dd.Trailer != nil {
			trailerFields = dd.Trailer.Fields
		}
		if msgType != "" {
			if msgDef, ok := dd.Messages[msgType]; ok {
				bodyFields = msgDef.Fields
			}
		}
	}

	if err := populateFieldMapFromJSON(&m.Body.FieldMap, raw.Body, bodyFields, dd); err != nil {
		return fmt.Errorf("error parsing FIX JSON Body: %w", err)
	}
	if err := populateFieldMapFromJSON(&m.Trailer.FieldMap, raw.Trailer, trailerFields, dd); err != nil {
		return fmt.Errorf("error parsing FIX JSON Trailer: %w", err)
	}

	return nil
}

// writeFieldMapJSON writes a FieldMap's fields as a JSON object into buf.
// contextFields (if non-nil) identifies which tags represent repeating groups in this message context.
func writeFieldMapJSON(buf *bytes.Buffer, fm *FieldMap, contextFields map[int]*datadictionary.FieldDef, dd *datadictionary.DataDictionary) {
	fm.rwLock.Lock()
	defer fm.rwLock.Unlock()

	buf.WriteByte('{')
	first := true

	for _, tag := range fm.sortedTags() {
		if tag == tagCheckSum {
			continue
		}

		f, ok := fm.tagLookup[tag]
		if !ok {
			continue
		}

		name := fieldNameOrTag(tag, dd)

		var groupDef *datadictionary.FieldDef
		if contextFields != nil {
			if fd, ok2 := contextFields[int(tag)]; ok2 && fd.IsGroup() {
				groupDef = fd
			}
		}

		if !first {
			buf.WriteByte(',')
		}
		first = false

		// Write JSON key.
		nameBytes, _ := json.Marshal(name)
		buf.Write(nameBytes)
		buf.WriteByte(':')

		if groupDef != nil {
			writeGroupJSON(buf, f, groupDef, dd)
		} else {
			valBytes, _ := json.Marshal(string(f[0].value))
			buf.Write(valBytes)
		}
	}

	buf.WriteByte('}')
}

// writeGroupJSON serializes a repeating-group field as a JSON array of objects.
// f is the raw field slice stored in the FieldMap (count tag + all group member tags).
func writeGroupJSON(buf *bytes.Buffer, f field, groupDef *datadictionary.FieldDef, dd *datadictionary.DataDictionary) {
	template := buildGroupTemplate(groupDef)
	rg := NewRepeatingGroup(Tag(groupDef.Tag()), template)
	_, _ = rg.Read(f)

	subFields := groupSubFields(groupDef)

	buf.WriteByte('[')
	for i, grp := range rg.groups {
		if i > 0 {
			buf.WriteByte(',')
		}
		writeFieldMapJSON(buf, &grp.FieldMap, subFields, dd)
	}
	buf.WriteByte(']')
}

// populateFieldMapFromJSON reads fields from a JSON object and sets them on fm.
// contextFields (if non-nil) identifies which tags are repeating groups in this context.
func populateFieldMapFromJSON(fm *FieldMap, data json.RawMessage, contextFields map[int]*datadictionary.FieldDef, dd *datadictionary.DataDictionary) error {
	if len(data) == 0 {
		return nil
	}

	var rawFields map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawFields); err != nil {
		return err
	}

	for name, rawValue := range rawFields {
		tag, err := resolveTag(name, dd)
		if err != nil {
			continue // skip unknown fields
		}

		// JSON arrays represent repeating groups.
		if len(rawValue) > 0 && rawValue[0] == '[' {
			var groupDef *datadictionary.FieldDef
			if contextFields != nil {
				if fd, ok := contextFields[int(tag)]; ok && fd.IsGroup() {
					groupDef = fd
				}
			}
			if groupDef == nil {
				continue // skip groups we cannot parse without a definition
			}

			var items []json.RawMessage
			if err := json.Unmarshal(rawValue, &items); err != nil {
				return fmt.Errorf("error parsing group %q: %w", name, err)
			}

			subFields := groupSubFields(groupDef)
			rg := NewRepeatingGroup(tag, buildGroupTemplate(groupDef))
			for _, item := range items {
				grp := rg.Add()
				if err := populateFieldMapFromJSON(&grp.FieldMap, item, subFields, dd); err != nil {
					return err
				}
			}
			fm.SetGroup(rg)
		} else {
			var value string
			if err := json.Unmarshal(rawValue, &value); err != nil {
				return fmt.Errorf("error parsing field %q value: %w", name, err)
			}
			fm.SetBytes(tag, []byte(value))
		}
	}

	return nil
}

// buildGroupTemplate constructs a GroupTemplate from a group FieldDef, handling nested groups recursively.
func buildGroupTemplate(groupDef *datadictionary.FieldDef) GroupTemplate {
	template := make(GroupTemplate, 0, len(groupDef.Fields))
	for _, childDef := range groupDef.Fields {
		if childDef.IsGroup() {
			template = append(template, NewRepeatingGroup(Tag(childDef.Tag()), buildGroupTemplate(childDef)))
		} else {
			template = append(template, GroupElement(Tag(childDef.Tag())))
		}
	}
	return template
}

// groupSubFields builds a tag→FieldDef map for the direct children of a group FieldDef.
func groupSubFields(groupDef *datadictionary.FieldDef) map[int]*datadictionary.FieldDef {
	sub := make(map[int]*datadictionary.FieldDef, len(groupDef.Fields))
	for _, fd := range groupDef.Fields {
		sub[fd.Tag()] = fd
	}
	return sub
}

// fieldNameOrTag returns the human-readable field name from the data dictionary,
// or the numeric tag string when the dictionary is absent or the tag is unknown.
func fieldNameOrTag(tag Tag, dd *datadictionary.DataDictionary) string {
	if dd != nil {
		if ft, ok := dd.FieldTypeByTag[int(tag)]; ok {
			return ft.Name()
		}
	}
	return strconv.Itoa(int(tag))
}

// resolveTag maps a JSON field name to a FIX Tag.
// It first checks the data dictionary by name, then falls back to parsing a numeric string.
func resolveTag(name string, dd *datadictionary.DataDictionary) (Tag, error) {
	if dd != nil {
		if ft, ok := dd.FieldTypeByName[name]; ok {
			return Tag(ft.Tag()), nil
		}
	}
	n, err := strconv.Atoi(name)
	if err != nil {
		return 0, fmt.Errorf("unknown field %q", name)
	}
	return Tag(n), nil
}
