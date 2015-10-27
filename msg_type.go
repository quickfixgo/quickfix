package quickfix

//IsAdminMessageType returns true if the message type is a sesion level message.
func isAdminMessageType(m string) bool {
	switch m {
	case "0", "A", "1", "2", "3", "4", "5":
		return true
	}

	return false
}
