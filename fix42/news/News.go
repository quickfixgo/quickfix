//Package news msg type = B.
package news

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a News wrapper for the generic Message type
type Message struct {
	message.Message
}

//OrigTime is a non-required field for News.
func (m Message) OrigTime() (*field.OrigTimeField, errors.MessageRejectError) {
	f := &field.OrigTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTime reads a OrigTime from News.
func (m Message) GetOrigTime(f *field.OrigTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Urgency is a non-required field for News.
func (m Message) Urgency() (*field.UrgencyField, errors.MessageRejectError) {
	f := &field.UrgencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUrgency reads a Urgency from News.
func (m Message) GetUrgency(f *field.UrgencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Headline is a required field for News.
func (m Message) Headline() (*field.HeadlineField, errors.MessageRejectError) {
	f := &field.HeadlineField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHeadline reads a Headline from News.
func (m Message) GetHeadline(f *field.HeadlineField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedHeadlineLen is a non-required field for News.
func (m Message) EncodedHeadlineLen() (*field.EncodedHeadlineLenField, errors.MessageRejectError) {
	f := &field.EncodedHeadlineLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedHeadlineLen reads a EncodedHeadlineLen from News.
func (m Message) GetEncodedHeadlineLen(f *field.EncodedHeadlineLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedHeadline is a non-required field for News.
func (m Message) EncodedHeadline() (*field.EncodedHeadlineField, errors.MessageRejectError) {
	f := &field.EncodedHeadlineField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedHeadline reads a EncodedHeadline from News.
func (m Message) GetEncodedHeadline(f *field.EncodedHeadlineField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for News.
func (m Message) NoRoutingIDs() (*field.NoRoutingIDsField, errors.MessageRejectError) {
	f := &field.NoRoutingIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from News.
func (m Message) GetNoRoutingIDs(f *field.NoRoutingIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for News.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from News.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LinesOfText is a required field for News.
func (m Message) LinesOfText() (*field.LinesOfTextField, errors.MessageRejectError) {
	f := &field.LinesOfTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLinesOfText reads a LinesOfText from News.
func (m Message) GetLinesOfText(f *field.LinesOfTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//URLLink is a non-required field for News.
func (m Message) URLLink() (*field.URLLinkField, errors.MessageRejectError) {
	f := &field.URLLinkField{}
	err := m.Body.Get(f)
	return f, err
}

//GetURLLink reads a URLLink from News.
func (m Message) GetURLLink(f *field.URLLinkField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for News.
func (m Message) RawDataLength() (*field.RawDataLengthField, errors.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from News.
func (m Message) GetRawDataLength(f *field.RawDataLengthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for News.
func (m Message) RawData() (*field.RawDataField, errors.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from News.
func (m Message) GetRawData(f *field.RawDataField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds News messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for News.
func Builder(
	headline *field.HeadlineField,
	linesoftext *field.LinesOfTextField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header().Set(field.NewMsgType("B"))
	builder.Body().Set(headline)
	builder.Body().Set(linesoftext)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX42, "B", r
}
