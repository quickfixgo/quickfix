package quotqualgrp

//NoQuoteQualifiers is a repeating group in QuotQualGrp
type NoQuoteQualifiers struct {
	//QuoteQualifier is a non-required field for NoQuoteQualifiers.
	QuoteQualifier *string `fix:"695"`
}

//Component is a fix50sp1 QuotQualGrp Component
type Component struct {
	//NoQuoteQualifiers is a non-required field for QuotQualGrp.
	NoQuoteQualifiers []NoQuoteQualifiers `fix:"735,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoQuoteQualifiers(v []NoQuoteQualifiers) { m.NoQuoteQualifiers = v }
