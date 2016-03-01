//Package marketdefinitionrequest msg type = BT.
package marketdefinitionrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a MarketDefinitionRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"BT"`
	Header     fixt11.Header
	//MarketReqID is a required field for MarketDefinitionRequest.
	MarketReqID string `fix:"1393"`
	//SubscriptionRequestType is a required field for MarketDefinitionRequest.
	SubscriptionRequestType string `fix:"263"`
	//MarketID is a non-required field for MarketDefinitionRequest.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for MarketDefinitionRequest.
	MarketSegmentID *string `fix:"1300"`
	//ParentMktSegmID is a non-required field for MarketDefinitionRequest.
	ParentMktSegmID *string `fix:"1325"`
	Trailer         fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "BT", r
}
