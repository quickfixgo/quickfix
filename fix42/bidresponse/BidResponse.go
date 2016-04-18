//Package bidresponse msg type = l.
package bidresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
)

//NoBidComponents is a repeating group in BidResponse
type NoBidComponents struct {
	//Commission is a required field for NoBidComponents.
	Commission float64 `fix:"12"`
	//CommType is a required field for NoBidComponents.
	CommType string `fix:"13"`
	//ListID is a non-required field for NoBidComponents.
	ListID *string `fix:"66"`
	//Country is a non-required field for NoBidComponents.
	Country *string `fix:"421"`
	//Side is a non-required field for NoBidComponents.
	Side *string `fix:"54"`
	//Price is a non-required field for NoBidComponents.
	Price *float64 `fix:"44"`
	//PriceType is a non-required field for NoBidComponents.
	PriceType *int `fix:"423"`
	//FairValue is a non-required field for NoBidComponents.
	FairValue *float64 `fix:"406"`
	//NetGrossInd is a non-required field for NoBidComponents.
	NetGrossInd *int `fix:"430"`
	//SettlmntTyp is a non-required field for NoBidComponents.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for NoBidComponents.
	FutSettDate *string `fix:"64"`
	//TradingSessionID is a non-required field for NoBidComponents.
	TradingSessionID *string `fix:"336"`
	//Text is a non-required field for NoBidComponents.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoBidComponents.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoBidComponents.
	EncodedText *string `fix:"355"`
}

//NewNoBidComponents returns an initialized NoBidComponents instance
func NewNoBidComponents(commission float64, commtype string) *NoBidComponents {
	var m NoBidComponents
	m.SetCommission(commission)
	m.SetCommType(commtype)
	return &m
}

func (m *NoBidComponents) SetCommission(v float64)      { m.Commission = v }
func (m *NoBidComponents) SetCommType(v string)         { m.CommType = v }
func (m *NoBidComponents) SetListID(v string)           { m.ListID = &v }
func (m *NoBidComponents) SetCountry(v string)          { m.Country = &v }
func (m *NoBidComponents) SetSide(v string)             { m.Side = &v }
func (m *NoBidComponents) SetPrice(v float64)           { m.Price = &v }
func (m *NoBidComponents) SetPriceType(v int)           { m.PriceType = &v }
func (m *NoBidComponents) SetFairValue(v float64)       { m.FairValue = &v }
func (m *NoBidComponents) SetNetGrossInd(v int)         { m.NetGrossInd = &v }
func (m *NoBidComponents) SetSettlmntTyp(v string)      { m.SettlmntTyp = &v }
func (m *NoBidComponents) SetFutSettDate(v string)      { m.FutSettDate = &v }
func (m *NoBidComponents) SetTradingSessionID(v string) { m.TradingSessionID = &v }
func (m *NoBidComponents) SetText(v string)             { m.Text = &v }
func (m *NoBidComponents) SetEncodedTextLen(v int)      { m.EncodedTextLen = &v }
func (m *NoBidComponents) SetEncodedText(v string)      { m.EncodedText = &v }

//Message is a BidResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"l"`
	fix42.Header
	//BidID is a non-required field for BidResponse.
	BidID *string `fix:"390"`
	//ClientBidID is a non-required field for BidResponse.
	ClientBidID *string `fix:"391"`
	//NoBidComponents is a required field for BidResponse.
	NoBidComponents []NoBidComponents `fix:"420"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetBidID(v string)                      { m.BidID = &v }
func (m *Message) SetClientBidID(v string)                { m.ClientBidID = &v }
func (m *Message) SetNoBidComponents(v []NoBidComponents) { m.NoBidComponents = v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX42, "l", r
}
