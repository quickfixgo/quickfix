//Package marketdefinition msg type = BU.
package marketdefinition

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp2/basetradingrules"
	"github.com/quickfixgo/quickfix/fix50sp2/execinstrules"
	"github.com/quickfixgo/quickfix/fix50sp2/ordtyperules"
	"github.com/quickfixgo/quickfix/fix50sp2/timeinforcerules"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a MarketDefinition FIX Message
type Message struct {
	FIXMsgType string `fix:"BU"`
	Header     fixt11.Header
	//MarketReportID is a required field for MarketDefinition.
	MarketReportID string `fix:"1394"`
	//MarketReqID is a non-required field for MarketDefinition.
	MarketReqID *string `fix:"1393"`
	//MarketID is a required field for MarketDefinition.
	MarketID string `fix:"1301"`
	//MarketSegmentID is a non-required field for MarketDefinition.
	MarketSegmentID *string `fix:"1300"`
	//MarketSegmentDesc is a non-required field for MarketDefinition.
	MarketSegmentDesc *string `fix:"1396"`
	//EncodedMktSegmDescLen is a non-required field for MarketDefinition.
	EncodedMktSegmDescLen *int `fix:"1397"`
	//EncodedMktSegmDesc is a non-required field for MarketDefinition.
	EncodedMktSegmDesc *string `fix:"1398"`
	//ParentMktSegmID is a non-required field for MarketDefinition.
	ParentMktSegmID *string `fix:"1325"`
	//Currency is a non-required field for MarketDefinition.
	Currency *string `fix:"15"`
	//BaseTradingRules Component
	BaseTradingRules basetradingrules.Component
	//OrdTypeRules Component
	OrdTypeRules ordtyperules.Component
	//TimeInForceRules Component
	TimeInForceRules timeinforcerules.Component
	//ExecInstRules Component
	ExecInstRules execinstrules.Component
	//TransactTime is a non-required field for MarketDefinition.
	TransactTime *time.Time `fix:"60"`
	//Text is a non-required field for MarketDefinition.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for MarketDefinition.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for MarketDefinition.
	EncodedText *string `fix:"355"`
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
	return enum.ApplVerID_FIX50SP2, "BU", r
}
