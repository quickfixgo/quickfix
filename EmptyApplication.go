package quickfix

type EmptyApplication struct{}

var _ Application = EmptyApplication{}

func (e EmptyApplication) OnCreate(_ SessionID) {
}

func (e EmptyApplication) OnLogon(_ SessionID) {
}

func (e EmptyApplication) OnLogout(_ SessionID) {
}

func (e EmptyApplication) ToAdmin(_ *Message, _ SessionID) {
}

func (e EmptyApplication) ToApp(_ *Message, _ SessionID) error {
	return nil
}

func (e EmptyApplication) FromAdmin(_ *Message, _ SessionID) MessageRejectError {
	return nil
}

func (e EmptyApplication) FromApp(_ *Message, _ SessionID) MessageRejectError {
	return nil
}
