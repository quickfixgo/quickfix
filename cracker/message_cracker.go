//Package cracker provides message routing utilities
package cracker

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
	"github.com/cbusbey/quickfixgo/tag"

	"github.com/cbusbey/quickfixgo/fix40"
	"github.com/cbusbey/quickfixgo/fix41"
	"github.com/cbusbey/quickfixgo/fix42"
	"github.com/cbusbey/quickfixgo/fix43"
	"github.com/cbusbey/quickfixgo/fix44"
	"github.com/cbusbey/quickfixgo/fix50"
	"github.com/cbusbey/quickfixgo/fixt11"
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

func Crack(msg quickfixgo.Message, sessionID quickfixgo.SessionID, router MessageRouter) quickfixgo.MessageReject {
	beginString := new(field.BeginString)
	msg.Header.Get(beginString)
	return tryCrack(beginString.Value, msg, sessionID, router)
}

func tryCrack(beginString string, msg quickfixgo.Message, sessionID quickfixgo.SessionID, router MessageRouter) quickfixgo.MessageReject {
	switch beginString {
	case quickfixgo.BeginString_FIX40:
		return fix40.Crack(msg, sessionID, router)
	case quickfixgo.BeginString_FIX41:
		return fix41.Crack(msg, sessionID, router)
	case quickfixgo.BeginString_FIX42:
		return fix42.Crack(msg, sessionID, router)
	case quickfixgo.BeginString_FIX43:
		return fix43.Crack(msg, sessionID, router)
	case quickfixgo.BeginString_FIX44:
		return fix44.Crack(msg, sessionID, router)
	case quickfixgo.BeginString_FIX50:
		return fix50.Crack(msg, sessionID, router)
	case quickfixgo.BeginString_FIXT11:

		msgType := new(quickfixgo.StringValue)
		if msg.Header.GetField(tag.MsgType, msgType); quickfixgo.IsAdminMessageType(msgType.Value) {
			return fixt11.Crack(msg, sessionID, router)
		} else {

			applVerId := sessionID.DefaultApplVerID
			applVerIDField := new(field.ApplVerID)
			if err := msg.Header.Get(applVerIDField); err == nil {
				applVerId = applVerIDField.Value
			}

			switch applVerId {
			case quickfixgo.ApplVerID_FIX40:
				beginString = quickfixgo.BeginString_FIX40
			case quickfixgo.ApplVerID_FIX41:
				beginString = quickfixgo.BeginString_FIX41
			case quickfixgo.ApplVerID_FIX42:
				beginString = quickfixgo.BeginString_FIX42
			case quickfixgo.ApplVerID_FIX43:
				beginString = quickfixgo.BeginString_FIX43
			case quickfixgo.ApplVerID_FIX44:
				beginString = quickfixgo.BeginString_FIX44
			case quickfixgo.ApplVerID_FIX50, quickfixgo.ApplVerID_FIX50SP1, quickfixgo.ApplVerID_FIX50SP2:
				beginString = quickfixgo.BeginString_FIX50
			default:
				beginString = ""
			}

			return tryCrack(beginString, msg, sessionID, router)
		}
	}

	panic("Invalid BeginString")
}
