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
	"fmt"
	"math"
	"time"

	"github.com/quickfixgo/quickfix/datadictionary"
)

// Header is first section of a FIX Message.
type Header struct{ FieldMap }

// msgparser contains message parsing vars needed to parse a string into a message.
type msgParser struct {
	msg                     *Message
	transportDataDictionary *datadictionary.DataDictionary
	appDataDictionary       *datadictionary.DataDictionary
	rawBytes                []byte
	fieldIndex              int
	parsedFieldBytes        *TagValue
	trailerBytes            []byte
	foundBody               bool
	foundTrailer            bool
}

// in the message header, the first 3 tags in the message header must be 8,9,35.
func headerFieldOrdering(i, j Tag) bool {
	var ordering = func(t Tag) uint32 {
		switch t {
		case tagBeginString:
			return 1
		case tagBodyLength:
			return 2
		case tagMsgType:
			return 3
		}

		return math.MaxUint32
	}

	orderi := ordering(i)
	orderj := ordering(j)

	switch {
	case orderi < orderj:
		return true
	case orderi > orderj:
		return false
	}

	return i < j
}

// Init initializes the Header instance.
func (h *Header) Init() {
	h.initWithOrdering(headerFieldOrdering)
}

// Body is the primary application section of a FIX message.
type Body struct{ FieldMap }

// Init initializes the FIX message.
func (b *Body) Init() {
	b.init()
}

// Trailer is the last section of a FIX message.
type Trailer struct{ FieldMap }

// In the trailer, CheckSum (tag 10) must be last.
func trailerFieldOrdering(i, j Tag) bool {
	switch {
	case i == tagCheckSum:
		return false
	case j == tagCheckSum:
		return true
	}

	return i > j
}

// Init initializes the FIX message.
func (t *Trailer) Init() {
	t.initWithOrdering(trailerFieldOrdering)
}

// Message is a FIX Message abstraction.
type Message struct {
	Header  Header
	Trailer Trailer
	Body    Body

	// ReceiveTime is the time that this message was read from the socket connection.
	ReceiveTime time.Time

	rawMessage *bytes.Buffer

	// Slice of Bytes corresponding to the message body.
	bodyBytes []byte

	// Field bytes as they appear in the raw message.
	fields []TagValue
}

// ToMessage returns the message itself.
func (m *Message) ToMessage() *Message { return m }

// parseError is returned when bytes cannot be parsed as a FIX message.
type parseError struct {
	OrigError string
}

func (e parseError) Error() string { return fmt.Sprintf("error parsing message: %s", e.OrigError) }

// NewMessage returns a newly initialized Message instance.
func NewMessage() *Message {
	m := new(Message)
	m.Header.Init()
	m.Body.Init()
	m.Trailer.Init()

	return m
}

// CopyInto erases the dest messages and copies the currency message content
// into it.
func (m *Message) CopyInto(to *Message) {
	m.Header.CopyInto(&to.Header.FieldMap)
	m.Body.CopyInto(&to.Body.FieldMap)
	m.Trailer.CopyInto(&to.Trailer.FieldMap)

	to.ReceiveTime = m.ReceiveTime
	to.bodyBytes = make([]byte, len(m.bodyBytes))
	copy(to.bodyBytes, m.bodyBytes)
	to.fields = make([]TagValue, len(m.fields))
	for i := range to.fields {
		to.fields[i].init(m.fields[i].tag, m.fields[i].value)
	}
}

// ParseMessage constructs a Message from a byte slice wrapping a FIX message.
func ParseMessage(msg *Message, rawMessage *bytes.Buffer) (err error) {
	return ParseMessageWithDataDictionary(msg, rawMessage, nil, nil)
}

// ParseMessageWithDataDictionary constructs a Message from a byte slice wrapping a FIX message using an optional session and application DataDictionary for reference.
func ParseMessageWithDataDictionary(
	msg *Message,
	rawMessage *bytes.Buffer,
	transportDataDictionary *datadictionary.DataDictionary,
	appDataDictionary *datadictionary.DataDictionary,
) (err error) {
	// Create msgparser before we go any further.
	mp := &msgParser{
		msg:                     msg,
		transportDataDictionary: transportDataDictionary,
		appDataDictionary:       appDataDictionary,
	}
	mp.msg.rawMessage = rawMessage
	mp.rawBytes = rawMessage.Bytes()

	return doParsing(mp)
}

// doParsing executes the message parsing process.
func doParsing(mp *msgParser) (err error) {
	mp.msg.Header.rwLock.Lock()
	defer mp.msg.Header.rwLock.Unlock()
	mp.msg.Body.rwLock.Lock()
	defer mp.msg.Body.rwLock.Unlock()
	mp.msg.Trailer.rwLock.Lock()
	defer mp.msg.Trailer.rwLock.Unlock()

	// Initialize for parsing.
	mp.msg.Header.clearNoLock()
	mp.msg.Body.clearNoLock()
	mp.msg.Trailer.clearNoLock()

	// Allocate expected message fields in one chunk.
	fieldCount := bytes.Count(mp.rawBytes, []byte{'\001'})
	if fieldCount == 0 {
		return parseError{OrigError: fmt.Sprintf("No Fields detected in %s", string(mp.rawBytes))}
	}
	if cap(mp.msg.fields) < fieldCount {
		mp.msg.fields = make([]TagValue, fieldCount)
	} else {
		mp.msg.fields = mp.msg.fields[0:fieldCount]
	}

	// Message must start with begin string, body length, msg type.
	// Get begin string.
	if mp.rawBytes, err = extractSpecificField(&mp.msg.fields[mp.fieldIndex], tagBeginString, mp.rawBytes); err != nil {
		return
	}
	mp.msg.Header.add(mp.msg.fields[mp.fieldIndex : mp.fieldIndex+1])

	// Get body length.
	mp.fieldIndex++
	mp.parsedFieldBytes = &mp.msg.fields[mp.fieldIndex]
	if mp.rawBytes, err = extractSpecificField(mp.parsedFieldBytes, tagBodyLength, mp.rawBytes); err != nil {
		return
	}
	mp.msg.Header.add(mp.msg.fields[mp.fieldIndex : mp.fieldIndex+1])

	// Get msg type.
	mp.fieldIndex++
	mp.parsedFieldBytes = &mp.msg.fields[mp.fieldIndex]
	if mp.rawBytes, err = extractSpecificField(mp.parsedFieldBytes, tagMsgType, mp.rawBytes); err != nil {
		return
	}
	mp.msg.Header.add(mp.msg.fields[mp.fieldIndex : mp.fieldIndex+1])

	// Start parsing.
	mp.fieldIndex++
	xmlDataLen := 0
	xmlDataMsg := false
	mp.trailerBytes = []byte{}
	mp.foundBody = false
	mp.foundTrailer = false
	for {
		mp.parsedFieldBytes = &mp.msg.fields[mp.fieldIndex]
		if xmlDataLen > 0 {
			mp.rawBytes, err = extractXMLDataField(mp.parsedFieldBytes, mp.rawBytes, xmlDataLen)
			xmlDataLen = 0
			xmlDataMsg = true
		} else {
			mp.rawBytes, err = extractField(mp.parsedFieldBytes, mp.rawBytes)
		}
		if err != nil {
			return
		}

		switch {
		case isHeaderField(mp.parsedFieldBytes.tag, mp.transportDataDictionary):
			mp.msg.Header.add(mp.msg.fields[mp.fieldIndex : mp.fieldIndex+1])
		case isTrailerField(mp.parsedFieldBytes.tag, mp.transportDataDictionary):
			mp.msg.Trailer.add(mp.msg.fields[mp.fieldIndex : mp.fieldIndex+1])
			mp.foundTrailer = true
		case isNumInGroupField(mp.msg, []Tag{mp.parsedFieldBytes.tag}, mp.appDataDictionary):
			parseGroup(mp, []Tag{mp.parsedFieldBytes.tag})
		default:
			mp.foundBody = true
			mp.trailerBytes = mp.rawBytes
			mp.msg.Body.add(mp.msg.fields[mp.fieldIndex : mp.fieldIndex+1])
		}
		if mp.parsedFieldBytes.tag == tagCheckSum {
			break
		}

		if !mp.foundBody {
			mp.msg.bodyBytes = mp.rawBytes
		}

		if mp.parsedFieldBytes.tag == tagXMLDataLen {
			xmlDataLen, _ = mp.msg.Header.getIntNoLock(tagXMLDataLen)
		}
		mp.fieldIndex++
	}

	// This will happen if there are no fields in the body
	if mp.foundTrailer && !mp.foundBody {
		mp.trailerBytes = mp.rawBytes
		mp.msg.bodyBytes = nil
	}

	// Body length would only be larger than trailer if fields out of order.
	if len(mp.msg.bodyBytes) > len(mp.trailerBytes) {
		mp.msg.bodyBytes = mp.msg.bodyBytes[:len(mp.msg.bodyBytes)-len(mp.trailerBytes)]
	}

	length := 0
	for _, field := range mp.msg.fields {
		switch field.tag {
		case tagBeginString, tagBodyLength, tagCheckSum: // Tags do not contribute to length.
		default:
			length += field.length()
		}
	}

	bodyLength, err := mp.msg.Header.getIntNoLock(tagBodyLength)
	if err != nil {
		err = parseError{OrigError: err.Error()}
	} else if length != bodyLength && !xmlDataMsg {
		err = parseError{OrigError: fmt.Sprintf("Incorrect Message Length, expected %d, got %d", bodyLength, length)}
	}

	return
}

// parseGroup iterates through a repeating group to maintain correct order of those fields.
func parseGroup(mp *msgParser, tags []Tag) {
	mp.foundBody = true
	dm := mp.msg.fields[mp.fieldIndex : mp.fieldIndex+1]
	fields := getGroupFields(mp.msg, tags, mp.appDataDictionary)

	for {
		mp.fieldIndex++
		mp.parsedFieldBytes = &mp.msg.fields[mp.fieldIndex]
		mp.rawBytes, _ = extractField(mp.parsedFieldBytes, mp.rawBytes)
		mp.trailerBytes = mp.rawBytes

		// Is this field a member for the group.
		if isGroupMember(mp.parsedFieldBytes.tag, fields) {
			// Is this field a nested repeating group.
			if isNumInGroupField(mp.msg, append(tags, mp.parsedFieldBytes.tag), mp.appDataDictionary) {
				dm = append(dm, *mp.parsedFieldBytes)
				tags = append(tags, mp.parsedFieldBytes.tag)
				fields = getGroupFields(mp.msg, tags, mp.appDataDictionary)
				continue
			}
			// Add the field member to the group.
			dm = append(dm, *mp.parsedFieldBytes)
		} else if isHeaderField(mp.parsedFieldBytes.tag, mp.transportDataDictionary) {
			// Found a header tag for some reason..
			mp.msg.Body.add(dm)
			mp.msg.Header.add(mp.msg.fields[mp.fieldIndex : mp.fieldIndex+1])
			break
		} else if isTrailerField(mp.parsedFieldBytes.tag, mp.transportDataDictionary) {
			// Found the trailer at the end of the message.
			mp.msg.Body.add(dm)
			mp.msg.Trailer.add(mp.msg.fields[mp.fieldIndex : mp.fieldIndex+1])
			mp.foundTrailer = true
			break
		} else {
			// Found a body field outside the group.
			searchTags := []Tag{mp.parsedFieldBytes.tag}
			// Is this a new group not inside the existing group.
			if isNumInGroupField(mp.msg, searchTags, mp.appDataDictionary) {
				// Add the current repeating group.
				mp.msg.Body.add(dm)
				// Cycle again with the new group.
				dm = mp.msg.fields[mp.fieldIndex : mp.fieldIndex+1]
				fields = getGroupFields(mp.msg, searchTags, mp.appDataDictionary)
				continue
			}
			if len(tags) > 1 {
				searchTags = tags[:len(tags)-1]
			}
			// Did this tag occur after a nested group and belongs to the parent group.
			if isNumInGroupField(mp.msg, searchTags, mp.appDataDictionary) {
				// Add the field member to the group.
				dm = append(dm, *mp.parsedFieldBytes)
				// Continue parsing the parent group.
				fields = getGroupFields(mp.msg, searchTags, mp.appDataDictionary)
				continue
			}
			// Add the repeating group.
			mp.msg.Body.add(dm)
			// Add the next body field.
			mp.msg.Body.add(mp.msg.fields[mp.fieldIndex : mp.fieldIndex+1])

			break
		}
	}
}

// isNumInGroupField evaluates if this tag is the start of a repeating group.
// tags slice will contain multiple tags if the tag in question is found while processing a group already.
func isNumInGroupField(msg *Message, tags []Tag, appDataDictionary *datadictionary.DataDictionary) bool {
	if appDataDictionary != nil {
		msgt, err := msg.msgTypeNoLock()
		if err != nil {
			return false
		}
		mm, ok := appDataDictionary.Messages[msgt]
		if ok {
			fields := mm.Fields
			for idx, tag := range tags {
				fd, ok := fields[int(tag)]
				if ok {
					if idx == len(tags)-1 {
						if len(fd.Fields) > 0 {
							return true
						}
					} else {
						// Map nested fields.
						newFields := make(map[int]*datadictionary.FieldDef)
						for _, ff := range fd.Fields {
							newFields[ff.Tag()] = ff
						}
						fields = newFields
					}
				}
			}
		}
	}
	return false
}

// getGroupFields gets the relevant fields for parsing a repeating group if this tag is the start of a repeating group.
// tags slice will contain multiple tags if the tag in question is found while processing a group already.
func getGroupFields(msg *Message, tags []Tag, appDataDictionary *datadictionary.DataDictionary) (fields []*datadictionary.FieldDef) {
	if appDataDictionary != nil {
		msgt, err := msg.msgTypeNoLock()
		if err != nil {
			return
		}
		mm, ok := appDataDictionary.Messages[msgt]
		if ok {
			fields := mm.Fields
			for idx, tag := range tags {
				fd, ok := fields[int(tag)]
				if ok {
					if idx == len(tags)-1 {
						if len(fd.Fields) > 0 {
							return fd.Fields
						}
					} else {
						// Map nested fields.
						newFields := make(map[int]*datadictionary.FieldDef)
						for _, ff := range fd.Fields {
							newFields[ff.Tag()] = ff
						}
						fields = newFields
					}
				}
			}
		}
	}
	return
}

// isGroupMember evaluates if this tag belongs to a repeating group.
func isGroupMember(tag Tag, fields []*datadictionary.FieldDef) bool {
	for _, f := range fields {
		if f.Tag() == int(tag) {
			return true
		}
	}
	return false
}

func isHeaderField(tag Tag, dataDict *datadictionary.DataDictionary) bool {
	if tag.IsHeader() {
		return true
	}

	if dataDict == nil {
		return false
	}

	_, ok := dataDict.Header.Fields[int(tag)]
	return ok
}

func isTrailerField(tag Tag, dataDict *datadictionary.DataDictionary) bool {
	if tag.IsTrailer() {
		return true
	}

	if dataDict == nil {
		return false
	}

	_, ok := dataDict.Trailer.Fields[int(tag)]
	return ok
}

// MsgType returns MsgType (tag 35) field's value.
func (m *Message) MsgType() (string, MessageRejectError) {
	return m.Header.GetString(tagMsgType)
}

func (m *Message) msgTypeNoLock() (string, MessageRejectError) {
	return m.Header.getStringNoLock(tagMsgType)
}

// IsMsgTypeOf returns true if the Header contains MsgType (tag 35) field and its value is the specified one.
func (m *Message) IsMsgTypeOf(msgType string) bool {
	if v, err := m.MsgType(); err == nil {
		return v == msgType
	}
	return false
}

// reverseRoute returns a message builder with routing header fields initialized as the reverse of this message.
func (m *Message) reverseRoute() *Message {
	reverseMsg := NewMessage()

	copyFunc := func(src Tag, dest Tag) {
		var field FIXString
		if m.Header.GetField(src, &field) == nil {
			if len(field) != 0 {
				reverseMsg.Header.SetField(dest, field)
			}
		}
	}

	copyFunc(tagSenderCompID, tagTargetCompID)
	copyFunc(tagSenderSubID, tagTargetSubID)
	copyFunc(tagSenderLocationID, tagTargetLocationID)

	copyFunc(tagTargetCompID, tagSenderCompID)
	copyFunc(tagTargetSubID, tagSenderSubID)
	copyFunc(tagTargetLocationID, tagSenderLocationID)

	copyFunc(tagOnBehalfOfCompID, tagDeliverToCompID)
	copyFunc(tagOnBehalfOfSubID, tagDeliverToSubID)
	copyFunc(tagDeliverToCompID, tagOnBehalfOfCompID)
	copyFunc(tagDeliverToSubID, tagOnBehalfOfSubID)

	// Tags added in 4.1.
	var beginString FIXString
	if m.Header.GetField(tagBeginString, &beginString) == nil {
		if string(beginString) != BeginStringFIX40 {
			copyFunc(tagOnBehalfOfLocationID, tagDeliverToLocationID)
			copyFunc(tagDeliverToLocationID, tagOnBehalfOfLocationID)
		}
	}

	return reverseMsg
}

func extractSpecificField(field *TagValue, expectedTag Tag, buffer []byte) (remBuffer []byte, err error) {
	remBuffer, err = extractField(field, buffer)
	switch {
	case err != nil:
		return
	case field.tag != expectedTag:
		err = parseError{OrigError: fmt.Sprintf("extractSpecificField: Fields out of order, expected %d, got %d", expectedTag, field.tag)}
		return
	}

	return
}

func extractXMLDataField(parsedFieldBytes *TagValue, buffer []byte, dataLen int) (remBytes []byte, err error) {
	endIndex := bytes.IndexByte(buffer, '=')
	if endIndex == -1 {
		err = parseError{OrigError: "extractField: No Trailing Delim in " + string(buffer)}
		remBytes = buffer
		return
	}
	endIndex += dataLen + 1

	err = parsedFieldBytes.parse(buffer[:endIndex+1])
	return buffer[(endIndex + 1):], err
}

func extractField(parsedFieldBytes *TagValue, buffer []byte) (remBytes []byte, err error) {
	endIndex := bytes.IndexByte(buffer, '\001')
	if endIndex == -1 {
		err = parseError{OrigError: "extractField: No Trailing Delim in " + string(buffer)}
		remBytes = buffer
		return
	}

	err = parsedFieldBytes.parse(buffer[:endIndex+1])
	return buffer[(endIndex + 1):], err
}

func (m *Message) Bytes() []byte {
	if m.rawMessage != nil {
		return m.rawMessage.Bytes()
	}

	return m.build()
}

func (m *Message) String() string {
	if m.rawMessage != nil {
		return m.rawMessage.String()
	}

	return string(m.build())
}

func formatCheckSum(value int) string {
	return fmt.Sprintf("%03d", value)
}

// Build constructs a []byte from a Message instance.
func (m *Message) build() []byte {
	m.cook()

	var b bytes.Buffer
	m.Header.write(&b)
	m.Body.write(&b)
	m.Trailer.write(&b)
	return b.Bytes()
}

// Constructs a []byte from a Message instance, using the given bodyBytes.
// This is a workaround for the fact that we currently rely on the generated Message types to properly serialize/deserialize RepeatingGroups.
// In other words, we cannot go from bytes to a Message then back to bytes, which is exactly what we need to do in the case of a Resend.
// This func lets us pull the Message from the Store, parse it, update the Header, and then build it back into bytes using the original Body.
// Note: The only standard non-Body group is NoHops.  If that is used in the Header, this workaround may fail.
func (m *Message) buildWithBodyBytes(bodyBytes []byte) []byte {
	m.cook()

	var b bytes.Buffer
	m.Header.write(&b)
	b.Write(bodyBytes)
	m.Trailer.write(&b)
	return b.Bytes()
}

func (m *Message) cook() {
	bodyLength := m.Header.length() + m.Body.length() + m.Trailer.length()
	m.Header.SetInt(tagBodyLength, bodyLength)
	checkSum := (m.Header.total() + m.Body.total() + m.Trailer.total()) % 256
	m.Trailer.SetString(tagCheckSum, formatCheckSum(checkSum))
}
