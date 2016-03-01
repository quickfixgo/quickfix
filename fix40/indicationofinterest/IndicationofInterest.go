//Package indicationofinterest msg type = 6.
package indicationofinterest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
	"time"
)

//Message is a IndicationofInterest FIX Message
type Message struct {
	FIXMsgType string `fix:"6"`
	Header     fix40.Header
	//IOIid is a required field for IndicationofInterest.
	IOIid int `fix:"23"`
	//IOITransType is a required field for IndicationofInterest.
	IOITransType string `fix:"28"`
	//IOIRefID is a non-required field for IndicationofInterest.
	IOIRefID *int `fix:"26"`
	//Symbol is a required field for IndicationofInterest.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for IndicationofInterest.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for IndicationofInterest.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for IndicationofInterest.
	IDSource *string `fix:"22"`
	//Issuer is a non-required field for IndicationofInterest.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for IndicationofInterest.
	SecurityDesc *string `fix:"107"`
	//Side is a required field for IndicationofInterest.
	Side string `fix:"54"`
	//IOIShares is a required field for IndicationofInterest.
	IOIShares string `fix:"27"`
	//Price is a non-required field for IndicationofInterest.
	Price *float64 `fix:"44"`
	//Currency is a non-required field for IndicationofInterest.
	Currency *string `fix:"15"`
	//ValidUntilTime is a non-required field for IndicationofInterest.
	ValidUntilTime *time.Time `fix:"62"`
	//IOIQltyInd is a non-required field for IndicationofInterest.
	IOIQltyInd *string `fix:"25"`
	//IOIOthSvc is a non-required field for IndicationofInterest.
	IOIOthSvc *string `fix:"24"`
	//IOINaturalFlag is a non-required field for IndicationofInterest.
	IOINaturalFlag *string `fix:"130"`
	//IOIQualifier is a non-required field for IndicationofInterest.
	IOIQualifier *string `fix:"104"`
	//Text is a non-required field for IndicationofInterest.
	Text    *string `fix:"58"`
	Trailer fix40.Trailer
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
	return enum.BeginStringFIX40, "6", r
}
