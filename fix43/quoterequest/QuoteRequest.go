//Package quoterequest msg type = R.
package quoterequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix43/stipulations"
	"github.com/quickfixgo/quickfix/fix43/yielddata"
	"time"
)

//NoRelatedSym is a repeating group in QuoteRequest
type NoRelatedSym struct {
	//Instrument is a required component for NoRelatedSym.
	instrument.Instrument
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
	//Stipulations is a non-required component for NoRelatedSym.
	Stipulations *stipulations.Stipulations
	//Side is a non-required field for NoRelatedSym.
	Side *string `fix:"54"`
	//QuantityType is a non-required field for NoRelatedSym.
	QuantityType *int `fix:"465"`
	//OrderQty is a non-required field for NoRelatedSym.
	OrderQty *float64 `fix:"38"`
	//CashOrderQty is a non-required field for NoRelatedSym.
	CashOrderQty *float64 `fix:"152"`
	//SettlmntTyp is a non-required field for NoRelatedSym.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for NoRelatedSym.
	FutSettDate *string `fix:"64"`
	//OrdType is a non-required field for NoRelatedSym.
	OrdType *string `fix:"40"`
	//FutSettDate2 is a non-required field for NoRelatedSym.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoRelatedSym.
	OrderQty2 *float64 `fix:"192"`
	//ExpireTime is a non-required field for NoRelatedSym.
	ExpireTime *time.Time `fix:"126"`
	//TransactTime is a non-required field for NoRelatedSym.
	TransactTime *time.Time `fix:"60"`
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//SpreadOrBenchmarkCurveData is a non-required component for NoRelatedSym.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//PriceType is a non-required field for NoRelatedSym.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NoRelatedSym.
	Price *float64 `fix:"44"`
	//Price2 is a non-required field for NoRelatedSym.
	Price2 *float64 `fix:"640"`
	//YieldData is a non-required component for NoRelatedSym.
	YieldData *yielddata.YieldData
}

//NewNoRelatedSym returns an initialized NoRelatedSym instance
func NewNoRelatedSym(instrument instrument.Instrument) *NoRelatedSym {
	var m NoRelatedSym
	m.SetInstrument(instrument)
	return &m
}

func (m *NoRelatedSym) SetInstrument(v instrument.Instrument)       { m.Instrument = v }
func (m *NoRelatedSym) SetPrevClosePx(v float64)                    { m.PrevClosePx = &v }
func (m *NoRelatedSym) SetQuoteRequestType(v int)                   { m.QuoteRequestType = &v }
func (m *NoRelatedSym) SetQuoteType(v int)                          { m.QuoteType = &v }
func (m *NoRelatedSym) SetTradingSessionID(v string)                { m.TradingSessionID = &v }
func (m *NoRelatedSym) SetTradingSessionSubID(v string)             { m.TradingSessionSubID = &v }
func (m *NoRelatedSym) SetTradeOriginationDate(v string)            { m.TradeOriginationDate = &v }
func (m *NoRelatedSym) SetStipulations(v stipulations.Stipulations) { m.Stipulations = &v }
func (m *NoRelatedSym) SetSide(v string)                            { m.Side = &v }
func (m *NoRelatedSym) SetQuantityType(v int)                       { m.QuantityType = &v }
func (m *NoRelatedSym) SetOrderQty(v float64)                       { m.OrderQty = &v }
func (m *NoRelatedSym) SetCashOrderQty(v float64)                   { m.CashOrderQty = &v }
func (m *NoRelatedSym) SetSettlmntTyp(v string)                     { m.SettlmntTyp = &v }
func (m *NoRelatedSym) SetFutSettDate(v string)                     { m.FutSettDate = &v }
func (m *NoRelatedSym) SetOrdType(v string)                         { m.OrdType = &v }
func (m *NoRelatedSym) SetFutSettDate2(v string)                    { m.FutSettDate2 = &v }
func (m *NoRelatedSym) SetOrderQty2(v float64)                      { m.OrderQty2 = &v }
func (m *NoRelatedSym) SetExpireTime(v time.Time)                   { m.ExpireTime = &v }
func (m *NoRelatedSym) SetTransactTime(v time.Time)                 { m.TransactTime = &v }
func (m *NoRelatedSym) SetCurrency(v string)                        { m.Currency = &v }
func (m *NoRelatedSym) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *NoRelatedSym) SetPriceType(v int)                 { m.PriceType = &v }
func (m *NoRelatedSym) SetPrice(v float64)                 { m.Price = &v }
func (m *NoRelatedSym) SetPrice2(v float64)                { m.Price2 = &v }
func (m *NoRelatedSym) SetYieldData(v yielddata.YieldData) { m.YieldData = &v }

//Message is a QuoteRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"R"`
	fix43.Header
	//QuoteReqID is a required field for QuoteRequest.
	QuoteReqID string `fix:"131"`
	//RFQReqID is a non-required field for QuoteRequest.
	RFQReqID *string `fix:"644"`
	//NoRelatedSym is a required field for QuoteRequest.
	NoRelatedSym []NoRelatedSym `fix:"146"`
	//Text is a non-required field for QuoteRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for QuoteRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for QuoteRequest.
	EncodedText *string `fix:"355"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteReqID(v string)           { m.QuoteReqID = v }
func (m *Message) SetRFQReqID(v string)             { m.RFQReqID = &v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
func (m *Message) SetText(v string)                 { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)          { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)          { m.EncodedText = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX43, "R", r
}
