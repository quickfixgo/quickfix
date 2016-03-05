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

func (m *NoRoutingIDs) SetRoutingType(v int)  { m.RoutingType = &v }
func (m *NoRoutingIDs) SetRoutingID(v string) { m.RoutingID = &v }

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

func (m *NoRelatedSym) SetRelatdSym(v string)           { m.RelatdSym = &v }
func (m *NoRelatedSym) SetSymbolSfx(v string)           { m.SymbolSfx = &v }
func (m *NoRelatedSym) SetSecurityID(v string)          { m.SecurityID = &v }
func (m *NoRelatedSym) SetIDSource(v string)            { m.IDSource = &v }
func (m *NoRelatedSym) SetSecurityType(v string)        { m.SecurityType = &v }
func (m *NoRelatedSym) SetMaturityMonthYear(v string)   { m.MaturityMonthYear = &v }
func (m *NoRelatedSym) SetMaturityDay(v int)            { m.MaturityDay = &v }
func (m *NoRelatedSym) SetPutOrCall(v int)              { m.PutOrCall = &v }
func (m *NoRelatedSym) SetStrikePrice(v float64)        { m.StrikePrice = &v }
func (m *NoRelatedSym) SetOptAttribute(v string)        { m.OptAttribute = &v }
func (m *NoRelatedSym) SetContractMultiplier(v float64) { m.ContractMultiplier = &v }
func (m *NoRelatedSym) SetCouponRate(v float64)         { m.CouponRate = &v }
func (m *NoRelatedSym) SetSecurityExchange(v string)    { m.SecurityExchange = &v }
func (m *NoRelatedSym) SetIssuer(v string)              { m.Issuer = &v }
func (m *NoRelatedSym) SetEncodedIssuerLen(v int)       { m.EncodedIssuerLen = &v }
func (m *NoRelatedSym) SetEncodedIssuer(v string)       { m.EncodedIssuer = &v }
func (m *NoRelatedSym) SetSecurityDesc(v string)        { m.SecurityDesc = &v }
func (m *NoRelatedSym) SetEncodedSecurityDescLen(v int) { m.EncodedSecurityDescLen = &v }
func (m *NoRelatedSym) SetEncodedSecurityDesc(v string) { m.EncodedSecurityDesc = &v }

//LinesOfText is a repeating group in Email
type LinesOfText struct {
	//Text is a required field for LinesOfText.
	Text string `fix:"58"`
	//EncodedTextLen is a non-required field for LinesOfText.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for LinesOfText.
	EncodedText *string `fix:"355"`
}

func (m *LinesOfText) SetText(v string)        { m.Text = v }
func (m *LinesOfText) SetEncodedTextLen(v int) { m.EncodedTextLen = &v }
func (m *LinesOfText) SetEncodedText(v string) { m.EncodedText = &v }

//Message is a Email FIX Message
type Message struct {
	FIXMsgType string `fix:"C"`
	fix42.Header
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
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetEmailThreadID(v string)        { m.EmailThreadID = v }
func (m *Message) SetEmailType(v string)            { m.EmailType = v }
func (m *Message) SetOrigTime(v time.Time)          { m.OrigTime = &v }
func (m *Message) SetSubject(v string)              { m.Subject = v }
func (m *Message) SetEncodedSubjectLen(v int)       { m.EncodedSubjectLen = &v }
func (m *Message) SetEncodedSubject(v string)       { m.EncodedSubject = &v }
func (m *Message) SetNoRoutingIDs(v []NoRoutingIDs) { m.NoRoutingIDs = v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
func (m *Message) SetOrderID(v string)              { m.OrderID = &v }
func (m *Message) SetClOrdID(v string)              { m.ClOrdID = &v }
func (m *Message) SetLinesOfText(v []LinesOfText)   { m.LinesOfText = v }
func (m *Message) SetRawDataLength(v int)           { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)              { m.RawData = &v }

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
	return enum.BeginStringFIX42, "C", r
}
