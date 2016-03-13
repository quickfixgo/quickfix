package partydetail

import (
	"github.com/quickfixgo/quickfix/fix50sp2/contextparties"
	"github.com/quickfixgo/quickfix/fix50sp2/partyaltids"
	"github.com/quickfixgo/quickfix/fix50sp2/ptyssubgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/risklimits"
)

//PartyDetail is a fix50sp2 Component
type PartyDetail struct {
	//PartyID is a required field for PartyDetail.
	PartyID string `fix:"448"`
	//PartyIDSource is a required field for PartyDetail.
	PartyIDSource string `fix:"447"`
	//PartyRole is a required field for PartyDetail.
	PartyRole int `fix:"452"`
	//PtysSubGrp Component
	ptyssubgrp.PtysSubGrp
	//PartyAltIDs Component
	partyaltids.PartyAltIDs
	//ContextParties Component
	contextparties.ContextParties
	//RiskLimits Component
	risklimits.RiskLimits
}

func (m *PartyDetail) SetPartyID(v string)       { m.PartyID = v }
func (m *PartyDetail) SetPartyIDSource(v string) { m.PartyIDSource = v }
func (m *PartyDetail) SetPartyRole(v int)        { m.PartyRole = v }
