//Package advertisement msg type = 7.
package advertisement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a Advertisement wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//AdvId is a required field for Advertisement.
func (m Message) AdvId() (*field.AdvIdField, quickfix.MessageRejectError) {
	f := &field.AdvIdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAdvId reads a AdvId from Advertisement.
func (m Message) GetAdvId(f *field.AdvIdField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AdvTransType is a required field for Advertisement.
func (m Message) AdvTransType() (*field.AdvTransTypeField, quickfix.MessageRejectError) {
	f := &field.AdvTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAdvTransType reads a AdvTransType from Advertisement.
func (m Message) GetAdvTransType(f *field.AdvTransTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AdvRefID is a non-required field for Advertisement.
func (m Message) AdvRefID() (*field.AdvRefIDField, quickfix.MessageRejectError) {
	f := &field.AdvRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAdvRefID reads a AdvRefID from Advertisement.
func (m Message) GetAdvRefID(f *field.AdvRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for Advertisement.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from Advertisement.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for Advertisement.
func (m Message) SymbolSfx() (*field.SymbolSfxField, quickfix.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from Advertisement.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for Advertisement.
func (m Message) SecurityID() (*field.SecurityIDField, quickfix.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from Advertisement.
func (m Message) GetSecurityID(f *field.SecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for Advertisement.
func (m Message) IDSource() (*field.IDSourceField, quickfix.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from Advertisement.
func (m Message) GetIDSource(f *field.IDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for Advertisement.
func (m Message) Issuer() (*field.IssuerField, quickfix.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from Advertisement.
func (m Message) GetIssuer(f *field.IssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for Advertisement.
func (m Message) SecurityDesc() (*field.SecurityDescField, quickfix.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from Advertisement.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AdvSide is a required field for Advertisement.
func (m Message) AdvSide() (*field.AdvSideField, quickfix.MessageRejectError) {
	f := &field.AdvSideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAdvSide reads a AdvSide from Advertisement.
func (m Message) GetAdvSide(f *field.AdvSideField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Shares is a required field for Advertisement.
func (m Message) Shares() (*field.SharesField, quickfix.MessageRejectError) {
	f := &field.SharesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetShares reads a Shares from Advertisement.
func (m Message) GetShares(f *field.SharesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for Advertisement.
func (m Message) Price() (*field.PriceField, quickfix.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from Advertisement.
func (m Message) GetPrice(f *field.PriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for Advertisement.
func (m Message) Currency() (*field.CurrencyField, quickfix.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from Advertisement.
func (m Message) GetCurrency(f *field.CurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for Advertisement.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from Advertisement.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for Advertisement.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Advertisement.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for Advertisement.
func New(
	advid *field.AdvIdField,
	advtranstype *field.AdvTransTypeField,
	symbol *field.SymbolField,
	advside *field.AdvSideField,
	shares *field.SharesField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX40))
	builder.Header.Set(field.NewMsgType("7"))
	builder.Body.Set(advid)
	builder.Body.Set(advtranstype)
	builder.Body.Set(symbol)
	builder.Body.Set(advside)
	builder.Body.Set(shares)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX40, "7", r
}
