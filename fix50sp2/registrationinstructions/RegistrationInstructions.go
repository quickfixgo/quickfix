//Package registrationinstructions msg type = o.
package registrationinstructions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/rgstdistinstgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/rgstdtlsgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a RegistrationInstructions FIX Message
type Message struct {
	FIXMsgType string `fix:"o"`
	Header     fixt11.Header
	//RegistID is a required field for RegistrationInstructions.
	RegistID string `fix:"513"`
	//RegistTransType is a required field for RegistrationInstructions.
	RegistTransType string `fix:"514"`
	//RegistRefID is a required field for RegistrationInstructions.
	RegistRefID string `fix:"508"`
	//ClOrdID is a non-required field for RegistrationInstructions.
	ClOrdID *string `fix:"11"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for RegistrationInstructions.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for RegistrationInstructions.
	AcctIDSource *int `fix:"660"`
	//RegistAcctType is a non-required field for RegistrationInstructions.
	RegistAcctType *string `fix:"493"`
	//TaxAdvantageType is a non-required field for RegistrationInstructions.
	TaxAdvantageType *int `fix:"495"`
	//OwnershipType is a non-required field for RegistrationInstructions.
	OwnershipType *string `fix:"517"`
	//RgstDtlsGrp Component
	RgstDtlsGrp rgstdtlsgrp.Component
	//RgstDistInstGrp Component
	RgstDistInstGrp rgstdistinstgrp.Component
	Trailer         fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "o", r
}
