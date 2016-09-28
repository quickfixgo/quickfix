package quickfix

import "bytes"

var msgTypeHeartbeat = []byte("0")
var msgTypeLogon = []byte("A")
var msgTypeTestRequest = []byte("1")
var msgTypeResendRequest = []byte("2")
var msgTypeReject = []byte("3")
var msgTypeSequenceReset = []byte("4")
var msgTypeLogout = []byte("5")

//isAdminMessageType returns true if the message type is a session level message.
func isAdminMessageType(m []byte) bool {
	switch {
	case bytes.Equal(msgTypeHeartbeat, m),
		bytes.Equal(msgTypeLogon, m),
		bytes.Equal(msgTypeTestRequest, m),
		bytes.Equal(msgTypeResendRequest, m),
		bytes.Equal(msgTypeReject, m),
		bytes.Equal(msgTypeSequenceReset, m),
		bytes.Equal(msgTypeLogout, m):
		return true
	}

	return false
}
