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

//Message is a BidResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"l"`
	Header     fix42.Header
	//BidID is a non-required field for BidResponse.
	BidID *string `fix:"390"`
	//ClientBidID is a non-required field for BidResponse.
	ClientBidID *string `fix:"391"`
	//NoBidComponents is a required field for BidResponse.
	NoBidComponents []NoBidComponents `fix:"420"`
	Trailer         fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
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
