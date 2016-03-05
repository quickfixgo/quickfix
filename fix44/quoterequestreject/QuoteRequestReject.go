//Package quoterequestreject msg type = AG.
package quoterequestreject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/legstipulations"
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
	"github.com/quickfixgo/quickfix/fix44/orderqtydata"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/stipulations"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fix44/yielddata"
	"time"
)

//NoRelatedSym is a repeating group in QuoteRequestReject
type NoRelatedSym struct {
	//Instrument Component
	instrument.Instrument
	//FinancingDetails Component
	financingdetails.FinancingDetails
	//NoUnderlyings is a non-required field for NoRelatedSym.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//PrevClosePx is a non-required field for NoRelatedSym.
	PrevClosePx *float64 `fix:"140"`
	//QuoteRequestType is a non-required field for NoRelatedSym.
	QuoteRequestType *int `fix:"303"`
	//QuoteType is a non-required field for NoRelatedSym.
	QuoteType *int `fix:"537"`
	//TradingSessionID is a non-required field for NoRelatedSym.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoRelatedSym.
	TradingSessionSubID *string `fix:"625"`
	//TradeOriginationDate is a non-required field for NoRelatedSym.
	TradeOriginationDate *string `fix:"229"`
	//Side is a non-required field for NoRelatedSym.
	Side *string `fix:"54"`
	//QtyType is a non-required field for NoRelatedSym.
	QtyType *int `fix:"854"`
	//OrderQtyData Component
	orderqtydata.OrderQtyData
	//SettlType is a non-required field for NoRelatedSym.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NoRelatedSym.
	SettlDate *string `fix:"64"`
	//SettlDate2 is a non-required field for NoRelatedSym.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoRelatedSym.
	OrderQty2 *float64 `fix:"192"`
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//Stipulations Component
	stipulations.Stipulations
	//Account is a non-required field for NoRelatedSym.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NoRelatedSym.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NoRelatedSym.
	AccountType *int `fix:"581"`
	//NoLegs is a non-required field for NoRelatedSym.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoQuoteQualifiers is a non-required field for NoRelatedSym.
	NoQuoteQualifiers []NoQuoteQualifiers `fix:"735,omitempty"`
	//QuotePriceType is a non-required field for NoRelatedSym.
	QuotePriceType *int `fix:"692"`
	//OrdType is a non-required field for NoRelatedSym.
	OrdType *string `fix:"40"`
	//ExpireTime is a non-required field for NoRelatedSym.
	ExpireTime *time.Time `fix:"126"`
	//TransactTime is a non-required field for NoRelatedSym.
	TransactTime *time.Time `fix:"60"`
	//SpreadOrBenchmarkCurveData Component
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//PriceType is a non-required field for NoRelatedSym.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NoRelatedSym.
	Price *float64 `fix:"44"`
	//Price2 is a non-required field for NoRelatedSym.
	Price2 *float64 `fix:"640"`
	//YieldData Component
	yielddata.YieldData
	//Parties Component
	parties.Parties
}

func (m *NoRelatedSym) SetNoUnderlyings(v []NoUnderlyings)         { m.NoUnderlyings = v }
func (m *NoRelatedSym) SetPrevClosePx(v float64)                   { m.PrevClosePx = &v }
func (m *NoRelatedSym) SetQuoteRequestType(v int)                  { m.QuoteRequestType = &v }
func (m *NoRelatedSym) SetQuoteType(v int)                         { m.QuoteType = &v }
func (m *NoRelatedSym) SetTradingSessionID(v string)               { m.TradingSessionID = &v }
func (m *NoRelatedSym) SetTradingSessionSubID(v string)            { m.TradingSessionSubID = &v }
func (m *NoRelatedSym) SetTradeOriginationDate(v string)           { m.TradeOriginationDate = &v }
func (m *NoRelatedSym) SetSide(v string)                           { m.Side = &v }
func (m *NoRelatedSym) SetQtyType(v int)                           { m.QtyType = &v }
func (m *NoRelatedSym) SetSettlType(v string)                      { m.SettlType = &v }
func (m *NoRelatedSym) SetSettlDate(v string)                      { m.SettlDate = &v }
func (m *NoRelatedSym) SetSettlDate2(v string)                     { m.SettlDate2 = &v }
func (m *NoRelatedSym) SetOrderQty2(v float64)                     { m.OrderQty2 = &v }
func (m *NoRelatedSym) SetCurrency(v string)                       { m.Currency = &v }
func (m *NoRelatedSym) SetAccount(v string)                        { m.Account = &v }
func (m *NoRelatedSym) SetAcctIDSource(v int)                      { m.AcctIDSource = &v }
func (m *NoRelatedSym) SetAccountType(v int)                       { m.AccountType = &v }
func (m *NoRelatedSym) SetNoLegs(v []NoLegs)                       { m.NoLegs = v }
func (m *NoRelatedSym) SetNoQuoteQualifiers(v []NoQuoteQualifiers) { m.NoQuoteQualifiers = v }
func (m *NoRelatedSym) SetQuotePriceType(v int)                    { m.QuotePriceType = &v }
func (m *NoRelatedSym) SetOrdType(v string)                        { m.OrdType = &v }
func (m *NoRelatedSym) SetExpireTime(v time.Time)                  { m.ExpireTime = &v }
func (m *NoRelatedSym) SetTransactTime(v time.Time)                { m.TransactTime = &v }
func (m *NoRelatedSym) SetPriceType(v int)                         { m.PriceType = &v }
func (m *NoRelatedSym) SetPrice(v float64)                         { m.Price = &v }
func (m *NoRelatedSym) SetPrice2(v float64)                        { m.Price2 = &v }

//NoUnderlyings is a repeating group in NoRelatedSym
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
}

//NoLegs is a repeating group in NoRelatedSym
type NoLegs struct {
	//InstrumentLeg Component
	instrumentleg.InstrumentLeg
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
	//LegStipulations Component
	legstipulations.LegStipulations
	//NestedParties Component
	nestedparties.NestedParties
	//LegBenchmarkCurveData Component
	legbenchmarkcurvedata.LegBenchmarkCurveData
}

func (m *NoLegs) SetLegQty(v float64)      { m.LegQty = &v }
func (m *NoLegs) SetLegSwapType(v int)     { m.LegSwapType = &v }
func (m *NoLegs) SetLegSettlType(v string) { m.LegSettlType = &v }
func (m *NoLegs) SetLegSettlDate(v string) { m.LegSettlDate = &v }

//NoQuoteQualifiers is a repeating group in NoRelatedSym
type NoQuoteQualifiers struct {
	//QuoteQualifier is a non-required field for NoQuoteQualifiers.
	QuoteQualifier *string `fix:"695"`
}

func (m *NoQuoteQualifiers) SetQuoteQualifier(v string) { m.QuoteQualifier = &v }

//Message is a QuoteRequestReject FIX Message
type Message struct {
	FIXMsgType string `fix:"AG"`
	fix44.Header
	//QuoteReqID is a required field for QuoteRequestReject.
	QuoteReqID string `fix:"131"`
	//RFQReqID is a non-required field for QuoteRequestReject.
	RFQReqID *string `fix:"644"`
	//QuoteRequestRejectReason is a required field for QuoteRequestReject.
	QuoteRequestRejectReason int `fix:"658"`
	//NoRelatedSym is a required field for QuoteRequestReject.
	NoRelatedSym []NoRelatedSym `fix:"146"`
	//Text is a non-required field for QuoteRequestReject.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for QuoteRequestReject.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for QuoteRequestReject.
	EncodedText *string `fix:"355"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteReqID(v string)            { m.QuoteReqID = v }
func (m *Message) SetRFQReqID(v string)              { m.RFQReqID = &v }
func (m *Message) SetQuoteRequestRejectReason(v int) { m.QuoteRequestRejectReason = v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym)  { m.NoRelatedSym = v }
func (m *Message) SetText(v string)                  { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)           { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)           { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "AG", r
}
