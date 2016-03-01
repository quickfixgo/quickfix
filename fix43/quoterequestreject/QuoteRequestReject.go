//Package quoterequestreject msg type = AG.
package quoterequestreject

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

//NoRelatedSym is a repeating group in QuoteRequestReject
type NoRelatedSym struct {
	//Instrument Component
	Instrument instrument.Component
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
	//Stipulations Component
	Stipulations stipulations.Component
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
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//PriceType is a non-required field for NoRelatedSym.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NoRelatedSym.
	Price *float64 `fix:"44"`
	//Price2 is a non-required field for NoRelatedSym.
	Price2 *float64 `fix:"640"`
	//YieldData Component
	YieldData yielddata.Component
}

//Message is a QuoteRequestReject FIX Message
type Message struct {
	FIXMsgType string `fix:"AG"`
	Header     fix43.Header
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
	Trailer     fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.BeginStringFIX43, "AG", r
}
