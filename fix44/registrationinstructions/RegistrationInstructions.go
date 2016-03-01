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
	NestedParties nestedparties.Component
	//OwnerType is a non-required field for NoRegistDtls.
	OwnerType *int `fix:"522"`
	//DateOfBirth is a non-required field for NoRegistDtls.
	DateOfBirth *string `fix:"486"`
	//InvestorCountryOfResidence is a non-required field for NoRegistDtls.
	InvestorCountryOfResidence *string `fix:"475"`
}

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

//Message is a RegistrationInstructions FIX Message
type Message struct {
	FIXMsgType string `fix:"o"`
	Header     fix44.Header
	//RegistID is a required field for RegistrationInstructions.
	RegistID string `fix:"513"`
	//RegistTransType is a required field for RegistrationInstructions.
	RegistTransType string `fix:"514"`
	//RegistRefID is a required field for RegistrationInstructions.
	RegistRefID string `fix:"508"`
	//ClOrdID is a non-required field for RegistrationInstructions.
	ClOrdID *string `fix:"11"`
	//Parties Component
	Parties parties.Component
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
	Trailer        fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
