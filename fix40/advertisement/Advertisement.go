//Package advertisement msg type = 7.
package advertisement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
	"time"
)

//Message is a Advertisement FIX Message
type Message struct {
	FIXMsgType string `fix:"7"`
	Header     fix40.Header
	//AdvId is a required field for Advertisement.
	AdvId int `fix:"2"`
	//AdvTransType is a required field for Advertisement.
	AdvTransType string `fix:"5"`
	//AdvRefID is a non-required field for Advertisement.
	AdvRefID *int `fix:"3"`
	//Symbol is a required field for Advertisement.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for Advertisement.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for Advertisement.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for Advertisement.
	IDSource *string `fix:"22"`
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
	//TransactTime is a non-required field for Advertisement.
	TransactTime *time.Time `fix:"60"`
	//Text is a non-required field for Advertisement.
	Text    *string `fix:"58"`
	Trailer fix40.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetAdvId(v int)              { m.AdvId = v }
func (m *Message) SetAdvTransType(v string)    { m.AdvTransType = v }
func (m *Message) SetAdvRefID(v int)           { m.AdvRefID = &v }
func (m *Message) SetSymbol(v string)          { m.Symbol = v }
func (m *Message) SetSymbolSfx(v string)       { m.SymbolSfx = &v }
func (m *Message) SetSecurityID(v string)      { m.SecurityID = &v }
func (m *Message) SetIDSource(v string)        { m.IDSource = &v }
func (m *Message) SetIssuer(v string)          { m.Issuer = &v }
func (m *Message) SetSecurityDesc(v string)    { m.SecurityDesc = &v }
func (m *Message) SetAdvSide(v string)         { m.AdvSide = v }
func (m *Message) SetShares(v int)             { m.Shares = v }
func (m *Message) SetPrice(v float64)          { m.Price = &v }
func (m *Message) SetCurrency(v string)        { m.Currency = &v }
func (m *Message) SetTransactTime(v time.Time) { m.TransactTime = &v }
func (m *Message) SetText(v string)            { m.Text = &v }

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
	return enum.BeginStringFIX40, "7", r
}
