//Package registrationinstructionsresponse msg type = p.
package registrationinstructionsresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/parties"
)

//Message is a RegistrationInstructionsResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"p"`
	fix43.Header
	//RegistID is a required field for RegistrationInstructionsResponse.
	RegistID string `fix:"513"`
	//RegistTransType is a required field for RegistrationInstructionsResponse.
	RegistTransType string `fix:"514"`
	//RegistRefID is a required field for RegistrationInstructionsResponse.
	RegistRefID string `fix:"508"`
	//ClOrdID is a non-required field for RegistrationInstructionsResponse.
	ClOrdID *string `fix:"11"`
	//Parties is a non-required component for RegistrationInstructionsResponse.
	Parties *parties.Parties
	//Account is a non-required field for RegistrationInstructionsResponse.
	Account *string `fix:"1"`
	//RegistStatus is a required field for RegistrationInstructionsResponse.
	RegistStatus string `fix:"506"`
	//RegistRejReasonCode is a non-required field for RegistrationInstructionsResponse.
	RegistRejReasonCode *int `fix:"507"`
	//RegistRejReasonText is a non-required field for RegistrationInstructionsResponse.
	RegistRejReasonText *string `fix:"496"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetRegistID(v string)            { m.RegistID = v }
func (m *Message) SetRegistTransType(v string)     { m.RegistTransType = v }
func (m *Message) SetRegistRefID(v string)         { m.RegistRefID = v }
func (m *Message) SetClOrdID(v string)             { m.ClOrdID = &v }
func (m *Message) SetParties(v parties.Parties)    { m.Parties = &v }
func (m *Message) SetAccount(v string)             { m.Account = &v }
func (m *Message) SetRegistStatus(v string)        { m.RegistStatus = v }
func (m *Message) SetRegistRejReasonCode(v int)    { m.RegistRejReasonCode = &v }
func (m *Message) SetRegistRejReasonText(v string) { m.RegistRejReasonText = &v }

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
	return enum.BeginStringFIX43, "p", r
}
