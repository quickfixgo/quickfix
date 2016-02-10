package fix40

//Trailer is the fix40 Trailer type
type Trailer struct {
	//SignatureLength is a non-required field for Trailer.
	SignatureLength *int `fix:"93"`
	//Signature is a non-required field for Trailer.
	Signature *string `fix:"89"`
	//CheckSum is a required field for Trailer.
	CheckSum string `fix:"10"`
}
