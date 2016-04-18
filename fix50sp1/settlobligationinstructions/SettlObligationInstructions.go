package settlobligationinstructions

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/settldetails"
	"time"
)

//New returns an initialized SettlObligationInstructions instance
func New() *SettlObligationInstructions {
	var m SettlObligationInstructions
	return &m
}

//NoSettlOblig is a repeating group in SettlObligationInstructions
type NoSettlOblig struct {
	//NetGrossInd is a non-required field for NoSettlOblig.
	NetGrossInd *int `fix:"430"`
	//SettlObligID is a non-required field for NoSettlOblig.
	SettlObligID *string `fix:"1161"`
	//SettlObligTransType is a non-required field for NoSettlOblig.
	SettlObligTransType *string `fix:"1162"`
	//SettlObligRefID is a non-required field for NoSettlOblig.
	SettlObligRefID *string `fix:"1163"`
	//CcyAmt is a non-required field for NoSettlOblig.
	CcyAmt *float64 `fix:"1157"`
	//SettlCurrAmt is a non-required field for NoSettlOblig.
	SettlCurrAmt *float64 `fix:"119"`
	//Currency is a non-required field for NoSettlOblig.
	Currency *string `fix:"15"`
	//SettlCurrency is a non-required field for NoSettlOblig.
	SettlCurrency *string `fix:"120"`
	//SettlCurrFxRate is a non-required field for NoSettlOblig.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlDate is a non-required field for NoSettlOblig.
	SettlDate *string `fix:"64"`
	//Instrument is a non-required component for NoSettlOblig.
	Instrument *instrument.Instrument
	//Parties is a non-required component for NoSettlOblig.
	Parties *parties.Parties
	//EffectiveTime is a non-required field for NoSettlOblig.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireTime is a non-required field for NoSettlOblig.
	ExpireTime *time.Time `fix:"126"`
	//LastUpdateTime is a non-required field for NoSettlOblig.
	LastUpdateTime *time.Time `fix:"779"`
	//SettlDetails is a non-required component for NoSettlOblig.
	SettlDetails *settldetails.SettlDetails
}

//NewNoSettlOblig returns an initialized NoSettlOblig instance
func NewNoSettlOblig() *NoSettlOblig {
	var m NoSettlOblig
	return &m
}

func (m *NoSettlOblig) SetNetGrossInd(v int)                        { m.NetGrossInd = &v }
func (m *NoSettlOblig) SetSettlObligID(v string)                    { m.SettlObligID = &v }
func (m *NoSettlOblig) SetSettlObligTransType(v string)             { m.SettlObligTransType = &v }
func (m *NoSettlOblig) SetSettlObligRefID(v string)                 { m.SettlObligRefID = &v }
func (m *NoSettlOblig) SetCcyAmt(v float64)                         { m.CcyAmt = &v }
func (m *NoSettlOblig) SetSettlCurrAmt(v float64)                   { m.SettlCurrAmt = &v }
func (m *NoSettlOblig) SetCurrency(v string)                        { m.Currency = &v }
func (m *NoSettlOblig) SetSettlCurrency(v string)                   { m.SettlCurrency = &v }
func (m *NoSettlOblig) SetSettlCurrFxRate(v float64)                { m.SettlCurrFxRate = &v }
func (m *NoSettlOblig) SetSettlDate(v string)                       { m.SettlDate = &v }
func (m *NoSettlOblig) SetInstrument(v instrument.Instrument)       { m.Instrument = &v }
func (m *NoSettlOblig) SetParties(v parties.Parties)                { m.Parties = &v }
func (m *NoSettlOblig) SetEffectiveTime(v time.Time)                { m.EffectiveTime = &v }
func (m *NoSettlOblig) SetExpireTime(v time.Time)                   { m.ExpireTime = &v }
func (m *NoSettlOblig) SetLastUpdateTime(v time.Time)               { m.LastUpdateTime = &v }
func (m *NoSettlOblig) SetSettlDetails(v settldetails.SettlDetails) { m.SettlDetails = &v }

//SettlObligationInstructions is a fix50sp1 Component
type SettlObligationInstructions struct {
	//NoSettlOblig is a non-required field for SettlObligationInstructions.
	NoSettlOblig []NoSettlOblig `fix:"1165,omitempty"`
}

func (m *SettlObligationInstructions) SetNoSettlOblig(v []NoSettlOblig) { m.NoSettlOblig = v }
