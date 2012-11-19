package quickfixgo

type Dictionary interface {
  GetString(setting SessionSetting) (string, bool)
  SetString(setting SessionSetting, val string)

  GetInt(setting SessionSetting) (int, bool)
  SetInt(setting SessionSetting, val int)
}

type dictionary struct {
  stringSettings map[SessionSetting] string
  intSettings map[SessionSetting] int
}

func (d * dictionary) GetString(setting SessionSetting) (val string, ok bool) {
  val,ok=d.stringSettings[setting]

  return
}

func (d * dictionary) SetString(setting SessionSetting, val string) {
  d.stringSettings[setting] = val
}

func (d* dictionary) GetInt(setting SessionSetting) (val int, ok bool) {
  val,ok=d.intSettings[setting]

  return
}

func (d * dictionary) SetInt(setting SessionSetting, val int) {
  d.intSettings[setting] = val
}


func NewDictionary() Dictionary {
  d:=dictionary{}
  d.stringSettings=make(map[SessionSetting] string)
  d.intSettings=make(map[SessionSetting] int)

  return &d
}
