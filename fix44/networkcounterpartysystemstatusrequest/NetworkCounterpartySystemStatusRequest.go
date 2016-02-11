//Package networkcounterpartysystemstatusrequest msg type = BC.
package networkcounterpartysystemstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
)

//NoCompIDs is a repeating group in NetworkCounterpartySystemStatusRequest
type NoCompIDs struct {
	//RefCompID is a non-required field for NoCompIDs.
	RefCompID *string `fix:"930"`
	//RefSubID is a non-required field for NoCompIDs.
	RefSubID *string `fix:"931"`
	//LocationID is a non-required field for NoCompIDs.
	LocationID *string `fix:"283"`
	//DeskID is a non-required field for NoCompIDs.
	DeskID *string `fix:"284"`
}

//Message is a NetworkCounterpartySystemStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"BC"`
	Header     fix44.Header
	//NetworkRequestType is a required field for NetworkCounterpartySystemStatusRequest.
	NetworkRequestType int `fix:"935"`
	//NetworkRequestID is a required field for NetworkCounterpartySystemStatusRequest.
	NetworkRequestID string `fix:"933"`
	//NoCompIDs is a non-required field for NetworkCounterpartySystemStatusRequest.
	NoCompIDs []NoCompIDs `fix:"936,omitempty"`
	Trailer   fix44.Trailer
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
	return enum.BeginStringFIX44, "BC", r
}
