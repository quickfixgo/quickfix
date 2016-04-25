//Package registrationinstructions msg type = o.
package registrationinstructions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/rgstdistinstgrp"
	"github.com/quickfixgo/quickfix/fix50/rgstdtlsgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a RegistrationInstructions FIX Message
type Message struct {
	FIXMsgType string `fix:"o"`
	fixt11.Header
	//RegistID is a required field for RegistrationInstructions.
	RegistID string `fix:"513"`
	//RegistTransType is a required field for RegistrationInstructions.
	RegistTransType string `fix:"514"`
	//RegistRefID is a required field for RegistrationInstructions.
	RegistRefID string `fix:"508"`
	//ClOrdID is a non-required field for RegistrationInstructions.
	ClOrdID *string `fix:"11"`
	//Parties is a non-required component for RegistrationInstructions.
	Parties *parties.Parties
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
	//RgstDtlsGrp is a non-required component for RegistrationInstructions.
	RgstDtlsGrp *rgstdtlsgrp.RgstDtlsGrp
	//RgstDistInstGrp is a non-required component for RegistrationInstructions.
	RgstDistInstGrp *rgstdistinstgrp.RgstDistInstGrp
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized RegistrationInstructions instance
func New(registid string, registtranstype string, registrefid string) *Message {
	var m Message
	m.SetRegistID(registid)
	m.SetRegistTransType(registtranstype)
	m.SetRegistRefID(registrefid)
	return &m
}

func (m *Message) SetRegistID(v string)                                 { m.RegistID = v }
func (m *Message) SetRegistTransType(v string)                          { m.RegistTransType = v }
func (m *Message) SetRegistRefID(v string)                              { m.RegistRefID = v }
func (m *Message) SetClOrdID(v string)                                  { m.ClOrdID = &v }
func (m *Message) SetParties(v parties.Parties)                         { m.Parties = &v }
func (m *Message) SetAccount(v string)                                  { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)                                { m.AcctIDSource = &v }
func (m *Message) SetRegistAcctType(v string)                           { m.RegistAcctType = &v }
func (m *Message) SetTaxAdvantageType(v int)                            { m.TaxAdvantageType = &v }
func (m *Message) SetOwnershipType(v string)                            { m.OwnershipType = &v }
func (m *Message) SetRgstDtlsGrp(v rgstdtlsgrp.RgstDtlsGrp)             { m.RgstDtlsGrp = &v }
func (m *Message) SetRgstDistInstGrp(v rgstdistinstgrp.RgstDistInstGrp) { m.RgstDistInstGrp = &v }

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
	return enum.ApplVerID_FIX50, "o", r
}
