package message

import (
	"time"
)

type IntField interface {
	Field

	//field value as an integer
	IntValue() int
}

type StringField interface {
	Field
}

type UTCTimestampField interface {
	Field

	//field value as timestamp
	UTCTimestampValue() time.Time
}
