package networkcounterpartysystemstatusrequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/tag"
)

//NetworkCounterpartySystemStatusRequest is the fix44 NetworkCounterpartySystemStatusRequest type, MsgType = BC
type NetworkCounterpartySystemStatusRequest struct {
	fix44.Header
	quickfix.Body
	fix44.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a NetworkCounterpartySystemStatusRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) NetworkCounterpartySystemStatusRequest {
	return NetworkCounterpartySystemStatusRequest{
		Header:      fix44.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix44.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m NetworkCounterpartySystemStatusRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a NetworkCounterpartySystemStatusRequest initialized with the required fields for NetworkCounterpartySystemStatusRequest
func New(networkrequesttype field.NetworkRequestTypeField, networkrequestid field.NetworkRequestIDField) (m NetworkCounterpartySystemStatusRequest) {
	m.Header = fix44.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("BC"))
	m.Header.Set(field.NewBeginString("FIX.4.4"))
	m.Set(networkrequesttype)
	m.Set(networkrequestid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg NetworkCounterpartySystemStatusRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "BC", r
}

//SetNetworkRequestID sets NetworkRequestID, Tag 933
func (m NetworkCounterpartySystemStatusRequest) SetNetworkRequestID(v string) {
	m.Set(field.NewNetworkRequestID(v))
}

//SetNetworkRequestType sets NetworkRequestType, Tag 935
func (m NetworkCounterpartySystemStatusRequest) SetNetworkRequestType(v int) {
	m.Set(field.NewNetworkRequestType(v))
}

//SetNoCompIDs sets NoCompIDs, Tag 936
func (m NetworkCounterpartySystemStatusRequest) SetNoCompIDs(f NoCompIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetNetworkRequestID gets NetworkRequestID, Tag 933
func (m NetworkCounterpartySystemStatusRequest) GetNetworkRequestID() (f field.NetworkRequestIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNetworkRequestType gets NetworkRequestType, Tag 935
func (m NetworkCounterpartySystemStatusRequest) GetNetworkRequestType() (f field.NetworkRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoCompIDs gets NoCompIDs, Tag 936
func (m NetworkCounterpartySystemStatusRequest) GetNoCompIDs() (NoCompIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoCompIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNetworkRequestID returns true if NetworkRequestID is present, Tag 933
func (m NetworkCounterpartySystemStatusRequest) HasNetworkRequestID() bool {
	return m.Has(tag.NetworkRequestID)
}

//HasNetworkRequestType returns true if NetworkRequestType is present, Tag 935
func (m NetworkCounterpartySystemStatusRequest) HasNetworkRequestType() bool {
	return m.Has(tag.NetworkRequestType)
}

//HasNoCompIDs returns true if NoCompIDs is present, Tag 936
func (m NetworkCounterpartySystemStatusRequest) HasNoCompIDs() bool {
	return m.Has(tag.NoCompIDs)
}

//NoCompIDs is a repeating group element, Tag 936
type NoCompIDs struct {
	quickfix.Group
}

//SetRefCompID sets RefCompID, Tag 930
func (m NoCompIDs) SetRefCompID(v string) {
	m.Set(field.NewRefCompID(v))
}

//SetRefSubID sets RefSubID, Tag 931
func (m NoCompIDs) SetRefSubID(v string) {
	m.Set(field.NewRefSubID(v))
}

//SetLocationID sets LocationID, Tag 283
func (m NoCompIDs) SetLocationID(v string) {
	m.Set(field.NewLocationID(v))
}

//SetDeskID sets DeskID, Tag 284
func (m NoCompIDs) SetDeskID(v string) {
	m.Set(field.NewDeskID(v))
}

//GetRefCompID gets RefCompID, Tag 930
func (m NoCompIDs) GetRefCompID() (f field.RefCompIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRefSubID gets RefSubID, Tag 931
func (m NoCompIDs) GetRefSubID() (f field.RefSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocationID gets LocationID, Tag 283
func (m NoCompIDs) GetLocationID() (f field.LocationIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeskID gets DeskID, Tag 284
func (m NoCompIDs) GetDeskID() (f field.DeskIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasRefCompID returns true if RefCompID is present, Tag 930
func (m NoCompIDs) HasRefCompID() bool {
	return m.Has(tag.RefCompID)
}

//HasRefSubID returns true if RefSubID is present, Tag 931
func (m NoCompIDs) HasRefSubID() bool {
	return m.Has(tag.RefSubID)
}

//HasLocationID returns true if LocationID is present, Tag 283
func (m NoCompIDs) HasLocationID() bool {
	return m.Has(tag.LocationID)
}

//HasDeskID returns true if DeskID is present, Tag 284
func (m NoCompIDs) HasDeskID() bool {
	return m.Has(tag.DeskID)
}

//NoCompIDsRepeatingGroup is a repeating group, Tag 936
type NoCompIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoCompIDsRepeatingGroup returns an initialized, NoCompIDsRepeatingGroup
func NewNoCompIDsRepeatingGroup() NoCompIDsRepeatingGroup {
	return NoCompIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoCompIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RefCompID), quickfix.GroupElement(tag.RefSubID), quickfix.GroupElement(tag.LocationID), quickfix.GroupElement(tag.DeskID)})}
}

//Add create and append a new NoCompIDs to this group
func (m NoCompIDsRepeatingGroup) Add() NoCompIDs {
	g := m.RepeatingGroup.Add()
	return NoCompIDs{g}
}

//Get returns the ith NoCompIDs in the NoCompIDsRepeatinGroup
func (m NoCompIDsRepeatingGroup) Get(i int) NoCompIDs {
	return NoCompIDs{m.RepeatingGroup.Get(i)}
}
