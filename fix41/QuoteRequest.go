package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteRequest msg type = R.
type QuoteRequest struct {
	message.Message
}

//QuoteRequestBuilder builds QuoteRequest messages.
type QuoteRequestBuilder struct {
	message.MessageBuilder
}

//CreateQuoteRequestBuilder returns an initialized QuoteRequestBuilder with specified required fields.
func CreateQuoteRequestBuilder(
	quotereqid field.QuoteReqID,
	symbol field.Symbol) QuoteRequestBuilder {
	var builder QuoteRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("R"))
	builder.Body.Set(quotereqid)
	builder.Body.Set(symbol)
	return builder
}

//QuoteReqID is a required field for QuoteRequest.
func (m QuoteRequest) QuoteReqID() (*field.QuoteReqID, errors.MessageRejectError) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteRequest.
func (m QuoteRequest) GetQuoteReqID(f *field.QuoteReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for QuoteRequest.
func (m QuoteRequest) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from QuoteRequest.
func (m QuoteRequest) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for QuoteRequest.
func (m QuoteRequest) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from QuoteRequest.
func (m QuoteRequest) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for QuoteRequest.
func (m QuoteRequest) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from QuoteRequest.
func (m QuoteRequest) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for QuoteRequest.
func (m QuoteRequest) IDSource() (*field.IDSource, errors.MessageRejectError) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from QuoteRequest.
func (m QuoteRequest) GetIDSource(f *field.IDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for QuoteRequest.
func (m QuoteRequest) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from QuoteRequest.
func (m QuoteRequest) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for QuoteRequest.
func (m QuoteRequest) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from QuoteRequest.
func (m QuoteRequest) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for QuoteRequest.
func (m QuoteRequest) MaturityDay() (*field.MaturityDay, errors.MessageRejectError) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from QuoteRequest.
func (m QuoteRequest) GetMaturityDay(f *field.MaturityDay) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for QuoteRequest.
func (m QuoteRequest) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from QuoteRequest.
func (m QuoteRequest) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for QuoteRequest.
func (m QuoteRequest) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from QuoteRequest.
func (m QuoteRequest) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for QuoteRequest.
func (m QuoteRequest) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from QuoteRequest.
func (m QuoteRequest) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for QuoteRequest.
func (m QuoteRequest) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from QuoteRequest.
func (m QuoteRequest) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for QuoteRequest.
func (m QuoteRequest) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from QuoteRequest.
func (m QuoteRequest) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for QuoteRequest.
func (m QuoteRequest) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from QuoteRequest.
func (m QuoteRequest) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for QuoteRequest.
func (m QuoteRequest) PrevClosePx() (*field.PrevClosePx, errors.MessageRejectError) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from QuoteRequest.
func (m QuoteRequest) GetPrevClosePx(f *field.PrevClosePx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for QuoteRequest.
func (m QuoteRequest) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from QuoteRequest.
func (m QuoteRequest) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for QuoteRequest.
func (m QuoteRequest) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from QuoteRequest.
func (m QuoteRequest) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for QuoteRequest.
func (m QuoteRequest) FutSettDate() (*field.FutSettDate, errors.MessageRejectError) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from QuoteRequest.
func (m QuoteRequest) GetFutSettDate(f *field.FutSettDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a non-required field for QuoteRequest.
func (m QuoteRequest) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from QuoteRequest.
func (m QuoteRequest) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate2 is a non-required field for QuoteRequest.
func (m QuoteRequest) FutSettDate2() (*field.FutSettDate2, errors.MessageRejectError) {
	f := new(field.FutSettDate2)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate2 reads a FutSettDate2 from QuoteRequest.
func (m QuoteRequest) GetFutSettDate2(f *field.FutSettDate2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for QuoteRequest.
func (m QuoteRequest) OrderQty2() (*field.OrderQty2, errors.MessageRejectError) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from QuoteRequest.
func (m QuoteRequest) GetOrderQty2(f *field.OrderQty2) errors.MessageRejectError {
	return m.Body.Get(f)
}
