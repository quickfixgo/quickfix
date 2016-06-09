package streamassignmentreportack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//StreamAssignmentReportACK is the fix50sp2 StreamAssignmentReportACK type, MsgType = CE
type StreamAssignmentReportACK struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
}

//FromMessage creates a StreamAssignmentReportACK from a quickfix.Message instance
func FromMessage(m quickfix.Message) StreamAssignmentReportACK {
	return StreamAssignmentReportACK{
		Header:  fixt11.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fixt11.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m StreamAssignmentReportACK) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a StreamAssignmentReportACK initialized with the required fields for StreamAssignmentReportACK
func New(streamasgnacktype field.StreamAsgnAckTypeField, streamasgnrptid field.StreamAsgnRptIDField) (m StreamAssignmentReportACK) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("CE"))
	m.Set(streamasgnacktype)
	m.Set(streamasgnrptid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg StreamAssignmentReportACK, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
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
func (m StreamAssignmentReportACK) SetStreamAsgnRejReason(v int) {
	m.Set(field.NewStreamAsgnRejReason(v))
}

//SetStreamAsgnAckType sets StreamAsgnAckType, Tag 1503
func (m StreamAssignmentReportACK) SetStreamAsgnAckType(v int) {
	m.Set(field.NewStreamAsgnAckType(v))
}

//GetText gets Text, Tag 58
func (m StreamAssignmentReportACK) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m StreamAssignmentReportACK) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m StreamAssignmentReportACK) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStreamAsgnRptID gets StreamAsgnRptID, Tag 1501
func (m StreamAssignmentReportACK) GetStreamAsgnRptID() (f field.StreamAsgnRptIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStreamAsgnRejReason gets StreamAsgnRejReason, Tag 1502
func (m StreamAssignmentReportACK) GetStreamAsgnRejReason() (f field.StreamAsgnRejReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStreamAsgnAckType gets StreamAsgnAckType, Tag 1503
func (m StreamAssignmentReportACK) GetStreamAsgnAckType() (f field.StreamAsgnAckTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
