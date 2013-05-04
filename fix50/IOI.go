package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type IOI struct {
	quickfixgo.Message
}

func (m *IOI) IOIID() (*field.IOIID, error) {
	f := new(field.IOIID)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) IOITransType() (*field.IOITransType, error) {
	f := new(field.IOITransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) IOIRefID() (*field.IOIRefID, error) {
	f := new(field.IOIRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) IOIQty() (*field.IOIQty, error) {
	f := new(field.IOIQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) ValidUntilTime() (*field.ValidUntilTime, error) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) IOIQltyInd() (*field.IOIQltyInd, error) {
	f := new(field.IOIQltyInd)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) IOINaturalFlag() (*field.IOINaturalFlag, error) {
	f := new(field.IOINaturalFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *IOI) URLLink() (*field.URLLink, error) {
	f := new(field.URLLink)
	err := m.Body.Get(f)
	return f, err
}
