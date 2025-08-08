// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

type routeKey struct {
	FIXVersion string
	MsgType    string
}

// FIX ApplVerID string values.
const (
	ApplVerIDFIX27    = "0"
	ApplVerIDFIX30    = "1"
	ApplVerIDFIX40    = "2"
	ApplVerIDFIX41    = "3"
	ApplVerIDFIX42    = "4"
	ApplVerIDFIX43    = "5"
	ApplVerIDFIX44    = "6"
	ApplVerIDFIX50    = "7"
	ApplVerIDFIX50SP1 = "8"
	ApplVerIDFIX50SP2 = "9"
)

// A MessageRoute is a function that can process a fromApp/fromAdmin callback.
type MessageRoute func(msg *Message, sessionID SessionID) MessageRejectError

// A MessageRouter is a mutex for MessageRoutes.
type MessageRouter struct {
	routes map[routeKey]MessageRoute
	*Registry
}
type RouterOption func(*MessageRouter)

func WithMessageRouterRegistry(registry *Registry) RouterOption {
	return func(r *MessageRouter) {
		r.Registry = registry
	}
}

// NewMessageRouter returns an initialized MessageRouter instance.
func NewMessageRouter(opts ...RouterOption) *MessageRouter {
	router := MessageRouter{routes: make(map[routeKey]MessageRoute)}
	router.Registry = defaultRegistry
	for _, opt := range opts {
		opt(&router)
	}
	return &router
}

// AddRoute adds a route to the MessageRouter instance keyed to begin string and msgType.
func (c MessageRouter) AddRoute(beginString string, msgType string, router MessageRoute) {
	c.routes[routeKey{beginString, msgType}] = router
}

// Route may be called from the fromApp/fromAdmin callbacks. Messages that cannot be routed will be rejected with UnsupportedMessageType.
func (c MessageRouter) Route(msg *Message, sessionID SessionID) MessageRejectError {
	beginString, err := msg.Header.GetBytes(tagBeginString)
	if err != nil {
		return nil
	}

	msgType, err := msg.Header.GetBytes(tagMsgType)
	if err != nil {
		return err
	}

	return c.tryRoute(string(beginString), string(msgType), msg, sessionID)
}

func (c MessageRouter) tryRoute(beginString string, msgType string, msg *Message, sessionID SessionID) MessageRejectError {
	fixVersion := beginString
	isAdminMsg := isAdminMessageType([]byte(msgType))

	if beginString == BeginStringFIXT11 && !isAdminMsg {
		var applVerID FIXString
		if err := msg.Header.GetField(tagApplVerID, &applVerID); err != nil {
			session, _ := c.lookupSession(sessionID)
			applVerID = FIXString(session.TargetDefaultApplicationVersionID())
		}

		switch applVerID {
		case ApplVerIDFIX40:
			fixVersion = BeginStringFIX40
		case ApplVerIDFIX41:
			fixVersion = BeginStringFIX41
		case ApplVerIDFIX42:
			fixVersion = BeginStringFIX42
		case ApplVerIDFIX43:
			fixVersion = BeginStringFIX43
		case ApplVerIDFIX44:
			fixVersion = BeginStringFIX44
		default:
			fixVersion = string(applVerID)
		}
	}

	if route, ok := c.routes[routeKey{fixVersion, msgType}]; ok {
		return route(msg, sessionID)
	}

	switch {
	case isAdminMsg || msgType == "j":
		return nil
	}

	return UnsupportedMessageType()
}
