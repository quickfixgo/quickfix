//Package advertisement msg type = 7.
package advertisement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//Message is a Advertisement FIX Message
type Message struct {
	FIXMsgType string `fix:"7"`
	Header     fix42.Header
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
	//ContractMultiplier is a non-required field for Advertisement.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for Advertisement.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for Advertisement.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for Advertisement.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for Advertisement.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for Advertisement.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for Advertisement.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for Advertisement.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for Advertisement.
	EncodedSecurityDesc *string `fix:"351"`
	//AdvSide is a required field for Advertisement.
	AdvSide string `fix:"4"`
	//Shares is a required field for Advertisement.
	Shares float64 `fix:"53"`
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
	//EncodedTextLen is a non-required field for Advertisement.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for Advertisement.
	EncodedText *string `fix:"355"`
	//URLLink is a non-required field for Advertisement.
	URLLink *string `fix:"149"`
	//LastMkt is a non-required field for Advertisement.
	LastMkt *string `fix:"30"`
	//TradingSessionID is a non-required field for Advertisement.
	TradingSessionID *string `fix:"336"`
	Trailer          fix42.Trailer
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
	return enum.BeginStringFIX42, "7", r
}
