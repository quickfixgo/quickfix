package fix40

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Advertisement msg type = 7.
type Advertisement struct {
	message.Message
}

//AdvertisementBuilder builds Advertisement messages.
type AdvertisementBuilder struct {
	message.MessageBuilder
}

//NewAdvertisementBuilder returns an initialized AdvertisementBuilder with specified required fields.
func NewAdvertisementBuilder(
	advid field.AdvId,
	advtranstype field.AdvTransType,
	symbol field.Symbol,
	advside field.AdvSide,
	shares field.Shares) *AdvertisementBuilder {
	builder := new(AdvertisementBuilder)
	builder.Body.Set(advid)
	builder.Body.Set(advtranstype)
	builder.Body.Set(symbol)
	builder.Body.Set(advside)
	builder.Body.Set(shares)
	return builder
}

//AdvId is a required field for Advertisement.
func (m *Advertisement) AdvId() (*field.AdvId, error) {
	f := new(field.AdvId)
	err := m.Body.Get(f)
	return f, err
}

//AdvTransType is a required field for Advertisement.
func (m *Advertisement) AdvTransType() (*field.AdvTransType, error) {
	f := new(field.AdvTransType)
	err := m.Body.Get(f)
	return f, err
}

//AdvRefID is a non-required field for Advertisement.
func (m *Advertisement) AdvRefID() (*field.AdvRefID, error) {
	f := new(field.AdvRefID)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a required field for Advertisement.
func (m *Advertisement) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for Advertisement.
func (m *Advertisement) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for Advertisement.
func (m *Advertisement) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//IDSource is a non-required field for Advertisement.
func (m *Advertisement) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for Advertisement.
func (m *Advertisement) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for Advertisement.
func (m *Advertisement) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//AdvSide is a required field for Advertisement.
func (m *Advertisement) AdvSide() (*field.AdvSide, error) {
	f := new(field.AdvSide)
	err := m.Body.Get(f)
	return f, err
}

//Shares is a required field for Advertisement.
func (m *Advertisement) Shares() (*field.Shares, error) {
	f := new(field.Shares)
	err := m.Body.Get(f)
	return f, err
}

//Price is a non-required field for Advertisement.
func (m *Advertisement) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for Advertisement.
func (m *Advertisement) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for Advertisement.
func (m *Advertisement) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for Advertisement.
func (m *Advertisement) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
