//Package ordercancelrequest msg type = F.
package ordercancelrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//Message is a OrderCancelRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"F"`
	Header     fix42.Header
	//OrigClOrdID is a required field for OrderCancelRequest.
	OrigClOrdID string `fix:"41"`
	//OrderID is a non-required field for OrderCancelRequest.
	OrderID *string `fix:"37"`
	//ClOrdID is a required field for OrderCancelRequest.
	ClOrdID string `fix:"11"`
	//ListID is a non-required field for OrderCancelRequest.
	ListID *string `fix:"66"`
	//Account is a non-required field for OrderCancelRequest.
	Account *string `fix:"1"`
	//ClientID is a non-required field for OrderCancelRequest.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for OrderCancelRequest.
	ExecBroker *string `fix:"76"`
	//Symbol is a required field for OrderCancelRequest.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for OrderCancelRequest.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for OrderCancelRequest.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for OrderCancelRequest.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for OrderCancelRequest.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for OrderCancelRequest.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for OrderCancelRequest.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for OrderCancelRequest.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for OrderCancelRequest.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for OrderCancelRequest.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for OrderCancelRequest.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for OrderCancelRequest.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for OrderCancelRequest.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for OrderCancelRequest.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for OrderCancelRequest.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for OrderCancelRequest.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for OrderCancelRequest.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for OrderCancelRequest.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for OrderCancelRequest.
	EncodedSecurityDesc *string `fix:"351"`
	//Side is a required field for OrderCancelRequest.
	Side string `fix:"54"`
	//TransactTime is a required field for OrderCancelRequest.
	TransactTime time.Time `fix:"60"`
	//OrderQty is a non-required field for OrderCancelRequest.
	OrderQty *float64 `fix:"38"`
	//CashOrderQty is a non-required field for OrderCancelRequest.
	CashOrderQty *float64 `fix:"152"`
	//ComplianceID is a non-required field for OrderCancelRequest.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for OrderCancelRequest.
	SolicitedFlag *bool `fix:"377"`
	//Text is a non-required field for OrderCancelRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderCancelRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderCancelRequest.
	EncodedText *string `fix:"355"`
	Trailer     fix42.Trailer
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
	return enum.BeginStringFIX42, "F", r
}
