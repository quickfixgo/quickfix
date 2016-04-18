package collinqqualgrp

//New returns an initialized CollInqQualGrp instance
func New() *CollInqQualGrp {
	var m CollInqQualGrp
	return &m
}

//NoCollInquiryQualifier is a repeating group in CollInqQualGrp
type NoCollInquiryQualifier struct {
	//CollInquiryQualifier is a non-required field for NoCollInquiryQualifier.
	CollInquiryQualifier *int `fix:"896"`
}

//NewNoCollInquiryQualifier returns an initialized NoCollInquiryQualifier instance
func NewNoCollInquiryQualifier() *NoCollInquiryQualifier {
	var m NoCollInquiryQualifier
	return &m
}

func (m *NoCollInquiryQualifier) SetCollInquiryQualifier(v int) { m.CollInquiryQualifier = &v }

//CollInqQualGrp is a fix50 Component
type CollInqQualGrp struct {
	//NoCollInquiryQualifier is a non-required field for CollInqQualGrp.
	NoCollInquiryQualifier []NoCollInquiryQualifier `fix:"938,omitempty"`
}

func (m *CollInqQualGrp) SetNoCollInquiryQualifier(v []NoCollInquiryQualifier) {
	m.NoCollInquiryQualifier = v
}
