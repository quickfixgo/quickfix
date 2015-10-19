//Package quoterequest msg type = R.
package quoterequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a QuoteRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//QuoteReqID is a required field for QuoteRequest.
func (m Message) QuoteReqID() (*field.QuoteReqIDField, quickfix.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteRequest.
func (m Message) GetQuoteReqID(f *field.QuoteReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for QuoteRequest.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from QuoteRequest.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for QuoteRequest.
func (m Message) SymbolSfx() (*field.SymbolSfxField, quickfix.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from QuoteRequest.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for QuoteRequest.
func (m Message) SecurityID() (*field.SecurityIDField, quickfix.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from QuoteRequest.
func (m Message) GetSecurityID(f *field.SecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for QuoteRequest.
func (m Message) IDSource() (*field.IDSourceField, quickfix.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from QuoteRequest.
func (m Message) GetIDSource(f *field.IDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for QuoteRequest.
func (m Message) Issuer() (*field.IssuerField, quickfix.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from QuoteRequest.
func (m Message) GetIssuer(f *field.IssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for QuoteRequest.
func (m Message) SecurityDesc() (*field.SecurityDescField, quickfix.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from QuoteRequest.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for QuoteRequest.
func (m Message) PrevClosePx() (*field.PrevClosePxField, quickfix.MessageRejectError) {
	f := &field.PrevClosePxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from QuoteRequest.
func (m Message) GetPrevClosePx(f *field.PrevClosePxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for QuoteRequest.
func (m Message) Side() (*field.SideField, quickfix.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from QuoteRequest.
func (m Message) GetSide(f *field.SideField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for QuoteRequest.
func (m Message) OrderQty() (*field.OrderQtyField, quickfix.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from QuoteRequest.
func (m Message) GetOrderQty(f *field.OrderQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds QuoteRequest messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for QuoteRequest.
func Builder(
	quotereqid *field.QuoteReqIDField,
	symbol *field.SymbolField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header.Set(field.NewMsgType("R"))
	builder.Body.Set(quotereqid)
	builder.Body.Set(symbol)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX40, "R", r
}
