//Package securitytypes msg type = w.
package securitytypes

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp2/sectypesgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a SecurityTypes FIX Message
type Message struct {
	FIXMsgType string `fix:"w"`
	Header     fixt11.Header
	//SecurityReqID is a required field for SecurityTypes.
	SecurityReqID string `fix:"320"`
	//SecurityResponseID is a required field for SecurityTypes.
	SecurityResponseID string `fix:"322"`
	//SecurityResponseType is a required field for SecurityTypes.
	SecurityResponseType int `fix:"323"`
	//TotNoSecurityTypes is a non-required field for SecurityTypes.
	TotNoSecurityTypes *int `fix:"557"`
	//LastFragment is a non-required field for SecurityTypes.
	LastFragment *bool `fix:"893"`
	//SecTypesGrp Component
	SecTypesGrp sectypesgrp.Component
	//Text is a non-required field for SecurityTypes.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityTypes.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityTypes.
	EncodedText *string `fix:"355"`
	//TradingSessionID is a non-required field for SecurityTypes.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityTypes.
	TradingSessionSubID *string `fix:"625"`
	//SubscriptionRequestType is a non-required field for SecurityTypes.
	SubscriptionRequestType *string `fix:"263"`
	//MarketID is a non-required field for SecurityTypes.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for SecurityTypes.
	MarketSegmentID *string `fix:"1300"`
	//ApplicationSequenceControl Component
	ApplicationSequenceControl applicationsequencecontrol.Component
	Trailer                    fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.ApplVerID_FIX50SP2, "w", r
}
