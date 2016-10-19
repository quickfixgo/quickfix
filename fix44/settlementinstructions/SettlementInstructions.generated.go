package settlementinstructions

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/tag"
)

//SettlementInstructions is the fix44 SettlementInstructions type, MsgType = T
type SettlementInstructions struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

//FromMessage creates a SettlementInstructions from a quickfix.Message instance
func FromMessage(m *quickfix.Message) SettlementInstructions {
	return SettlementInstructions{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m SettlementInstructions) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a SettlementInstructions initialized with the required fields for SettlementInstructions
func New(settlinstmsgid field.SettlInstMsgIDField, settlinstmode field.SettlInstModeField, transacttime field.TransactTimeField) (m SettlementInstructions) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("T"))
	m.Set(settlinstmsgid)
	m.Set(settlinstmode)
	m.Set(transacttime)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg SettlementInstructions, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "T", r
}

//SetClOrdID sets ClOrdID, Tag 11
func (m SettlementInstructions) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetText sets Text, Tag 58
func (m SettlementInstructions) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m SettlementInstructions) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlInstMode sets SettlInstMode, Tag 160
func (m SettlementInstructions) SetSettlInstMode(v enum.SettlInstMode) {
	m.Set(field.NewSettlInstMode(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m SettlementInstructions) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m SettlementInstructions) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetSettlInstMsgID sets SettlInstMsgID, Tag 777
func (m SettlementInstructions) SetSettlInstMsgID(v string) {
	m.Set(field.NewSettlInstMsgID(v))
}

//SetNoSettlInst sets NoSettlInst, Tag 778
func (m SettlementInstructions) SetNoSettlInst(f NoSettlInstRepeatingGroup) {
	m.SetGroup(f)
}

//SetSettlInstReqID sets SettlInstReqID, Tag 791
func (m SettlementInstructions) SetSettlInstReqID(v string) {
	m.Set(field.NewSettlInstReqID(v))
}

//SetSettlInstReqRejCode sets SettlInstReqRejCode, Tag 792
func (m SettlementInstructions) SetSettlInstReqRejCode(v enum.SettlInstReqRejCode) {
	m.Set(field.NewSettlInstReqRejCode(v))
}

//GetClOrdID gets ClOrdID, Tag 11
func (m SettlementInstructions) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m SettlementInstructions) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m SettlementInstructions) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlInstMode gets SettlInstMode, Tag 160
func (m SettlementInstructions) GetSettlInstMode() (v enum.SettlInstMode, err quickfix.MessageRejectError) {
	var f field.SettlInstModeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m SettlementInstructions) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m SettlementInstructions) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlInstMsgID gets SettlInstMsgID, Tag 777
func (m SettlementInstructions) GetSettlInstMsgID() (v string, err quickfix.MessageRejectError) {
	var f field.SettlInstMsgIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoSettlInst gets NoSettlInst, Tag 778
func (m SettlementInstructions) GetNoSettlInst() (NoSettlInstRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlInstRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSettlInstReqID gets SettlInstReqID, Tag 791
func (m SettlementInstructions) GetSettlInstReqID() (v string, err quickfix.MessageRejectError) {
	var f field.SettlInstReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlInstReqRejCode gets SettlInstReqRejCode, Tag 792
func (m SettlementInstructions) GetSettlInstReqRejCode() (v enum.SettlInstReqRejCode, err quickfix.MessageRejectError) {
	var f field.SettlInstReqRejCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m SettlementInstructions) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasText returns true if Text is present, Tag 58
func (m SettlementInstructions) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m SettlementInstructions) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlInstMode returns true if SettlInstMode is present, Tag 160
func (m SettlementInstructions) HasSettlInstMode() bool {
	return m.Has(tag.SettlInstMode)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m SettlementInstructions) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m SettlementInstructions) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasSettlInstMsgID returns true if SettlInstMsgID is present, Tag 777
func (m SettlementInstructions) HasSettlInstMsgID() bool {
	return m.Has(tag.SettlInstMsgID)
}

//HasNoSettlInst returns true if NoSettlInst is present, Tag 778
func (m SettlementInstructions) HasNoSettlInst() bool {
	return m.Has(tag.NoSettlInst)
}

//HasSettlInstReqID returns true if SettlInstReqID is present, Tag 791
func (m SettlementInstructions) HasSettlInstReqID() bool {
	return m.Has(tag.SettlInstReqID)
}

//HasSettlInstReqRejCode returns true if SettlInstReqRejCode is present, Tag 792
func (m SettlementInstructions) HasSettlInstReqRejCode() bool {
	return m.Has(tag.SettlInstReqRejCode)
}

//NoSettlInst is a repeating group element, Tag 778
type NoSettlInst struct {
	*quickfix.Group
}

//SetSettlInstID sets SettlInstID, Tag 162
func (m NoSettlInst) SetSettlInstID(v string) {
	m.Set(field.NewSettlInstID(v))
}

//SetSettlInstTransType sets SettlInstTransType, Tag 163
func (m NoSettlInst) SetSettlInstTransType(v enum.SettlInstTransType) {
	m.Set(field.NewSettlInstTransType(v))
}

//SetSettlInstRefID sets SettlInstRefID, Tag 214
func (m NoSettlInst) SetSettlInstRefID(v string) {
	m.Set(field.NewSettlInstRefID(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m NoSettlInst) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetSide sets Side, Tag 54
func (m NoSettlInst) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

//SetProduct sets Product, Tag 460
func (m NoSettlInst) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NoSettlInst) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetCFICode sets CFICode, Tag 461
func (m NoSettlInst) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetEffectiveTime sets EffectiveTime, Tag 168
func (m NoSettlInst) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m NoSettlInst) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetLastUpdateTime sets LastUpdateTime, Tag 779
func (m NoSettlInst) SetLastUpdateTime(v time.Time) {
	m.Set(field.NewLastUpdateTime(v))
}

//SetSettlDeliveryType sets SettlDeliveryType, Tag 172
func (m NoSettlInst) SetSettlDeliveryType(v enum.SettlDeliveryType) {
	m.Set(field.NewSettlDeliveryType(v))
}

//SetStandInstDbType sets StandInstDbType, Tag 169
func (m NoSettlInst) SetStandInstDbType(v enum.StandInstDbType) {
	m.Set(field.NewStandInstDbType(v))
}

//SetStandInstDbName sets StandInstDbName, Tag 170
func (m NoSettlInst) SetStandInstDbName(v string) {
	m.Set(field.NewStandInstDbName(v))
}

//SetStandInstDbID sets StandInstDbID, Tag 171
func (m NoSettlInst) SetStandInstDbID(v string) {
	m.Set(field.NewStandInstDbID(v))
}

//SetNoDlvyInst sets NoDlvyInst, Tag 85
func (m NoSettlInst) SetNoDlvyInst(f NoDlvyInstRepeatingGroup) {
	m.SetGroup(f)
}

//SetPaymentMethod sets PaymentMethod, Tag 492
func (m NoSettlInst) SetPaymentMethod(v enum.PaymentMethod) {
	m.Set(field.NewPaymentMethod(v))
}

//SetPaymentRef sets PaymentRef, Tag 476
func (m NoSettlInst) SetPaymentRef(v string) {
	m.Set(field.NewPaymentRef(v))
}

//SetCardHolderName sets CardHolderName, Tag 488
func (m NoSettlInst) SetCardHolderName(v string) {
	m.Set(field.NewCardHolderName(v))
}

//SetCardNumber sets CardNumber, Tag 489
func (m NoSettlInst) SetCardNumber(v string) {
	m.Set(field.NewCardNumber(v))
}

//SetCardStartDate sets CardStartDate, Tag 503
func (m NoSettlInst) SetCardStartDate(v string) {
	m.Set(field.NewCardStartDate(v))
}

//SetCardExpDate sets CardExpDate, Tag 490
func (m NoSettlInst) SetCardExpDate(v string) {
	m.Set(field.NewCardExpDate(v))
}

//SetCardIssNum sets CardIssNum, Tag 491
func (m NoSettlInst) SetCardIssNum(v string) {
	m.Set(field.NewCardIssNum(v))
}

//SetPaymentDate sets PaymentDate, Tag 504
func (m NoSettlInst) SetPaymentDate(v string) {
	m.Set(field.NewPaymentDate(v))
}

//SetPaymentRemitterID sets PaymentRemitterID, Tag 505
func (m NoSettlInst) SetPaymentRemitterID(v string) {
	m.Set(field.NewPaymentRemitterID(v))
}

//GetSettlInstID gets SettlInstID, Tag 162
func (m NoSettlInst) GetSettlInstID() (v string, err quickfix.MessageRejectError) {
	var f field.SettlInstIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlInstTransType gets SettlInstTransType, Tag 163
func (m NoSettlInst) GetSettlInstTransType() (v enum.SettlInstTransType, err quickfix.MessageRejectError) {
	var f field.SettlInstTransTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlInstRefID gets SettlInstRefID, Tag 214
func (m NoSettlInst) GetSettlInstRefID() (v string, err quickfix.MessageRejectError) {
	var f field.SettlInstRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m NoSettlInst) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSide gets Side, Tag 54
func (m NoSettlInst) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetProduct gets Product, Tag 460
func (m NoSettlInst) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m NoSettlInst) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCFICode gets CFICode, Tag 461
func (m NoSettlInst) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEffectiveTime gets EffectiveTime, Tag 168
func (m NoSettlInst) GetEffectiveTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.EffectiveTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m NoSettlInst) GetExpireTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ExpireTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastUpdateTime gets LastUpdateTime, Tag 779
func (m NoSettlInst) GetLastUpdateTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.LastUpdateTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlDeliveryType gets SettlDeliveryType, Tag 172
func (m NoSettlInst) GetSettlDeliveryType() (v enum.SettlDeliveryType, err quickfix.MessageRejectError) {
	var f field.SettlDeliveryTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStandInstDbType gets StandInstDbType, Tag 169
func (m NoSettlInst) GetStandInstDbType() (v enum.StandInstDbType, err quickfix.MessageRejectError) {
	var f field.StandInstDbTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStandInstDbName gets StandInstDbName, Tag 170
func (m NoSettlInst) GetStandInstDbName() (v string, err quickfix.MessageRejectError) {
	var f field.StandInstDbNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStandInstDbID gets StandInstDbID, Tag 171
func (m NoSettlInst) GetStandInstDbID() (v string, err quickfix.MessageRejectError) {
	var f field.StandInstDbIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoDlvyInst gets NoDlvyInst, Tag 85
func (m NoSettlInst) GetNoDlvyInst() (NoDlvyInstRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDlvyInstRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetPaymentMethod gets PaymentMethod, Tag 492
func (m NoSettlInst) GetPaymentMethod() (v enum.PaymentMethod, err quickfix.MessageRejectError) {
	var f field.PaymentMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPaymentRef gets PaymentRef, Tag 476
func (m NoSettlInst) GetPaymentRef() (v string, err quickfix.MessageRejectError) {
	var f field.PaymentRefField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCardHolderName gets CardHolderName, Tag 488
func (m NoSettlInst) GetCardHolderName() (v string, err quickfix.MessageRejectError) {
	var f field.CardHolderNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCardNumber gets CardNumber, Tag 489
func (m NoSettlInst) GetCardNumber() (v string, err quickfix.MessageRejectError) {
	var f field.CardNumberField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCardStartDate gets CardStartDate, Tag 503
func (m NoSettlInst) GetCardStartDate() (v string, err quickfix.MessageRejectError) {
	var f field.CardStartDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCardExpDate gets CardExpDate, Tag 490
func (m NoSettlInst) GetCardExpDate() (v string, err quickfix.MessageRejectError) {
	var f field.CardExpDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCardIssNum gets CardIssNum, Tag 491
func (m NoSettlInst) GetCardIssNum() (v string, err quickfix.MessageRejectError) {
	var f field.CardIssNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPaymentDate gets PaymentDate, Tag 504
func (m NoSettlInst) GetPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.PaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPaymentRemitterID gets PaymentRemitterID, Tag 505
func (m NoSettlInst) GetPaymentRemitterID() (v string, err quickfix.MessageRejectError) {
	var f field.PaymentRemitterIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasSettlInstID returns true if SettlInstID is present, Tag 162
func (m NoSettlInst) HasSettlInstID() bool {
	return m.Has(tag.SettlInstID)
}

//HasSettlInstTransType returns true if SettlInstTransType is present, Tag 163
func (m NoSettlInst) HasSettlInstTransType() bool {
	return m.Has(tag.SettlInstTransType)
}

//HasSettlInstRefID returns true if SettlInstRefID is present, Tag 214
func (m NoSettlInst) HasSettlInstRefID() bool {
	return m.Has(tag.SettlInstRefID)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m NoSettlInst) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasSide returns true if Side is present, Tag 54
func (m NoSettlInst) HasSide() bool {
	return m.Has(tag.Side)
}

//HasProduct returns true if Product is present, Tag 460
func (m NoSettlInst) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NoSettlInst) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m NoSettlInst) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasEffectiveTime returns true if EffectiveTime is present, Tag 168
func (m NoSettlInst) HasEffectiveTime() bool {
	return m.Has(tag.EffectiveTime)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m NoSettlInst) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasLastUpdateTime returns true if LastUpdateTime is present, Tag 779
func (m NoSettlInst) HasLastUpdateTime() bool {
	return m.Has(tag.LastUpdateTime)
}

//HasSettlDeliveryType returns true if SettlDeliveryType is present, Tag 172
func (m NoSettlInst) HasSettlDeliveryType() bool {
	return m.Has(tag.SettlDeliveryType)
}

//HasStandInstDbType returns true if StandInstDbType is present, Tag 169
func (m NoSettlInst) HasStandInstDbType() bool {
	return m.Has(tag.StandInstDbType)
}

//HasStandInstDbName returns true if StandInstDbName is present, Tag 170
func (m NoSettlInst) HasStandInstDbName() bool {
	return m.Has(tag.StandInstDbName)
}

//HasStandInstDbID returns true if StandInstDbID is present, Tag 171
func (m NoSettlInst) HasStandInstDbID() bool {
	return m.Has(tag.StandInstDbID)
}

//HasNoDlvyInst returns true if NoDlvyInst is present, Tag 85
func (m NoSettlInst) HasNoDlvyInst() bool {
	return m.Has(tag.NoDlvyInst)
}

//HasPaymentMethod returns true if PaymentMethod is present, Tag 492
func (m NoSettlInst) HasPaymentMethod() bool {
	return m.Has(tag.PaymentMethod)
}

//HasPaymentRef returns true if PaymentRef is present, Tag 476
func (m NoSettlInst) HasPaymentRef() bool {
	return m.Has(tag.PaymentRef)
}

//HasCardHolderName returns true if CardHolderName is present, Tag 488
func (m NoSettlInst) HasCardHolderName() bool {
	return m.Has(tag.CardHolderName)
}

//HasCardNumber returns true if CardNumber is present, Tag 489
func (m NoSettlInst) HasCardNumber() bool {
	return m.Has(tag.CardNumber)
}

//HasCardStartDate returns true if CardStartDate is present, Tag 503
func (m NoSettlInst) HasCardStartDate() bool {
	return m.Has(tag.CardStartDate)
}

//HasCardExpDate returns true if CardExpDate is present, Tag 490
func (m NoSettlInst) HasCardExpDate() bool {
	return m.Has(tag.CardExpDate)
}

//HasCardIssNum returns true if CardIssNum is present, Tag 491
func (m NoSettlInst) HasCardIssNum() bool {
	return m.Has(tag.CardIssNum)
}

//HasPaymentDate returns true if PaymentDate is present, Tag 504
func (m NoSettlInst) HasPaymentDate() bool {
	return m.Has(tag.PaymentDate)
}

//HasPaymentRemitterID returns true if PaymentRemitterID is present, Tag 505
func (m NoSettlInst) HasPaymentRemitterID() bool {
	return m.Has(tag.PaymentRemitterID)
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	*quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

//SetNoPartySubIDs sets NoPartySubIDs, Tag 802
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartySubIDs gets NoPartySubIDs, Tag 802
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasPartyID returns true if PartyID is present, Tag 448
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

//HasPartyIDSource returns true if PartyIDSource is present, Tag 447
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

//HasPartyRole returns true if PartyRole is present, Tag 452
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

//HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

//NoPartySubIDs is a repeating group element, Tag 802
type NoPartySubIDs struct {
	*quickfix.Group
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//SetPartySubIDType sets PartySubIDType, Tag 803
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartySubIDType gets PartySubIDType, Tag 803
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//HasPartySubIDType returns true if PartySubIDType is present, Tag 803
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

//NoPartySubIDsRepeatingGroup is a repeating group, Tag 802
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)})}
}

//Add create and append a new NoPartySubIDs to this group
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

//Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoPartyIDs to this group
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

//Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoDlvyInst is a repeating group element, Tag 85
type NoDlvyInst struct {
	*quickfix.Group
}

//SetSettlInstSource sets SettlInstSource, Tag 165
func (m NoDlvyInst) SetSettlInstSource(v enum.SettlInstSource) {
	m.Set(field.NewSettlInstSource(v))
}

//SetDlvyInstType sets DlvyInstType, Tag 787
func (m NoDlvyInst) SetDlvyInstType(v enum.DlvyInstType) {
	m.Set(field.NewDlvyInstType(v))
}

//SetNoSettlPartyIDs sets NoSettlPartyIDs, Tag 781
func (m NoDlvyInst) SetNoSettlPartyIDs(f NoSettlPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetSettlInstSource gets SettlInstSource, Tag 165
func (m NoDlvyInst) GetSettlInstSource() (v enum.SettlInstSource, err quickfix.MessageRejectError) {
	var f field.SettlInstSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDlvyInstType gets DlvyInstType, Tag 787
func (m NoDlvyInst) GetDlvyInstType() (v enum.DlvyInstType, err quickfix.MessageRejectError) {
	var f field.DlvyInstTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoSettlPartyIDs gets NoSettlPartyIDs, Tag 781
func (m NoDlvyInst) GetNoSettlPartyIDs() (NoSettlPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasSettlInstSource returns true if SettlInstSource is present, Tag 165
func (m NoDlvyInst) HasSettlInstSource() bool {
	return m.Has(tag.SettlInstSource)
}

//HasDlvyInstType returns true if DlvyInstType is present, Tag 787
func (m NoDlvyInst) HasDlvyInstType() bool {
	return m.Has(tag.DlvyInstType)
}

//HasNoSettlPartyIDs returns true if NoSettlPartyIDs is present, Tag 781
func (m NoDlvyInst) HasNoSettlPartyIDs() bool {
	return m.Has(tag.NoSettlPartyIDs)
}

//NoSettlPartyIDs is a repeating group element, Tag 781
type NoSettlPartyIDs struct {
	*quickfix.Group
}

//SetSettlPartyID sets SettlPartyID, Tag 782
func (m NoSettlPartyIDs) SetSettlPartyID(v string) {
	m.Set(field.NewSettlPartyID(v))
}

//SetSettlPartyIDSource sets SettlPartyIDSource, Tag 783
func (m NoSettlPartyIDs) SetSettlPartyIDSource(v string) {
	m.Set(field.NewSettlPartyIDSource(v))
}

//SetSettlPartyRole sets SettlPartyRole, Tag 784
func (m NoSettlPartyIDs) SetSettlPartyRole(v int) {
	m.Set(field.NewSettlPartyRole(v))
}

//SetNoSettlPartySubIDs sets NoSettlPartySubIDs, Tag 801
func (m NoSettlPartyIDs) SetNoSettlPartySubIDs(f NoSettlPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetSettlPartyID gets SettlPartyID, Tag 782
func (m NoSettlPartyIDs) GetSettlPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.SettlPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlPartyIDSource gets SettlPartyIDSource, Tag 783
func (m NoSettlPartyIDs) GetSettlPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.SettlPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlPartyRole gets SettlPartyRole, Tag 784
func (m NoSettlPartyIDs) GetSettlPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.SettlPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoSettlPartySubIDs gets NoSettlPartySubIDs, Tag 801
func (m NoSettlPartyIDs) GetNoSettlPartySubIDs() (NoSettlPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasSettlPartyID returns true if SettlPartyID is present, Tag 782
func (m NoSettlPartyIDs) HasSettlPartyID() bool {
	return m.Has(tag.SettlPartyID)
}

//HasSettlPartyIDSource returns true if SettlPartyIDSource is present, Tag 783
func (m NoSettlPartyIDs) HasSettlPartyIDSource() bool {
	return m.Has(tag.SettlPartyIDSource)
}

//HasSettlPartyRole returns true if SettlPartyRole is present, Tag 784
func (m NoSettlPartyIDs) HasSettlPartyRole() bool {
	return m.Has(tag.SettlPartyRole)
}

//HasNoSettlPartySubIDs returns true if NoSettlPartySubIDs is present, Tag 801
func (m NoSettlPartyIDs) HasNoSettlPartySubIDs() bool {
	return m.Has(tag.NoSettlPartySubIDs)
}

//NoSettlPartySubIDs is a repeating group element, Tag 801
type NoSettlPartySubIDs struct {
	*quickfix.Group
}

//SetSettlPartySubID sets SettlPartySubID, Tag 785
func (m NoSettlPartySubIDs) SetSettlPartySubID(v string) {
	m.Set(field.NewSettlPartySubID(v))
}

//SetSettlPartySubIDType sets SettlPartySubIDType, Tag 786
func (m NoSettlPartySubIDs) SetSettlPartySubIDType(v int) {
	m.Set(field.NewSettlPartySubIDType(v))
}

//GetSettlPartySubID gets SettlPartySubID, Tag 785
func (m NoSettlPartySubIDs) GetSettlPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.SettlPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlPartySubIDType gets SettlPartySubIDType, Tag 786
func (m NoSettlPartySubIDs) GetSettlPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.SettlPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasSettlPartySubID returns true if SettlPartySubID is present, Tag 785
func (m NoSettlPartySubIDs) HasSettlPartySubID() bool {
	return m.Has(tag.SettlPartySubID)
}

//HasSettlPartySubIDType returns true if SettlPartySubIDType is present, Tag 786
func (m NoSettlPartySubIDs) HasSettlPartySubIDType() bool {
	return m.Has(tag.SettlPartySubIDType)
}

//NoSettlPartySubIDsRepeatingGroup is a repeating group, Tag 801
type NoSettlPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSettlPartySubIDsRepeatingGroup returns an initialized, NoSettlPartySubIDsRepeatingGroup
func NewNoSettlPartySubIDsRepeatingGroup() NoSettlPartySubIDsRepeatingGroup {
	return NoSettlPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSettlPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlPartySubID), quickfix.GroupElement(tag.SettlPartySubIDType)})}
}

//Add create and append a new NoSettlPartySubIDs to this group
func (m NoSettlPartySubIDsRepeatingGroup) Add() NoSettlPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoSettlPartySubIDs{g}
}

//Get returns the ith NoSettlPartySubIDs in the NoSettlPartySubIDsRepeatinGroup
func (m NoSettlPartySubIDsRepeatingGroup) Get(i int) NoSettlPartySubIDs {
	return NoSettlPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoSettlPartyIDsRepeatingGroup is a repeating group, Tag 781
type NoSettlPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSettlPartyIDsRepeatingGroup returns an initialized, NoSettlPartyIDsRepeatingGroup
func NewNoSettlPartyIDsRepeatingGroup() NoSettlPartyIDsRepeatingGroup {
	return NoSettlPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSettlPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlPartyID), quickfix.GroupElement(tag.SettlPartyIDSource), quickfix.GroupElement(tag.SettlPartyRole), NewNoSettlPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoSettlPartyIDs to this group
func (m NoSettlPartyIDsRepeatingGroup) Add() NoSettlPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoSettlPartyIDs{g}
}

//Get returns the ith NoSettlPartyIDs in the NoSettlPartyIDsRepeatinGroup
func (m NoSettlPartyIDsRepeatingGroup) Get(i int) NoSettlPartyIDs {
	return NoSettlPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoDlvyInstRepeatingGroup is a repeating group, Tag 85
type NoDlvyInstRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoDlvyInstRepeatingGroup returns an initialized, NoDlvyInstRepeatingGroup
func NewNoDlvyInstRepeatingGroup() NoDlvyInstRepeatingGroup {
	return NoDlvyInstRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoDlvyInst,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlInstSource), quickfix.GroupElement(tag.DlvyInstType), NewNoSettlPartyIDsRepeatingGroup()})}
}

//Add create and append a new NoDlvyInst to this group
func (m NoDlvyInstRepeatingGroup) Add() NoDlvyInst {
	g := m.RepeatingGroup.Add()
	return NoDlvyInst{g}
}

//Get returns the ith NoDlvyInst in the NoDlvyInstRepeatinGroup
func (m NoDlvyInstRepeatingGroup) Get(i int) NoDlvyInst {
	return NoDlvyInst{m.RepeatingGroup.Get(i)}
}

//NoSettlInstRepeatingGroup is a repeating group, Tag 778
type NoSettlInstRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSettlInstRepeatingGroup returns an initialized, NoSettlInstRepeatingGroup
func NewNoSettlInstRepeatingGroup() NoSettlInstRepeatingGroup {
	return NoSettlInstRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSettlInst,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlInstID), quickfix.GroupElement(tag.SettlInstTransType), quickfix.GroupElement(tag.SettlInstRefID), NewNoPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.Product), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.CFICode), quickfix.GroupElement(tag.EffectiveTime), quickfix.GroupElement(tag.ExpireTime), quickfix.GroupElement(tag.LastUpdateTime), quickfix.GroupElement(tag.SettlDeliveryType), quickfix.GroupElement(tag.StandInstDbType), quickfix.GroupElement(tag.StandInstDbName), quickfix.GroupElement(tag.StandInstDbID), NewNoDlvyInstRepeatingGroup(), quickfix.GroupElement(tag.PaymentMethod), quickfix.GroupElement(tag.PaymentRef), quickfix.GroupElement(tag.CardHolderName), quickfix.GroupElement(tag.CardNumber), quickfix.GroupElement(tag.CardStartDate), quickfix.GroupElement(tag.CardExpDate), quickfix.GroupElement(tag.CardIssNum), quickfix.GroupElement(tag.PaymentDate), quickfix.GroupElement(tag.PaymentRemitterID)})}
}

//Add create and append a new NoSettlInst to this group
func (m NoSettlInstRepeatingGroup) Add() NoSettlInst {
	g := m.RepeatingGroup.Add()
	return NoSettlInst{g}
}

//Get returns the ith NoSettlInst in the NoSettlInstRepeatinGroup
func (m NoSettlInstRepeatingGroup) Get(i int) NoSettlInst {
	return NoSettlInst{m.RepeatingGroup.Get(i)}
}
