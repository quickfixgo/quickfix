package marketdatarequestreject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/gen/enum"
	"github.com/quickfixgo/quickfix/gen/field"
	"github.com/quickfixgo/quickfix/gen/fixt11"
	"github.com/quickfixgo/quickfix/gen/tag"
)

// MarketDataRequestReject is the fix50 MarketDataRequestReject type, MsgType = Y.
type MarketDataRequestReject struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a MarketDataRequestReject from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) MarketDataRequestReject {
	return MarketDataRequestReject{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m MarketDataRequestReject) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a MarketDataRequestReject initialized with the required fields for MarketDataRequestReject.
func New(mdreqid field.MDReqIDField) (m MarketDataRequestReject) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("Y"))
	m.Set(mdreqid)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg MarketDataRequestReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "Y", r
}

// SetText sets Text, Tag 58.
func (m MarketDataRequestReject) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetMDReqID sets MDReqID, Tag 262.
func (m MarketDataRequestReject) SetMDReqID(v string) {
	m.Set(field.NewMDReqID(v))
}

// SetMDReqRejReason sets MDReqRejReason, Tag 281.
func (m MarketDataRequestReject) SetMDReqRejReason(v enum.MDReqRejReason) {
	m.Set(field.NewMDReqRejReason(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m MarketDataRequestReject) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m MarketDataRequestReject) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetNoAltMDSource sets NoAltMDSource, Tag 816.
func (m MarketDataRequestReject) SetNoAltMDSource(f NoAltMDSourceRepeatingGroup) {
	m.SetGroup(f)
}

// GetText gets Text, Tag 58.
func (m MarketDataRequestReject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDReqID gets MDReqID, Tag 262.
func (m MarketDataRequestReject) GetMDReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MDReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMDReqRejReason gets MDReqRejReason, Tag 281.
func (m MarketDataRequestReject) GetMDReqRejReason() (v enum.MDReqRejReason, err quickfix.MessageRejectError) {
	var f field.MDReqRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m MarketDataRequestReject) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m MarketDataRequestReject) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoAltMDSource gets NoAltMDSource, Tag 816.
func (m MarketDataRequestReject) GetNoAltMDSource() (NoAltMDSourceRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAltMDSourceRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasText returns true if Text is present, Tag 58.
func (m MarketDataRequestReject) HasText() bool {
	return m.Has(tag.Text)
}

// HasMDReqID returns true if MDReqID is present, Tag 262.
func (m MarketDataRequestReject) HasMDReqID() bool {
	return m.Has(tag.MDReqID)
}

// HasMDReqRejReason returns true if MDReqRejReason is present, Tag 281.
func (m MarketDataRequestReject) HasMDReqRejReason() bool {
	return m.Has(tag.MDReqRejReason)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m MarketDataRequestReject) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m MarketDataRequestReject) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasNoAltMDSource returns true if NoAltMDSource is present, Tag 816.
func (m MarketDataRequestReject) HasNoAltMDSource() bool {
	return m.Has(tag.NoAltMDSource)
}

// NoAltMDSource is a repeating group element, Tag 816.
type NoAltMDSource struct {
	*quickfix.Group
}

// SetAltMDSourceID sets AltMDSourceID, Tag 817.
func (m NoAltMDSource) SetAltMDSourceID(v string) {
	m.Set(field.NewAltMDSourceID(v))
}

// GetAltMDSourceID gets AltMDSourceID, Tag 817.
func (m NoAltMDSource) GetAltMDSourceID() (v string, err quickfix.MessageRejectError) {
	var f field.AltMDSourceIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAltMDSourceID returns true if AltMDSourceID is present, Tag 817.
func (m NoAltMDSource) HasAltMDSourceID() bool {
	return m.Has(tag.AltMDSourceID)
}

// NoAltMDSourceRepeatingGroup is a repeating group, Tag 816.
type NoAltMDSourceRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoAltMDSourceRepeatingGroup returns an initialized, NoAltMDSourceRepeatingGroup.
func NewNoAltMDSourceRepeatingGroup() NoAltMDSourceRepeatingGroup {
	return NoAltMDSourceRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAltMDSource,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AltMDSourceID)})}
}

// Add create and append a new NoAltMDSource to this group.
func (m NoAltMDSourceRepeatingGroup) Add() NoAltMDSource {
	g := m.RepeatingGroup.Add()
	return NoAltMDSource{g}
}

// Get returns the ith NoAltMDSource in the NoAltMDSourceRepeatinGroup.
func (m NoAltMDSourceRepeatingGroup) Get(i int) NoAltMDSource {
	return NoAltMDSource{m.RepeatingGroup.Get(i)}
}
