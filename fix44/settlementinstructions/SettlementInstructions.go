//Package settlementinstructions msg type = T.
package settlementinstructions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/settlinstructionsdata"
	"time"
)

//NoSettlInst is a repeating group in SettlementInstructions
type NoSettlInst struct {
	//SettlInstID is a non-required field for NoSettlInst.
	SettlInstID *string `fix:"162"`
	//SettlInstTransType is a non-required field for NoSettlInst.
	SettlInstTransType *string `fix:"163"`
	//SettlInstRefID is a non-required field for NoSettlInst.
	SettlInstRefID *string `fix:"214"`
	//Parties Component
	parties.Parties
	//Side is a non-required field for NoSettlInst.
	Side *string `fix:"54"`
	//Product is a non-required field for NoSettlInst.
	Product *int `fix:"460"`
	//SecurityType is a non-required field for NoSettlInst.
	SecurityType *string `fix:"167"`
	//CFICode is a non-required field for NoSettlInst.
	CFICode *string `fix:"461"`
	//EffectiveTime is a non-required field for NoSettlInst.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireTime is a non-required field for NoSettlInst.
	ExpireTime *time.Time `fix:"126"`
	//LastUpdateTime is a non-required field for NoSettlInst.
	LastUpdateTime *time.Time `fix:"779"`
	//SettlInstructionsData Component
	settlinstructionsdata.SettlInstructionsData
	//PaymentMethod is a non-required field for NoSettlInst.
	PaymentMethod *int `fix:"492"`
	//PaymentRef is a non-required field for NoSettlInst.
	PaymentRef *string `fix:"476"`
	//CardHolderName is a non-required field for NoSettlInst.
	CardHolderName *string `fix:"488"`
	//CardNumber is a non-required field for NoSettlInst.
	CardNumber *string `fix:"489"`
	//CardStartDate is a non-required field for NoSettlInst.
	CardStartDate *string `fix:"503"`
	//CardExpDate is a non-required field for NoSettlInst.
	CardExpDate *string `fix:"490"`
	//CardIssNum is a non-required field for NoSettlInst.
	CardIssNum *string `fix:"491"`
	//PaymentDate is a non-required field for NoSettlInst.
	PaymentDate *string `fix:"504"`
	//PaymentRemitterID is a non-required field for NoSettlInst.
	PaymentRemitterID *string `fix:"505"`
}

func (m *NoSettlInst) SetSettlInstID(v string)        { m.SettlInstID = &v }
func (m *NoSettlInst) SetSettlInstTransType(v string) { m.SettlInstTransType = &v }
func (m *NoSettlInst) SetSettlInstRefID(v string)     { m.SettlInstRefID = &v }
func (m *NoSettlInst) SetSide(v string)               { m.Side = &v }
func (m *NoSettlInst) SetProduct(v int)               { m.Product = &v }
func (m *NoSettlInst) SetSecurityType(v string)       { m.SecurityType = &v }
func (m *NoSettlInst) SetCFICode(v string)            { m.CFICode = &v }
func (m *NoSettlInst) SetEffectiveTime(v time.Time)   { m.EffectiveTime = &v }
func (m *NoSettlInst) SetExpireTime(v time.Time)      { m.ExpireTime = &v }
func (m *NoSettlInst) SetLastUpdateTime(v time.Time)  { m.LastUpdateTime = &v }
func (m *NoSettlInst) SetPaymentMethod(v int)         { m.PaymentMethod = &v }
func (m *NoSettlInst) SetPaymentRef(v string)         { m.PaymentRef = &v }
func (m *NoSettlInst) SetCardHolderName(v string)     { m.CardHolderName = &v }
func (m *NoSettlInst) SetCardNumber(v string)         { m.CardNumber = &v }
func (m *NoSettlInst) SetCardStartDate(v string)      { m.CardStartDate = &v }
func (m *NoSettlInst) SetCardExpDate(v string)        { m.CardExpDate = &v }
func (m *NoSettlInst) SetCardIssNum(v string)         { m.CardIssNum = &v }
func (m *NoSettlInst) SetPaymentDate(v string)        { m.PaymentDate = &v }
func (m *NoSettlInst) SetPaymentRemitterID(v string)  { m.PaymentRemitterID = &v }

//Message is a SettlementInstructions FIX Message
type Message struct {
	FIXMsgType string `fix:"T"`
	fix44.Header
	//SettlInstMsgID is a required field for SettlementInstructions.
	SettlInstMsgID string `fix:"777"`
	//SettlInstReqID is a non-required field for SettlementInstructions.
	SettlInstReqID *string `fix:"791"`
	//SettlInstMode is a required field for SettlementInstructions.
	SettlInstMode string `fix:"160"`
	//SettlInstReqRejCode is a non-required field for SettlementInstructions.
	SettlInstReqRejCode *int `fix:"792"`
	//Text is a non-required field for SettlementInstructions.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SettlementInstructions.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SettlementInstructions.
	EncodedText *string `fix:"355"`
	//ClOrdID is a non-required field for SettlementInstructions.
	ClOrdID *string `fix:"11"`
	//TransactTime is a required field for SettlementInstructions.
	TransactTime time.Time `fix:"60"`
	//NoSettlInst is a non-required field for SettlementInstructions.
	NoSettlInst []NoSettlInst `fix:"778,omitempty"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetSettlInstMsgID(v string)     { m.SettlInstMsgID = v }
func (m *Message) SetSettlInstReqID(v string)     { m.SettlInstReqID = &v }
func (m *Message) SetSettlInstMode(v string)      { m.SettlInstMode = v }
func (m *Message) SetSettlInstReqRejCode(v int)   { m.SettlInstReqRejCode = &v }
func (m *Message) SetText(v string)               { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)        { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)        { m.EncodedText = &v }
func (m *Message) SetClOrdID(v string)            { m.ClOrdID = &v }
func (m *Message) SetTransactTime(v time.Time)    { m.TransactTime = v }
func (m *Message) SetNoSettlInst(v []NoSettlInst) { m.NoSettlInst = v }

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
	return enum.BeginStringFIX44, "T", r
}
