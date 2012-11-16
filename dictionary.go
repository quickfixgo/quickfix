package quickfixgo

type Dictionary interface {
  GetString(setting SessionSetting) (bool, string)
}
