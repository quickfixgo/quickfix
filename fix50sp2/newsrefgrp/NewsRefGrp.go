package newsrefgrp

func New() *NewsRefGrp {
	var m NewsRefGrp
	return &m
}

//NoNewsRefIDs is a repeating group in NewsRefGrp
type NoNewsRefIDs struct {
	//NewsRefID is a non-required field for NoNewsRefIDs.
	NewsRefID *string `fix:"1476"`
	//NewsRefType is a non-required field for NoNewsRefIDs.
	NewsRefType *int `fix:"1477"`
}

func (m *NoNewsRefIDs) SetNewsRefID(v string) { m.NewsRefID = &v }
func (m *NoNewsRefIDs) SetNewsRefType(v int)  { m.NewsRefType = &v }

//NewsRefGrp is a fix50sp2 Component
type NewsRefGrp struct {
	//NoNewsRefIDs is a non-required field for NewsRefGrp.
	NoNewsRefIDs []NoNewsRefIDs `fix:"1475,omitempty"`
}

func (m *NewsRefGrp) SetNoNewsRefIDs(v []NoNewsRefIDs) { m.NoNewsRefIDs = v }
