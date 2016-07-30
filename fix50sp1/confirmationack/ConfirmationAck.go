package confirmationack

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//ConfirmationAck is the fix50sp1 ConfirmationAck type, MsgType = AU
type ConfirmationAck struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a ConfirmationAck from a quickfix.Message instance
func FromMessage(m quickfix.Message) ConfirmationAck {
	return ConfirmationAck{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ConfirmationAck) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a ConfirmationAck initialized with the required fields for ConfirmationAck
func New(confirmid field.ConfirmIDField, tradedate field.TradeDateField, transacttime field.TransactTimeField, affirmstatus field.AffirmStatusField) (m ConfirmationAck) {
	m.Header = fixt11.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("AU"))
	m.Header.Set(field.NewBeginString("8"))
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
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "8", "AU", r
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
func (m ConfirmationAck) SetMatchStatus(v string) {
	m.Set(field.NewMatchStatus(v))
}

//SetConfirmID sets ConfirmID, Tag 664
func (m ConfirmationAck) SetConfirmID(v string) {
	m.Set(field.NewConfirmID(v))
}

//SetConfirmRejReason sets ConfirmRejReason, Tag 774
func (m ConfirmationAck) SetConfirmRejReason(v int) {
	m.Set(field.NewConfirmRejReason(v))
}

//SetAffirmStatus sets AffirmStatus, Tag 940
func (m ConfirmationAck) SetAffirmStatus(v int) {
	m.Set(field.NewAffirmStatus(v))
}

//GetText gets Text, Tag 58
func (m ConfirmationAck) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m ConfirmationAck) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m ConfirmationAck) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m ConfirmationAck) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m ConfirmationAck) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMatchStatus gets MatchStatus, Tag 573
func (m ConfirmationAck) GetMatchStatus() (f field.MatchStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetConfirmID gets ConfirmID, Tag 664
func (m ConfirmationAck) GetConfirmID() (f field.ConfirmIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetConfirmRejReason gets ConfirmRejReason, Tag 774
func (m ConfirmationAck) GetConfirmRejReason() (f field.ConfirmRejReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAffirmStatus gets AffirmStatus, Tag 940
func (m ConfirmationAck) GetAffirmStatus() (f field.AffirmStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
