package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/gen"
)

var (
	fieldMap     map[string]int
	fieldTypeMap map[string]*datadictionary.FieldType
	sortedTags   []string
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: generate-fields [flags] <path to data dictionary> ... \n")
	flag.PrintDefaults()
	os.Exit(2)
}

func genEnums() {
	fileOut := "package enum\n"

	for _, fieldName := range sortedTags {
		fieldType, _ := fieldTypeMap[fieldName]
		if len(fieldType.Enums) == 0 {
			continue
		}

		sortedEnums := make([]string, len(fieldType.Enums))
		i := 0
		for enum := range fieldType.Enums {
			sortedEnums[i] = enum
			i++
		}
		sort.Strings(sortedEnums)

		fileOut += fmt.Sprintf("//Enum values for %v\n", fieldName)
		fileOut += "const(\n"
		for _, enumVal := range sortedEnums {
			enum, _ := fieldType.Enums[enumVal]
			fileOut += fmt.Sprintf("%v_%v = \"%v\"\n", fieldName, enum.Description, enum.Value)
		}
		fileOut += ")\n"
	}

	gen.WriteFile("enum/enums.go", fileOut)
}

func genFields() {
	fileOut := "package field\n"
	fileOut += "import(\n"
	fileOut += "\"github.com/quickfixgo/quickfix\"\n"
	fileOut += "\"" + gen.GetImportPathRoot() + "/tag\"\n"
	fileOut += ")\n"

	for _, tag := range sortedTags {
		field := fieldTypeMap[tag]

		baseType := ""
		goType := ""
		switch field.Type {
		case "MULTIPLESTRINGVALUE", "MULTIPLEVALUESTRING":
			fallthrough
		case "MULTIPLECHARVALUE":
			fallthrough
		case "CHAR":
			fallthrough
		case "CURRENCY":
			fallthrough
		case "DATA":
			fallthrough
		case "MONTHYEAR":
			fallthrough
		case "LOCALMKTDATE":
			fallthrough
		case "TIME":
			fallthrough
		case "DATE":
			fallthrough
		case "EXCHANGE":
			fallthrough
		case "LANGUAGE":
			fallthrough
		case "XMLDATA":
			fallthrough
		case "COUNTRY":
			fallthrough
		case "UTCTIMEONLY":
			fallthrough
		case "UTCDATE":
			fallthrough
		case "UTCDATEONLY":
			fallthrough
		case "TZTIMEONLY":
			fallthrough
		case "TZTIMESTAMP":
			fallthrough
		case "STRING":
			baseType = "FIXString"
			goType = "string"

		case "BOOLEAN":
			baseType = "FIXBoolean"
			goType = "bool"

		case "LENGTH":
			fallthrough
		case "DAYOFMONTH":
			fallthrough
		case "NUMINGROUP":
			fallthrough
		case "SEQNUM":
			fallthrough
		case "INT":
			baseType = "FIXInt"
			goType = "int"

		case "UTCTIMESTAMP":
			baseType = "FIXUTCTimestamp"

		case "QTY":
			fallthrough
		case "QUANTITY":
			fallthrough
		case "AMT":
			fallthrough
		case "PRICE":
			fallthrough
		case "PRICEOFFSET":
			fallthrough
		case "PERCENTAGE":
			fallthrough
		case "FLOAT":
			baseType = "FIXFloat"
			goType = "float64"

		default:
			fmt.Printf("Unknown type '%v' for tag '%v'\n", field.Type, tag)
		}

		fileOut += fmt.Sprintf("//%vField is a %v field\n", field.Name(), field.Type)
		fileOut += fmt.Sprintf("type %vField struct { quickfix.%v }\n", field.Name(), baseType)
		fileOut += fmt.Sprintf("//Tag returns tag.%v (%v)\n", field.Name(), field.Tag())
		fileOut += fmt.Sprintf("func (f %vField) Tag() quickfix.Tag {return tag.%v}\n", field.Name(), field.Name())

		switch goType {
		case "bool", "int", "float64", "string":
			fileOut += fmt.Sprintf("//New%v returns a new %vField initialized with val\n", field.Name(), field.Name())
			fileOut += fmt.Sprintf("func New%v(val %v) *%vField {\n", field.Name(), goType, field.Name())
			fileOut += fmt.Sprintf("return &%vField{quickfix.%v(val)}\n", field.Name(), baseType)
			fileOut += "}\n"
		}
	}

	gen.WriteFile("field/fields.go", fileOut)
}

func genTags() {
	fileOut := "package tag\n"
	fileOut += "import(\"github.com/quickfixgo/quickfix\")\n"

	fileOut += "const (\n"
	for _, tag := range sortedTags {
		fileOut += fmt.Sprintf("%v quickfix.Tag = %v\n", tag, fieldMap[tag])
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
	fieldTypeMap = make(map[string]*datadictionary.FieldType)

	for _, dataDict := range flag.Args() {
		spec, err := datadictionary.Parse(dataDict)

		if err != nil {
			fmt.Println(dataDict)
			panic(err)
		}

		for _, field := range spec.FieldTypeByTag {
			fieldMap[field.Name()] = field.Tag()

			if oldField, ok := fieldTypeMap[field.Name()]; ok {
				//merge old enums with new
				if len(oldField.Enums) > 0 && field.Enums == nil {
					field.Enums = make(map[string]datadictionary.Enum)
				}

				for enumVal, enum := range oldField.Enums {
					if _, ok := field.Enums[enumVal]; !ok {
						//Verify an existing enum doesn't have the same description. Keep newer enum
						okToKeepEnum := true
						for _, newEnum := range field.Enums {
							if newEnum.Description == enum.Description {
								okToKeepEnum = false
								break
							}
						}

						if okToKeepEnum {
							field.Enums[enumVal] = enum
						}
					}
				}
			}

			fieldTypeMap[field.Name()] = field
		}
	}

	sortedTags = make([]string, len(fieldMap))
	i := 0
	for f := range fieldMap {
		sortedTags[i] = f
		i++
	}
	sort.Strings(sortedTags)

	genTags()
	genFields()
	genEnums()
}
