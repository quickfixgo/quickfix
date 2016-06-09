package marketdefinitionrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//MarketDefinitionRequest is the fix50sp2 MarketDefinitionRequest type, MsgType = BT
type MarketDefinitionRequest struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
}

//FromMessage creates a MarketDefinitionRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) MarketDefinitionRequest {
	return MarketDefinitionRequest{
		Header:  fixt11.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fixt11.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m MarketDefinitionRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a MarketDefinitionRequest initialized with the required fields for MarketDefinitionRequest
func New(marketreqid field.MarketReqIDField, subscriptionrequesttype field.SubscriptionRequestTypeField) (m MarketDefinitionRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("BT"))
	m.Set(marketreqid)
	m.Set(subscriptionrequesttype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg MarketDefinitionRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "BT", r
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m MarketDefinitionRequest) SetSubscriptionRequestType(v string) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetMarketSegmentID sets MarketSegmentID, Tag 1300
func (m MarketDefinitionRequest) SetMarketSegmentID(v string) {
	m.Set(field.NewMarketSegmentID(v))
}

//SetMarketID sets MarketID, Tag 1301
func (m MarketDefinitionRequest) SetMarketID(v string) {
	m.Set(field.NewMarketID(v))
}

//SetParentMktSegmID sets ParentMktSegmID, Tag 1325
func (m MarketDefinitionRequest) SetParentMktSegmID(v string) {
	m.Set(field.NewParentMktSegmID(v))
}

//SetMarketReqID sets MarketReqID, Tag 1393
func (m MarketDefinitionRequest) SetMarketReqID(v string) {
	m.Set(field.NewMarketReqID(v))
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m MarketDefinitionRequest) GetSubscriptionRequestType() (f field.SubscriptionRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarketSegmentID gets MarketSegmentID, Tag 1300
func (m MarketDefinitionRequest) GetMarketSegmentID() (f field.MarketSegmentIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarketID gets MarketID, Tag 1301
func (m MarketDefinitionRequest) GetMarketID() (f field.MarketIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetParentMktSegmID gets ParentMktSegmID, Tag 1325
func (m MarketDefinitionRequest) GetParentMktSegmID() (f field.ParentMktSegmIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarketReqID gets MarketReqID, Tag 1393
func (m MarketDefinitionRequest) GetMarketReqID() (f field.MarketReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263
func (m MarketDefinitionRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

//HasMarketSegmentID returns true if MarketSegmentID is present, Tag 1300
func (m MarketDefinitionRequest) HasMarketSegmentID() bool {
	return m.Has(tag.MarketSegmentID)
}

//HasMarketID returns true if MarketID is present, Tag 1301
func (m MarketDefinitionRequest) HasMarketID() bool {
	return m.Has(tag.MarketID)
}

//HasParentMktSegmID returns true if ParentMktSegmID is present, Tag 1325
func (m MarketDefinitionRequest) HasParentMktSegmID() bool {
	return m.Has(tag.ParentMktSegmID)
}

//HasMarketReqID returns true if MarketReqID is present, Tag 1393
func (m MarketDefinitionRequest) HasMarketReqID() bool {
	return m.Has(tag.MarketReqID)
}
