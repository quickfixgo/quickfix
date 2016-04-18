//Package advertisement msg type = 7.
package advertisement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
	"time"
)

//Message is a Advertisement FIX Message
type Message struct {
	FIXMsgType string `fix:"7"`
	fix41.Header
	//AdvId is a required field for Advertisement.
	AdvId string `fix:"2"`
	//AdvTransType is a required field for Advertisement.
	AdvTransType string `fix:"5"`
	//AdvRefID is a non-required field for Advertisement.
	AdvRefID *string `fix:"3"`
	//Symbol is a required field for Advertisement.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for Advertisement.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for Advertisement.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for Advertisement.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for Advertisement.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for Advertisement.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for Advertisement.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for Advertisement.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for Advertisement.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for Advertisement.
	OptAttribute *string `fix:"206"`
	//SecurityExchange is a non-required field for Advertisement.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for Advertisement.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for Advertisement.
	SecurityDesc *string `fix:"107"`
	//AdvSide is a required field for Advertisement.
	AdvSide string `fix:"4"`
	//Shares is a required field for Advertisement.
	Shares int `fix:"53"`
	//Price is a non-required field for Advertisement.
	Price *float64 `fix:"44"`
	//Currency is a non-required field for Advertisement.
	Currency *string `fix:"15"`
	//TradeDate is a non-required field for Advertisement.
	TradeDate *string `fix:"75"`
	//TransactTime is a non-required field for Advertisement.
	TransactTime *time.Time `fix:"60"`
	//Text is a non-required field for Advertisement.
	Text *string `fix:"58"`
	//URLLink is a non-required field for Advertisement.
	URLLink *string `fix:"149"`
	//LastMkt is a non-required field for Advertisement.
	LastMkt *string `fix:"30"`
	fix41.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized Advertisement instance
func New(advid string, advtranstype string, symbol string, advside string, shares int) *Message {
	var m Message
	m.SetAdvId(advid)
	m.SetAdvTransType(advtranstype)
	m.SetSymbol(symbol)
	m.SetAdvSide(advside)
	m.SetShares(shares)
	return &m
}

func (m *Message) SetAdvId(v string)             { m.AdvId = v }
func (m *Message) SetAdvTransType(v string)      { m.AdvTransType = v }
func (m *Message) SetAdvRefID(v string)          { m.AdvRefID = &v }
func (m *Message) SetSymbol(v string)            { m.Symbol = v }
func (m *Message) SetSymbolSfx(v string)         { m.SymbolSfx = &v }
func (m *Message) SetSecurityID(v string)        { m.SecurityID = &v }
func (m *Message) SetIDSource(v string)          { m.IDSource = &v }
func (m *Message) SetSecurityType(v string)      { m.SecurityType = &v }
func (m *Message) SetMaturityMonthYear(v string) { m.MaturityMonthYear = &v }
func (m *Message) SetMaturityDay(v int)          { m.MaturityDay = &v }
func (m *Message) SetPutOrCall(v int)            { m.PutOrCall = &v }
func (m *Message) SetStrikePrice(v float64)      { m.StrikePrice = &v }
func (m *Message) SetOptAttribute(v string)      { m.OptAttribute = &v }
func (m *Message) SetSecurityExchange(v string)  { m.SecurityExchange = &v }
func (m *Message) SetIssuer(v string)            { m.Issuer = &v }
func (m *Message) SetSecurityDesc(v string)      { m.SecurityDesc = &v }
func (m *Message) SetAdvSide(v string)           { m.AdvSide = v }
func (m *Message) SetShares(v int)               { m.Shares = v }
func (m *Message) SetPrice(v float64)            { m.Price = &v }
func (m *Message) SetCurrency(v string)          { m.Currency = &v }
func (m *Message) SetTradeDate(v string)         { m.TradeDate = &v }
func (m *Message) SetTransactTime(v time.Time)   { m.TransactTime = &v }
func (m *Message) SetText(v string)              { m.Text = &v }
func (m *Message) SetURLLink(v string)           { m.URLLink = &v }
func (m *Message) SetLastMkt(v string)           { m.LastMkt = &v }

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
	return enum.BeginStringFIX41, "7", r
}
