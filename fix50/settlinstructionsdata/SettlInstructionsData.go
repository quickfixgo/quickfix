package settlinstructionsdata

import (
	"github.com/quickfixgo/quickfix/fix50/settlparties"
)

//NoDlvyInst is a repeating group in SettlInstructionsData
type NoDlvyInst struct {
	//SettlInstSource is a non-required field for NoDlvyInst.
	SettlInstSource *string `fix:"165"`
	//DlvyInstType is a non-required field for NoDlvyInst.
	DlvyInstType *string `fix:"787"`
	//SettlParties Component
	settlparties.SettlParties
}

//SettlInstructionsData is a fix50 Component
type SettlInstructionsData struct {
	//SettlDeliveryType is a non-required field for SettlInstructionsData.
	SettlDeliveryType *int `fix:"172"`
	//StandInstDbType is a non-required field for SettlInstructionsData.
	StandInstDbType *int `fix:"169"`
	//StandInstDbName is a non-required field for SettlInstructionsData.
	StandInstDbName *string `fix:"170"`
	//StandInstDbID is a non-required field for SettlInstructionsData.
	StandInstDbID *string `fix:"171"`
	//NoDlvyInst is a non-required field for SettlInstructionsData.
	NoDlvyInst []NoDlvyInst `fix:"85,omitempty"`
}

func (m *SettlInstructionsData) SetSettlDeliveryType(v int)   { m.SettlDeliveryType = &v }
func (m *SettlInstructionsData) SetStandInstDbType(v int)     { m.StandInstDbType = &v }
func (m *SettlInstructionsData) SetStandInstDbName(v string)  { m.StandInstDbName = &v }
func (m *SettlInstructionsData) SetStandInstDbID(v string)    { m.StandInstDbID = &v }
func (m *SettlInstructionsData) SetNoDlvyInst(v []NoDlvyInst) { m.NoDlvyInst = v }
