package fixt11

//Trailer is the fixt11 Trailer type
type Trailer struct {
	//SignatureLength is a non-required field for Trailer.
	SignatureLength *int `fix:"93"`
	//Signature is a non-required field for Trailer.
	Signature *string `fix:"89"`
	//CheckSum is a required field for Trailer.
	CheckSum string `fix:"10"`
}

func (m *Trailer) SetSignatureLength(v int) { m.SignatureLength = &v }
func (m *Trailer) SetSignature(v string)    { m.Signature = &v }
func (m *Trailer) SetCheckSum(v string)     { m.CheckSum = v }
