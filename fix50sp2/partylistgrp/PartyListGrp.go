package partylistgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/partydetail"
	"github.com/quickfixgo/quickfix/fix50sp2/relatedpartygrp"
)

//NoPartyList is a repeating group in PartyListGrp
type NoPartyList struct {
	//PartyDetail Component
	partydetail.PartyDetail
	//RelatedPartyGrp Component
	relatedpartygrp.RelatedPartyGrp
}

//PartyListGrp is a fix50sp2 Component
type PartyListGrp struct {
	//NoPartyList is a non-required field for PartyListGrp.
	NoPartyList []NoPartyList `fix:"1513,omitempty"`
}

func (m *PartyListGrp) SetNoPartyList(v []NoPartyList) { m.NoPartyList = v }
