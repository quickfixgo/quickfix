//Package tradecapturereportrequest msg type = AD.
package tradecapturereportrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentextension"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"time"
)

//NoUnderlyings is a repeating group in TradeCaptureReportRequest
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//NoLegs is a repeating group in TradeCaptureReportRequest
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//NoDates is a repeating group in TradeCaptureReportRequest
type NoDates struct {
	//TradeDate is a non-required field for NoDates.
	TradeDate *string `fix:"75"`
	//TransactTime is a non-required field for NoDates.
	TransactTime *time.Time `fix:"60"`
}

func (m *NoDates) SetTradeDate(v string)       { m.TradeDate = &v }
func (m *NoDates) SetTransactTime(v time.Time) { m.TransactTime = &v }

//Message is a TradeCaptureReportRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AD"`
	Header     fix44.Header
	//TradeRequestID is a required field for TradeCaptureReportRequest.
	TradeRequestID string `fix:"568"`
	//TradeRequestType is a required field for TradeCaptureReportRequest.
	TradeRequestType int `fix:"569"`
	//SubscriptionRequestType is a non-required field for TradeCaptureReportRequest.
	SubscriptionRequestType *string `fix:"263"`
	//TradeReportID is a non-required field for TradeCaptureReportRequest.
	TradeReportID *string `fix:"571"`
	//SecondaryTradeReportID is a non-required field for TradeCaptureReportRequest.
	SecondaryTradeReportID *string `fix:"818"`
	//ExecID is a non-required field for TradeCaptureReportRequest.
	ExecID *string `fix:"17"`
	//ExecType is a non-required field for TradeCaptureReportRequest.
	ExecType *string `fix:"150"`
	//OrderID is a non-required field for TradeCaptureReportRequest.
	OrderID *string `fix:"37"`
	//ClOrdID is a non-required field for TradeCaptureReportRequest.
	ClOrdID *string `fix:"11"`
	//MatchStatus is a non-required field for TradeCaptureReportRequest.
	MatchStatus *string `fix:"573"`
	//TrdType is a non-required field for TradeCaptureReportRequest.
	TrdType *int `fix:"828"`
	//TrdSubType is a non-required field for TradeCaptureReportRequest.
	TrdSubType *int `fix:"829"`
	//TransferReason is a non-required field for TradeCaptureReportRequest.
	TransferReason *string `fix:"830"`
	//SecondaryTrdType is a non-required field for TradeCaptureReportRequest.
	SecondaryTrdType *int `fix:"855"`
	//TradeLinkID is a non-required field for TradeCaptureReportRequest.
	TradeLinkID *string `fix:"820"`
	//TrdMatchID is a non-required field for TradeCaptureReportRequest.
	TrdMatchID *string `fix:"880"`
	//Parties Component
	Parties parties.Component
	//Instrument Component
	Instrument instrument.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//NoUnderlyings is a non-required field for TradeCaptureReportRequest.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//NoLegs is a non-required field for TradeCaptureReportRequest.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoDates is a non-required field for TradeCaptureReportRequest.
	NoDates []NoDates `fix:"580,omitempty"`
	//ClearingBusinessDate is a non-required field for TradeCaptureReportRequest.
	ClearingBusinessDate *string `fix:"715"`
	//TradingSessionID is a non-required field for TradeCaptureReportRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for TradeCaptureReportRequest.
	TradingSessionSubID *string `fix:"625"`
	//TimeBracket is a non-required field for TradeCaptureReportRequest.
	TimeBracket *string `fix:"943"`
	//Side is a non-required field for TradeCaptureReportRequest.
	Side *string `fix:"54"`
	//MultiLegReportingType is a non-required field for TradeCaptureReportRequest.
	MultiLegReportingType *string `fix:"442"`
	//TradeInputSource is a non-required field for TradeCaptureReportRequest.
	TradeInputSource *string `fix:"578"`
	//TradeInputDevice is a non-required field for TradeCaptureReportRequest.
	TradeInputDevice *string `fix:"579"`
	//ResponseTransportType is a non-required field for TradeCaptureReportRequest.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for TradeCaptureReportRequest.
	ResponseDestination *string `fix:"726"`
	//Text is a non-required field for TradeCaptureReportRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for TradeCaptureReportRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for TradeCaptureReportRequest.
	EncodedText *string `fix:"355"`
	Trailer     fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetTradeRequestID(v string)          { m.TradeRequestID = v }
func (m *Message) SetTradeRequestType(v int)           { m.TradeRequestType = v }
func (m *Message) SetSubscriptionRequestType(v string) { m.SubscriptionRequestType = &v }
func (m *Message) SetTradeReportID(v string)           { m.TradeReportID = &v }
func (m *Message) SetSecondaryTradeReportID(v string)  { m.SecondaryTradeReportID = &v }
func (m *Message) SetExecID(v string)                  { m.ExecID = &v }
func (m *Message) SetExecType(v string)                { m.ExecType = &v }
func (m *Message) SetOrderID(v string)                 { m.OrderID = &v }
func (m *Message) SetClOrdID(v string)                 { m.ClOrdID = &v }
func (m *Message) SetMatchStatus(v string)             { m.MatchStatus = &v }
func (m *Message) SetTrdType(v int)                    { m.TrdType = &v }
func (m *Message) SetTrdSubType(v int)                 { m.TrdSubType = &v }
func (m *Message) SetTransferReason(v string)          { m.TransferReason = &v }
func (m *Message) SetSecondaryTrdType(v int)           { m.SecondaryTrdType = &v }
func (m *Message) SetTradeLinkID(v string)             { m.TradeLinkID = &v }
func (m *Message) SetTrdMatchID(v string)              { m.TrdMatchID = &v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)  { m.NoUnderlyings = v }
func (m *Message) SetNoLegs(v []NoLegs)                { m.NoLegs = v }
func (m *Message) SetNoDates(v []NoDates)              { m.NoDates = v }
func (m *Message) SetClearingBusinessDate(v string)    { m.ClearingBusinessDate = &v }
func (m *Message) SetTradingSessionID(v string)        { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)     { m.TradingSessionSubID = &v }
func (m *Message) SetTimeBracket(v string)             { m.TimeBracket = &v }
func (m *Message) SetSide(v string)                    { m.Side = &v }
func (m *Message) SetMultiLegReportingType(v string)   { m.MultiLegReportingType = &v }
func (m *Message) SetTradeInputSource(v string)        { m.TradeInputSource = &v }
func (m *Message) SetTradeInputDevice(v string)        { m.TradeInputDevice = &v }
func (m *Message) SetResponseTransportType(v int)      { m.ResponseTransportType = &v }
func (m *Message) SetResponseDestination(v string)     { m.ResponseDestination = &v }
func (m *Message) SetText(v string)                    { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)             { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)             { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "AD", r
}
