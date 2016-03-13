package relatedpartygrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/partyrelationships"
	"github.com/quickfixgo/quickfix/fix50sp2/relatedpartydetail"
)

//NoRelatedPartyIDs is a repeating group in RelatedPartyGrp
type NoRelatedPartyIDs struct {
	//RelatedPartyDetail is a non-required component for NoRelatedPartyIDs.
	RelatedPartyDetail *relatedpartydetail.RelatedPartyDetail
	//PartyRelationships is a non-required component for NoRelatedPartyIDs.
	PartyRelationships *partyrelationships.PartyRelationships
}

//RelatedPartyGrp is a fix50sp2 Component
type RelatedPartyGrp struct {
	//NoRelatedPartyIDs is a non-required field for RelatedPartyGrp.
	NoRelatedPartyIDs []NoRelatedPartyIDs `fix:"1562,omitempty"`
}

func (m *RelatedPartyGrp) SetNoRelatedPartyIDs(v []NoRelatedPartyIDs) { m.NoRelatedPartyIDs = v }
