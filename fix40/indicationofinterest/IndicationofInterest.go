//Package indicationofinterest msg type = 6.
package indicationofinterest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a IndicationofInterest wrapper for the generic Message type
type Message struct {
	message.Message
}

//IOIid is a required field for IndicationofInterest.
func (m Message) IOIid() (*field.IOIidField, errors.MessageRejectError) {
	f := &field.IOIidField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIid reads a IOIid from IndicationofInterest.
func (m Message) GetIOIid(f *field.IOIidField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOITransType is a required field for IndicationofInterest.
func (m Message) IOITransType() (*field.IOITransTypeField, errors.MessageRejectError) {
	f := &field.IOITransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOITransType reads a IOITransType from IndicationofInterest.
func (m Message) GetIOITransType(f *field.IOITransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIRefID is a non-required field for IndicationofInterest.
func (m Message) IOIRefID() (*field.IOIRefIDField, errors.MessageRejectError) {
	f := &field.IOIRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIRefID reads a IOIRefID from IndicationofInterest.
func (m Message) GetIOIRefID(f *field.IOIRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for IndicationofInterest.
func (m Message) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from IndicationofInterest.
func (m Message) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for IndicationofInterest.
func (m Message) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from IndicationofInterest.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for IndicationofInterest.
func (m Message) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from IndicationofInterest.
func (m Message) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for IndicationofInterest.
func (m Message) IDSource() (*field.IDSourceField, errors.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from IndicationofInterest.
func (m Message) GetIDSource(f *field.IDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for IndicationofInterest.
func (m Message) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from IndicationofInterest.
func (m Message) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for IndicationofInterest.
func (m Message) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from IndicationofInterest.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for IndicationofInterest.
func (m Message) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from IndicationofInterest.
func (m Message) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIShares is a required field for IndicationofInterest.
func (m Message) IOIShares() (*field.IOISharesField, errors.MessageRejectError) {
	f := &field.IOISharesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIShares reads a IOIShares from IndicationofInterest.
func (m Message) GetIOIShares(f *field.IOISharesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for IndicationofInterest.
func (m Message) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from IndicationofInterest.
func (m Message) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for IndicationofInterest.
func (m Message) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from IndicationofInterest.
func (m Message) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValidUntilTime is a non-required field for IndicationofInterest.
func (m Message) ValidUntilTime() (*field.ValidUntilTimeField, errors.MessageRejectError) {
	f := &field.ValidUntilTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValidUntilTime reads a ValidUntilTime from IndicationofInterest.
func (m Message) GetValidUntilTime(f *field.ValidUntilTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIQltyInd is a non-required field for IndicationofInterest.
func (m Message) IOIQltyInd() (*field.IOIQltyIndField, errors.MessageRejectError) {
	f := &field.IOIQltyIndField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIQltyInd reads a IOIQltyInd from IndicationofInterest.
func (m Message) GetIOIQltyInd(f *field.IOIQltyIndField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIOthSvc is a non-required field for IndicationofInterest.
func (m Message) IOIOthSvc() (*field.IOIOthSvcField, errors.MessageRejectError) {
	f := &field.IOIOthSvcField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIOthSvc reads a IOIOthSvc from IndicationofInterest.
func (m Message) GetIOIOthSvc(f *field.IOIOthSvcField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOINaturalFlag is a non-required field for IndicationofInterest.
func (m Message) IOINaturalFlag() (*field.IOINaturalFlagField, errors.MessageRejectError) {
	f := &field.IOINaturalFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOINaturalFlag reads a IOINaturalFlag from IndicationofInterest.
func (m Message) GetIOINaturalFlag(f *field.IOINaturalFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIQualifier is a non-required field for IndicationofInterest.
func (m Message) IOIQualifier() (*field.IOIQualifierField, errors.MessageRejectError) {
	f := &field.IOIQualifierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIQualifier reads a IOIQualifier from IndicationofInterest.
func (m Message) GetIOIQualifier(f *field.IOIQualifierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for IndicationofInterest.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from IndicationofInterest.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds IndicationofInterest messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for IndicationofInterest.
func Builder(
	ioiid *field.IOIidField,
	ioitranstype *field.IOITransTypeField,
	symbol *field.SymbolField,
	side *field.SideField,
	ioishares *field.IOISharesField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header().Set(field.NewMsgType("6"))
	builder.Body().Set(ioiid)
	builder.Body().Set(ioitranstype)
	builder.Body().Set(symbol)
	builder.Body().Set(side)
	builder.Body().Set(ioishares)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX40, "6", r
}
