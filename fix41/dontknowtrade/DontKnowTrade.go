//Package dontknowtrade msg type = Q.
package dontknowtrade

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
)

//Message is a DontKnowTrade FIX Message
type Message struct {
	FIXMsgType string `fix:"Q"`
	Header     fix41.Header
	//OrderID is a non-required field for DontKnowTrade.
	OrderID *string `fix:"37"`
	//ExecID is a non-required field for DontKnowTrade.
	ExecID *string `fix:"17"`
	//DKReason is a required field for DontKnowTrade.
	DKReason string `fix:"127"`
	//Symbol is a required field for DontKnowTrade.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for DontKnowTrade.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for DontKnowTrade.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for DontKnowTrade.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for DontKnowTrade.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for DontKnowTrade.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for DontKnowTrade.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for DontKnowTrade.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for DontKnowTrade.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for DontKnowTrade.
	OptAttribute *string `fix:"206"`
	//SecurityExchange is a non-required field for DontKnowTrade.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for DontKnowTrade.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for DontKnowTrade.
	SecurityDesc *string `fix:"107"`
	//Side is a required field for DontKnowTrade.
	Side string `fix:"54"`
	//OrderQty is a non-required field for DontKnowTrade.
	OrderQty *int `fix:"38"`
	//CashOrderQty is a non-required field for DontKnowTrade.
	CashOrderQty *float64 `fix:"152"`
	//LastShares is a non-required field for DontKnowTrade.
	LastShares *int `fix:"32"`
	//LastPx is a non-required field for DontKnowTrade.
	LastPx *float64 `fix:"31"`
	//Text is a non-required field for DontKnowTrade.
	Text    *string `fix:"58"`
	Trailer fix41.Trailer
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
	return enum.BeginStringFIX41, "Q", r
}
