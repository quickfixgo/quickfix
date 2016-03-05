package applidrequestgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/nestedparties"
)

//NoApplIDs is a repeating group in ApplIDRequestGrp
type NoApplIDs struct {
	//RefApplID is a non-required field for NoApplIDs.
	RefApplID *string `fix:"1355"`
	//ApplBegSeqNum is a non-required field for NoApplIDs.
	ApplBegSeqNum *int `fix:"1182"`
	//ApplEndSeqNum is a non-required field for NoApplIDs.
	ApplEndSeqNum *int `fix:"1183"`
	//NestedParties Component
	nestedparties.NestedParties
	//RefApplReqID is a non-required field for NoApplIDs.
	RefApplReqID *string `fix:"1433"`
}

//ApplIDRequestGrp is a fix50sp2 Component
type ApplIDRequestGrp struct {
	//NoApplIDs is a non-required field for ApplIDRequestGrp.
	NoApplIDs []NoApplIDs `fix:"1351,omitempty"`
}

func (m *ApplIDRequestGrp) SetNoApplIDs(v []NoApplIDs) { m.NoApplIDs = v }
