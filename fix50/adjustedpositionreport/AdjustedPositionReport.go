//Package adjustedpositionreport msg type = BL.
package adjustedpositionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/positionqty"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a AdjustedPositionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BL"`
	Header     fixt11.Header
	//PosMaintRptID is a required field for AdjustedPositionReport.
	PosMaintRptID string `fix:"721"`
	//PosReqType is a non-required field for AdjustedPositionReport.
	PosReqType *int `fix:"724"`
	//ClearingBusinessDate is a required field for AdjustedPositionReport.
	ClearingBusinessDate string `fix:"715"`
	//SettlSessID is a non-required field for AdjustedPositionReport.
	SettlSessID *string `fix:"716"`
	//Parties Component
	Parties parties.Component
	//PositionQty Component
	PositionQty positionqty.Component
	//Instrument Component
	Instrument instrument.Component
	//SettlPrice is a non-required field for AdjustedPositionReport.
	SettlPrice *float64 `fix:"730"`
	//PriorSettlPrice is a non-required field for AdjustedPositionReport.
	PriorSettlPrice *float64 `fix:"734"`
	//PosMaintRptRefID is a non-required field for AdjustedPositionReport.
	PosMaintRptRefID *string `fix:"714"`
	Trailer          fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetPosMaintRptID(v string)        { m.PosMaintRptID = v }
func (m *Message) SetPosReqType(v int)              { m.PosReqType = &v }
func (m *Message) SetClearingBusinessDate(v string) { m.ClearingBusinessDate = v }
func (m *Message) SetSettlSessID(v string)          { m.SettlSessID = &v }
func (m *Message) SetSettlPrice(v float64)          { m.SettlPrice = &v }
func (m *Message) SetPriorSettlPrice(v float64)     { m.PriorSettlPrice = &v }
func (m *Message) SetPosMaintRptRefID(v string)     { m.PosMaintRptRefID = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX50, "BL", r
}
