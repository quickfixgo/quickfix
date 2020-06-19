/*
Package config declares application and session settings for QuickFIX/Go

BeginString

Version of FIX this session should use. Valid values:

      FIXT.1.1
      FIX.4.4
      FIX.4.3
      FIX.4.2
      FIX.4.1
      FIX.4.0

SenderCompID

Your ID as associated with this FIX session. Value is	case-sensitive alpha-numeric string.

SenderSubID

(Optional) Your subID as associated with this FIX session.  Value is case-sensitive alpha-numeric string.

SenderLocationID

(Optional) Your locationID as associated with this FIX session.  Value is case-sensitive alpha-numeric string.

TargetCompID

Counter parties ID as associated with this FIX session. Value is case-sensitive alpha-numeric string.

TargetSubID

(Optional) Counterparty's subID as associated with this FIX session.  Value is case-sensitive alpha-numeric string.

TargetLocationID

(Optional) Counterparty's locationID as associated with this FIX session. Value is case-sensitive alpha-numeric string.

SessionQualifier

Additional qualifier to disambiguate otherwise identical sessions. Value is	case-sensitive alpha-numeric string.

DefaultApplVerID

Required only for FIXT 1.1 (and newer). Ignored for earlier transport versions. Specifies the default application version ID for the session. This can either be the ApplVerID enum (see the ApplVerID field) or the BeginString for the default version.	Valid Values:

  FIX.5.0SP2
  FIX.5.0SP1
  FIX.5.0
  FIX.4.4
  FIX.4.3
  FIX.4.2
  FIX.4.1
  FIX.4.0
  9
  8
  7
  6
  5
  4
  3
  2

TimeZone

Time zone for this session; if specified, the session start and end will be converted from this zone to UTC.  Valid Values:

 IANA Time zone ID (America/New_York, Asia/Tokyo, Europe/London, etc.)
 Local (The Zone on host)

Defaults to UTC.

StartTime

Time of day that this FIX session becomes activated. Valid Values:

 time in the format of HH:MM:SS, time is represented in time zone configured by TimeZone

EndTime

Time of day that this FIX session becomes deactivated. Valid Values:

 time in the format of HH:MM:SS, time is represented in time zone configured by TimeZone

StartDay

For week long sessions, the starting day of week for the session. Use in combination with StartTime. Valid Values:

 Full day of week in English, or 3 letter abbreviation (i.e. Monday and Mon are valid)

EndDay

For week long sessions, the ending day of week for the session. Use in combination with EndTime. Valid Values:

 Full day of week in English, or 3 letter abbreviation (i.e. Monday and Mon are valid)

EnableLastMsgSeqNumProcessed

Add the last message sequence number processed in the header (optional tag 369).  Valid Values:
 Y
 N

Defaults to N.

ResendRequestChunkSize

Setting to limit the size of a resend request in case of missing messages. This is useful when the remote FIX engine does not allow to ask for more than n message for a ResendRequest.  E.g. if the ResendRequestChunkSize is set to 5 and a gap of 7 messages is detected, a first resend request will be sent for 5 messages. When this gap has been filled, another resend request for 2 messages will be sent. If the ResendRequestChunkSize is set to 0, only one ResendRequest for all the missing messages will be sent. Value must be positive integer. Defaults to 0 (disables splitting).

ResetOnLogon

Determines if sequence numbers should be reset when receiving a logon request. Acceptors only.	Valid Values:
 Y
 N

Defaults to N.

ResetOnLogout

Determines if sequence numbers should be reset to 1 after a normal logout termination. Valid Values:
 Y
 N

Defaults to N.

ResetOnDisconnect

Determines if sequence numbers should be reset to 1 after an abnormal termination. Valid Values:
 Y
 N

Defaults to N.

RefreshOnLogon

Determines if session state should be restored from persistence layer when logging on. Useful for creating hot failover sessions. Valid Values:
 Y
 N

Defaults to N.

TimeStampPrecision

Determines precision for timestamps in (Orig)SendingTime fields that are sent out. Only available for FIX.4.2 and greater, FIX versions earlier than FIX.4.2 will use timestamp resolution in seconds. Valid Values:
 SECONDS
 MILLIS
 MICROS
 NANOS

Defaults to MILLIS.

Validation

The following settings are specific to message validation.

DataDictionary

XML definition file for validating incoming FIX messages. If no DataDictionary is supplied, only basic message validation will be done.

This setting should only be used with FIX transport versions older than FIXT.1.1. See TransportDataDictionary and AppDataDictionary for FIXT.1.1 settings. Value must be a valid XML data dictionary file. QuickFIX/Go comes with the following defaults in the spec directory

  FIX44.xml
  FIX43.xml
  FIX42.xml
  FIX41.xml
  FIX40.xml

TransportDataDictionary

XML definition file for validating admin (transport) messages. This setting is only valid for FIXT.1.1 (or newer) sessions.  See DataDictionary for older transport versions (FIX.4.0 - FIX.4.4) for additional information. Value must be a valid XML data dictionary file, QuickFIX/Go comes with the following defaults in the spec directory

 FIXT1.1.xml

AppDataDictionary

XML definition file for validating application messages. This setting is only valid for FIXT.1.1 (or newer) sessions. See DataDictionary for older transport versions (FIX.4.0 - FIX.4.4) for additional information.

This setting supports the possibility of a custom application data dictionary for each session. This setting would only be used with FIXT 1.1 and new transport protocols. This setting can be used as a prefix to specify multiple application dictionaries for the FIXT transport. For example:

 DefaultApplVerID=FIX.4.2
 # For default application version ID
 AppDataDictionary=FIX42.xml
 # For nondefault application version ID
 # Use BeginString suffix for app version
 AppDataDictionary.FIX.4.4=FIX44.xml

Value must be a valid XML data dictionary file. QuickFIX/Go comes with the following defaults in the spec directory

 FIX50SP2.xml
 FIX50SP1.xml
 FIX50.xml
 FIX44.xml
 FIX43.xml
 FIX42.xml
 FIX41.xml
 FIX40.xml

ValidateFieldsOutOfOrder

If set to N, fields that are out of order (i.e. body fields in the header, or header fields in the body) will not be rejected. Useful for connecting to systems which do not properly order fields. Valid Values:
 Y
 N

Defaults to Y.

RejectInvalidMessage

If RejectInvalidMessage is set to N, zero errors will be thrown on reception of message that fails data dictionary validation. Valid Values:
 Y
 N

Defaults to Y.

CheckLatency

If set to Y, messages must be received from the counterparty within a defined number of seconds. It is useful to turn this off if a system uses localtime for it's timestamps instead of GMT. Valid Values:
 Y
 N

Defaults to Y.

MaxLatency

If CheckLatency is set to Y, this defines the number of seconds latency allowed for a message to be processed. Value must be positive integer.

Defaults to 120.

ReconnectInterval

Time between reconnection attempts in seconds. Only used for initiators.    Value must be positive integer.

Defaults to 30

LogoutTimeout

Session setting for logout timeout in seconds. Only used for initiators. Value must be positive integer.

Defaults to 2

LogonTimeout

Session setting for logon timeout in seconds. Only used for initiators. Value must be positive integer.

Defaults to 10

HeartBtInt

Heartbeat interval in seconds. Only used for initiators.	Value must be positive integer.

SocketConnectPort

Socket port for connecting to a session. Only used for initiators. Must be positive integer

SocketConnectHost

Host to connect to. Only used for initiators. Value must be a valid IPv4 or IPv6 address or a domain name

SocketConnectPort<n>

Alternate socket ports for connecting to a session for failover, where n is a positive integer. (i.e.) SocketConnectPort1, SocketConnectPort2... must be consecutive and have a matching SocketConnectHost[n]. Value must be a positive integer.

SocketConnectHost<n>

Alternate socket hosts for connecting to a session for failover, where n is a positive integer. (i.e.) SocketConnectHost1, SocketConnectHost2... must be consecutive and have a matching SocketConnectPort[n]. Value must be a valid IPv4 or IPv6 address or a domain name

SocketTimeout

Duration of timeout for TLS handshake. Only used for initiators.

Example Values:
 SocketTimeout=30s # 30 seconds
 SocketTimeout=60m # 60 minutes

Defaults to 0(means nothing timeout).

SocketAcceptHost

Socket host address for listening on incoming connections, only used for acceptors. By default acceptors listen on all available interfaces.

SocketAcceptPort

Socket port for listening to incoming connections, only used for acceptors. Value must be a positive integer, valid open socket port.

SocketPrivateKeyFile

Private key to use for secure TLS connections.  Must be used with SocketCertificateFile.

SocketCertificateFile

Certificate to use for secure TLS connections. Must be used with SocketPrivateKeyFile.

SocketCAFile

Optional root CA to use for secure TLS connections. For acceptors, client certificates will be verified against this CA.  For initiators, clients will use the CA to verify the server certificate. If not configurated, initiators will verify the server certificate using the host's root CA set.

SocketServerName

The expected server name on a returned certificate, unless SocketInsecureSkipVerify is true. This is for the TLS Server Name Indication extension. Initiator only.

SocketMinimumTLSVersion

Specify the Minimum TLS version to use when creating a secure connection. The valid choices are SSL30, TLS10, TLS11, TLS12. Defaults to TLS12.

SocketUseSSL

Use SSL for initiators even if client certificates are not present. If set to N or omitted, TLS will not be used if SocketPrivateKeyFile or SocketCertificateFile are not supplied.

ProxyType

Proxy type. Valid Values:
 socks

ProxyHost

Proxy server IP address in the format of x.x.x.x or a domain name

ProxyPort

Proxy server port

ProxyUser

Proxy user

ProxyPassword

Proxy password

UseTCPProxy

Use TCP proxy for servers listening behind HAProxy of Amazon ELB load balancers. The server can then receive the address of the client instead of the load balancer's.	Valid Values:
 Y
 N

PersistMessages

If set to N, no messages will be persisted. This will force QuickFIX/Go to always send GapFills instead of resending messages. Use this if you know you never want to resend a message. Useful for market data streams.  Valid Values:
 Y
 N

Defaults to Y.

FileLogPath

Directory to store logs.	Value must be valid directory for storing files, application must have write access.

FileStorePath

Directory to store sequence number and message files.  Only used with FileStoreFactory.

MongoStoreConnection

The MongoDB connection URL to use (see https://godoc.org/github.com/globalsign/mgo#Dial for the URL Format).  Only used with MongoStoreFactory.

MongoStoreDatabase

The MongoDB-specific name of the database to use.  Only used with MongoStoreFactory.

SQLStoreDriver

The name of the database driver to use (see https://github.com/golang/go/wiki/SQLDrivers for the list of available drivers).  Only used with SqlStoreFactory.

SQLStoreDataSourceName

The driver-specific data source name of the database to use.  Only used with SqlStoreFactory.

SQLStoreConnMaxLifetime

SetConnMaxLifetime sets the maximum duration of time that a database connection may be reused (see https://golang.org/pkg/database/sql/#DB.SetConnMaxLifetime).  Defaults to zero, which causes connections to be reused forever.  Only used with SqlStoreFactory.

If your database server has a config option to close inactive connections after some duration (e.g. MySQL "wait_timeout"), set SQLConnMaxLifetime to a value less than that duration.

Example Values:
 SQLConnMaxLifetime=14400s # 14400 seconds
 SQLConnMaxLifetime=2h45m  # 2 hours and 45 minutes
*/
package config
