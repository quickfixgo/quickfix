package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/gen"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: generate-fields [flags] <path to data dictionary> ... \n")
	flag.PrintDefaults()
	os.Exit(2)
}

func genEnums(fieldTypes []*datadictionary.FieldType) error {
	fileOut := "package enum\n"

	for _, fieldType := range fieldTypes {
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

		fileOut += fmt.Sprintf("//Enum values for %v\n", fieldType.Name())
		fileOut += "const(\n"
		for _, enumVal := range sortedEnums {
			enum, _ := fieldType.Enums[enumVal]
			fileOut += fmt.Sprintf("%v_%v = \"%v\"\n", fieldType.Name(), enum.Description, enum.Value)
		}
		fileOut += ")\n"
	}

	return gen.WriteFile("enum/enums.go", fileOut)
}

func genFields(fieldTypes []*datadictionary.FieldType) error {
	fileOut := "package field\n"
	fileOut += "import(\n"
	fileOut += "\"github.com/quickfixgo/quickfix\"\n"
	fileOut += "\"" + gen.GetImportPathRoot() + "/tag\"\n"
	fileOut += ")\n"

	for _, field := range fieldTypes {

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
			return fmt.Errorf("Unknown type '%v' for tag '%v'\n", field.Type, field.Tag())
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

	return gen.WriteFile("field/fields.go", fileOut)
}

func genTags(fieldTypes []*datadictionary.FieldType) error {
	fileOut := "package tag\n"
	fileOut += "import(\"github.com/quickfixgo/quickfix\")\n"

	fileOut += "const (\n"
	for _, f := range fieldTypes {
		fileOut += fmt.Sprintf("%v quickfix.Tag = %v\n", f.Name(), f.Tag())
	}
	fileOut += ")\n"

	return gen.WriteFile("tag/tag_numbers.go", fileOut)
}

type byFieldName []*datadictionary.FieldType

func (n byFieldName) Len() int           { return len(n) }
func (n byFieldName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n byFieldName) Less(i, j int) bool { return n[i].Name() < n[j].Name() }

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
	}

	fieldTypeMap := make(map[string]*datadictionary.FieldType)

	for _, dataDict := range flag.Args() {
		spec, err := datadictionary.Parse(dataDict)

		if err != nil {
			fmt.Println(dataDict)
			panic(err)
		}

		for _, field := range spec.FieldTypeByTag {
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

	fieldTypes := make([]*datadictionary.FieldType, len(fieldTypeMap))
	i := 0
	for _, fieldType := range fieldTypeMap {
		fieldTypes[i] = fieldType
		i++
	}

	sort.Sort(byFieldName(fieldTypes))

	var h gen.ErrorHandler
	h.Handle(genTags(fieldTypes))
	h.Handle(genFields(fieldTypes))
	h.Handle(genEnums(fieldTypes))
	os.Exit(h.ReturnCode)
}
