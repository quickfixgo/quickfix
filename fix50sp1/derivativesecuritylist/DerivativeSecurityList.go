//Package derivativesecuritylist msg type = AA.
package derivativesecuritylist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp1/derivativesecuritydefinition"
	"github.com/quickfixgo/quickfix/fix50sp1/relsymderivsecgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a DerivativeSecurityList FIX Message
type Message struct {
	FIXMsgType string `fix:"AA"`
	Header     fixt11.Header
	//SecurityReqID is a non-required field for DerivativeSecurityList.
	SecurityReqID *string `fix:"320"`
	//SecurityResponseID is a required field for DerivativeSecurityList.
	SecurityResponseID string `fix:"322"`
	//SecurityRequestResult is a non-required field for DerivativeSecurityList.
	SecurityRequestResult *int `fix:"560"`
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//TotNoRelatedSym is a non-required field for DerivativeSecurityList.
	TotNoRelatedSym *int `fix:"393"`
	//LastFragment is a non-required field for DerivativeSecurityList.
	LastFragment *bool `fix:"893"`
	//RelSymDerivSecGrp Component
	RelSymDerivSecGrp relsymderivsecgrp.Component
	//DerivativeSecurityDefinition Component
	DerivativeSecurityDefinition derivativesecuritydefinition.Component
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
	return enum.ApplVerID_FIX50SP1, "AA", r
}
