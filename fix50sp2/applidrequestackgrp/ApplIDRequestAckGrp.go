package applidrequestackgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties"
)

func New() *ApplIDRequestAckGrp {
	var m ApplIDRequestAckGrp
	return &m
}

//NoApplIDs is a repeating group in ApplIDRequestAckGrp
type NoApplIDs struct {
	//RefApplID is a non-required field for NoApplIDs.
	RefApplID *string `fix:"1355"`
	//ApplBegSeqNum is a non-required field for NoApplIDs.
	ApplBegSeqNum *int `fix:"1182"`
	//ApplEndSeqNum is a non-required field for NoApplIDs.
	ApplEndSeqNum *int `fix:"1183"`
	//RefApplLastSeqNum is a non-required field for NoApplIDs.
	RefApplLastSeqNum *int `fix:"1357"`
	//ApplResponseError is a non-required field for NoApplIDs.
	ApplResponseError *int `fix:"1354"`
	//NestedParties is a non-required component for NoApplIDs.
	NestedParties *nestedparties.NestedParties
	//RefApplReqID is a non-required field for NoApplIDs.
	RefApplReqID *string `fix:"1433"`
}

func (m *NoApplIDs) SetRefApplID(v string)                          { m.RefApplID = &v }
func (m *NoApplIDs) SetApplBegSeqNum(v int)                         { m.ApplBegSeqNum = &v }
func (m *NoApplIDs) SetApplEndSeqNum(v int)                         { m.ApplEndSeqNum = &v }
func (m *NoApplIDs) SetRefApplLastSeqNum(v int)                     { m.RefApplLastSeqNum = &v }
func (m *NoApplIDs) SetApplResponseError(v int)                     { m.ApplResponseError = &v }
func (m *NoApplIDs) SetNestedParties(v nestedparties.NestedParties) { m.NestedParties = &v }
func (m *NoApplIDs) SetRefApplReqID(v string)                       { m.RefApplReqID = &v }

//ApplIDRequestAckGrp is a fix50sp2 Component
type ApplIDRequestAckGrp struct {
	//NoApplIDs is a non-required field for ApplIDRequestAckGrp.
	NoApplIDs []NoApplIDs `fix:"1351,omitempty"`
}

func (m *ApplIDRequestAckGrp) SetNoApplIDs(v []NoApplIDs) { m.NoApplIDs = v }
