package partydetail

import (
	"github.com/quickfixgo/quickfix/fix50sp2/altptyssubgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/contextptyssubgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/riskinstrumentscope"
	"github.com/quickfixgo/quickfix/fix50sp2/riskwarninglevels"
)

//NoPartySubIDs is a repeating group in PartyDetail
type NoPartySubIDs struct {
	//PartySubID is a non-required field for NoPartySubIDs.
	PartySubID *string `fix:"523"`
	//PartySubIDType is a non-required field for NoPartySubIDs.
	PartySubIDType *int `fix:"803"`
}

//NoPartyAltIDs is a repeating group in PartyDetail
type NoPartyAltIDs struct {
	//PartyAltID is a non-required field for NoPartyAltIDs.
	PartyAltID *string `fix:"1517"`
	//PartyAltIDSource is a non-required field for NoPartyAltIDs.
	PartyAltIDSource *string `fix:"1518"`
	//AltPtysSubGrp Component
	AltPtysSubGrp altptyssubgrp.Component
}

//NoContextPartyIDs is a repeating group in PartyDetail
type NoContextPartyIDs struct {
	//ContextPartyID is a non-required field for NoContextPartyIDs.
	ContextPartyID *string `fix:"1523"`
	//ContextPartyIDSource is a non-required field for NoContextPartyIDs.
	ContextPartyIDSource *string `fix:"1524"`
	//ContextPartyRole is a non-required field for NoContextPartyIDs.
	ContextPartyRole *int `fix:"1525"`
	//ContextPtysSubGrp Component
	ContextPtysSubGrp contextptyssubgrp.Component
}

//NoRiskLimits is a repeating group in PartyDetail
type NoRiskLimits struct {
	//RiskLimitType is a non-required field for NoRiskLimits.
	RiskLimitType *int `fix:"1530"`
	//RiskLimitAmount is a non-required field for NoRiskLimits.
	RiskLimitAmount *float64 `fix:"1531"`
	//RiskLimitCurrency is a non-required field for NoRiskLimits.
	RiskLimitCurrency *string `fix:"1532"`
	//RiskLimitPlatform is a non-required field for NoRiskLimits.
	RiskLimitPlatform *string `fix:"1533"`
	//RiskInstrumentScope Component
	RiskInstrumentScope riskinstrumentscope.Component
	//RiskWarningLevels Component
	RiskWarningLevels riskwarninglevels.Component
}

//Component is a fix50sp2 PartyDetail Component
type Component struct {
	//PartyID is a required field for PartyDetail.
	PartyID string `fix:"448"`
	//PartyIDSource is a required field for PartyDetail.
	PartyIDSource string `fix:"447"`
	//PartyRole is a required field for PartyDetail.
	PartyRole int `fix:"452"`
	//NoPartySubIDs is a non-required field for PartyDetail.
	NoPartySubIDs []NoPartySubIDs `fix:"802,omitempty"`
	//NoPartyAltIDs is a non-required field for PartyDetail.
	NoPartyAltIDs []NoPartyAltIDs `fix:"1516,omitempty"`
	//NoContextPartyIDs is a non-required field for PartyDetail.
	NoContextPartyIDs []NoContextPartyIDs `fix:"1522,omitempty"`
	//NoRiskLimits is a non-required field for PartyDetail.
	NoRiskLimits []NoRiskLimits `fix:"1529,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetPartyID(v string)                        { m.PartyID = v }
func (m *Component) SetPartyIDSource(v string)                  { m.PartyIDSource = v }
func (m *Component) SetPartyRole(v int)                         { m.PartyRole = v }
func (m *Component) SetNoPartySubIDs(v []NoPartySubIDs)         { m.NoPartySubIDs = v }
func (m *Component) SetNoPartyAltIDs(v []NoPartyAltIDs)         { m.NoPartyAltIDs = v }
func (m *Component) SetNoContextPartyIDs(v []NoContextPartyIDs) { m.NoContextPartyIDs = v }
func (m *Component) SetNoRiskLimits(v []NoRiskLimits)           { m.NoRiskLimits = v }
