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
