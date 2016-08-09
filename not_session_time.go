package quickfix

type notSessionTime struct{ latentState }

func (notSessionTime) IsSessionTime() bool { return false }
