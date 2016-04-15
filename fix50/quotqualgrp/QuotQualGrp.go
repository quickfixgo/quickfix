package quotqualgrp

func New() *QuotQualGrp {
	var m QuotQualGrp
	return &m
}

//NoQuoteQualifiers is a repeating group in QuotQualGrp
type NoQuoteQualifiers struct {
	//QuoteQualifier is a non-required field for NoQuoteQualifiers.
	QuoteQualifier *string `fix:"695"`
}

func (m *NoQuoteQualifiers) SetQuoteQualifier(v string) { m.QuoteQualifier = &v }

//QuotQualGrp is a fix50 Component
type QuotQualGrp struct {
	//NoQuoteQualifiers is a non-required field for QuotQualGrp.
	NoQuoteQualifiers []NoQuoteQualifiers `fix:"735,omitempty"`
}

func (m *QuotQualGrp) SetNoQuoteQualifiers(v []NoQuoteQualifiers) { m.NoQuoteQualifiers = v }
