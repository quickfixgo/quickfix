package securitytypes

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//SecurityTypes is the fix43 SecurityTypes type, MsgType = w
type SecurityTypes struct {
	fix43.Header
	*quickfix.Body
	fix43.Trailer
	Message *quickfix.Message
}

//FromMessage creates a SecurityTypes from a quickfix.Message instance
func FromMessage(m *quickfix.Message) SecurityTypes {
	return SecurityTypes{
		Header:  fix43.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix43.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m SecurityTypes) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a SecurityTypes initialized with the required fields for SecurityTypes
func New(securityreqid field.SecurityReqIDField, securityresponseid field.SecurityResponseIDField, securityresponsetype field.SecurityResponseTypeField) (m SecurityTypes) {
	m.Message = quickfix.NewMessage()
	m.Header = fix43.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("w"))
	m.Set(securityreqid)
	m.Set(securityresponseid)
	m.Set(securityresponsetype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg SecurityTypes, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "w", r
}

//SetText sets Text, Tag 58
func (m SecurityTypes) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m SecurityTypes) SetSubscriptionRequestType(v enum.SubscriptionRequestType) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetSecurityReqID sets SecurityReqID, Tag 320
func (m SecurityTypes) SetSecurityReqID(v string) {
	m.Set(field.NewSecurityReqID(v))
}

//SetSecurityResponseID sets SecurityResponseID, Tag 322
func (m SecurityTypes) SetSecurityResponseID(v string) {
	m.Set(field.NewSecurityResponseID(v))
}

//SetSecurityResponseType sets SecurityResponseType, Tag 323
func (m SecurityTypes) SetSecurityResponseType(v enum.SecurityResponseType) {
	m.Set(field.NewSecurityResponseType(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m SecurityTypes) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m SecurityTypes) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m SecurityTypes) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetTotalNumSecurityTypes sets TotalNumSecurityTypes, Tag 557
func (m SecurityTypes) SetTotalNumSecurityTypes(v int) {
	m.Set(field.NewTotalNumSecurityTypes(v))
}

//SetNoSecurityTypes sets NoSecurityTypes, Tag 558
func (m SecurityTypes) SetNoSecurityTypes(f NoSecurityTypesRepeatingGroup) {
	m.SetGroup(f)
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m SecurityTypes) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

//GetText gets Text, Tag 58
func (m SecurityTypes) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m SecurityTypes) GetSubscriptionRequestType() (v enum.SubscriptionRequestType, err quickfix.MessageRejectError) {
	var f field.SubscriptionRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityReqID gets SecurityReqID, Tag 320
func (m SecurityTypes) GetSecurityReqID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityResponseID gets SecurityResponseID, Tag 322
func (m SecurityTypes) GetSecurityResponseID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityResponseIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityResponseType gets SecurityResponseType, Tag 323
func (m SecurityTypes) GetSecurityResponseType() (v enum.SecurityResponseType, err quickfix.MessageRejectError) {
	var f field.SecurityResponseTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m SecurityTypes) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m SecurityTypes) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m SecurityTypes) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotalNumSecurityTypes gets TotalNumSecurityTypes, Tag 557
func (m SecurityTypes) GetTotalNumSecurityTypes() (v int, err quickfix.MessageRejectError) {
	var f field.TotalNumSecurityTypesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoSecurityTypes gets NoSecurityTypes, Tag 558
func (m SecurityTypes) GetNoSecurityTypes() (NoSecurityTypesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityTypesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m SecurityTypes) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasText returns true if Text is present, Tag 58
func (m SecurityTypes) HasText() bool {
	return m.Has(tag.Text)
}

//HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263
func (m SecurityTypes) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

//HasSecurityReqID returns true if SecurityReqID is present, Tag 320
func (m SecurityTypes) HasSecurityReqID() bool {
	return m.Has(tag.SecurityReqID)
}

//HasSecurityResponseID returns true if SecurityResponseID is present, Tag 322
func (m SecurityTypes) HasSecurityResponseID() bool {
	return m.Has(tag.SecurityResponseID)
}

//HasSecurityResponseType returns true if SecurityResponseType is present, Tag 323
func (m SecurityTypes) HasSecurityResponseType() bool {
	return m.Has(tag.SecurityResponseType)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m SecurityTypes) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m SecurityTypes) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m SecurityTypes) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasTotalNumSecurityTypes returns true if TotalNumSecurityTypes is present, Tag 557
func (m SecurityTypes) HasTotalNumSecurityTypes() bool {
	return m.Has(tag.TotalNumSecurityTypes)
}

//HasNoSecurityTypes returns true if NoSecurityTypes is present, Tag 558
func (m SecurityTypes) HasNoSecurityTypes() bool {
	return m.Has(tag.NoSecurityTypes)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m SecurityTypes) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//NoSecurityTypes is a repeating group element, Tag 558
type NoSecurityTypes struct {
	*quickfix.Group
}

//SetSecurityType sets SecurityType, Tag 167
func (m NoSecurityTypes) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetProduct sets Product, Tag 460
func (m NoSecurityTypes) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m NoSecurityTypes) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//GetSecurityType gets SecurityType, Tag 167
func (m NoSecurityTypes) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetProduct gets Product, Tag 460
func (m NoSecurityTypes) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCFICode gets CFICode, Tag 461
func (m NoSecurityTypes) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NoSecurityTypes) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasProduct returns true if Product is present, Tag 460
func (m NoSecurityTypes) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m NoSecurityTypes) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//NoSecurityTypesRepeatingGroup is a repeating group, Tag 558
type NoSecurityTypesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSecurityTypesRepeatingGroup returns an initialized, NoSecurityTypesRepeatingGroup
func NewNoSecurityTypesRepeatingGroup() NoSecurityTypesRepeatingGroup {
	return NoSecurityTypesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityTypes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.Product), quickfix.GroupElement(tag.CFICode)})}
}

//Add create and append a new NoSecurityTypes to this group
func (m NoSecurityTypesRepeatingGroup) Add() NoSecurityTypes {
	g := m.RepeatingGroup.Add()
	return NoSecurityTypes{g}
}

//Get returns the ith NoSecurityTypes in the NoSecurityTypesRepeatinGroup
func (m NoSecurityTypesRepeatingGroup) Get(i int) NoSecurityTypes {
	return NoSecurityTypes{m.RepeatingGroup.Get(i)}
}
