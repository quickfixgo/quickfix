package listexecute

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//ListExecute is the fix43 ListExecute type, MsgType = L
type ListExecute struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a ListExecute from a quickfix.Message instance
func FromMessage(m quickfix.Message) ListExecute {
	return ListExecute{
		Header:      fix43.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix43.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ListExecute) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a ListExecute initialized with the required fields for ListExecute
func New(listid field.ListIDField, transacttime field.TransactTimeField) (m ListExecute) {
	m.Header = fix43.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("L"))
	m.Set(listid)
	m.Set(transacttime)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ListExecute, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "L", r
}

//SetText sets Text, Tag 58
func (m ListExecute) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m ListExecute) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetListID sets ListID, Tag 66
func (m ListExecute) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m ListExecute) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m ListExecute) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetBidID sets BidID, Tag 390
func (m ListExecute) SetBidID(v string) {
	m.Set(field.NewBidID(v))
}

//SetClientBidID sets ClientBidID, Tag 391
func (m ListExecute) SetClientBidID(v string) {
	m.Set(field.NewClientBidID(v))
}

//GetText gets Text, Tag 58
func (m ListExecute) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m ListExecute) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m ListExecute) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m ListExecute) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m ListExecute) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidID gets BidID, Tag 390
func (m ListExecute) GetBidID() (f field.BidIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientBidID gets ClientBidID, Tag 391
func (m ListExecute) GetClientBidID() (f field.ClientBidIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m ListExecute) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m ListExecute) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasListID returns true if ListID is present, Tag 66
func (m ListExecute) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m ListExecute) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m ListExecute) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasBidID returns true if BidID is present, Tag 390
func (m ListExecute) HasBidID() bool {
	return m.Has(tag.BidID)
}

//HasClientBidID returns true if ClientBidID is present, Tag 391
func (m ListExecute) HasClientBidID() bool {
	return m.Has(tag.ClientBidID)
}
