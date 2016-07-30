package email

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//Email is the fix40 Email type, MsgType = C
type Email struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a Email from a quickfix.Message instance
func FromMessage(m quickfix.Message) Email {
	return Email{
		Header:      fix40.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix40.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Email) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a Email initialized with the required fields for Email
func New(emailtype field.EmailTypeField, linesoftext field.LinesOfTextField, text field.TextField) (m Email) {
	m.Header = fix40.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("C"))
	m.Set(emailtype)
	m.Set(linesoftext)
	m.Set(text)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Email, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "C", r
}

//SetClOrdID sets ClOrdID, Tag 11
func (m Email) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetLinesOfText sets LinesOfText, Tag 33
func (m Email) SetLinesOfText(v int) {
	m.Set(field.NewLinesOfText(v))
}

//SetOrderID sets OrderID, Tag 37
func (m Email) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrigTime sets OrigTime, Tag 42
func (m Email) SetOrigTime(v time.Time) {
	m.Set(field.NewOrigTime(v))
}

//SetRelatdSym sets RelatdSym, Tag 46
func (m Email) SetRelatdSym(v string) {
	m.Set(field.NewRelatdSym(v))
}

//SetText sets Text, Tag 58
func (m Email) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEmailType sets EmailType, Tag 94
func (m Email) SetEmailType(v string) {
	m.Set(field.NewEmailType(v))
}

//SetRawDataLength sets RawDataLength, Tag 95
func (m Email) SetRawDataLength(v int) {
	m.Set(field.NewRawDataLength(v))
}

//SetRawData sets RawData, Tag 96
func (m Email) SetRawData(v string) {
	m.Set(field.NewRawData(v))
}

//GetClOrdID gets ClOrdID, Tag 11
func (m Email) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLinesOfText gets LinesOfText, Tag 33
func (m Email) GetLinesOfText() (f field.LinesOfTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m Email) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigTime gets OrigTime, Tag 42
func (m Email) GetOrigTime() (f field.OrigTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRelatdSym gets RelatdSym, Tag 46
func (m Email) GetRelatdSym() (f field.RelatdSymField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m Email) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEmailType gets EmailType, Tag 94
func (m Email) GetEmailType() (f field.EmailTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRawDataLength gets RawDataLength, Tag 95
func (m Email) GetRawDataLength() (f field.RawDataLengthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRawData gets RawData, Tag 96
func (m Email) GetRawData() (f field.RawDataField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m Email) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasLinesOfText returns true if LinesOfText is present, Tag 33
func (m Email) HasLinesOfText() bool {
	return m.Has(tag.LinesOfText)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m Email) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasOrigTime returns true if OrigTime is present, Tag 42
func (m Email) HasOrigTime() bool {
	return m.Has(tag.OrigTime)
}

//HasRelatdSym returns true if RelatdSym is present, Tag 46
func (m Email) HasRelatdSym() bool {
	return m.Has(tag.RelatdSym)
}

//HasText returns true if Text is present, Tag 58
func (m Email) HasText() bool {
	return m.Has(tag.Text)
}

//HasEmailType returns true if EmailType is present, Tag 94
func (m Email) HasEmailType() bool {
	return m.Has(tag.EmailType)
}

//HasRawDataLength returns true if RawDataLength is present, Tag 95
func (m Email) HasRawDataLength() bool {
	return m.Has(tag.RawDataLength)
}

//HasRawData returns true if RawData is present, Tag 96
func (m Email) HasRawData() bool {
	return m.Has(tag.RawData)
}
