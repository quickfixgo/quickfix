//Package adjustedpositionreport msg type = BL.
package adjustedpositionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a AdjustedPositionReport wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//PosMaintRptID is a required field for AdjustedPositionReport.
func (m Message) PosMaintRptID() (*field.PosMaintRptIDField, quickfix.MessageRejectError) {
	f := &field.PosMaintRptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptID reads a PosMaintRptID from AdjustedPositionReport.
func (m Message) GetPosMaintRptID(f *field.PosMaintRptIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqType is a non-required field for AdjustedPositionReport.
func (m Message) PosReqType() (*field.PosReqTypeField, quickfix.MessageRejectError) {
	f := &field.PosReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqType reads a PosReqType from AdjustedPositionReport.
func (m Message) GetPosReqType(f *field.PosReqTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a required field for AdjustedPositionReport.
func (m Message) ClearingBusinessDate() (*field.ClearingBusinessDateField, quickfix.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from AdjustedPositionReport.
func (m Message) GetClearingBusinessDate(f *field.ClearingBusinessDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for AdjustedPositionReport.
func (m Message) SettlSessID() (*field.SettlSessIDField, quickfix.MessageRejectError) {
	f := &field.SettlSessIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from AdjustedPositionReport.
func (m Message) GetSettlSessID(f *field.SettlSessIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AdjustedPositionReport.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AdjustedPositionReport.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPositions is a non-required field for AdjustedPositionReport.
func (m Message) NoPositions() (*field.NoPositionsField, quickfix.MessageRejectError) {
	f := &field.NoPositionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPositions reads a NoPositions from AdjustedPositionReport.
func (m Message) GetNoPositions(f *field.NoPositionsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for AdjustedPositionReport.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from AdjustedPositionReport.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlPrice is a non-required field for AdjustedPositionReport.
func (m Message) SettlPrice() (*field.SettlPriceField, quickfix.MessageRejectError) {
	f := &field.SettlPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlPrice reads a SettlPrice from AdjustedPositionReport.
func (m Message) GetSettlPrice(f *field.SettlPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriorSettlPrice is a non-required field for AdjustedPositionReport.
func (m Message) PriorSettlPrice() (*field.PriorSettlPriceField, quickfix.MessageRejectError) {
	f := &field.PriorSettlPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriorSettlPrice reads a PriorSettlPrice from AdjustedPositionReport.
func (m Message) GetPriorSettlPrice(f *field.PriorSettlPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PosMaintRptRefID is a non-required field for AdjustedPositionReport.
func (m Message) PosMaintRptRefID() (*field.PosMaintRptRefIDField, quickfix.MessageRejectError) {
	f := &field.PosMaintRptRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptRefID reads a PosMaintRptRefID from AdjustedPositionReport.
func (m Message) GetPosMaintRptRefID(f *field.PosMaintRptRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds AdjustedPositionReport messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for AdjustedPositionReport.
func Builder(
	posmaintrptid *field.PosMaintRptIDField,
	clearingbusinessdate *field.ClearingBusinessDateField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("BL"))
	builder.Body().Set(posmaintrptid)
	builder.Body().Set(clearingbusinessdate)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "BL", r
}
