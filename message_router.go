package quickfix

import (
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/enum"
	"github.com/quickfixgo/quickfix/fix/field"
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
	beginString := &field.BeginStringField{}
	err := msg.Header.Get(beginString)

	if err != nil {
		return err
	}

	msgType := &field.MsgTypeField{}
	err = msg.Header.Get(msgType)

	if err != nil {
		return err
	}

	return c.tryRoute(beginString, msgType, msg, sessionID)
}

func (c MessageRouter) tryRoute(beginString *field.BeginStringField, msgType *field.MsgTypeField, msg Message, sessionID SessionID) MessageRejectError {

	if beginString.Value == fix.BeginString_FIXT11 && !fix.IsAdminMessageType(msgType.Value) {
		applVerID := &field.ApplVerIDField{}
		if err := msg.Header.Get(applVerID); err != nil {
			session, _ := LookupSession(sessionID)
			applVerID.Value = session.TargetDefaultApplicationVersionID()
		}

		switch applVerID.Value {
		case enum.ApplVerID_FIX40:
			beginString.Value = fix.BeginString_FIX40
		case enum.ApplVerID_FIX41:
			beginString.Value = fix.BeginString_FIX41
		case enum.ApplVerID_FIX42:
			beginString.Value = fix.BeginString_FIX42
		case enum.ApplVerID_FIX43:
			beginString.Value = fix.BeginString_FIX43
		case enum.ApplVerID_FIX44:
			beginString.Value = fix.BeginString_FIX44
		case enum.ApplVerID_FIX50, enum.ApplVerID_FIX50SP1, enum.ApplVerID_FIX50SP2:
			beginString.Value = fix.BeginString_FIX50
		}

		return c.tryRoute(beginString, msgType, msg, sessionID)
	}

	route, ok := c.routes[routeKey{beginString.Value, msgType.Value}]

	if !ok {
		return UnsupportedMessageType()
	}

	return route(msg, sessionID)
}

var defaultRouter *MessageRouter

func init() {
	defaultRouter = NewMessageRouter()
}

//AddRoute assigns a message route to the default MessageRouter.
func AddRoute(beginString string, msgType string, route MessageRoute) {
	defaultRouter.AddRoute(beginString, msgType, route)
}

//Route may be called from the fromApp/fromAdmin callbacks to route to the default MessageRouter instance.
func Route(msg Message, sessionID SessionID) MessageRejectError {
	return defaultRouter.Route(msg, sessionID)
}
