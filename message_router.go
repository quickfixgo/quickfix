package quickfix

import (
	"github.com/quickfixgo/quickfix/enum"
)

type routeKey struct {
	FIXVersion string
	MsgType    string
}

//A MessageRoute is a function that can process a fromApp/fromAdmin callback
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

	msgType, err := msg.MsgType()
	if err != nil {
		return err
	}

	return c.tryRoute(string(beginString), string(msgType), msg, sessionID)
}

func (c MessageRouter) tryRoute(beginString string, msgType string, msg Message, sessionID SessionID) MessageRejectError {
	fixVersion := beginString
	if beginString == enum.BeginStringFIXT11 && !isAdminMessageType(msgType) {
		var applVerID FIXString
		if err := msg.Header.GetField(tagApplVerID, &applVerID); err != nil {
			session, _ := lookupSession(sessionID)
			applVerID = FIXString(session.TargetDefaultApplicationVersionID())
		}

		switch enum.ApplVerID(applVerID) {
		case enum.ApplVerID_FIX40:
			fixVersion = enum.BeginStringFIX40
		case enum.ApplVerID_FIX41:
			fixVersion = enum.BeginStringFIX41
		case enum.ApplVerID_FIX42:
			fixVersion = enum.BeginStringFIX42
		case enum.ApplVerID_FIX43:
			fixVersion = enum.BeginStringFIX43
		case enum.ApplVerID_FIX44:
			fixVersion = enum.BeginStringFIX44
		default:
			fixVersion = string(applVerID)
		}
	}

	if route, ok := c.routes[routeKey{fixVersion, msgType}]; ok {
		return route(msg, sessionID)
	}

	switch {
	case isAdminMessageType(msgType) || msgType == "j":
		return nil
	}

	return UnsupportedMessageType()
}
