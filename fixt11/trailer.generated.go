package fixt11

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/tag"
)

//Trailer is the fixt11 Trailer type
type Trailer struct {
	quickfix.Trailer
}

//SetCheckSum sets CheckSum, Tag 10
func (t Trailer) SetCheckSum(v string) {
	t.Set(field.NewCheckSum(v))
}

//SetSignature sets Signature, Tag 89
func (t Trailer) SetSignature(v string) {
	t.Set(field.NewSignature(v))
}

//SetSignatureLength sets SignatureLength, Tag 93
func (t Trailer) SetSignatureLength(v int) {
	t.Set(field.NewSignatureLength(v))
}

//GetCheckSum gets CheckSum, Tag 10
func (t Trailer) GetCheckSum() (f field.CheckSumField, err quickfix.MessageRejectError) {
	err = t.Get(&f)
	return
}

//GetSignature gets Signature, Tag 89
func (t Trailer) GetSignature() (f field.SignatureField, err quickfix.MessageRejectError) {
	err = t.Get(&f)
	return
}

//GetSignatureLength gets SignatureLength, Tag 93
func (t Trailer) GetSignatureLength() (f field.SignatureLengthField, err quickfix.MessageRejectError) {
	err = t.Get(&f)
	return
}

//HasCheckSum returns true if CheckSum is present, Tag 10
func (t Trailer) HasCheckSum() bool {
	return t.Has(tag.CheckSum)
}

//HasSignature returns true if Signature is present, Tag 89
func (t Trailer) HasSignature() bool {
	return t.Has(tag.Signature)
}

//HasSignatureLength returns true if SignatureLength is present, Tag 93
func (t Trailer) HasSignatureLength() bool {
	return t.Has(tag.SignatureLength)
}
