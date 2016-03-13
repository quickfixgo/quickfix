package settlinstructionsdata

import (
	"github.com/quickfixgo/quickfix/fix50sp1/dlvyinstgrp"
)

//SettlInstructionsData is a fix50sp1 Component
type SettlInstructionsData struct {
	//SettlDeliveryType is a non-required field for SettlInstructionsData.
	SettlDeliveryType *int `fix:"172"`
	//StandInstDbType is a non-required field for SettlInstructionsData.
	StandInstDbType *int `fix:"169"`
	//StandInstDbName is a non-required field for SettlInstructionsData.
	StandInstDbName *string `fix:"170"`
	//StandInstDbID is a non-required field for SettlInstructionsData.
	StandInstDbID *string `fix:"171"`
	//DlvyInstGrp is a non-required component for SettlInstructionsData.
	DlvyInstGrp *dlvyinstgrp.DlvyInstGrp
}

func (m *SettlInstructionsData) SetSettlDeliveryType(v int)               { m.SettlDeliveryType = &v }
func (m *SettlInstructionsData) SetStandInstDbType(v int)                 { m.StandInstDbType = &v }
func (m *SettlInstructionsData) SetStandInstDbName(v string)              { m.StandInstDbName = &v }
func (m *SettlInstructionsData) SetStandInstDbID(v string)                { m.StandInstDbID = &v }
func (m *SettlInstructionsData) SetDlvyInstGrp(v dlvyinstgrp.DlvyInstGrp) { m.DlvyInstGrp = &v }
