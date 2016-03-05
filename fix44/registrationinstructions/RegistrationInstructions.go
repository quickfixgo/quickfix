//Package registrationinstructions msg type = o.
package registrationinstructions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
	"github.com/quickfixgo/quickfix/fix44/parties"
)

//NoRegistDtls is a repeating group in RegistrationInstructions
type NoRegistDtls struct {
	//RegistDtls is a non-required field for NoRegistDtls.
	RegistDtls *string `fix:"509"`
	//RegistEmail is a non-required field for NoRegistDtls.
	RegistEmail *string `fix:"511"`
	//MailingDtls is a non-required field for NoRegistDtls.
	MailingDtls *string `fix:"474"`
	//MailingInst is a non-required field for NoRegistDtls.
	MailingInst *string `fix:"482"`
	//NestedParties Component
	nestedparties.NestedParties
	//OwnerType is a non-required field for NoRegistDtls.
	OwnerType *int `fix:"522"`
	//DateOfBirth is a non-required field for NoRegistDtls.
	DateOfBirth *string `fix:"486"`
	//InvestorCountryOfResidence is a non-required field for NoRegistDtls.
	InvestorCountryOfResidence *string `fix:"475"`
}

func (m *NoRegistDtls) SetRegistDtls(v string)                 { m.RegistDtls = &v }
func (m *NoRegistDtls) SetRegistEmail(v string)                { m.RegistEmail = &v }
func (m *NoRegistDtls) SetMailingDtls(v string)                { m.MailingDtls = &v }
func (m *NoRegistDtls) SetMailingInst(v string)                { m.MailingInst = &v }
func (m *NoRegistDtls) SetOwnerType(v int)                     { m.OwnerType = &v }
func (m *NoRegistDtls) SetDateOfBirth(v string)                { m.DateOfBirth = &v }
func (m *NoRegistDtls) SetInvestorCountryOfResidence(v string) { m.InvestorCountryOfResidence = &v }

//NoDistribInsts is a repeating group in RegistrationInstructions
type NoDistribInsts struct {
	//DistribPaymentMethod is a non-required field for NoDistribInsts.
	DistribPaymentMethod *int `fix:"477"`
	//DistribPercentage is a non-required field for NoDistribInsts.
	DistribPercentage *float64 `fix:"512"`
	//CashDistribCurr is a non-required field for NoDistribInsts.
	CashDistribCurr *string `fix:"478"`
	//CashDistribAgentName is a non-required field for NoDistribInsts.
	CashDistribAgentName *string `fix:"498"`
	//CashDistribAgentCode is a non-required field for NoDistribInsts.
	CashDistribAgentCode *string `fix:"499"`
	//CashDistribAgentAcctNumber is a non-required field for NoDistribInsts.
	CashDistribAgentAcctNumber *string `fix:"500"`
	//CashDistribPayRef is a non-required field for NoDistribInsts.
	CashDistribPayRef *string `fix:"501"`
	//CashDistribAgentAcctName is a non-required field for NoDistribInsts.
	CashDistribAgentAcctName *string `fix:"502"`
}

func (m *NoDistribInsts) SetDistribPaymentMethod(v int)          { m.DistribPaymentMethod = &v }
func (m *NoDistribInsts) SetDistribPercentage(v float64)         { m.DistribPercentage = &v }
func (m *NoDistribInsts) SetCashDistribCurr(v string)            { m.CashDistribCurr = &v }
func (m *NoDistribInsts) SetCashDistribAgentName(v string)       { m.CashDistribAgentName = &v }
func (m *NoDistribInsts) SetCashDistribAgentCode(v string)       { m.CashDistribAgentCode = &v }
func (m *NoDistribInsts) SetCashDistribAgentAcctNumber(v string) { m.CashDistribAgentAcctNumber = &v }
func (m *NoDistribInsts) SetCashDistribPayRef(v string)          { m.CashDistribPayRef = &v }
func (m *NoDistribInsts) SetCashDistribAgentAcctName(v string)   { m.CashDistribAgentAcctName = &v }

//Message is a RegistrationInstructions FIX Message
type Message struct {
	FIXMsgType string `fix:"o"`
	fix44.Header
	//RegistID is a required field for RegistrationInstructions.
	RegistID string `fix:"513"`
	//RegistTransType is a required field for RegistrationInstructions.
	RegistTransType string `fix:"514"`
	//RegistRefID is a required field for RegistrationInstructions.
	RegistRefID string `fix:"508"`
	//ClOrdID is a non-required field for RegistrationInstructions.
	ClOrdID *string `fix:"11"`
	//Parties Component
	parties.Parties
	//Account is a non-required field for RegistrationInstructions.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for RegistrationInstructions.
	AcctIDSource *int `fix:"660"`
	//RegistAcctType is a non-required field for RegistrationInstructions.
	RegistAcctType *string `fix:"493"`
	//TaxAdvantageType is a non-required field for RegistrationInstructions.
	TaxAdvantageType *int `fix:"495"`
	//OwnershipType is a non-required field for RegistrationInstructions.
	OwnershipType *string `fix:"517"`
	//NoRegistDtls is a non-required field for RegistrationInstructions.
	NoRegistDtls []NoRegistDtls `fix:"473,omitempty"`
	//NoDistribInsts is a non-required field for RegistrationInstructions.
	NoDistribInsts []NoDistribInsts `fix:"510,omitempty"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetRegistID(v string)                 { m.RegistID = v }
func (m *Message) SetRegistTransType(v string)          { m.RegistTransType = v }
func (m *Message) SetRegistRefID(v string)              { m.RegistRefID = v }
func (m *Message) SetClOrdID(v string)                  { m.ClOrdID = &v }
func (m *Message) SetAccount(v string)                  { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)                { m.AcctIDSource = &v }
func (m *Message) SetRegistAcctType(v string)           { m.RegistAcctType = &v }
func (m *Message) SetTaxAdvantageType(v int)            { m.TaxAdvantageType = &v }
func (m *Message) SetOwnershipType(v string)            { m.OwnershipType = &v }
func (m *Message) SetNoRegistDtls(v []NoRegistDtls)     { m.NoRegistDtls = v }
func (m *Message) SetNoDistribInsts(v []NoDistribInsts) { m.NoDistribInsts = v }

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
	return enum.BeginStringFIX44, "o", r
}
