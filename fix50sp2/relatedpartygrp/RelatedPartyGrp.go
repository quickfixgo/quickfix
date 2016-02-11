package relatedpartygrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/partyrelationships"
	"github.com/quickfixgo/quickfix/fix50sp2/relatedpartydetail"
)

//NoRelatedPartyIDs is a repeating group in RelatedPartyGrp
type NoRelatedPartyIDs struct {
	//RelatedPartyDetail Component
	RelatedPartyDetail relatedpartydetail.Component
	//PartyRelationships Component
	PartyRelationships partyrelationships.Component
}

//Component is a fix50sp2 RelatedPartyGrp Component
type Component struct {
	//NoRelatedPartyIDs is a non-required field for RelatedPartyGrp.
	NoRelatedPartyIDs []NoRelatedPartyIDs `fix:"1562,omitempty"`
}

func New() *Component { return new(Component) }
