package datadictionary

import (
	"encoding/xml"
	"reflect"
	"testing"
)

var cachedXMLDoc *XMLDoc

func xmlDoc() (*XMLDoc, error) {
	if cachedXMLDoc != nil {
		return cachedXMLDoc, nil
	}

	data := `
		<fix major='4' type='FIX' servicepack='0' minor='3'>
			<header>
				<field name='BeginString' required='Y' />
				<group name='NoHops' required='N'>
				 <field name='HopCompID' required='N' />
				 <field name='HopSendingTime' required='N' />
				 <field name='HopRefID' required='N' />
				</group>
			</header>
		 <messages>
			<message name='Heartbeat' msgcat='admin' msgtype='0'>
			 <field name='TestReqID' required='N' />
			</message>
  <message name='IOI' msgcat='app' msgtype='6'>
		 <field name='IOIid' required='Y' />
		 <field name='IOITransType' required='Y' />
		 <field name='IOIRefID' required='N' />
		 <component name='Instrument' required='Y' />
   <group name='NoRoutingIDs' required='N'>
    <field name='RoutingType' required='N' />
    <field name='RoutingID' required='N' />
   </group>

	 </message>
  <message name='NewOrderSingle' msgcat='app' msgtype='D'>
   <field name='ClOrdID' required='Y' />
   <field name='SecondaryClOrdID' required='N' />
   <field name='ClOrdLinkID' required='N' />
   <component name='Parties' required='N' />
   <field name='TradeOriginationDate' required='N' />
   <field name='Account' required='N' />
   <field name='AccountType' required='N' />
   <field name='DayBookingInst' required='N' />
   <field name='BookingUnit' required='N' />
   <field name='PreallocMethod' required='N' />
   <group name='NoAllocs' required='N'>
    <field name='AllocAccount' required='N' />
    <field name='IndividualAllocID' required='N' />
    <component name='NestedParties' required='N' />
    <field name='AllocQty' required='N' />
   </group>
   <field name='SettlmntTyp' required='N' />
   <field name='FutSettDate' required='N' />
   <field name='CashMargin' required='N' />
   <field name='ClearingFeeIndicator' required='N' />
   <field name='HandlInst' required='Y' />
   <field name='ExecInst' required='N' />
   <field name='MinQty' required='N' />
   <field name='MaxFloor' required='N' />
   <field name='ExDestination' required='N' />
   <group name='NoTradingSessions' required='N'>
    <field name='TradingSessionID' required='N' />
    <field name='TradingSessionSubID' required='N' />
   </group>
	 </message>


			</messages>

			<trailer>
				<field name='SignatureLength' required='N' />
				<field name='Signature' required='N' />
				<field name='CheckSum' required='Y' />
			</trailer>
		</fix>
	`

	cachedXMLDoc = new(XMLDoc)
	err := xml.Unmarshal([]byte(data), cachedXMLDoc)

	return cachedXMLDoc, err
}

func TestBoilerPlate(t *testing.T) {
	doc, err := xmlDoc()

	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	var tests = []struct {
		Value         interface{}
		ExpectedValue interface{}
	}{
		{doc.Type, "FIX"},
		{doc.Major, "4"},
		{doc.Minor, "3"},
		{doc.ServicePack, 0},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(test.Value, test.ExpectedValue) {
			t.Errorf("Expected %v got %v", test.ExpectedValue, test.Value)
		}
	}
}

func TestComponentMembers(t *testing.T) {
	doc, err := xmlDoc()

	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if doc.Header == nil {
		t.Fatal("Header is nil")
	}

	var tests = []struct {
		Value        *XMLComponentMember
		XMLNameLocal string
		Name         string
		Required     bool
	}{
		{doc.Header.Members[0], "field", "BeginString", true},
		{doc.Header.Members[1], "group", "NoHops", false},
		{doc.Header.Members[1].Members[0], "field", "HopCompID", false},
		{doc.Trailer.Members[0], "field", "SignatureLength", false},
		{doc.Messages[0].Members[0], "field", "TestReqID", false},
		{doc.Messages[1].Members[3], "component", "Instrument", true},
		{doc.Messages[1].Members[4], "group", "NoRoutingIDs", false},
		{doc.Messages[1].Members[4].Members[0], "field", "RoutingType", false},
		{doc.Messages[1].Members[4].Members[1], "field", "RoutingID", false},
	}

	for _, test := range tests {
		if test.Value.XMLName.Local != test.XMLNameLocal {
			t.Errorf("%v: Expected XMLName Local %v got %v", test.Name, test.XMLNameLocal, test.Value.XMLName.Local)
		}

		if test.Value.Name != test.Name {
			t.Errorf("%v: Expected Name %v got %v", test.Name, test.Name, test.Value.Name)
		}

		if test.Value.isRequired() != test.Required {
			t.Errorf("%v: Expected Required %v got %v", test.Name, test.Required, test.Value.isRequired())
		}
	}
}

func TestMessages(t *testing.T) {
	doc, err := xmlDoc()
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	var tests = []struct {
		Value   *XMLComponent
		Name    string
		MsgCat  string
		MsgType string
	}{
		{doc.Messages[0], "Heartbeat", "admin", "0"},
	}

	for _, test := range tests {
		if test.Value.Name != test.Name {
			t.Errorf("Expected Name %v got %v", test.Name, test.Value.Name)
		}

		if test.Value.MsgCat != test.MsgCat {
			t.Errorf("Expected MsgCat %v got %v", test.MsgCat, test.Value.MsgCat)
		}

		if test.Value.MsgType != test.MsgType {
			t.Errorf("Expected MsgType %v got %v", test.MsgType, test.Value.MsgType)
		}

	}
}
