package quickfix

import (
	"time"
)

type GormSessions struct {
	BeginString      string    `gorm:"column:beginstring;primaryKey;type:varchar(8)"`
	SenderCompId     string    `gorm:"column:sendercompid;primaryKey;type:varchar(64)"`
	SenderSubId      string    `gorm:"column:sendersubid;primaryKey;type:varchar(64)"`
	SenderLocId      string    `gorm:"column:senderlocid;primaryKey;type:varchar(64)"`
	TargetCompId     string    `gorm:"column:targetcompid;primaryKey;type:varchar(64)"`
	TargetSubId      string    `gorm:"column:targetsubid;primaryKey;type:varchar(64)"`
	TargetLocId      string    `gorm:"column:targetlocid;primaryKey;type:varchar(64)"`
	SessionQualifier string    `gorm:"column:session_qualifier;primaryKey;type:varchar(64)"`
	CreationTime     time.Time `gorm:"column:creation_time"`
	IncomingSeqNum   int64     `gorm:"column:incoming_seqnum"`
	OutgoingSeqNum   int64     `gorm:"column:outgoing_seqnum"`
}

func (g GormSessions) TableName() string {
	return "sessions"
}

type GormMessages struct {
	BeginString      string    `gorm:"column:beginstring;primaryKey;type:varchar(8)"`
	SenderCompId     string    `gorm:"column:sendercompid;primaryKey;type:varchar(64)"`
	SenderSubId      string    `gorm:"column:sendersubid;primaryKey;type:varchar(64)"`
	SenderLocId      string    `gorm:"column:senderlocid;primaryKey;type:varchar(64)"`
	TargetCompId     string    `gorm:"column:targetcompid;primaryKey;type:varchar(64)"`
	TargetSubId      string    `gorm:"column:targetsubid;primaryKey;type:varchar(64)"`
	TargetLocId      string    `gorm:"column:targetlocid;primaryKey;type:varchar(64)"`
	SessionQualifier string    `gorm:"column:session_qualifier;primaryKey;type:varchar(64)"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	Message          string    `gorm:"column:message;type:text"`
	MsgSeqNum        int64     `gorm:"column:msgseqnum;primaryKey"`
}

func (g GormMessages) TableName() string {
	return "messages"
}
