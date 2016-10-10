package confirmationack

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/tag"
)

//ConfirmationAck is the fix44 ConfirmationAck type, MsgType = AU
type ConfirmationAck struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

//FromMessage creates a ConfirmationAck from a quickfix.Message instance
func FromMessage(m *quickfix.Message) ConfirmationAck {
	return ConfirmationAck{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ConfirmationAck) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a ConfirmationAck initialized with the required fields for ConfirmationAck
func New(confirmid field.ConfirmIDField, tradedate field.TradeDateField, transacttime field.TransactTimeField, affirmstatus field.AffirmStatusField) (m ConfirmationAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("AU"))
	m.Set(confirmid)
	m.Set(tradedate)
	m.Set(transacttime)
	m.Set(affirmstatus)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ConfirmationAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "AU", r
}

//SetText sets Text, Tag 58
func (m ConfirmationAck) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m ConfirmationAck) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m ConfirmationAck) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m ConfirmationAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m ConfirmationAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetMatchStatus sets MatchStatus, Tag 573
func (m ConfirmationAck) SetMatchStatus(v enum.MatchStatus) {
	m.Set(field.NewMatchStatus(v))
}

//SetConfirmID sets ConfirmID, Tag 664
func (m ConfirmationAck) SetConfirmID(v string) {
	m.Set(field.NewConfirmID(v))
}

//SetConfirmRejReason sets ConfirmRejReason, Tag 774
func (m ConfirmationAck) SetConfirmRejReason(v enum.ConfirmRejReason) {
	m.Set(field.NewConfirmRejReason(v))
}

//SetAffirmStatus sets AffirmStatus, Tag 940
func (m ConfirmationAck) SetAffirmStatus(v enum.AffirmStatus) {
	m.Set(field.NewAffirmStatus(v))
}

//GetText gets Text, Tag 58
func (m ConfirmationAck) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m ConfirmationAck) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m ConfirmationAck) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m ConfirmationAck) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m ConfirmationAck) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMatchStatus gets MatchStatus, Tag 573
func (m ConfirmationAck) GetMatchStatus() (v enum.MatchStatus, err quickfix.MessageRejectError) {
	var f field.MatchStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetConfirmID gets ConfirmID, Tag 664
func (m ConfirmationAck) GetConfirmID() (v string, err quickfix.MessageRejectError) {
	var f field.ConfirmIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetConfirmRejReason gets ConfirmRejReason, Tag 774
func (m ConfirmationAck) GetConfirmRejReason() (v enum.ConfirmRejReason, err quickfix.MessageRejectError) {
	var f field.ConfirmRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAffirmStatus gets AffirmStatus, Tag 940
func (m ConfirmationAck) GetAffirmStatus() (v enum.AffirmStatus, err quickfix.MessageRejectError) {
	var f field.AffirmStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasText returns true if Text is present, Tag 58
func (m ConfirmationAck) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m ConfirmationAck) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m ConfirmationAck) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m ConfirmationAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m ConfirmationAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasMatchStatus returns true if MatchStatus is present, Tag 573
func (m ConfirmationAck) HasMatchStatus() bool {
	return m.Has(tag.MatchStatus)
}

//HasConfirmID returns true if ConfirmID is present, Tag 664
func (m ConfirmationAck) HasConfirmID() bool {
	return m.Has(tag.ConfirmID)
}

//HasConfirmRejReason returns true if ConfirmRejReason is present, Tag 774
func (m ConfirmationAck) HasConfirmRejReason() bool {
	return m.Has(tag.ConfirmRejReason)
}

//HasAffirmStatus returns true if AffirmStatus is present, Tag 940
func (m ConfirmationAck) HasAffirmStatus() bool {
	return m.Has(tag.AffirmStatus)
}
