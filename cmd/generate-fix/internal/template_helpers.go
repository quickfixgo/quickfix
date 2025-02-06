package internal

import (
	"fmt"

	"github.com/quickfixgo/quickfix/datadictionary"
)

func checkIfDecimalImportRequiredForFields(fTypes []*datadictionary.FieldType) (ok bool, err error) {
	var t string
	for _, fType := range fTypes {
		t, err = quickfixType(fType)
		if err != nil {
			return
		}

		if t == "FIXDecimal" {
			return true, nil
		}
	}

	return
}

func checkIfTimeImportRequiredForFields(fTypes []*datadictionary.FieldType) (ok bool, err error) {
	var t string
	for _, fType := range fTypes {
		t, err = quickfixType(fType)
		if err != nil {
			return
		}

		var vt string
		if vt, err = quickfixValueType(t); err != nil {
			return
		}

		if vt == "time.Time" {
			return true, nil
		}
	}

	return
}

func checkFieldDecimalRequired(f *datadictionary.FieldDef) (required bool, err error) {
	var globalType *datadictionary.FieldType
	if globalType, err = getGlobalFieldType(f); err != nil {
		return
	}

	var t string
	if t, err = quickfixType(globalType); err != nil {
		return
	}

	if t == "FIXDecimal" {
		required = true
		return
	}

	for _, groupField := range f.Fields {
		if required, err = checkFieldDecimalRequired(groupField); required || err != nil {
			return
		}
	}

	return
}

func checkFieldTimeRequired(f *datadictionary.FieldDef) (required bool, err error) {
	var globalType *datadictionary.FieldType
	if globalType, err = getGlobalFieldType(f); err != nil {
		return
	}

	var t string
	if t, err = quickfixType(globalType); err != nil {
		return
	}

	var vt string
	if vt, err = quickfixValueType(t); err != nil {
		return
	}

	if vt == "time.Time" {
		required = true
		return
	}

	for _, groupField := range f.Fields {
		if required, err = checkFieldTimeRequired(groupField); required || err != nil {
			return
		}
	}

	return
}

func collectStandardImports(m *datadictionary.MessageDef) (imports []string, err error) {
	var timeRequired bool
	for _, f := range m.Fields {
		if !timeRequired {
			if timeRequired, err = checkFieldTimeRequired(f); err != nil {
				return
			}
		}

		if timeRequired {
			break
		}
	}

	if timeRequired {
		imports = append(imports, "time")
	}

	return
}

func collectExtraImports(m *datadictionary.MessageDef) (imports []string, err error) {
	var decimalRequired bool
	for _, f := range m.Fields {
		if !decimalRequired {
			if decimalRequired, err = checkFieldDecimalRequired(f); err != nil {
				return
			}
		}

		if decimalRequired {
			break
		}
	}

	if decimalRequired {
		imports = append(imports, "github.com/shopspring/decimal")
	}

	return
}

func checkIfEnumImportRequired(m *datadictionary.MessageDef) (required bool, err error) {
	for _, f := range m.Fields {
		required, err = checkFieldEnumRequired(f)
		if err != nil || required {
			return
		}
	}

	return
}

func checkFieldEnumRequired(f *datadictionary.FieldDef) (required bool, err error) {
	var globalType *datadictionary.FieldType
	if globalType, err = getGlobalFieldType(f); err != nil {
		return
	}

	if globalType.Enums != nil && 0 < len(globalType.Enums) {
		var t string
		if t, err = quickfixType(globalType); err != nil {
			return
		}
		if t != "FIXBoolean" {
			required = true
			return
		}
	}

	for _, groupField := range f.Fields {
		if required, err = checkFieldEnumRequired(groupField); required || err != nil {
			return
		}
	}

	return
}

func quickfixValueType(quickfixType string) (goType string, err error) {
	switch quickfixType {
	case "FIXString":
		goType = "string"
	case "FIXBoolean":
		goType = "bool"
	case "FIXInt":
		goType = "int"
	case "FIXUTCTimestamp":
		goType = "time.Time"
	case "FIXFloat":
		goType = "float64"
	case "FIXDecimal":
		goType = "decimal.Decimal"
	default:
		err = fmt.Errorf("Unknown QuickFIX Type: %v", quickfixType)
	}

	return
}

func quickfixType(field *datadictionary.FieldType) (quickfixType string, err error) {
	// Define mappings of FIX field types to QuickFIX types
	stringTypes := map[string]bool{
		"MULTIPLESTRINGVALUE": true, "MULTIPLEVALUESTRING": true, "MULTIPLECHARVALUE": true,
		"CHAR": true, "CURRENCY": true, "DATA": true, "MONTHYEAR": true,
		"LOCALMKTTIME": true, "LOCALMKTDATE": true, "TIME": true, "DATE": true,
		"EXCHANGE": true, "LANGUAGE": true, "XMLDATA": true, "COUNTRY": true,
		"UTCTIMEONLY": true, "UTCDATE": true, "UTCDATEONLY": true, "TZTIMEONLY": true,
		"TZTIMESTAMP": true, "XID": true, "XIDREF": true, "STRING": true,
	}
	intTypes := map[string]bool{
		"LENGTH": true, "DAYOFMONTH": true, "NUMINGROUP": true,
		"SEQNUM": true, "TAGNUM": true, "INT": true,
	}

	floatTypes := map[string]bool{
		"QTY": true, "QUANTITY": true, "AMT": true,
		"PRICE": true, "PRICEOFFSET": true, "PERCENTAGE": true, "FLOAT": true,
	}

	switch {
	case stringTypes[field.Type]:
		return "FIXString", nil
	case intTypes[field.Type]:
		return "FIXInt", nil
	case field.Type == "BOOLEAN":
		return "FIXBoolean", nil
	case field.Type == "UTCTIMESTAMP":
		return "FIXUTCTimestamp", nil
	case floatTypes[field.Type]:
		if *useFloat {
			return "FIXFloat", nil
		}
		return "FIXDecimal", nil
	default:
		return "", fmt.Errorf("Unknown type '%v' for tag '%v'\n", field.Type, field.Tag())
	}
}

func requiredFields(m *datadictionary.MessageDef) (required []*datadictionary.FieldDef) {
	for _, part := range m.RequiredParts() {
		if part.Required() {
			switch pType := part.(type) {
			case *datadictionary.FieldDef:
				if !pType.IsGroup() {
					required = append(required, pType)
				}
			case *datadictionary.Component:
				for _, f := range pType.RequiredFields() {
					if !f.IsGroup() {
						required = append(required, f)
					}
				}
			}
		}
	}

	return
}

func beginString(spec *datadictionary.DataDictionary) string {
	if spec.FIXType == "FIXT" || spec.Major == 5 {
		return "FIXT.1.1"
	}

	return fmt.Sprintf("FIX.%v.%v", spec.Major, spec.Minor)
}

func routerBeginString(spec *datadictionary.DataDictionary) (routerBeginString string) {
	switch {
	case spec.FIXType == "FIXT":
		routerBeginString = "FIXT.1.1"
	case spec.Major != 5 && spec.ServicePack == 0:
		routerBeginString = fmt.Sprintf("FIX.%v.%v", spec.Major, spec.Minor)
		// ApplVerID enums.
	case spec.Major == 2:
		routerBeginString = "0"
	case spec.Major == 3:
		routerBeginString = "1"
	case spec.Major == 4 && spec.Minor == 0:
		routerBeginString = "2"
	case spec.Major == 4 && spec.Minor == 1:
		routerBeginString = "3"
	case spec.Major == 4 && spec.Minor == 2:
		routerBeginString = "4"
	case spec.Major == 4 && spec.Minor == 3:
		routerBeginString = "5"
	case spec.Major == 4 && spec.Minor == 4:
		routerBeginString = "6"
	case spec.Major == 5 && spec.Minor == 0 && spec.ServicePack == 0:
		routerBeginString = "7"
	case spec.Major == 5 && spec.Minor == 0 && spec.ServicePack == 1:
		routerBeginString = "8"
	case spec.Major == 5 && spec.Minor == 0 && spec.ServicePack == 2:
		routerBeginString = "9"
	}

	return
}
