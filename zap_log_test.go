package quickfix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewZapLogFactory(t *testing.T) {
	f := NewZapLogFactory("pre", true)
	assert.Equal(t, "pre", f.prefix)
	assert.True(t, f.messages)
}

func TestReplaceEos(t *testing.T) {
	var preFormat, postFormat []byte
	preFormat = []byte("8=FIXT.1.1\u00019=74\u000135=CW\u000134=114\u000149=MIXFI\u000152=20221227-13:51:21.209\u000156=SELL1\u0001117=ID09876\u0001297=0\u000110=078\u0001")
	postFormat = []byte("8=FIXT.1.1|9=74|35=CW|34=114|49=MIXFI|52=20221227-13:51:21.209|56=SELL1|117=ID09876|297=0|10=078|")

	res := ReplaceEos(preFormat)
	assert.Equal(t, postFormat, res)
}

func TestZapLogFactory_Create(t *testing.T) {
	f := NewZapLogFactory("fix ", false)
	l, err := f.Create()
	assert.Nil(t, err, "no error from Create()")
	zl, ok := l.(ZapLog)
	assert.True(t, ok, "logger is of type ZapLog")
	assert.Equal(t, "fix GLOBAL", zl.name, "logger name")
}

func TestZapLogFactory_CreateSessionLog(t *testing.T) {

	f := NewZapLogFactory("", false)
	sessionId := SessionID{
		BeginString:      "",
		TargetCompID:     "TOCOMP",
		TargetSubID:      "",
		TargetLocationID: "",
		SenderCompID:     "FROMCOMP",
		SenderSubID:      "",
		SenderLocationID: "",
		Qualifier:        "",
	}
	l, err := f.CreateSessionLog(sessionId)
	assert.Nil(t, err, "no error from Create()")
	zl, ok := l.(ZapLog)
	assert.True(t, ok, "logger is of type ZapLog")
	assert.Equal(t, "FROMCOMP:TOCOMP", zl.name, "logger name")
}
