package collinqqualgrp

//NoCollInquiryQualifier is a repeating group in CollInqQualGrp
type NoCollInquiryQualifier struct {
	//CollInquiryQualifier is a non-required field for NoCollInquiryQualifier.
	CollInquiryQualifier *int `fix:"896"`
}

//Component is a fix50 CollInqQualGrp Component
type Component struct {
	//NoCollInquiryQualifier is a non-required field for CollInqQualGrp.
	NoCollInquiryQualifier []NoCollInquiryQualifier `fix:"938,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoCollInquiryQualifier(v []NoCollInquiryQualifier) {
	m.NoCollInquiryQualifier = v
}
