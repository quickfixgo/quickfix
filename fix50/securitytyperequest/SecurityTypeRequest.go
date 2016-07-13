package securitytyperequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//SecurityTypeRequest is the fix50 SecurityTypeRequest type, MsgType = v
type SecurityTypeRequest struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a SecurityTypeRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) SecurityTypeRequest {
	return SecurityTypeRequest{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m SecurityTypeRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a SecurityTypeRequest initialized with the required fields for SecurityTypeRequest
func New(securityreqid field.SecurityReqIDField) (m SecurityTypeRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("v"))
	m.Header.Set(field.NewBeginString("7"))
	m.Set(securityreqid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg SecurityTypeRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "v", r
}

//SetText sets Text, Tag 58
func (m SecurityTypeRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m SecurityTypeRequest) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetSecurityReqID sets SecurityReqID, Tag 320
func (m SecurityTypeRequest) SetSecurityReqID(v string) {
	m.Set(field.NewSecurityReqID(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m SecurityTypeRequest) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m SecurityTypeRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m SecurityTypeRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetProduct sets Product, Tag 460
func (m SecurityTypeRequest) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m SecurityTypeRequest) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m SecurityTypeRequest) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//GetText gets Text, Tag 58
func (m SecurityTypeRequest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m SecurityTypeRequest) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityReqID gets SecurityReqID, Tag 320
func (m SecurityTypeRequest) GetSecurityReqID() (f field.SecurityReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m SecurityTypeRequest) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m SecurityTypeRequest) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m SecurityTypeRequest) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProduct gets Product, Tag 460
func (m SecurityTypeRequest) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m SecurityTypeRequest) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m SecurityTypeRequest) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m SecurityTypeRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m SecurityTypeRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasSecurityReqID returns true if SecurityReqID is present, Tag 320
func (m SecurityTypeRequest) HasSecurityReqID() bool {
	return m.Has(tag.SecurityReqID)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m SecurityTypeRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m SecurityTypeRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m SecurityTypeRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasProduct returns true if Product is present, Tag 460
func (m SecurityTypeRequest) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m SecurityTypeRequest) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m SecurityTypeRequest) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}
