//Package news msg type = B.
package news

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a News wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//OrigTime is a non-required field for News.
func (m Message) OrigTime() (*field.OrigTimeField, quickfix.MessageRejectError) {
	f := &field.OrigTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTime reads a OrigTime from News.
func (m Message) GetOrigTime(f *field.OrigTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Urgency is a non-required field for News.
func (m Message) Urgency() (*field.UrgencyField, quickfix.MessageRejectError) {
	f := &field.UrgencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUrgency reads a Urgency from News.
func (m Message) GetUrgency(f *field.UrgencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Headline is a required field for News.
func (m Message) Headline() (*field.HeadlineField, quickfix.MessageRejectError) {
	f := &field.HeadlineField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHeadline reads a Headline from News.
func (m Message) GetHeadline(f *field.HeadlineField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for News.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from News.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LinesOfText is a required field for News.
func (m Message) LinesOfText() (*field.LinesOfTextField, quickfix.MessageRejectError) {
	f := &field.LinesOfTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLinesOfText reads a LinesOfText from News.
func (m Message) GetLinesOfText(f *field.LinesOfTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//URLLink is a non-required field for News.
func (m Message) URLLink() (*field.URLLinkField, quickfix.MessageRejectError) {
	f := &field.URLLinkField{}
	err := m.Body.Get(f)
	return f, err
}

//GetURLLink reads a URLLink from News.
func (m Message) GetURLLink(f *field.URLLinkField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for News.
func (m Message) RawDataLength() (*field.RawDataLengthField, quickfix.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from News.
func (m Message) GetRawDataLength(f *field.RawDataLengthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for News.
func (m Message) RawData() (*field.RawDataField, quickfix.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from News.
func (m Message) GetRawData(f *field.RawDataField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds News messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for News.
func Builder(
	headline *field.HeadlineField,
	linesoftext *field.LinesOfTextField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX41))
	builder.Header().Set(field.NewMsgType("B"))
	builder.Body().Set(headline)
	builder.Body().Set(linesoftext)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX41, "B", r
}
