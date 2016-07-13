package logon

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//Logon is the fix41 Logon type, MsgType = A
type Logon struct {
	fix41.Header
	quickfix.Body
	fix41.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a Logon from a quickfix.Message instance
func FromMessage(m quickfix.Message) Logon {
	return Logon{
		Header:      fix41.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix41.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Logon) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a Logon initialized with the required fields for Logon
func New(encryptmethod field.EncryptMethodField, heartbtint field.HeartBtIntField) (m Logon) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("A"))
	m.Header.Set(field.NewBeginString("FIX.4.1"))
	m.Set(encryptmethod)
	m.Set(heartbtint)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Logon, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "A", r
}

//SetRawDataLength sets RawDataLength, Tag 95
func (m Logon) SetRawDataLength(v int) {
	m.Set(field.NewRawDataLength(v))
}

//SetRawData sets RawData, Tag 96
func (m Logon) SetRawData(v string) {
	m.Set(field.NewRawData(v))
}

//SetEncryptMethod sets EncryptMethod, Tag 98
func (m Logon) SetEncryptMethod(v int) {
	m.Set(field.NewEncryptMethod(v))
}

//SetHeartBtInt sets HeartBtInt, Tag 108
func (m Logon) SetHeartBtInt(v int) {
	m.Set(field.NewHeartBtInt(v))
}

//SetResetSeqNumFlag sets ResetSeqNumFlag, Tag 141
func (m Logon) SetResetSeqNumFlag(v bool) {
	m.Set(field.NewResetSeqNumFlag(v))
}

//GetRawDataLength gets RawDataLength, Tag 95
func (m Logon) GetRawDataLength() (f field.RawDataLengthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRawData gets RawData, Tag 96
func (m Logon) GetRawData() (f field.RawDataField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncryptMethod gets EncryptMethod, Tag 98
func (m Logon) GetEncryptMethod() (f field.EncryptMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetHeartBtInt gets HeartBtInt, Tag 108
func (m Logon) GetHeartBtInt() (f field.HeartBtIntField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetResetSeqNumFlag gets ResetSeqNumFlag, Tag 141
func (m Logon) GetResetSeqNumFlag() (f field.ResetSeqNumFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasRawDataLength returns true if RawDataLength is present, Tag 95
func (m Logon) HasRawDataLength() bool {
	return m.Has(tag.RawDataLength)
}

//HasRawData returns true if RawData is present, Tag 96
func (m Logon) HasRawData() bool {
	return m.Has(tag.RawData)
}

//HasEncryptMethod returns true if EncryptMethod is present, Tag 98
func (m Logon) HasEncryptMethod() bool {
	return m.Has(tag.EncryptMethod)
}

//HasHeartBtInt returns true if HeartBtInt is present, Tag 108
func (m Logon) HasHeartBtInt() bool {
	return m.Has(tag.HeartBtInt)
}

//HasResetSeqNumFlag returns true if ResetSeqNumFlag is present, Tag 141
func (m Logon) HasResetSeqNumFlag() bool {
	return m.Has(tag.ResetSeqNumFlag)
}
