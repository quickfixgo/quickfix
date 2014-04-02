package datadictionary

import (
	"encoding/xml"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type XMLTests struct {
	doc *XMLDoc
}

var _ = Suite(&XMLTests{})

func (s *XMLTests) SetUpTest(c *C) {
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

	s.doc = new(XMLDoc)
	err := xml.Unmarshal([]byte(data), s.doc)
	c.Check(err, IsNil)
}

func (s *XMLTests) TestBoilerPlate(c *C) {
	c.Check(s.doc.Type, Equals, "FIX")
	c.Check(s.doc.Major, Equals, 4)
	c.Check(s.doc.Minor, Equals, 3)
	c.Check(s.doc.ServicePack, Equals, 0)
}

func (s *XMLTests) TestHeader(c *C) {
	c.Check(s.doc.Header, NotNil)
	BeginString := s.doc.Header.Members[0]
	c.Check(BeginString.XMLName.Local, Equals, "field")
	c.Check(BeginString.Name, Equals, "BeginString")
	c.Check(BeginString.Required, Equals, "Y")

	NoHops := s.doc.Header.Members[1]
	c.Check(NoHops.XMLName.Local, Equals, "group")
	c.Check(NoHops.Name, Equals, "NoHops")
	c.Check(NoHops.Required, Equals, "N")

	HopCompID := NoHops.Members[0]
	c.Check(HopCompID.XMLName.Local, Equals, "field")
	c.Check(HopCompID.Name, Equals, "HopCompID")
	c.Check(HopCompID.Required, Equals, "N")
}

func (s *XMLTests) TestTrailer(c *C) {
	c.Check(s.doc.Trailer, NotNil)
	SignatureLength := s.doc.Trailer.Members[0]
	c.Check(SignatureLength.XMLName.Local, Equals, "field")
	c.Check(SignatureLength.Name, Equals, "SignatureLength")
	c.Check(SignatureLength.Required, Equals, "N")
}

func (s *XMLTests) TestMessages(c *C) {
	c.Check(s.doc.Messages, NotNil)
	heartbeat := s.doc.Messages[0]
	c.Check(heartbeat.Name, Equals, "Heartbeat")
	c.Check(heartbeat.MsgCat, Equals, "admin")
	c.Check(heartbeat.MsgType, Equals, "0")

	testReqID := heartbeat.Members[0]
	c.Check(testReqID.Name, Equals, "TestReqID")

	ioi := s.doc.Messages[1]
	instrument := ioi.Members[3]
	c.Check(instrument.Name, Equals, "Instrument")
	c.Check(instrument.XMLName.Local, Equals, "component")

	noRoutingIDs := ioi.Members[4]
	c.Check(noRoutingIDs.Name, Equals, "NoRoutingIDs")
	c.Check(noRoutingIDs.XMLName.Local, Equals, "group")
	c.Check(noRoutingIDs.Members[0].Name, Equals, "RoutingType")
	c.Check(noRoutingIDs.Members[1].Name, Equals, "RoutingID")
}
