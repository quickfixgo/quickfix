package msgtypegrp

//NoMsgTypes is a repeating group in MsgTypeGrp
type NoMsgTypes struct {
	//RefMsgType is a non-required field for NoMsgTypes.
	RefMsgType *string `fix:"372"`
	//MsgDirection is a non-required field for NoMsgTypes.
	MsgDirection *string `fix:"385"`
	//RefApplVerID is a non-required field for NoMsgTypes.
	RefApplVerID *string `fix:"1130"`
	//RefCstmApplVerID is a non-required field for NoMsgTypes.
	RefCstmApplVerID *string `fix:"1131"`
}

//Component is a fixt11 MsgTypeGrp Component
type Component struct {
	//NoMsgTypes is a non-required field for MsgTypeGrp.
	NoMsgTypes []NoMsgTypes `fix:"384,omitempty"`
}

func New() *Component { return new(Component) }
