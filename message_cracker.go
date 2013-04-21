package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/tag"

	"github.com/cbusbey/quickfixgo/fix40"
	"github.com/cbusbey/quickfixgo/fix41"
	"github.com/cbusbey/quickfixgo/fix42"
	"github.com/cbusbey/quickfixgo/fix43"
	"github.com/cbusbey/quickfixgo/fix44"
	"github.com/cbusbey/quickfixgo/fix50"
	"github.com/cbusbey/quickfixgo/fixt11"

	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session"
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

func Crack(msg message.Message, sessionID session.ID, router MessageRouter) reject.MessageReject {
	BeginString, _ := msg.Header.StringValue(tag.BeginString)

	return tryCrack(BeginString, msg, sessionID, router)
}

func tryCrack(beginString string, msg message.Message, sessionID session.ID, router MessageRouter) reject.MessageReject {
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

		if msgType, _ := msg.Header.StringValue(tag.MsgType); session.IsAdminMessageType(msgType) {
			return fixt11.Crack(msg, sessionID, router)
		} else {

			applVerId := sessionID.DefaultApplVerID
			if ApplVerIdField, ok := msg.Header.StringValue(tag.ApplVerID); ok {
				applVerId = ApplVerIdField
			}

			switch applVerId {
			case fix.ApplVerID_FIX40:
				beginString = fix.BeginString_FIX40
			case fix.ApplVerID_FIX41:
				beginString = fix.BeginString_FIX41
			case fix.ApplVerID_FIX42:
				beginString = fix.BeginString_FIX42
			case fix.ApplVerID_FIX43:
				beginString = fix.BeginString_FIX43
			case fix.ApplVerID_FIX44:
				beginString = fix.BeginString_FIX44
			case fix.ApplVerID_FIX50, fix.ApplVerID_FIX50SP1, fix.ApplVerID_FIX50SP2:
				beginString = fix.BeginString_FIX50
			default:
				beginString = ""
			}

			return tryCrack(beginString, msg, sessionID, router)
		}
	}

	panic("Invalid BeginString")
}
