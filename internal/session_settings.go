package internal

import "time"

// SessionSettings stores all of the configuration for a given session.
type SessionSettings struct {
	ResetOnLogon                 bool
	RefreshOnLogon               bool
	ResetOnLogout                bool
	ResetOnDisconnect            bool
	HeartBtInt                   time.Duration
	HeartBtIntOverride           bool
	SessionTime                  *TimeRange
	InitiateLogon                bool
	ResendRequestChunkSize       int
	EnableLastMsgSeqNumProcessed bool
	EnableNextExpectedMsgSeqNum  bool
	SkipCheckLatency             bool
	MaxLatency                   time.Duration
	DisableMessagePersist        bool

	// Logon recovery relies on 789 alone, see config.NextExpectedMsgSeqNumRecovery.
	NextExpectedMsgSeqNumRecovery bool

	// Required on logon for FIX.T.1 messages.
	DefaultApplVerID string

	// Specific to initiators.
	ReconnectInterval    time.Duration
	LogoutTimeout        time.Duration
	LogonTimeout         time.Duration
	SocketWriteTimeout   time.Duration
	SocketReadTimeout    time.Duration
	SocketConnectAddress []string
}
