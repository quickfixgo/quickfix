//Package partydetailslistrequest msg type = CF.
package partydetailslistrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/partylistresponsetypegrp"
	"github.com/quickfixgo/quickfix/fix50sp2/partyrelationships"
	"github.com/quickfixgo/quickfix/fix50sp2/requestedpartyrolegrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a PartyDetailsListRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"CF"`
	Header     fixt11.Header
	//PartyDetailsListRequestID is a required field for PartyDetailsListRequest.
	PartyDetailsListRequestID string `fix:"1505"`
	//PartyListResponseTypeGrp Component
	PartyListResponseTypeGrp partylistresponsetypegrp.Component
	//Parties Component
	Parties parties.Component
	//RequestedPartyRoleGrp Component
	RequestedPartyRoleGrp requestedpartyrolegrp.Component
	//PartyRelationships Component
	PartyRelationships partyrelationships.Component
	//SubscriptionRequestType is a non-required field for PartyDetailsListRequest.
	SubscriptionRequestType *string `fix:"263"`
	//Text is a non-required field for PartyDetailsListRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for PartyDetailsListRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for PartyDetailsListRequest.
	EncodedText *string `fix:"355"`
	Trailer     fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "CF", r
}
