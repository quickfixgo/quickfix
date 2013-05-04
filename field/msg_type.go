package field

//Returns true if the message type is a session level type
func (m MsgType) IsAdminMessageType() bool {
	switch m.Value {
	case "0", "A", "1", "2", "3", "4", "5":
		return true
	}

	return false
}
