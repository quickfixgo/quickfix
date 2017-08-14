## 0.6.0 (August 14, 2017)

FEATURES

* CheckLatency [GH 241, 266]
* ResendRequestChunkSize [GH 243, 245]
* EnableLastMsgSeqNumProcessed  [GH 253, 256]

ENHANCEMENTS

* config.SocketTLSForTesting(bool) [GH 235]
* API Interface Enhancements [GH 251, 252, 254, 255, 257, 258, 259]
* Misc. Performance Optimizations [GH 260, 261, 263, 264, 265, 268, 270, 271, 272, 273, 274, 275]
* TLS Configuration [GH 279, 280]
* Use data dictionary for parsing fix messages [GH 281]

BUG FIXES

* Resend logon fix [GH 244]
* PossDup messages with seqnum too low should not be sent to FromCallbacks [GH 246, 247]
* Router should not reject admin messages or business rejects [GH 249, 250]
* Fixes file log output for incoming, outgoing [GH 262]
* message.String() returns rawMessage if set, builds otherwise [GH 269]
* Use timestamp with time zone for postgres sql [GH 286]


## 0.5.0 (September 1, 2016)

FEATURES

* Session Scheduling [GH 31, 195, 196, 197, 198, 200, 201, 202, 203, 204, 205, 211, 212, 218, 220]
* TimeZone configuration [GH 206]
* StartDay, EndDay for week long sessions [GH 207, 239]
* Support for connection over SSL [GH 63, 193]
* SocketConnectHost/Port<n> [GH 65, 217]
* ResetOnLogout Configuration enhancement [GH 67, 192]
* SocketAcceptAddress, ipv6 support [GH 83, 215]
* RefreshOnLogon [GH 214]

ENHANCEMENTS

* Reject Logon Support [GH 57, 225]
* FIX Decimal [GH 114, 223, 224, 227, 228]
* test refactoring, leveraging testify mock assertions [GH 186]
* KISS on registry, session management [GH 208]
* moving and renaming test jigs [GH 210]
* Config Enhancements [GH 216]
* travis uses go 1.7 [GH 219]
* adds 'generated.go' suffix to generated source [GH 221]
* vendored deps GH [GH 222, 226]
* renames SQLStore config settings [GH 229]
* Add FieldMap.{SetInt, SetString} [GH 236]


BUG FIXES

* DefaultApplVerID Configuration not translating enum names [GH 89, 213]
* Dropped issues in logout state [GH 187, 188]
* correctly increments in target seq num on logout for in session state [GH 189]
* SendToTarget should return an error if toApp returns DoNotSend [GH 190, 191]
* Logon fix [GH 194]
* onlogout calls depend on session state [GH 199]
* fixes bug in resend state where resend response is processed incomplete [GH 230]
* fixes logic around logon message with sequence number too high [GH 231]
* SequenceReset, Resent Messages not hitting ToAdmin/ToApp [GH 232, 233]
* Next MsgSeqNo after received ResetSeqNumFlag=Y [GH 238, 240]


## 0.4.0 (August 1, 2016)

FEATURES

* ReconnectInterval Initiator Configuration [GH 66, 161]

ENHANCEMENTS

* Generate Getters for messages [GH 118]
* Drop generate-* apps into cmd/ [GH 125, 150]
* Misc field type refactoring  [GH 145]
* Cmd gen [GH 146, 147]
* refactoring enum generation [GH 148]
* pipelining generation [GH 149]
* Sending a message to an unlogged-in session, results an error [GH 173, 182]
* adds event logging related to session events [GH 175]
* Error handling around session code enhancement [GH 176]
* refactoring of session code [GH 183]
* Gen header beginstring [GH 184]

BUG FIXES

* Do not incr target seq num when seq num too high  [GH 151]
* can't parse securitylist message [GH 153, 155] 
* Concurrent map operations on Acceptor.Stop()  [GH 156]
* Return requiredConfigurationMissing when "FileStorePath" not found [GH 157]
* Checks around HeartBtInt configuration [GH 158]
* Removed complexity around closing Initiator sessions [GH 159]
* Proper FIX logout sequence [GH 160]
* Session logic doesn't account for a failure when calling messageStore methods (e.g. persisting a message) [GH 162]
* Session event loop [GH 164]
* Session event loop follow up [GH 165]
* Handle OnLogon/Logout callbacks in user space [GH 167]
* Session deadlocks if both initiator and acceptor enter the resend state [GH 169, 179]
* Initiator.Stop() does not wait for Acceptor's logout response, causing a resend on next logon [GH 170, 172]
* Reset peer timer after logon [GH 171]
* Ensure OnLogon is called even if seq num too high [GH 174]
* increment target seq num on logout [GH 177]
* fixes bogus resend logic [GH 178]
* Session forgets it is in resend state [GH 180, 181]
* fixes donotsend logic [GH 185]

## 0.3.0 (June 3, 2016)

FEATURES

* ValidateFieldsOutOfOrder Configuration Option [GH 107]

ENHANCEMENTS

* Datadictionary tests [GH 108]
* Proposal to add public method to convert raw fix message to quickfix.Message [GH 110]
* Make gen package public [GH 113]
* Make gen import path relative [GH 119]
* Validator interface [GH 120]
* Expose IncorrectDataFormat [GH 121]
* Spelling, fmtness [GH 123]
* More verbose error text on conditionally required field BMR [GH 131]
* better error handling in gen package [GH 134]
* Include timestamp in messages file log [GH 135]
* extracts repeating group interface [GH 137]
* Header, Body, Trailer FieldMap types [GH 138]
* Datadictionary/Gen refactor [GH 140]
* Gen value timestamp [GH 141]
* Slight gen revert [GH 142]
* replaced regex with faster impl for float parsing [GH 143]
* Expose sql.DB's SetConnMaxLifetime() in settings [GH 144, 139]

BUG FIXES

* Routing Incorrect for FIXT [GH 101]
* Validation Error for component tag [GH 102]
* Unmarshal error for repeating Group [GH 103]
* Marshaler/Reflector tries to convert time.Time struct into fix format [GH 109]
* nil pointer panic if no config.DataDictionary specified [GH 127]
* fixes required group validation error [GH 129]
* RefTagID is not known to BusinessMessageReject type [GH 130]

## 0.2.0 (April 19, 2016)

FEATURES

* Persisted Store [GH 32]
* Initializer Constructors for Generated Messages [GH 69]
* Field Setters for generated messages [GH 70]
* Support for optional components, component setters [GH 79]
* Setters for repeating groups in components [GH 87]
* DB Store [GH 88]

ENHANCEMENTS

* Gen refactor [GH 78]
* Refactoring data dictionary pkg [GH 93]

BUG FIXES

* Initiator panic if connection closed [GH 59]
* New logs overrides old ones [GH 72]
* Session sending message timeout [GH 80]
* Potential FIX50 Market Data marshaling bug [GH 91]
* Allow settings values to contain equals signs [GH 97]
* Error when trying to unmarshal FIX message (FIX 5.0) [GH 99]

## 0.1.0 (February 21, 2016)

* Initial release
