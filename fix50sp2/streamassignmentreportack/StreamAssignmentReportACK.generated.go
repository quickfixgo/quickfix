package streamassignmentreportack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//StreamAssignmentReportACK is the fix50sp2 StreamAssignmentReportACK type, MsgType = CE
type StreamAssignmentReportACK struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

//FromMessage creates a StreamAssignmentReportACK from a quickfix.Message instance
func FromMessage(m *quickfix.Message) StreamAssignmentReportACK {
	return StreamAssignmentReportACK{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m StreamAssignmentReportACK) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a StreamAssignmentReportACK initialized with the required fields for StreamAssignmentReportACK
func New(streamasgnacktype field.StreamAsgnAckTypeField, streamasgnrptid field.StreamAsgnRptIDField) (m StreamAssignmentReportACK) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("CE"))
	m.Set(streamasgnacktype)
	m.Set(streamasgnrptid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg StreamAssignmentReportACK, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "CE", r
}

//SetText sets Text, Tag 58
func (m StreamAssignmentReportACK) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m StreamAssignmentReportACK) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m StreamAssignmentReportACK) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetStreamAsgnRptID sets StreamAsgnRptID, Tag 1501
func (m StreamAssignmentReportACK) SetStreamAsgnRptID(v string) {
	m.Set(field.NewStreamAsgnRptID(v))
}

//SetStreamAsgnRejReason sets StreamAsgnRejReason, Tag 1502
func (m StreamAssignmentReportACK) SetStreamAsgnRejReason(v enum.StreamAsgnRejReason) {
	m.Set(field.NewStreamAsgnRejReason(v))
}

//SetStreamAsgnAckType sets StreamAsgnAckType, Tag 1503
func (m StreamAssignmentReportACK) SetStreamAsgnAckType(v enum.StreamAsgnAckType) {
	m.Set(field.NewStreamAsgnAckType(v))
}

//GetText gets Text, Tag 58
func (m StreamAssignmentReportACK) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m StreamAssignmentReportACK) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m StreamAssignmentReportACK) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStreamAsgnRptID gets StreamAsgnRptID, Tag 1501
func (m StreamAssignmentReportACK) GetStreamAsgnRptID() (v string, err quickfix.MessageRejectError) {
	var f field.StreamAsgnRptIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStreamAsgnRejReason gets StreamAsgnRejReason, Tag 1502
func (m StreamAssignmentReportACK) GetStreamAsgnRejReason() (v enum.StreamAsgnRejReason, err quickfix.MessageRejectError) {
	var f field.StreamAsgnRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStreamAsgnAckType gets StreamAsgnAckType, Tag 1503
func (m StreamAssignmentReportACK) GetStreamAsgnAckType() (v enum.StreamAsgnAckType, err quickfix.MessageRejectError) {
	var f field.StreamAsgnAckTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasText returns true if Text is present, Tag 58
func (m StreamAssignmentReportACK) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m StreamAssignmentReportACK) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m StreamAssignmentReportACK) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasStreamAsgnRptID returns true if StreamAsgnRptID is present, Tag 1501
func (m StreamAssignmentReportACK) HasStreamAsgnRptID() bool {
	return m.Has(tag.StreamAsgnRptID)
}

//HasStreamAsgnRejReason returns true if StreamAsgnRejReason is present, Tag 1502
func (m StreamAssignmentReportACK) HasStreamAsgnRejReason() bool {
	return m.Has(tag.StreamAsgnRejReason)
}

//HasStreamAsgnAckType returns true if StreamAsgnAckType is present, Tag 1503
func (m StreamAssignmentReportACK) HasStreamAsgnAckType() bool {
	return m.Has(tag.StreamAsgnAckType)
}
