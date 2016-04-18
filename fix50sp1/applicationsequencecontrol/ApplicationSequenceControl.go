package applicationsequencecontrol

//New returns an initialized ApplicationSequenceControl instance
func New() *ApplicationSequenceControl {
	var m ApplicationSequenceControl
	return &m
}

//ApplicationSequenceControl is a fix50sp1 Component
type ApplicationSequenceControl struct {
	//ApplID is a non-required field for ApplicationSequenceControl.
	ApplID *string `fix:"1180"`
	//ApplSeqNum is a non-required field for ApplicationSequenceControl.
	ApplSeqNum *int `fix:"1181"`
	//ApplLastSeqNum is a non-required field for ApplicationSequenceControl.
	ApplLastSeqNum *int `fix:"1350"`
	//ApplResendFlag is a non-required field for ApplicationSequenceControl.
	ApplResendFlag *bool `fix:"1352"`
}

func (m *ApplicationSequenceControl) SetApplID(v string)       { m.ApplID = &v }
func (m *ApplicationSequenceControl) SetApplSeqNum(v int)      { m.ApplSeqNum = &v }
func (m *ApplicationSequenceControl) SetApplLastSeqNum(v int)  { m.ApplLastSeqNum = &v }
func (m *ApplicationSequenceControl) SetApplResendFlag(v bool) { m.ApplResendFlag = &v }
