package config

// NOTE: Additions to this file should be made to both config/doc.go and http://www.quickfixgo.org/docs/

// Const configuration settings
const (
	BeginString                  string = "BeginString"
	SenderCompID                 string = "SenderCompID"
	SenderSubID                  string = "SenderSubID"
	SenderLocationID             string = "SenderLocationID"
	TargetCompID                 string = "TargetCompID"
	TargetSubID                  string = "TargetSubID"
	TargetLocationID             string = "TargetLocationID"
	SessionQualifier             string = "SessionQualifier"
	SocketAcceptHost             string = "SocketAcceptHost"
	SocketAcceptPort             string = "SocketAcceptPort"
	SocketConnectHost            string = "SocketConnectHost"
	SocketConnectPort            string = "SocketConnectPort"
	SocketPrivateKeyFile         string = "SocketPrivateKeyFile"
	SocketCertificateFile        string = "SocketCertificateFile"
	SocketCAFile                 string = "SocketCAFile"
	SocketInsecureSkipVerify     string = "SocketInsecureSkipVerify"
	SocketMinimumTLSVersion      string = "SocketMinimumTLSVersion"
	DefaultApplVerID             string = "DefaultApplVerID"
	StartTime                    string = "StartTime"
	EndTime                      string = "EndTime"
	StartDay                     string = "StartDay"
	EndDay                       string = "EndDay"
	TimeZone                     string = "TimeZone"
	DataDictionary               string = "DataDictionary"
	TransportDataDictionary      string = "TransportDataDictionary"
	AppDataDictionary            string = "AppDataDictionary"
	ResetOnLogon                 string = "ResetOnLogon"
	RefreshOnLogon               string = "RefreshOnLogon"
	ResetOnLogout                string = "ResetOnLogout"
	ResetOnDisconnect            string = "ResetOnDisconnect"
	ReconnectInterval            string = "ReconnectInterval"
	HeartBtInt                   string = "HeartBtInt"
	FileLogPath                  string = "FileLogPath"
	FileStorePath                string = "FileStorePath"
	SQLStoreDriver               string = "SQLStoreDriver"
	SQLStoreDataSourceName       string = "SQLStoreDataSourceName"
	SQLStoreConnMaxLifetime      string = "SQLStoreConnMaxLifetime"
	SQLStoreConnMaxIdle          string = "SQLStoreConnMaxIdle"
	SQLStoreConnMaxOpen          string = "SQLStoreConnMaxOpen"
	ValidateFieldsOutOfOrder     string = "ValidateFieldsOutOfOrder"
	ResendRequestChunkSize       string = "ResendRequestChunkSize"
	EnableLastMsgSeqNumProcessed string = "EnableLastMsgSeqNumProcessed"
	CheckLatency                 string = "CheckLatency"
	TimeStampPrecision           string = "TimeStampPrecision"
	MaxLatency                   string = "MaxLatency"
	PersistMessages              string = "PersistMessages"
)
