package settlobligationinstructions

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/settldetails"
	"time"
)

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

//SettlObligationInstructions is a fix50sp1 Component
type SettlObligationInstructions struct {
	//NoSettlOblig is a non-required field for SettlObligationInstructions.
	NoSettlOblig []NoSettlOblig `fix:"1165,omitempty"`
}

func (m *SettlObligationInstructions) SetNoSettlOblig(v []NoSettlOblig) { m.NoSettlOblig = v }
