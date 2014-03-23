package main

import (
	"flag"
	"fmt"
	"github.com/cbusbey/quickfixgo/gen"
	"github.com/cbusbey/quickfixgo/spec"
	"os"
	"sort"
)

var (
	fieldMap     map[string]int
	fieldTypeMap map[string]spec.FieldType
	sortedTags   []string
	pkg          = "fix"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: generate-fields [flags] <path to data dictionary> ... \n")
	flag.PrintDefaults()
	os.Exit(2)
}

func genFields() {
	fileOut := "package field\n"
	fileOut += "import(\"github.com/cbusbey/quickfixgo/tag\")\n"

	for _, tag := range sortedTags {
		field := fieldTypeMap[tag]

		baseType := ""
		switch field.Type {
		case "STRING":
			baseType = "StringValue"
		case "MULTIPLESTRINGVALUE", "MULTIPLEVALUESTRING":
			baseType = "MultipleStringValue"
		case "MULTIPLECHARVALUE":
			baseType = "MultipleCharValue"
		case "CHAR":
			baseType = "CharValue"
		case "CURRENCY":
			baseType = "CurrencyValue"
		case "DATA":
			baseType = "DataValue"
		case "MONTHYEAR":
			baseType = "MonthYearValue"
		case "LOCALMKTDATE":
			baseType = "LocalMktDateValue"
		case "EXCHANGE":
			baseType = "ExchangeValue"
		case "LANGUAGE":
			baseType = "LanguageValue"
		case "XMLDATA":
			baseType = "XMLDataValue"
		case "COUNTRY":
			baseType = "CountryValue"
		case "UTCTIMEONLY":
			baseType = "UTCTimeOnlyValue"
		case "UTCDATEONLY":
			baseType = "UTCDateOnlyValue"
		case "TZTIMEONLY":
			baseType = "TZTimeOnlyValue"
		case "TZTIMESTAMP":
			baseType = "TZTimestampValue"
		case "BOOLEAN":
			baseType = "BooleanValue"
		case "INT":
			baseType = "IntValue"
		case "LENGTH":
			baseType = "LengthValue"
		case "DAYOFMONTH":
			baseType = "DayOfMonthValue"
		case "NUMINGROUP":
			baseType = "NumInGroupValue"
		case "SEQNUM":
			baseType = "SeqNumValue"
		case "UTCTIMESTAMP":
			baseType = "UTCTimestampValue"
		case "FLOAT":
			baseType = "FloatValue"
		case "QTY":
			baseType = "QtyValue"
		case "AMT":
			baseType = "AmtValue"
		case "PRICE":
			baseType = "PriceValue"
		case "PRICEOFFSET":
			baseType = "PriceOffsetValue"
		case "PERCENTAGE":
			baseType = "PercentageValue"
		default:
			fmt.Printf("Unknown type '%v' for tag '%v'\n", field.Type, tag)
		}

		fileOut += fmt.Sprintf("type %v struct { %v }\n", field.Name, baseType)
		fileOut += fmt.Sprintf("func (f %v) Tag() tag.Tag {return tag.%v}\n", field.Name, field.Name)
	}

	gen.WriteFile("field/fields.go", fileOut)
}

func genTags() {
	fileOut := "package tag\n"

	fileOut += "const (\n"
	for _, tag := range sortedTags {
		fileOut += fmt.Sprintf("%v Tag = %v\n", tag, fieldMap[tag])
	}
	fileOut += ")\n"

	gen.WriteFile("tag/tag_numbers.go", fileOut)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}

	fieldMap = make(map[string]int)
	fieldTypeMap = make(map[string]spec.FieldType)

	for _, dataDict := range flag.Args() {
		spec, err := spec.ParseFixSpec(dataDict)

		if err != nil {
			panic(err)
		}

		for _, field := range spec.FieldTypes {
			fieldMap[field.Name] = field.Number
			fieldTypeMap[field.Name] = field
		}
	}

	sortedTags = make([]string, len(fieldMap))
	i := 0
	for f, _ := range fieldMap {
		sortedTags[i] = f
		i++
	}
	sort.Strings(sortedTags)

	genTags()
	genFields()
}
