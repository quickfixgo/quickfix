package quickfix

import (
	"github.com/quickfixgo/quickfix/enum"
)

type routeKey struct {
	BeginString string
	MsgType     string
}

//A MessageRoute is a function can process a fromApp/fromAdmin callback
type MessageRoute func(msg Message, sessionID SessionID) MessageRejectError

//A MessageRouter is a mutex for MessageRoutes
type MessageRouter struct {
	routes map[routeKey]MessageRoute
}

//NewMessageRouter returns an initialized MessageRouter instance
func NewMessageRouter() *MessageRouter {
	return &MessageRouter{routes: make(map[routeKey]MessageRoute)}
}

//AddRoute adds a route to the MessageRouter instance keyed to begin string and msgType.
func (c MessageRouter) AddRoute(beginString string, msgType string, router MessageRoute) {
	c.routes[routeKey{beginString, msgType}] = router
}

//Route may be called from the fromApp/fromAdmin callbacks. Messages that cannot be routed will be rejected with UnsupportedMessageType.
func (c MessageRouter) Route(msg Message, sessionID SessionID) MessageRejectError {
	var beginString FIXString
	if err := msg.Header.GetField(tagBeginString, &beginString); err != nil {
		return nil
	}

	var msgType FIXString
	if err := msg.Header.GetField(tagMsgType, &msgType); err != nil {
		return err
	}

	return c.tryRoute(beginString, msgType, msg, sessionID)
}

func (c MessageRouter) tryRoute(beginString FIXString, msgType FIXString, msg Message, sessionID SessionID) MessageRejectError {

	if string(beginString) == enum.BeginStringFIXT11 && !isAdminMessageType(string(msgType)) {
		var applVerID FIXString
		if err := msg.Header.GetField(tagApplVerID, &applVerID); err != nil {
			session, _ := LookupSession(sessionID)
			applVerID = FIXString(session.TargetDefaultApplicationVersionID())
		}

		switch string(applVerID) {
		case enum.ApplVerID_FIX40:
			beginString = enum.BeginStringFIX40
		case enum.ApplVerID_FIX41:
			beginString = enum.BeginStringFIX41
		case enum.ApplVerID_FIX42:
			beginString = enum.BeginStringFIX42
		case enum.ApplVerID_FIX43:
			beginString = enum.BeginStringFIX43
		case enum.ApplVerID_FIX44:
			beginString = enum.BeginStringFIX44
		case enum.ApplVerID_FIX50, enum.ApplVerID_FIX50SP1, enum.ApplVerID_FIX50SP2:
			beginString = enum.BeginStringFIX50
		}

		return c.tryRoute(beginString, msgType, msg, sessionID)
	}

	route, ok := c.routes[routeKey{string(beginString), string(msgType)}]

	if !ok {
		return unsupportedMessageType()
	}

	return route(msg, sessionID)
}
