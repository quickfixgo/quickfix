package internal

//Event is an abstraction for session events
type Event int

const (
	//PeerTimeout indicates the session peer has become unresponsive
	PeerTimeout Event = iota
	//NeedHeartbeat indicates the session should send a heartbeat
	NeedHeartbeat
	//LogonTimeout indicates the peer has not sent a logon request
	LogonTimeout
	//LogoutTimeout indicates the peer has not sent a logout request
	LogoutTimeout
)
