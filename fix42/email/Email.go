//Package email msg type = C.
package email

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//NoRoutingIDs is a repeating group in Email
type NoRoutingIDs struct {
	//RoutingType is a non-required field for NoRoutingIDs.
	RoutingType *int `fix:"216"`
	//RoutingID is a non-required field for NoRoutingIDs.
	RoutingID *string `fix:"217"`
}

//NoRelatedSym is a repeating group in Email
type NoRelatedSym struct {
	//RelatdSym is a non-required field for NoRelatedSym.
	RelatdSym *string `fix:"46"`
	//SymbolSfx is a non-required field for NoRelatedSym.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for NoRelatedSym.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for NoRelatedSym.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for NoRelatedSym.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for NoRelatedSym.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for NoRelatedSym.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for NoRelatedSym.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for NoRelatedSym.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for NoRelatedSym.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for NoRelatedSym.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for NoRelatedSym.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for NoRelatedSym.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for NoRelatedSym.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for NoRelatedSym.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for NoRelatedSym.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for NoRelatedSym.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for NoRelatedSym.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for NoRelatedSym.
	EncodedSecurityDesc *string `fix:"351"`
}

//LinesOfText is a repeating group in Email
type LinesOfText struct {
	//Text is a required field for LinesOfText.
	Text string `fix:"58"`
	//EncodedTextLen is a non-required field for LinesOfText.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for LinesOfText.
	EncodedText *string `fix:"355"`
}

//Message is a Email FIX Message
type Message struct {
	FIXMsgType string `fix:"C"`
	Header     fix42.Header
	//EmailThreadID is a required field for Email.
	EmailThreadID string `fix:"164"`
	//EmailType is a required field for Email.
	EmailType string `fix:"94"`
	//OrigTime is a non-required field for Email.
	OrigTime *time.Time `fix:"42"`
	//Subject is a required field for Email.
	Subject string `fix:"147"`
	//EncodedSubjectLen is a non-required field for Email.
	EncodedSubjectLen *int `fix:"356"`
	//EncodedSubject is a non-required field for Email.
	EncodedSubject *string `fix:"357"`
	//NoRoutingIDs is a non-required field for Email.
	NoRoutingIDs []NoRoutingIDs `fix:"215,omitempty"`
	//NoRelatedSym is a non-required field for Email.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
	//OrderID is a non-required field for Email.
	OrderID *string `fix:"37"`
	//ClOrdID is a non-required field for Email.
	ClOrdID *string `fix:"11"`
	//LinesOfText is a required field for Email.
	LinesOfText []LinesOfText `fix:"33"`
	//RawDataLength is a non-required field for Email.
	RawDataLength *int `fix:"95"`
	//RawData is a non-required field for Email.
	RawData *string `fix:"96"`
	Trailer fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.BeginStringFIX42, "C", r
}
