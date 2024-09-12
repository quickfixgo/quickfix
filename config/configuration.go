/*
Package config declares settings for configuring QuickFIX/Go.

A quickfix acceptor or initiator can maintain multiple FIX
sessions. A FIX session is defined in QuickFIX as a unique
combination of a BeginString (the FIX version number), a
SenderCompID (your ID), and a TargetCompID (the
ID of your counterparty). A SessionQualifier can also be
used in rare cases to disambiguate otherwise identical sessions.

The quickfix.ParseSettings(reader io.Reader) func will pull settings
out of any stream, most commonly, a file stream.
If you decide to write your own components,
(storage for a particular database, a new kind of connector
etc...), you may also use the session settings to store settings
for your custom component.

A settings file is set up with two types of headings, a
[DEFAULT] heading and a [SESSION] heading.

[DEFAULT] is where you can define settings that all sessions use by default.

[SESSION] tells QuickFIX/Go that a new Session is being defined.

If you do not provide a setting that QuickFIX/Go needs, it will throw an error
telling you what setting is missing or improperly formatted.

# Sample Configuration Settings File:

	# default settings for sessions
	[DEFAULT]
	ConnectionType=initiator
	ReconnectInterval=60
	SenderCompID=TW

	# session definition
	[SESSION]
	# inherit ConnectionType, ReconnectInterval and SenderCompID from default
	BeginString=FIX.4.1
	TargetCompID=ARCA
	StartTime=12:30:00
	EndTime=23:30:00
	HeartBtInt=20
	SocketConnectPort=9823
	SocketConnectHost=123.123.123.123
	DataDictionary=somewhere/FIX41.xml

	[SESSION]
	BeginString=FIX.4.0
	TargetCompID=ISLD
	StartTime=12:00:00
	EndTime=23:00:00
	HeartBtInt=30
	SocketConnectPort=8323
	SocketConnectHost=23.23.23.23
	DataDictionary=somewhere/FIX40.xml

	[SESSION]
	BeginString=FIX.4.2
	TargetCompID=INCA
	StartTime=12:30:00
	EndTime=21:30:00
	# overide default setting for RecconnectInterval
	ReconnectInterval=30
	HeartBtInt=30
	SocketConnectPort=6523
	SocketConnectHost=3.3.3.3
	# (optional) alternate connection ports and hosts to cycle through on failover
	SocketConnectPort1=8392
	SocketConnectHost1=8.8.8.8
	SocketConnectPort2=2932
	SocketConnectHost2=12.12.12.12
	DataDictionary=somewhere/FIX42.xml
*/
package config

const (
	// Session settings.

	// BeginString is the version of FIX this session should use.
	//
	// Required: Yes
	//
	// Default: N/A
	//
	// Valid Values:
	// 	- FIXT.1.1
	//  - FIX.4.4
	//  - FIX.4.3
	//  - FIX.4.2
	//  - FIX.4.1
	//  - FIX.4.0
	BeginString string = "BeginString"

	// SenderCompID is your ID as associated with this FIX session.
	//
	// Required: Yes, unless configuring an acceptor with DynamicSessions=Y
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A case-sensitive alpha-numeric string.
	SenderCompID string = "SenderCompID"

	// SenderSubID is your subID as associated with this FIX session.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A case-sensitive alpha-numeric string.
	SenderSubID string = "SenderSubID"

	// SenderLocationID is your locationID as associated with this FIX session.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A case-sensitive alpha-numeric string.
	SenderLocationID string = "SenderLocationID"

	// TargetCompID is the counterparty's ID as associated with this FIX session.
	//
	// Required: Yes, unless configuring an acceptor with DynamicSessions=Y
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A case-sensitive alpha-numeric string.
	TargetCompID string = "TargetCompID"

	// TargetSubID is the counterparty's subID as associated with this FIX session.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A case-sensitive alpha-numeric string.
	TargetSubID string = "TargetSubID"

	// TargetLocationID is the counterparty's locationID as associated with this FIX session.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A case-sensitive alpha-numeric string.
	TargetLocationID string = "TargetLocationID"

	// SessionQualifier is an additional qualifier to disambiguate otherwise identical sessions.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A case-sensitive alpha-numeric string.
	SessionQualifier string = "SessionQualifier"

	// DefaultApplVerID specifies the default application version ID for the session.
	// This can either be the ApplVerID enum (see the ApplVerID field) or the BeginString for the default version.
	//
	// Required: Only for FIXT 1.1 (and newer). Ignored for earlier transport versions.
	//
	// Default: N/A
	//
	// Valid Values:
	//  - FIX.5.0SP2
	//  - FIX.5.0SP1
	//  - FIX.5.0
	//  - FIX.4.4
	//  - FIX.4.3
	//  - FIX.4.2
	//  - FIX.4.1
	//  - FIX.4.0
	//  - 9
	//  - 8
	//  - 7
	//  - 6
	//  - 5
	//  - 4
	//  - 3
	//  - 2
	DefaultApplVerID string = "DefaultApplVerID"

	// StartTime is the time of day that this FIX session becomes activated.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A time in the format of HH:MM:SS, time is represented in time zone configured by TimeZone
	StartTime string = "StartTime"

	// EndTime is the time of day that this FIX session becomes deactivated.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A time in the format of HH:MM:SS, time is represented in time zone configured by TimeZone
	EndTime string = "EndTime"

	// StartDay is the starting day of week for the session,
	// used to configure week-long sessions,
	// and can be used in combination with StartTime.
	// Incompatible with the Weekdays setting.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - Full day of week in English, or 3 letter abbreviation (i.e. Monday and Mon are valid)
	StartDay string = "StartDay"

	// EndDay is the ending day of week for the session,
	// used to configure week-long sessions,
	// and can be used in combination with EndTime.
	// Incompatible with the Weekdays setting.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - Full day of week in English, or 3 letter abbreviation (i.e. Monday and Mon are valid)
	EndDay string = "EndDay"

	// Weekdays is for daily sessions that are only active on specific days of the week.
	// Use in combination with StartTime and EndTime.
	// Incompatible with StartDay and EndDay.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - Comma delimited list of days of the week in English, or 3 letter abbreviation (e.g. "Monday,Tuesday,Wednesday" or "Mon,Tue,Wed" would both be valid values).
	Weekdays string = "Weekdays"

	// TimeZone sets the time zone for this session; if specified, StartTime and EndTime will be converted from this zone to UTC.
	// Times in messages will still be set to UTC as this is required by FIX specifications.
	//
	// Required: No
	//
	// Default: UTC
	//
	// Valid Values:
	//  - IANA Time zone ID (America/New_York, Asia/Tokyo, Europe/London, etc.)
	//  - Local (The zone on host)
	TimeZone string = "TimeZone"

	// TimeStampPrecision determines precision for timestamps in (Orig)SendingTime fields in outbound messages.
	// Only available for FIX.4.2 and greater, FIX versions earlier than FIX.4.2 will use timestamp resolution in seconds.
	//
	// Required: No
	//
	// Default: MILLIS
	//
	// Valid Values:
	//  - SECONDS
	//  - MILLIS
	//  - MICROS
	//  - NANOS
	TimeStampPrecision string = "TimeStampPrecision"

	// ResetOnLogon determines if sequence numbers should be reset when receiving a logon request.
	// Valid for FIX Acceptors only.
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	ResetOnLogon string = "ResetOnLogon"

	// RefreshOnLogon determines if session state should be restored from persistence layer when logging on.
	// Useful for creating hot failover sessions.
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	RefreshOnLogon string = "RefreshOnLogon"

	// ResetOnLogout determines if sequence numbers should be reset to 1 after a normal logout termination.
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	ResetOnLogout string = "ResetOnLogout"

	// ResetOnDisconnect determines if sequence numbers should be reset to 1 after an abnormal termination.
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	ResetOnDisconnect string = "ResetOnDisconnect"
)

const (
	// Validation settings.

	// DataDictionary is the path to an XML definition file for validating incoming FIX messages.
	// If this is not set, only basic message validation will be done.
	// If messages contain repeating groups, this is required to parse them correctly.
	// This setting should only be used with FIX transport versions older than FIXT.1.1.
	// See TransportDataDictionary and AppDataDictionary for FIXT.1.1 messages.
	// Value must be a path to a valid XML data dictionary file.
	//
	// QuickFIX/Go repo contains the following standard dictionaries in the spec/ directory:
	//  - FIX44.xml
	//  - FIX43.xml
	//  - FIX42.xml
	//  - FIX41.xml
	//  - FIX40.xml
	//
	// Required: No, but strongly recommended
	//
	// Default: No dictionary, and QuickFIX/Go does not attempt to load any standard dictionaries
	//
	// Valid Values:
	//  - A filepath to a XML file with read access.
	DataDictionary string = "DataDictionary"

	// TransportDataDictionary is the path to an XML definition file for validating admin (transport) messages.
	// This setting is only valid for FIXT.1.1 (or newer) sessions.
	// See DataDictionary for older transport versions (FIX.4.0 - FIX.4.4) for additional information.
	// Value must be a path to a valid XML data dictionary file.
	//
	// QuickFIX/Go repo contains the following standard dictionary in the spec/ directory
	//  - FIXT1.1.xml
	//
	// Required: No, but strongly recommended. Required if also using AppDataDictionary
	//
	// Default: No dictionary, and QuickFIX/Go does not attempt to load any standard dictionaries
	//
	// Valid Values:
	//  - A filepath to a XML file with read access.
	TransportDataDictionary string = "TransportDataDictionary"

	// AppDataDictionary is the path to an XML definition file for validating application messages.
	// This setting is only valid for FIXT.1.1 (or newer) sessions.
	// See DataDictionary for older transport versions (FIX.4.0 - FIX.4.4) for additional information.
	// Value must be a path to a valid XML data dictionary file.
	//
	// This setting supports the possibility of a custom application data dictionary for each session.
	// This setting can be used as a prefix to specify multiple application dictionaries for the FIXT transport.
	//
	// For example:
	//  DefaultApplVerID=FIX.4.2
	//  # For default application version ID
	//  AppDataDictionary=FIX42.xml
	//  # For nondefault application version ID
	//  # Use BeginString suffix for app version
	//  AppDataDictionary.FIX.4.4=FIX44.xml
	//
	// QuickFIX/Go repo contains the following standard dictionaries in the spec/ directory
	//  - FIX50SP2.xml
	//  - FIX50SP1.xml
	//  - FIX50.xml
	//  - FIX44.xml
	//  - FIX43.xml
	//  - FIX42.xml
	//  - FIX41.xml
	//  - FIX40.xml
	//
	// Required: No, but strongly recommended. Required if also using TransportDataDictionary
	//
	// Default: No dictionary, and QuickFIX/Go does not attempt to load any standard dictionaries
	//
	// Valid Values:
	//  - A filepath to a XML file with read access.
	AppDataDictionary string = "AppDataDictionary"

	// RejectInvalidMessage is set by detault to Y, meaning that on reception of a message
	// that fails data dictionary validation, a reject will be issued to the counter-party in responnse.
	//
	// Required: No
	//
	// Default: Y
	//
	// Valid Values:
	//  - Y
	//  - N
	RejectInvalidMessage string = "RejectInvalidMessage"

	// AllowUnknownMessageFields is set by default to N, meaning that non user-defined fields (field with tag < 5000)
	// will be rejected if they are not defined in the data dictionary,
	// or are present in messages they do not belong to.
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	AllowUnknownMessageFields string = "AllowUnknownMsgFields"

	// CheckUserDefinedFields if set to N, user-defined fields (field with tag >= 5000) will not be rejected
	// if they are not defined in the data dictionary, or are present in messages they do not belong to.
	//
	// Required: No
	//
	// Default: Y
	//
	// Valid Values:
	//  - Y
	//  - N
	CheckUserDefinedFields string = "ValidateUserDefinedFields"

	// ValidateFieldsOutOfOrder if set to N, fields that are out of order (i.e. body fields in the header, or header fields in the body)
	// will not be rejected. Useful for connecting to systems which do not properly order fields.
	//
	// Required: No
	//
	// Default: Y
	//
	// Valid Values:
	//  - Y
	//  - N
	ValidateFieldsOutOfOrder string = "ValidateFieldsOutOfOrder"

	// CheckLatency if set to Y, messages must be received from the counter-party within a defined number of seconds.
	// It is useful to turn this off if a system uses localtime for it's timestamps instead of GMT.
	//
	// Required: No
	//
	// Default: Y
	//
	// Valid Values:
	//  - Y
	//  - N
	CheckLatency string = "CheckLatency"

	// MaxLatency if CheckLatency is set to Y, this defines the number of seconds latency allowed for a message to be processed.
	// Value must be positive integer.
	//
	// Required: No
	//
	// Default: 120
	//
	// Valid Values:
	//  - Any positive integer
	MaxLatency string = "MaxLatency"
)

const (
	// Initiator-only settings.

	// ReconnectInterval defines the time between reconnection attempts in seconds.
	// Only used for initiators.
	// Value must be positive integer.
	//
	// Required: No
	//
	// Default: 30
	//
	// Valid Values:
	//  - Any positive integer
	ReconnectInterval string = "ReconnectInterval"

	// LogoutTimeout defines the number of seconds to wait for a logout response before disconnecting.
	// Only used for initiators.
	// Value must be positive integer.
	//
	// Required: No
	//
	// Default: 2
	//
	// Valid Values:
	//  - Any positive integer
	LogoutTimeout string = "LogoutTimeout"

	// LogonTimeout defines the number of seconds to wait for a logon response before disconnecting.
	// Only used for initiators.
	// Value must be positive integer.
	//
	// Required: No
	//
	// Default: 10
	//
	// Valid Values:
	//  - Any positive integer
	LogonTimeout string = "LogonTimeout"

	// HeartBtInt sets the FIX session heartbeat interval in seconds.
	// Only used for initiators (unless acceptor sets HeartBtIntOverride to Y).
	// Value must be positive integer.
	//
	// Required: Yes for initiators or acceptors with HeartBtIntOverride
	//
	// Default: None
	//
	// Valid Values:
	//  - Any positive integer
	HeartBtInt string = "HeartBtInt"

	// SocketConnectHost sets the host to attempt to connect to.
	// In config files you can also set SocketConnectHost<n> where n is a positive integer.
	// This allows for alternate socket hosts for connecting to a session for failover.
	// (i.e.) SocketConnectHost1, SocketConnectHost2... must be consecutive and have a matching SocketConnectPort<n>.
	//
	// Required: Yes for initiators
	//
	// Default: None
	//
	// Valid Values:
	//  - A valid IPv4 or IPv6 address or a domain name
	SocketConnectHost string = "SocketConnectHost"

	// SocketConnectPort sets the socket port for connecting to a session.
	// In config files you can also set SocketConnectPort<n> where n is a positive integer.
	// This allows for alternate socket ports for connecting to a session for failover.
	// (i.e.) SocketConnectPort1, SocketConnectPort2... must be consecutive and have a matching SocketConnectHost<n>.
	//
	// Required: Yes for initiators
	//
	// Default: None
	//
	// Valid Values:
	//  - A positive integer
	SocketConnectPort string = "SocketConnectPort"

	// SocketTimeout sets the duration of timeout for TLS handshake.
	// Only used for initiators.
	//
	// Example Values:
	//  - SocketTimeout=30s # 30 seconds
	//  - SocketTimeout=60m # 60 minutes
	//
	// Required: No
	//
	// Default: 0 (no timeout)
	//
	// Valid Values:
	//  - A valid go time.Duration
	SocketTimeout string = "SocketTimeout"

	// ProxyType sets the type of proxy server to connect to.
	// Only used for initiators.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - socks
	ProxyType string = "ProxyType"

	// ProxyHost provides the address of the proxy server to connect to.
	// Only used for initiators.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A valid IPv4 or IPv6 address or a domain name
	ProxyHost string = "ProxyHost"

	// ProxyPort provides the port of the proxy server to connect to.
	// Only used for initiators.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - Any positive integer
	ProxyPort string = "ProxyPort"

	// ProxyUser sets the username for the proxy server connection.
	// Only used for initiators.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - Any string
	ProxyUser string = "ProxyUser"

	// ProxyPassword sets the password for the proxy server connection.
	// Only used for initiators.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - Any string
	ProxyPassword string = "ProxyPassword"
)

const (
	// Acceptor-only settings.

	// SocketAcceptHost sets the address for listening on incoming connections.
	// Used for acceptors only.
	//
	// Common examples:
	//  - 127.0.0.1 (for testing on localhost)
	//  - 0.0.0.0 (explicitly specify every available interface)
	//
	// Required: No
	//
	// Default: Listens on all available network interfaces.
	//
	// Valid Values:
	//  - A valid IPv4 or IPv6 address or a domain name
	SocketAcceptHost string = "SocketAcceptHost"

	// SocketAcceptPort sets the socket port for listening to incoming connections.
	// Used for acceptors only.
	//
	// Required: Yes for acceptors
	//
	// Default: None
	//
	// Valid Values:
	//  - A positive integer, representing a valid open socket port
	SocketAcceptPort string = "SocketAcceptPort"

	// HeartBtIntOverride if set to Y, will use the HeartBtInt value in the acceptor's config file for the heartbeat interval rather than what the initiator dictates.
	// Used for acceptors only.
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	HeartBtIntOverride string = "HeartBtIntOverride"

	// UseTCPProxy if set to Y, use TCP proxy for servers listening behind HAProxy of Amazon ELB load balancers.
	// The server can then receive the address of the client instead of the load balancer's.
	// Used for acceptors only.
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	UseTCPProxy string = "UseTCPProxy"

	// DynamicSessions if set to Y, allows sessions to connect to this acceptor
	// without explicitly stating SenderCompID/TargetCompID in the config file for the acceptor.
	// Used for acceptors only.
	//
	// Use DynamicQualifier=Y if initiator client sessions are using the same SenderCompID.
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	DynamicSessions string = "DynamicSessions"

	// DynamicQualifier is used in conjunction with DynamicSessions.
	// If set to Y, allows sessions to connect to this acceptor
	// when initiator client sessions are using the same SenderCompID.
	// Used for acceptors only.
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	DynamicQualifier string = "DynamicQualifier"
)

const (
	// Security settings.

	// SocketPrivateKeyFile is the filepath for the private key to use for secure TLS connections.
	// Must be used with SocketCertificateFile.
	// Must contain PEM encoded data.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A filepath to a file with read access.
	SocketPrivateKeyFile string = "SocketPrivateKeyFile"

	// SocketCertificateFile is the filepath for the certificate to use for secure TLS connections.
	// Must be used with SocketPrivateKeyFile.
	// Must contain PEM encoded data.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A filepath to a file with read access.
	SocketCertificateFile string = "SocketCertificateFile"

	// SocketCAFile is an optional filepath for a root CA to use for secure TLS connections.
	// For acceptors, client certificates will be verified against this CA.
	// For initiators, clients will use the CA to verify the server certificate.
	// If not configurated, initiators will verify the server certificate using the host's root CA set.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A filepath to a file with read access.
	SocketCAFile string = "SocketCAFile"

	// SocketPrivateKeyBytes is an optional value containing raw bytes of a PEM
	// encoded private key to use for secure TLS communications.
	// Must be used with SocketCertificateBytes.
	// Must contain PEM encoded data.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - Raw bytes containing a valid PEM encoded private key.
	SocketPrivateKeyBytes string = "SocketPrivateKeyBytes"

	// SocketCertificateBytes is an optional value containing raw bytes of a PEM
	// encoded certificate to use for secure TLS communications.
	// Must be used with SocketPrivateKeyBytes.
	// Must contain PEM encoded data.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - Raw bytes containing a valid PEM encoded certificate.
	SocketCertificateBytes string = "SocketCertificateBytes"

	// SocketCABytes is an optional value containing raw bytes of a PEM encoded
	// root CA to use for secure TLS communications. For acceptors, client
	// certificates will be verified against this CA. For initiators, clients
	// will use the CA to verify the server certificate. If not configured,
	// initiators will verify the server certificates using the host's root CA
	// set.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - Raw bytes containing a valid PEM encoded CA.
	SocketCABytes string = "SocketCABytes"

	// SocketInsecureSkipVerify controls whether a client verifies the server's certificate chain and host name.
	// If SocketInsecureSkipVerify is set to Y, crypto/tls accepts any certificate presented by the server and any host name in that certificate.
	// In this mode, TLS is susceptible to machine-in-the-middle attacks unless custom verification is used.
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	SocketInsecureSkipVerify string = "SocketInsecureSkipVerify"

	// SocketServerName sets the expected server name on a returned certificate, unless SocketInsecureSkipVerify is true.
	// This is for the TLS Server Name Indication extension.
	// Only used for initiators.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - Any string
	SocketServerName string = "SocketServerName"

	// SocketMinimumTLSVersion specifies the minimum TLS version to use when creating a secure connection.
	//
	// Required: No
	//
	// Default: TLS12.
	//
	// Valid Values:
	//  - SSL30
	//  - TLS10
	//  - TLS11
	//  - TLS12
	SocketMinimumTLSVersion string = "SocketMinimumTLSVersion"

	// SocketUseSSL if set to Y, an initiator will use TLS even if client certificates are not present.
	// It is set to N by default, meaning TLS will not be used if SocketPrivateKeyFile or SocketCertificateFile are not supplied.
	//
	// This is used for initiators in the case where you do not need to provide client certificates, but still need to use
	// a TLS connection and should verify the server certificate(s).
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	SocketUseSSL string = "SocketUseSSL"
)

const (
	// Logging settings.

	// FileLogPath sets the directory path in which to write log files to.
	// This will create the directory path if it does not already exist.
	// FileLogPath is only relevant if also using quickfix.NewFileLogFactory(..) in code
	// when creating your LogFactory for your initiator or acceptor.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A valid path
	FileLogPath string = "FileLogPath"

	// SQLLogDriver sets the name of the database driver to use for application logs (see https://go.dev/wiki/SQLDrivers for the list of available drivers).
	// SQLLogDriver is only relevant if also using sql.NewLogFactory(..) in code
	// when creating your LogFactory for your initiator or acceptor.
	//
	// Required: Only if using a sql db as your Log
	//
	// Default: N/A
	//
	// Valid Values:
	//  - See https://go.dev/wiki/SQLDrivers
	SQLLogDriver string = "SQLLogDriver"

	// SQLLogDataSourceName sets the driver-specific data source name of the database to use for application logs.
	// This usually consists of at least a database name and connection information.
	// SQLLogDataSourceName is only relevant if also using sql.NewLogFactory(..) in code
	// when creating your LogFactory for your initiator or acceptor.
	//
	// See https://pkg.go.dev/database/sql#Open for more information.
	//
	// Required: Only if using a sql db as your Log.
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A string correspondinng to a datasource
	SQLLogDataSourceName string = "SQLLogDataSourceName"

	// SQLLogConnMaxLifetime sets the maximum duration of time that a database connection may be reused.
	// See https://pkg.go.dev/database/sql#DB.SetConnMaxLifetime for more information.
	//
	// If your database server has a config option to close inactive connections after some duration (e.g. MySQL "wait_timeout"),
	// set SQLLogConnMaxLifetime to a value less than that duration.
	//
	// Example Values:
	//  - SQLLogConnMaxLifetime=14400s # 14400 seconds
	//  - SQLLogConnMaxLifetime=2h45m  # 2 hours and 45 minutes
	//
	// SQLLogConnMaxLifetime is only relevant if also using sql.NewLogFactory(..) in code
	// when creating your LogFactory for your initiator or acceptor.
	//
	// Required: No
	//
	// Default: 0 (forever)
	//
	// Valid Values:
	//  - A valid go time.Duration
	SQLLogConnMaxLifetime string = "SQLLogConnMaxLifetime"

	// MongoLogConnection sets the MongoDB connection URL to use for application logs.
	//
	// See https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Connect for more information.
	//
	// MongoLogConnection is only relevant if also using mongo.NewLogFactory(..) in code
	// when creating your LogFactory for your initiator or acceptor.
	//
	// Required: Only if using MongoDB as your Log.
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A string representing a MongoDB connection
	MongoLogConnection string = "MongoLogConnection"

	// MongoLogDatabase sets the MongoDB-specific name of the database to use for application logs.
	//
	// See https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Connect for more information.
	//
	// MongoLogDatabase is only relevant if also using mongo.NewLogFactory(..) in code
	// when creating your LogFactory for your initiator or acceptor.
	//
	// Required: Only if using MongoDB as your Log.
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A string corresponding to a MongoDB database
	MongoLogDatabase string = "MongoLogDatabase"

	// MongoLogReplicaSet sets the MongoDB replica set to use for application logs.
	// This is optional.
	//
	// See https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Connect for more information.
	//
	// MongoLogReplicaSet is only relevant if also using mongo.NewLogFactory(..) in code
	// when creating your LogFactory for your initiator or acceptor.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A string corresponding to a MongoDB replica set
	MongoLogReplicaSet string = "MongoLogReplicaSet"
)

const (
	// Storage settings.

	// PersistMessages if set to N, no messages will be persisted.
	// This will force QuickFIX/Go to always send GapFills instead of resending messages.
	// Use this if you know you never want to resend a message.
	// This is useful for market data streams when logging all incoming messages is not important.
	//
	// Required: No
	//
	// Default: Y
	//
	// Valid Values:
	//  - Y
	//  - N
	PersistMessages string = "PersistMessages"

	// FileStorePath sets the directory path in which to write sequence number and message files.
	// This will create the directory path if it does not already exist.
	// FileStorePath is only relevant if also using file.NewStoreFactory(..) in code
	// when creating your MessageStoreFactory for your initiator or acceptor.
	//
	// Required: Only if using files as your MessageStore
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A valid path
	FileStorePath string = "FileStorePath"

	// FileStoreSync controls whether the FileStore syncs to the hard drive on every write.
	// It's safer to sync, but it's also much slower.
	// FileStoreSync is only relevant if also using file.NewStoreFactory(..) in code
	// when creating your MessageStoreFactory for your initiator or acceptor.
	//
	// Required: No
	//
	// Default: Y
	//
	// Valid Values:
	//  - Y
	//  - N
	FileStoreSync string = "FileStoreSync"

	// SQLStoreDriver sets the name of the database driver to use for message storage (see https://go.dev/wiki/SQLDrivers for the list of available drivers).
	// SQLStoreDriver is only relevant if also using sql.NewStoreFactory(..) in code
	// when creating your MessageStoreFactory for your initiator or acceptor.
	//
	// Required: Only if using a sql db as your MessageStore
	//
	// Default: N/A
	//
	// Valid Values:
	//  - See https://go.dev/wiki/SQLDrivers
	SQLStoreDriver string = "SQLStoreDriver"

	// SQLStoreDataSourceName sets the driver-specific data source name of the database to use for messagge storage.
	// This usually consists of at least a database name and connection information.
	// SQLStoreDataSourceName is only relevant if also using sql.NewStoreFactory(..) in code
	// when creating your MessageStoreFactory for your initiator or acceptor.
	//
	// See https://pkg.go.dev/database/sql#Open for more information.
	//
	// Required: Only if using a sql db as your MessageStore
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A string correspondinng to a datasource
	SQLStoreDataSourceName string = "SQLStoreDataSourceName"

	// SQLStoreConnMaxLifetime sets the maximum duration of time that a database connection may be reused.
	// See https://pkg.go.dev/database/sql#DB.SetConnMaxLifetime for more information.
	//
	// If your database server has a config option to close inactive connections after some duration (e.g. MySQL "wait_timeout"),
	// set SQLStoreConnMaxLifetime to a value less than that duration.
	//
	// Example Values:
	//  - SQLStoreConnMaxLifetime=14400s # 14400 seconds
	//  - SQLStoreConnMaxLifetime=2h45m  # 2 hours and 45 minutes
	//
	// SQLStoreConnMaxLifetime is only relevant if also using sql.NewStoreFactory(..) in code
	// when creating your MessageStoreFactory for your initiator or acceptor.
	//
	// Required: No
	//
	// Default: 0 (forever)
	//
	// Valid Values:
	//  - A valid go time.Duration
	SQLStoreConnMaxLifetime string = "SQLStoreConnMaxLifetime"

	// MongoStoreConnection sets the MongoDB connection URL to use for message storage.
	//
	// See https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Connect for more information.
	//
	// MongoStoreConnection is only relevant if also using mongo.NewStoreFactory(..) in code
	// when creating your MessageStoreFactory for your initiator or acceptor.
	//
	// Required: Only if using MongoDB as your MessageStore
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A string representing a MongoDB connection
	MongoStoreConnection string = "MongoStoreConnection"

	// MongoStoreDatabase sets the MongoDB-specific name of the database to use for message storage.
	//
	// See https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Connect for more information.
	//
	// MongoStoreDatabase is only relevant if also using mongo.NewStoreFactory(..) in code
	// when creating your MessageStoreFactory for your initiator or acceptor.
	//
	// Required: Only if using MongoDB as your MessageStore
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A string corresponding to a MongoDB database
	MongoStoreDatabase string = "MongoStoreDatabase"

	// MongoStoreReplicaSet sets the MongoDB replica set to use for message storage.
	// This is optional.
	//
	// See https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Connect for more information.
	//
	// MongoStoreReplicaSet is only relevant if also using mongo.NewStoreFactory(..) in code
	// when creating your MessageStoreFactory for your initiator or acceptor.
	//
	// Required: No
	//
	// Default: N/A
	//
	// Valid Values:
	//  - A string corresponding to a MongoDB replica set
	MongoStoreReplicaSet string = "MongoStoreReplicaSet"
)

const (
	// Misc settings.

	// ResendRequestChunkSize is a setting to limit the size of a resend request in case of missing messages.
	// This is useful when the remote FIX engine does not allow to ask for more than <n> messages for a ResendRequest.
	//
	// I.e. if the ResendRequestChunkSize is set to 5 and a gap of 7 messages is detected, a first resend request will be sent for 5 messages.
	// When this gap has been filled, another resend request for 2 messages will be sent.
	// If the ResendRequestChunkSize is set to 0, only one ResendRequest for all the missing messages will be sent.
	//
	// Required: No
	//
	// Default: 0 (do not chunk resends)
	//
	// Valid Values:
	//  - A positive integer
	ResendRequestChunkSize string = "ResendRequestChunkSize"

	// EnableLastMsgSeqNumProcessed tells the FIX engine to add the last message sequence number processed
	// to outgoing message headers (using optional tag 369).
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	EnableLastMsgSeqNumProcessed string = "EnableLastMsgSeqNumProcessed"

	// EnableNextExpectedMsgSeqNum tells the FIX engine to add tag NextExpectedMsgSeqNum (optional tag 789) on the
	// sent Logon message and use value of tag 789 on received Logon message to synchronize session.
	// This should not be enabled for FIX versions less than FIX.4.4.
	//
	// Required: No
	//
	// Default: N
	//
	// Valid Values:
	//  - Y
	//  - N
	EnableNextExpectedMsgSeqNum string = "EnableNextExpectedMsgSeqNum"
)
