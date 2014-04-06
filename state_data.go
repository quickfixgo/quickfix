package quickfix

import (
	"time"
)

type stateData struct {
	expectedSeqNum   int
	seqNum           int
	heartBeatTimeout time.Duration
}
