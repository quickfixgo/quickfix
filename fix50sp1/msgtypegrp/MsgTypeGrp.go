package msgtypegrp

func New() *MsgTypeGrp {
	var m MsgTypeGrp
	return &m
}

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
	//RefApplExtID is a non-required field for NoMsgTypes.
	RefApplExtID *int `fix:"1406"`
	//DefaultVerIndicator is a non-required field for NoMsgTypes.
	DefaultVerIndicator *bool `fix:"1410"`
}

func (m *NoMsgTypes) SetRefMsgType(v string)        { m.RefMsgType = &v }
func (m *NoMsgTypes) SetMsgDirection(v string)      { m.MsgDirection = &v }
func (m *NoMsgTypes) SetRefApplVerID(v string)      { m.RefApplVerID = &v }
func (m *NoMsgTypes) SetRefCstmApplVerID(v string)  { m.RefCstmApplVerID = &v }
func (m *NoMsgTypes) SetRefApplExtID(v int)         { m.RefApplExtID = &v }
func (m *NoMsgTypes) SetDefaultVerIndicator(v bool) { m.DefaultVerIndicator = &v }

//MsgTypeGrp is a fix50sp1 Component
type MsgTypeGrp struct {
	//NoMsgTypes is a non-required field for MsgTypeGrp.
	NoMsgTypes []NoMsgTypes `fix:"384,omitempty"`
}

func (m *MsgTypeGrp) SetNoMsgTypes(v []NoMsgTypes) { m.NoMsgTypes = v }
