//Package marketdatarequestreject msg type = Y.
package marketdatarequestreject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a MarketDataRequestReject wrapper for the generic Message type
type Message struct {
	message.Message
}

//MDReqID is a required field for MarketDataRequestReject.
func (m Message) MDReqID() (*field.MDReqIDField, errors.MessageRejectError) {
	f := &field.MDReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataRequestReject.
func (m Message) GetMDReqID(f *field.MDReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDReqRejReason is a non-required field for MarketDataRequestReject.
func (m Message) MDReqRejReason() (*field.MDReqRejReasonField, errors.MessageRejectError) {
	f := &field.MDReqRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqRejReason reads a MDReqRejReason from MarketDataRequestReject.
func (m Message) GetMDReqRejReason(f *field.MDReqRejReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAltMDSource is a non-required field for MarketDataRequestReject.
func (m Message) NoAltMDSource() (*field.NoAltMDSourceField, errors.MessageRejectError) {
	f := &field.NoAltMDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAltMDSource reads a NoAltMDSource from MarketDataRequestReject.
func (m Message) GetNoAltMDSource(f *field.NoAltMDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MarketDataRequestReject.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MarketDataRequestReject.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MarketDataRequestReject.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MarketDataRequestReject.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MarketDataRequestReject.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MarketDataRequestReject.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MarketDataRequestReject.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MarketDataRequestReject.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds MarketDataRequestReject messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for MarketDataRequestReject.
func Builder(
	mdreqid *field.MDReqIDField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("Y"))
	builder.Body().Set(mdreqid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "Y", r
}
