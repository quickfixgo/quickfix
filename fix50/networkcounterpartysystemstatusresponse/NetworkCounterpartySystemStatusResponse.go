package networkcounterpartysystemstatusresponse

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//NetworkCounterpartySystemStatusResponse is the fix50 NetworkCounterpartySystemStatusResponse type, MsgType = BD
type NetworkCounterpartySystemStatusResponse struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a NetworkCounterpartySystemStatusResponse from a quickfix.Message instance
func FromMessage(m quickfix.Message) NetworkCounterpartySystemStatusResponse {
	return NetworkCounterpartySystemStatusResponse{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m NetworkCounterpartySystemStatusResponse) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a NetworkCounterpartySystemStatusResponse initialized with the required fields for NetworkCounterpartySystemStatusResponse
func New(networkstatusresponsetype field.NetworkStatusResponseTypeField, networkresponseid field.NetworkResponseIDField) (m NetworkCounterpartySystemStatusResponse) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("BD"))
	m.Header.Set(field.NewBeginString("7"))
	m.Set(networkstatusresponsetype)
	m.Set(networkresponseid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg NetworkCounterpartySystemStatusResponse, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "BD", r
}

//SetNetworkResponseID sets NetworkResponseID, Tag 932
func (m NetworkCounterpartySystemStatusResponse) SetNetworkResponseID(v string) {
	m.Set(field.NewNetworkResponseID(v))
}

//SetNetworkRequestID sets NetworkRequestID, Tag 933
func (m NetworkCounterpartySystemStatusResponse) SetNetworkRequestID(v string) {
	m.Set(field.NewNetworkRequestID(v))
}

//SetLastNetworkResponseID sets LastNetworkResponseID, Tag 934
func (m NetworkCounterpartySystemStatusResponse) SetLastNetworkResponseID(v string) {
	m.Set(field.NewLastNetworkResponseID(v))
}

//SetNoCompIDs sets NoCompIDs, Tag 936
func (m NetworkCounterpartySystemStatusResponse) SetNoCompIDs(f NoCompIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNetworkStatusResponseType sets NetworkStatusResponseType, Tag 937
func (m NetworkCounterpartySystemStatusResponse) SetNetworkStatusResponseType(v int) {
	m.Set(field.NewNetworkStatusResponseType(v))
}

//GetNetworkResponseID gets NetworkResponseID, Tag 932
func (m NetworkCounterpartySystemStatusResponse) GetNetworkResponseID() (f field.NetworkResponseIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNetworkRequestID gets NetworkRequestID, Tag 933
func (m NetworkCounterpartySystemStatusResponse) GetNetworkRequestID() (f field.NetworkRequestIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastNetworkResponseID gets LastNetworkResponseID, Tag 934
func (m NetworkCounterpartySystemStatusResponse) GetLastNetworkResponseID() (f field.LastNetworkResponseIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoCompIDs gets NoCompIDs, Tag 936
func (m NetworkCounterpartySystemStatusResponse) GetNoCompIDs() (NoCompIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoCompIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNetworkStatusResponseType gets NetworkStatusResponseType, Tag 937
func (m NetworkCounterpartySystemStatusResponse) GetNetworkStatusResponseType() (f field.NetworkStatusResponseTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasNetworkResponseID returns true if NetworkResponseID is present, Tag 932
func (m NetworkCounterpartySystemStatusResponse) HasNetworkResponseID() bool {
	return m.Has(tag.NetworkResponseID)
}

//HasNetworkRequestID returns true if NetworkRequestID is present, Tag 933
func (m NetworkCounterpartySystemStatusResponse) HasNetworkRequestID() bool {
	return m.Has(tag.NetworkRequestID)
}

//HasLastNetworkResponseID returns true if LastNetworkResponseID is present, Tag 934
func (m NetworkCounterpartySystemStatusResponse) HasLastNetworkResponseID() bool {
	return m.Has(tag.LastNetworkResponseID)
}

//HasNoCompIDs returns true if NoCompIDs is present, Tag 936
func (m NetworkCounterpartySystemStatusResponse) HasNoCompIDs() bool {
	return m.Has(tag.NoCompIDs)
}

//HasNetworkStatusResponseType returns true if NetworkStatusResponseType is present, Tag 937
func (m NetworkCounterpartySystemStatusResponse) HasNetworkStatusResponseType() bool {
	return m.Has(tag.NetworkStatusResponseType)
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

//SetStatusValue sets StatusValue, Tag 928
func (m NoCompIDs) SetStatusValue(v int) {
	m.Set(field.NewStatusValue(v))
}

//SetStatusText sets StatusText, Tag 929
func (m NoCompIDs) SetStatusText(v string) {
	m.Set(field.NewStatusText(v))
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

//GetStatusValue gets StatusValue, Tag 928
func (m NoCompIDs) GetStatusValue() (f field.StatusValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStatusText gets StatusText, Tag 929
func (m NoCompIDs) GetStatusText() (f field.StatusTextField, err quickfix.MessageRejectError) {
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

//HasStatusValue returns true if StatusValue is present, Tag 928
func (m NoCompIDs) HasStatusValue() bool {
	return m.Has(tag.StatusValue)
}

//HasStatusText returns true if StatusText is present, Tag 929
func (m NoCompIDs) HasStatusText() bool {
	return m.Has(tag.StatusText)
}

//NoCompIDsRepeatingGroup is a repeating group, Tag 936
type NoCompIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoCompIDsRepeatingGroup returns an initialized, NoCompIDsRepeatingGroup
func NewNoCompIDsRepeatingGroup() NoCompIDsRepeatingGroup {
	return NoCompIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoCompIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RefCompID), quickfix.GroupElement(tag.RefSubID), quickfix.GroupElement(tag.LocationID), quickfix.GroupElement(tag.DeskID), quickfix.GroupElement(tag.StatusValue), quickfix.GroupElement(tag.StatusText)})}
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
