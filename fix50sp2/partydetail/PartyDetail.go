package partydetail

import (
	"github.com/quickfixgo/quickfix/fix50sp2/contextparties"
	"github.com/quickfixgo/quickfix/fix50sp2/partyaltids"
	"github.com/quickfixgo/quickfix/fix50sp2/ptyssubgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/risklimits"
)

func New(partyid string, partyidsource string, partyrole int) *PartyDetail {
	var m PartyDetail
	m.SetPartyID(partyid)
	m.SetPartyIDSource(partyidsource)
	m.SetPartyRole(partyrole)
	return &m
}

//PartyDetail is a fix50sp2 Component
type PartyDetail struct {
	//PartyID is a required field for PartyDetail.
	PartyID string `fix:"448"`
	//PartyIDSource is a required field for PartyDetail.
	PartyIDSource string `fix:"447"`
	//PartyRole is a required field for PartyDetail.
	PartyRole int `fix:"452"`
	//PtysSubGrp is a non-required component for PartyDetail.
	PtysSubGrp *ptyssubgrp.PtysSubGrp
	//PartyAltIDs is a non-required component for PartyDetail.
	PartyAltIDs *partyaltids.PartyAltIDs
	//ContextParties is a non-required component for PartyDetail.
	ContextParties *contextparties.ContextParties
	//RiskLimits is a non-required component for PartyDetail.
	RiskLimits *risklimits.RiskLimits
}

func (m *PartyDetail) SetPartyID(v string)                               { m.PartyID = v }
func (m *PartyDetail) SetPartyIDSource(v string)                         { m.PartyIDSource = v }
func (m *PartyDetail) SetPartyRole(v int)                                { m.PartyRole = v }
func (m *PartyDetail) SetPtysSubGrp(v ptyssubgrp.PtysSubGrp)             { m.PtysSubGrp = &v }
func (m *PartyDetail) SetPartyAltIDs(v partyaltids.PartyAltIDs)          { m.PartyAltIDs = &v }
func (m *PartyDetail) SetContextParties(v contextparties.ContextParties) { m.ContextParties = &v }
func (m *PartyDetail) SetRiskLimits(v risklimits.RiskLimits)             { m.RiskLimits = &v }
