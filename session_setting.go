package quickfixgo

type SessionSetting string

const (
    BeginString SessionSetting = "BeginString"
    SenderCompID SessionSetting = "SenderCompID"
    TargetCompID SessionSetting = "TargetCompID"
    SocketAcceptPort SessionSetting = "SocketAcceptPort"
    SocketConnectPort SessionSetting = "SocketConnectPort"
    )

