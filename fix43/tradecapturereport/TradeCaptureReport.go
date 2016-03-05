//Package tradecapturereport msg type = AE.
package tradecapturereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/commissiondata"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/orderqtydata"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"time"
)

//NoSides is a repeating group in TradeCaptureReport
type NoSides struct {
	//Side is a required field for NoSides.
	Side string `fix:"54"`
	//OrderID is a required field for NoSides.
	OrderID string `fix:"37"`
	//SecondaryOrderID is a non-required field for NoSides.
	SecondaryOrderID *string `fix:"198"`
	//ClOrdID is a non-required field for NoSides.
	ClOrdID *string `fix:"11"`
	//Parties Component
	parties.Parties
	//Account is a non-required field for NoSides.
	Account *string `fix:"1"`
	//AccountType is a non-required field for NoSides.
	AccountType *int `fix:"581"`
	//ProcessCode is a non-required field for NoSides.
	ProcessCode *string `fix:"81"`
	//OddLot is a non-required field for NoSides.
	OddLot *bool `fix:"575"`
	//NoClearingInstructions is a non-required field for NoSides.
	NoClearingInstructions []NoClearingInstructions `fix:"576,omitempty"`
	//ClearingFeeIndicator is a non-required field for NoSides.
	ClearingFeeIndicator *string `fix:"635"`
	//TradeInputSource is a non-required field for NoSides.
	TradeInputSource *string `fix:"578"`
	//TradeInputDevice is a non-required field for NoSides.
	TradeInputDevice *string `fix:"579"`
	//Currency is a non-required field for NoSides.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NoSides.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NoSides.
	SolicitedFlag *bool `fix:"377"`
	//OrderCapacity is a non-required field for NoSides.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NoSides.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NoSides.
	CustOrderCapacity *int `fix:"582"`
	//TransBkdTime is a non-required field for NoSides.
	TransBkdTime *time.Time `fix:"483"`
	//TradingSessionID is a non-required field for NoSides.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoSides.
	TradingSessionSubID *string `fix:"625"`
	//CommissionData Component
	commissiondata.CommissionData
	//GrossTradeAmt is a non-required field for NoSides.
	GrossTradeAmt *float64 `fix:"381"`
	//NumDaysInterest is a non-required field for NoSides.
	NumDaysInterest *int `fix:"157"`
	//ExDate is a non-required field for NoSides.
	ExDate *string `fix:"230"`
	//AccruedInterestRate is a non-required field for NoSides.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for NoSides.
	AccruedInterestAmt *float64 `fix:"159"`
	//Concession is a non-required field for NoSides.
	Concession *float64 `fix:"238"`
	//TotalTakedown is a non-required field for NoSides.
	TotalTakedown *float64 `fix:"237"`
	//NetMoney is a non-required field for NoSides.
	NetMoney *float64 `fix:"118"`
	//SettlCurrAmt is a non-required field for NoSides.
	SettlCurrAmt *float64 `fix:"119"`
	//SettlCurrency is a non-required field for NoSides.
	SettlCurrency *string `fix:"120"`
	//SettlCurrFxRate is a non-required field for NoSides.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlCurrFxRateCalc is a non-required field for NoSides.
	SettlCurrFxRateCalc *string `fix:"156"`
	//PositionEffect is a non-required field for NoSides.
	PositionEffect *string `fix:"77"`
	//Text is a non-required field for NoSides.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoSides.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoSides.
	EncodedText *string `fix:"355"`
	//MultiLegReportingType is a non-required field for NoSides.
	MultiLegReportingType *string `fix:"442"`
	//NoContAmts is a non-required field for NoSides.
	NoContAmts []NoContAmts `fix:"518,omitempty"`
	//NoMiscFees is a non-required field for NoSides.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
}

func (m *NoSides) SetSide(v string)                                     { m.Side = v }
func (m *NoSides) SetOrderID(v string)                                  { m.OrderID = v }
func (m *NoSides) SetSecondaryOrderID(v string)                         { m.SecondaryOrderID = &v }
func (m *NoSides) SetClOrdID(v string)                                  { m.ClOrdID = &v }
func (m *NoSides) SetAccount(v string)                                  { m.Account = &v }
func (m *NoSides) SetAccountType(v int)                                 { m.AccountType = &v }
func (m *NoSides) SetProcessCode(v string)                              { m.ProcessCode = &v }
func (m *NoSides) SetOddLot(v bool)                                     { m.OddLot = &v }
func (m *NoSides) SetNoClearingInstructions(v []NoClearingInstructions) { m.NoClearingInstructions = v }
func (m *NoSides) SetClearingFeeIndicator(v string)                     { m.ClearingFeeIndicator = &v }
func (m *NoSides) SetTradeInputSource(v string)                         { m.TradeInputSource = &v }
func (m *NoSides) SetTradeInputDevice(v string)                         { m.TradeInputDevice = &v }
func (m *NoSides) SetCurrency(v string)                                 { m.Currency = &v }
func (m *NoSides) SetComplianceID(v string)                             { m.ComplianceID = &v }
func (m *NoSides) SetSolicitedFlag(v bool)                              { m.SolicitedFlag = &v }
func (m *NoSides) SetOrderCapacity(v string)                            { m.OrderCapacity = &v }
func (m *NoSides) SetOrderRestrictions(v string)                        { m.OrderRestrictions = &v }
func (m *NoSides) SetCustOrderCapacity(v int)                           { m.CustOrderCapacity = &v }
func (m *NoSides) SetTransBkdTime(v time.Time)                          { m.TransBkdTime = &v }
func (m *NoSides) SetTradingSessionID(v string)                         { m.TradingSessionID = &v }
func (m *NoSides) SetTradingSessionSubID(v string)                      { m.TradingSessionSubID = &v }
func (m *NoSides) SetGrossTradeAmt(v float64)                           { m.GrossTradeAmt = &v }
func (m *NoSides) SetNumDaysInterest(v int)                             { m.NumDaysInterest = &v }
func (m *NoSides) SetExDate(v string)                                   { m.ExDate = &v }
func (m *NoSides) SetAccruedInterestRate(v float64)                     { m.AccruedInterestRate = &v }
func (m *NoSides) SetAccruedInterestAmt(v float64)                      { m.AccruedInterestAmt = &v }
func (m *NoSides) SetConcession(v float64)                              { m.Concession = &v }
func (m *NoSides) SetTotalTakedown(v float64)                           { m.TotalTakedown = &v }
func (m *NoSides) SetNetMoney(v float64)                                { m.NetMoney = &v }
func (m *NoSides) SetSettlCurrAmt(v float64)                            { m.SettlCurrAmt = &v }
func (m *NoSides) SetSettlCurrency(v string)                            { m.SettlCurrency = &v }
func (m *NoSides) SetSettlCurrFxRate(v float64)                         { m.SettlCurrFxRate = &v }
func (m *NoSides) SetSettlCurrFxRateCalc(v string)                      { m.SettlCurrFxRateCalc = &v }
func (m *NoSides) SetPositionEffect(v string)                           { m.PositionEffect = &v }
func (m *NoSides) SetText(v string)                                     { m.Text = &v }
func (m *NoSides) SetEncodedTextLen(v int)                              { m.EncodedTextLen = &v }
func (m *NoSides) SetEncodedText(v string)                              { m.EncodedText = &v }
func (m *NoSides) SetMultiLegReportingType(v string)                    { m.MultiLegReportingType = &v }
func (m *NoSides) SetNoContAmts(v []NoContAmts)                         { m.NoContAmts = v }
func (m *NoSides) SetNoMiscFees(v []NoMiscFees)                         { m.NoMiscFees = v }

//NoClearingInstructions is a repeating group in NoSides
type NoClearingInstructions struct {
	//ClearingInstruction is a non-required field for NoClearingInstructions.
	ClearingInstruction *int `fix:"577"`
}

func (m *NoClearingInstructions) SetClearingInstruction(v int) { m.ClearingInstruction = &v }

//NoContAmts is a repeating group in NoSides
type NoContAmts struct {
	//ContAmtType is a non-required field for NoContAmts.
	ContAmtType *int `fix:"519"`
	//ContAmtValue is a non-required field for NoContAmts.
	ContAmtValue *float64 `fix:"520"`
	//ContAmtCurr is a non-required field for NoContAmts.
	ContAmtCurr *string `fix:"521"`
}

func (m *NoContAmts) SetContAmtType(v int)      { m.ContAmtType = &v }
func (m *NoContAmts) SetContAmtValue(v float64) { m.ContAmtValue = &v }
func (m *NoContAmts) SetContAmtCurr(v string)   { m.ContAmtCurr = &v }

//NoMiscFees is a repeating group in NoSides
type NoMiscFees struct {
	//MiscFeeAmt is a non-required field for NoMiscFees.
	MiscFeeAmt *float64 `fix:"137"`
	//MiscFeeCurr is a non-required field for NoMiscFees.
	MiscFeeCurr *string `fix:"138"`
	//MiscFeeType is a non-required field for NoMiscFees.
	MiscFeeType *string `fix:"139"`
}

func (m *NoMiscFees) SetMiscFeeAmt(v float64) { m.MiscFeeAmt = &v }
func (m *NoMiscFees) SetMiscFeeCurr(v string) { m.MiscFeeCurr = &v }
func (m *NoMiscFees) SetMiscFeeType(v string) { m.MiscFeeType = &v }

//Message is a TradeCaptureReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AE"`
	fix43.Header
	//TradeReportID is a required field for TradeCaptureReport.
	TradeReportID string `fix:"571"`
	//TradeReportTransType is a non-required field for TradeCaptureReport.
	TradeReportTransType *string `fix:"487"`
	//TradeRequestID is a non-required field for TradeCaptureReport.
	TradeRequestID *string `fix:"568"`
	//ExecType is a required field for TradeCaptureReport.
	ExecType string `fix:"150"`
	//TradeReportRefID is a non-required field for TradeCaptureReport.
	TradeReportRefID *string `fix:"572"`
	//ExecID is a non-required field for TradeCaptureReport.
	ExecID *string `fix:"17"`
	//SecondaryExecID is a non-required field for TradeCaptureReport.
	SecondaryExecID *string `fix:"527"`
	//ExecRestatementReason is a non-required field for TradeCaptureReport.
	ExecRestatementReason *int `fix:"378"`
	//PreviouslyReported is a required field for TradeCaptureReport.
	PreviouslyReported bool `fix:"570"`
	//Instrument Component
	instrument.Instrument
	//OrderQtyData Component
	orderqtydata.OrderQtyData
	//LastQty is a required field for TradeCaptureReport.
	LastQty float64 `fix:"32"`
	//LastPx is a required field for TradeCaptureReport.
	LastPx float64 `fix:"31"`
	//LastSpotRate is a non-required field for TradeCaptureReport.
	LastSpotRate *float64 `fix:"194"`
	//LastForwardPoints is a non-required field for TradeCaptureReport.
	LastForwardPoints *float64 `fix:"195"`
	//LastMkt is a non-required field for TradeCaptureReport.
	LastMkt *string `fix:"30"`
	//TradeDate is a required field for TradeCaptureReport.
	TradeDate string `fix:"75"`
	//TransactTime is a required field for TradeCaptureReport.
	TransactTime time.Time `fix:"60"`
	//SettlmntTyp is a non-required field for TradeCaptureReport.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for TradeCaptureReport.
	FutSettDate *string `fix:"64"`
	//MatchStatus is a non-required field for TradeCaptureReport.
	MatchStatus *string `fix:"573"`
	//MatchType is a non-required field for TradeCaptureReport.
	MatchType *string `fix:"574"`
	//NoSides is a required field for TradeCaptureReport.
	NoSides []NoSides `fix:"552"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetTradeReportID(v string)        { m.TradeReportID = v }
func (m *Message) SetTradeReportTransType(v string) { m.TradeReportTransType = &v }
func (m *Message) SetTradeRequestID(v string)       { m.TradeRequestID = &v }
func (m *Message) SetExecType(v string)             { m.ExecType = v }
func (m *Message) SetTradeReportRefID(v string)     { m.TradeReportRefID = &v }
func (m *Message) SetExecID(v string)               { m.ExecID = &v }
func (m *Message) SetSecondaryExecID(v string)      { m.SecondaryExecID = &v }
func (m *Message) SetExecRestatementReason(v int)   { m.ExecRestatementReason = &v }
func (m *Message) SetPreviouslyReported(v bool)     { m.PreviouslyReported = v }
func (m *Message) SetLastQty(v float64)             { m.LastQty = v }
func (m *Message) SetLastPx(v float64)              { m.LastPx = v }
func (m *Message) SetLastSpotRate(v float64)        { m.LastSpotRate = &v }
func (m *Message) SetLastForwardPoints(v float64)   { m.LastForwardPoints = &v }
func (m *Message) SetLastMkt(v string)              { m.LastMkt = &v }
func (m *Message) SetTradeDate(v string)            { m.TradeDate = v }
func (m *Message) SetTransactTime(v time.Time)      { m.TransactTime = v }
func (m *Message) SetSettlmntTyp(v string)          { m.SettlmntTyp = &v }
func (m *Message) SetFutSettDate(v string)          { m.FutSettDate = &v }
func (m *Message) SetMatchStatus(v string)          { m.MatchStatus = &v }
func (m *Message) SetMatchType(v string)            { m.MatchType = &v }
func (m *Message) SetNoSides(v []NoSides)           { m.NoSides = v }

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
	return enum.BeginStringFIX43, "AE", r
}
