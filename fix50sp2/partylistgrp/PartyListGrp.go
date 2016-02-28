package partylistgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/partydetail"
	"github.com/quickfixgo/quickfix/fix50sp2/relatedpartygrp"
)

//NoPartyList is a repeating group in PartyListGrp
type NoPartyList struct {
	//PartyDetail Component
	PartyDetail partydetail.Component
	//RelatedPartyGrp Component
	RelatedPartyGrp relatedpartygrp.Component
}

//Component is a fix50sp2 PartyListGrp Component
type Component struct {
	//NoPartyList is a non-required field for PartyListGrp.
	NoPartyList []NoPartyList `fix:"1513,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoPartyList(v []NoPartyList) { m.NoPartyList = v }
