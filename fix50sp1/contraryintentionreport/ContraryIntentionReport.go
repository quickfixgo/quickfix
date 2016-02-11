//Package contraryintentionreport msg type = BO.
package contraryintentionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp1/expirationqty"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a ContraryIntentionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BO"`
	Header     fixt11.Header
	//ContIntRptID is a required field for ContraryIntentionReport.
	ContIntRptID string `fix:"977"`
	//TransactTime is a non-required field for ContraryIntentionReport.
	TransactTime *time.Time `fix:"60"`
	//LateIndicator is a non-required field for ContraryIntentionReport.
	LateIndicator *bool `fix:"978"`
	//InputSource is a non-required field for ContraryIntentionReport.
	InputSource *string `fix:"979"`
	//ClearingBusinessDate is a required field for ContraryIntentionReport.
	ClearingBusinessDate string `fix:"715"`
	//Parties Component
	Parties parties.Component
	//ExpirationQty Component
	ExpirationQty expirationqty.Component
	//Instrument Component
	Instrument instrument.Component
	//Text is a non-required field for ContraryIntentionReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ContraryIntentionReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ContraryIntentionReport.
	EncodedText *string `fix:"355"`
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//ApplicationSequenceControl Component
	ApplicationSequenceControl applicationsequencecontrol.Component
	Trailer                    fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "BO", r
}
