//Package cracker provides message routing utilities
package cracker

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/tag"

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

func Crack(msg quickfix.Message, sessionID quickfix.SessionID, router MessageRouter) quickfix.MessageReject {
	beginString := new(field.BeginString)
	msg.Header.Get(beginString)
	return tryCrack(beginString.Value, msg, sessionID, router)
}

func tryCrack(beginString string, msg quickfix.Message, sessionID quickfix.SessionID, router MessageRouter) quickfix.MessageReject {
	switch beginString {
	case quickfix.BeginString_FIX40:
		return fix40.Crack(msg, sessionID, router)
	case quickfix.BeginString_FIX41:
		return fix41.Crack(msg, sessionID, router)
	case quickfix.BeginString_FIX42:
		return fix42.Crack(msg, sessionID, router)
	case quickfix.BeginString_FIX43:
		return fix43.Crack(msg, sessionID, router)
	case quickfix.BeginString_FIX44:
		return fix44.Crack(msg, sessionID, router)
	case quickfix.BeginString_FIX50:
		return fix50.Crack(msg, sessionID, router)
	case quickfix.BeginString_FIXT11:

		msgType := new(quickfix.StringValue)
		if msg.Header.GetField(tag.MsgType, msgType); quickfix.IsAdminMessageType(msgType.Value) {
			return fixt11.Crack(msg, sessionID, router)
		} else {

			applVerId := sessionID.DefaultApplVerID
			applVerIDField := new(field.ApplVerID)
			if err := msg.Header.Get(applVerIDField); err == nil {
				applVerId = applVerIDField.Value
			}

			switch applVerId {
			case quickfix.ApplVerID_FIX40:
				beginString = quickfix.BeginString_FIX40
			case quickfix.ApplVerID_FIX41:
				beginString = quickfix.BeginString_FIX41
			case quickfix.ApplVerID_FIX42:
				beginString = quickfix.BeginString_FIX42
			case quickfix.ApplVerID_FIX43:
				beginString = quickfix.BeginString_FIX43
			case quickfix.ApplVerID_FIX44:
				beginString = quickfix.BeginString_FIX44
			case quickfix.ApplVerID_FIX50, quickfix.ApplVerID_FIX50SP1, quickfix.ApplVerID_FIX50SP2:
				beginString = quickfix.BeginString_FIX50
			default:
				beginString = ""
			}

			return tryCrack(beginString, msg, sessionID, router)
		}
	}

	panic("Invalid BeginString")
}
