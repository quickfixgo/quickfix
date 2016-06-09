package allocationack

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//AllocationAck is the fix43 AllocationAck type, MsgType = P
type AllocationAck struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
}

//FromMessage creates a AllocationAck from a quickfix.Message instance
func FromMessage(m quickfix.Message) AllocationAck {
	return AllocationAck{
		Header:  fix43.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fix43.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m AllocationAck) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a AllocationAck initialized with the required fields for AllocationAck
func New(allocid field.AllocIDField, tradedate field.TradeDateField, allocstatus field.AllocStatusField) (m AllocationAck) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("P"))
	m.Set(allocid)
	m.Set(tradedate)
	m.Set(allocstatus)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg AllocationAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "P", r
}

//SetText sets Text, Tag 58
func (m AllocationAck) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m AllocationAck) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetAllocID sets AllocID, Tag 70
func (m AllocationAck) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m AllocationAck) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetAllocStatus sets AllocStatus, Tag 87
func (m AllocationAck) SetAllocStatus(v int) {
	m.Set(field.NewAllocStatus(v))
}

//SetAllocRejCode sets AllocRejCode, Tag 88
func (m AllocationAck) SetAllocRejCode(v int) {
	m.Set(field.NewAllocRejCode(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m AllocationAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m AllocationAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m AllocationAck) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegalConfirm sets LegalConfirm, Tag 650
func (m AllocationAck) SetLegalConfirm(v bool) {
	m.Set(field.NewLegalConfirm(v))
}

//GetText gets Text, Tag 58
func (m AllocationAck) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m AllocationAck) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocID gets AllocID, Tag 70
func (m AllocationAck) GetAllocID() (f field.AllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m AllocationAck) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocStatus gets AllocStatus, Tag 87
func (m AllocationAck) GetAllocStatus() (f field.AllocStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocRejCode gets AllocRejCode, Tag 88
func (m AllocationAck) GetAllocRejCode() (f field.AllocRejCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m AllocationAck) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m AllocationAck) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m AllocationAck) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegalConfirm gets LegalConfirm, Tag 650
func (m AllocationAck) GetLegalConfirm() (f field.LegalConfirmField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m AllocationAck) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m AllocationAck) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m AllocationAck) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m AllocationAck) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasAllocStatus returns true if AllocStatus is present, Tag 87
func (m AllocationAck) HasAllocStatus() bool {
	return m.Has(tag.AllocStatus)
}

//HasAllocRejCode returns true if AllocRejCode is present, Tag 88
func (m AllocationAck) HasAllocRejCode() bool {
	return m.Has(tag.AllocRejCode)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m AllocationAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m AllocationAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m AllocationAck) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasLegalConfirm returns true if LegalConfirm is present, Tag 650
func (m AllocationAck) HasLegalConfirm() bool {
	return m.Has(tag.LegalConfirm)
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v string) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v int) {
	m.Set(field.NewPartyRole(v))
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartyIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (f field.PartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (f field.PartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (f field.PartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartyIDs) GetPartySubID() (f field.PartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
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

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartyIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), quickfix.GroupElement(tag.PartySubID)})}
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
