package fix42

import (
	"github.com/alpacahq/quickfix"
	"github.com/alpacahq/quickfix/field"
	"github.com/alpacahq/quickfix/tag"
)

//Trailer is the fix42 Trailer type
type Trailer struct {
	*quickfix.Trailer
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
func (t Trailer) GetCheckSum() (v string, err quickfix.MessageRejectError) {
	var f field.CheckSumField
	if err = t.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSignature gets Signature, Tag 89
func (t Trailer) GetSignature() (v string, err quickfix.MessageRejectError) {
	var f field.SignatureField
	if err = t.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSignatureLength gets SignatureLength, Tag 93
func (t Trailer) GetSignatureLength() (v int, err quickfix.MessageRejectError) {
	var f field.SignatureLengthField
	if err = t.Get(&f); err == nil {
		v = f.Value()
	}
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
