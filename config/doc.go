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

TargetCompID

Counter parties ID as associated with this FIX session. Value is case-sensitive alpha-numeric string.

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

ResetOnLogon

Determines if sequence numbers should be reset when receiving a logon request. Acceptors only.	Valid Values:
 Y
 N

Defaults to N.

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

ReconnectInterval

Time between reconnection attempts in seconds. Only used for initiators.    Value must be positive integer.

Defaults to 30

HeartBtInt

Heartbeat interval in seconds. Only used for initiators.	Value must be positive integer.

SocketConnectPort

Socket port for connecting to a session. Only used for initiators. Must be positive integer

SocketConnectHost

Host to connect to. Only used for initiators. Value must be valid IP address in the format of x.x.x.x or a domain name

SocketAcceptPort

Socket port for listening to incoming connections, Only used with for acceptors. Value must be	positive integer, valid open socket port.

FileLogPath

Directory to store logs.	Value must be valid directory for storing files, application must have write access.

FileStorePath

Directory to store sequence number and message files.  Only used with FileStoreFactory.

SQLDriver

The name of the database driver to use (see https://github.com/golang/go/wiki/SQLDrivers for the list of available drivers).  Only used with SqlStoreFactory.

SQLDataSourceName

The driver-specific data source name of the database to use.  Only used with SqlStoreFactory.

SQLConnMaxLifetime

SetConnMaxLifetime sets the maximum duration of time that a database connection may be reused (see https://golang.org/pkg/database/sql/#DB.SetConnMaxLifetime).  Defaults to zero, which causes connections to be reused forever.  Only used with SqlStoreFactory.

If your database server has a config option to close inactive connections after some duration (e.g. MySQL "wait_timeout"), set SQLConnMaxLifetime to a value less than that duration.

Example Values:
 SQLConnMaxLifetime=14400s # 14400 seconds
 SQLConnMaxLifetime=2h45m  # 2 hours and 45 minutes
*/
package config
