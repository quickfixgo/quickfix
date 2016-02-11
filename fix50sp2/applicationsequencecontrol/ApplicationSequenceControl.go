package applicationsequencecontrol

//Component is a fix50sp2 ApplicationSequenceControl Component
type Component struct {
	//ApplID is a non-required field for ApplicationSequenceControl.
	ApplID *string `fix:"1180"`
	//ApplSeqNum is a non-required field for ApplicationSequenceControl.
	ApplSeqNum *int `fix:"1181"`
	//ApplLastSeqNum is a non-required field for ApplicationSequenceControl.
	ApplLastSeqNum *int `fix:"1350"`
	//ApplResendFlag is a non-required field for ApplicationSequenceControl.
	ApplResendFlag *bool `fix:"1352"`
}

func New() *Component { return new(Component) }
