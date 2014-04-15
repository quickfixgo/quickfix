//Package cracker provides message routing utilities
package cracker

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/enum"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/fix/tag"
	"github.com/quickfixgo/quickfix/message"

	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix50"
	"github.com/quickfixgo/quickfix/fixt11"
)

type MessageRouter interface {
	fix40.MessageRouter
	fix41.MessageRouter
	fix42.MessageRouter
	fix43.MessageRouter
	fix44.MessageRouter
	fix50.MessageRouter
	fixt11.MessageRouter
}

type MessageCracker struct {
	fix40.FIX40MessageCracker
	fix41.FIX41MessageCracker
	fix42.FIX42MessageCracker
	fix43.FIX43MessageCracker
	fix44.FIX44MessageCracker
	fix50.FIX50MessageCracker
	fixt11.FIXT11MessageCracker
}

func Crack(msg message.Message, sessionID quickfix.SessionID, router MessageRouter) message.MessageReject {
	beginString := new(field.BeginString)
	msg.Header.Get(beginString)
	return tryCrack(beginString.Value, msg, sessionID, router)
}

func tryCrack(beginString string, msg message.Message, sessionID quickfix.SessionID, router MessageRouter) message.MessageReject {
	switch beginString {
	case fix.BeginString_FIX40:
		return fix40.Crack(msg, sessionID, router)
	case fix.BeginString_FIX41:
		return fix41.Crack(msg, sessionID, router)
	case fix.BeginString_FIX42:
		return fix42.Crack(msg, sessionID, router)
	case fix.BeginString_FIX43:
		return fix43.Crack(msg, sessionID, router)
	case fix.BeginString_FIX44:
		return fix44.Crack(msg, sessionID, router)
	case fix.BeginString_FIX50:
		return fix50.Crack(msg, sessionID, router)
	case fix.BeginString_FIXT11:

		msgType := new(message.StringValue)
		if msg.Header.GetField(tag.MsgType, msgType); fix.IsAdminMessageType(msgType.Value) {
			return fixt11.Crack(msg, sessionID, router)
		}

		applVerID := sessionID.DefaultApplVerID
		applVerIDField := new(field.ApplVerID)
		if err := msg.Header.Get(applVerIDField); err == nil {
			applVerID = applVerIDField.Value
		}

		switch applVerID {
		case enum.ApplVerID_FIX40:
			beginString = fix.BeginString_FIX40
		case enum.ApplVerID_FIX41:
			beginString = fix.BeginString_FIX41
		case enum.ApplVerID_FIX42:
			beginString = fix.BeginString_FIX42
		case enum.ApplVerID_FIX43:
			beginString = fix.BeginString_FIX43
		case enum.ApplVerID_FIX44:
			beginString = fix.BeginString_FIX44
		case enum.ApplVerID_FIX50, enum.ApplVerID_FIX50SP1, enum.ApplVerID_FIX50SP2:
			beginString = fix.BeginString_FIX50
		default:
			beginString = ""
		}

		return tryCrack(beginString, msg, sessionID, router)
	}

	panic("Invalid BeginString")
}
