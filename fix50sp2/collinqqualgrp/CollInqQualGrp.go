package collinqqualgrp

func New() *CollInqQualGrp {
	var m CollInqQualGrp
	return &m
}

//NoCollInquiryQualifier is a repeating group in CollInqQualGrp
type NoCollInquiryQualifier struct {
	//CollInquiryQualifier is a non-required field for NoCollInquiryQualifier.
	CollInquiryQualifier *int `fix:"896"`
}

func (m *NoCollInquiryQualifier) SetCollInquiryQualifier(v int) { m.CollInquiryQualifier = &v }

//CollInqQualGrp is a fix50sp2 Component
type CollInqQualGrp struct {
	//NoCollInquiryQualifier is a non-required field for CollInqQualGrp.
	NoCollInquiryQualifier []NoCollInquiryQualifier `fix:"938,omitempty"`
}

func (m *CollInqQualGrp) SetNoCollInquiryQualifier(v []NoCollInquiryQualifier) {
	m.NoCollInquiryQualifier = v
}
