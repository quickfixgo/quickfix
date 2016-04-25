//Package securitydefinitionrequest msg type = c.
package securitydefinitionrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//SubscriptionRequestType is a repeating group in SecurityDefinitionRequest
type SubscriptionRequestType struct {
	//LegOptionRatio is a non-required field for SubscriptionRequestType.
	LegOptionRatio *float64 `fix:"1017"`
	//LegPrice is a non-required field for SubscriptionRequestType.
	LegPrice *float64 `fix:"566"`
}

//NewSubscriptionRequestType returns an initialized SubscriptionRequestType instance
func NewSubscriptionRequestType() *SubscriptionRequestType {
	var m SubscriptionRequestType
	return &m
}

func (m *SubscriptionRequestType) SetLegOptionRatio(v float64) { m.LegOptionRatio = &v }
func (m *SubscriptionRequestType) SetLegPrice(v float64)       { m.LegPrice = &v }

//Message is a SecurityDefinitionRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"c"`
	fixt11.Header
	//SecurityReqID is a required field for SecurityDefinitionRequest.
	SecurityReqID string `fix:"320"`
	//SecurityRequestType is a required field for SecurityDefinitionRequest.
	SecurityRequestType int `fix:"321"`
	//Instrument is a non-required component for SecurityDefinitionRequest.
	Instrument *instrument.Instrument
	//InstrumentExtension is a non-required component for SecurityDefinitionRequest.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//UndInstrmtGrp is a non-required component for SecurityDefinitionRequest.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//Currency is a non-required field for SecurityDefinitionRequest.
	Currency *string `fix:"15"`
	//Text is a non-required field for SecurityDefinitionRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityDefinitionRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityDefinitionRequest.
	EncodedText *string `fix:"355"`
	//TradingSessionID is a non-required field for SecurityDefinitionRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityDefinitionRequest.
	TradingSessionSubID *string `fix:"625"`
	//InstrmtLegGrp is a non-required component for SecurityDefinitionRequest.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//ExpirationCycle is a non-required field for SecurityDefinitionRequest.
	ExpirationCycle *int `fix:"827"`
	//SubscriptionRequestType is a non-required field for SecurityDefinitionRequest.
	SubscriptionRequestType []SubscriptionRequestType `fix:"263,omitempty"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized SecurityDefinitionRequest instance
func New(securityreqid string, securityrequesttype int) *Message {
	var m Message
	m.SetSecurityReqID(securityreqid)
	m.SetSecurityRequestType(securityrequesttype)
	return &m
}

func (m *Message) SetSecurityReqID(v string)             { m.SecurityReqID = v }
func (m *Message) SetSecurityRequestType(v int)          { m.SecurityRequestType = v }
func (m *Message) SetInstrument(v instrument.Instrument) { m.Instrument = &v }
func (m *Message) SetInstrumentExtension(v instrumentextension.InstrumentExtension) {
	m.InstrumentExtension = &v
}
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
func (m *Message) SetCurrency(v string)                           { m.Currency = &v }
func (m *Message) SetText(v string)                               { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                        { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                        { m.EncodedText = &v }
func (m *Message) SetTradingSessionID(v string)                   { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)                { m.TradingSessionSubID = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp) { m.InstrmtLegGrp = &v }
func (m *Message) SetExpirationCycle(v int)                       { m.ExpirationCycle = &v }
func (m *Message) SetSubscriptionRequestType(v []SubscriptionRequestType) {
	m.SubscriptionRequestType = v
}

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
	return enum.ApplVerID_FIX50, "c", r
}
