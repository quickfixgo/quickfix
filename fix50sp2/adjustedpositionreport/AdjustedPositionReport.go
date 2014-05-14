//Package adjustedpositionreport msg type = BL.
package adjustedpositionreport

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

//Message is a AdjustedPositionReport wrapper for the generic Message type
type Message struct {
	message.Message
}

//PosMaintRptID is a required field for AdjustedPositionReport.
func (m Message) PosMaintRptID() (*field.PosMaintRptIDField, errors.MessageRejectError) {
	f := &field.PosMaintRptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptID reads a PosMaintRptID from AdjustedPositionReport.
func (m Message) GetPosMaintRptID(f *field.PosMaintRptIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqType is a non-required field for AdjustedPositionReport.
func (m Message) PosReqType() (*field.PosReqTypeField, errors.MessageRejectError) {
	f := &field.PosReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqType reads a PosReqType from AdjustedPositionReport.
func (m Message) GetPosReqType(f *field.PosReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a required field for AdjustedPositionReport.
func (m Message) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from AdjustedPositionReport.
func (m Message) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for AdjustedPositionReport.
func (m Message) SettlSessID() (*field.SettlSessIDField, errors.MessageRejectError) {
	f := &field.SettlSessIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from AdjustedPositionReport.
func (m Message) GetSettlSessID(f *field.SettlSessIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AdjustedPositionReport.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AdjustedPositionReport.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPositions is a non-required field for AdjustedPositionReport.
func (m Message) NoPositions() (*field.NoPositionsField, errors.MessageRejectError) {
	f := &field.NoPositionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPositions reads a NoPositions from AdjustedPositionReport.
func (m Message) GetNoPositions(f *field.NoPositionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for AdjustedPositionReport.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from AdjustedPositionReport.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlPrice is a non-required field for AdjustedPositionReport.
func (m Message) SettlPrice() (*field.SettlPriceField, errors.MessageRejectError) {
	f := &field.SettlPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlPrice reads a SettlPrice from AdjustedPositionReport.
func (m Message) GetSettlPrice(f *field.SettlPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriorSettlPrice is a non-required field for AdjustedPositionReport.
func (m Message) PriorSettlPrice() (*field.PriorSettlPriceField, errors.MessageRejectError) {
	f := &field.PriorSettlPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriorSettlPrice reads a PriorSettlPrice from AdjustedPositionReport.
func (m Message) GetPriorSettlPrice(f *field.PriorSettlPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosMaintRptRefID is a non-required field for AdjustedPositionReport.
func (m Message) PosMaintRptRefID() (*field.PosMaintRptRefIDField, errors.MessageRejectError) {
	f := &field.PosMaintRptRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptRefID reads a PosMaintRptRefID from AdjustedPositionReport.
func (m Message) GetPosMaintRptRefID(f *field.PosMaintRptRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds AdjustedPositionReport messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for AdjustedPositionReport.
func Builder(
	posmaintrptid *field.PosMaintRptIDField,
	clearingbusinessdate *field.ClearingBusinessDateField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("BL"))
	builder.Body().Set(posmaintrptid)
	builder.Body().Set(clearingbusinessdate)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "BL", r
}
