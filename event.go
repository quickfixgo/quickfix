package quickfix

type event int

const (
	peerTimeout event = iota
	needHeartbeat
	logonTimeout
	logoutTimeout
)
