package main

import (
	"flag"
	"fmt"
	"github.com/quickfixgo/quickfix/_gen"
	"github.com/quickfixgo/quickfix/datadictionary"
	"os"
	"path"
	"strconv"
	"strings"
)

var (
	pkg     string
	fixSpec *datadictionary.DataDictionary
	imports map[string]bool
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: generate-components [flags] <path to data dictionary>\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func initPackage() {
	pkg = strings.ToLower(fixSpec.FIXType) + strconv.Itoa(fixSpec.Major) + strconv.Itoa(fixSpec.Minor)

	if fixSpec.ServicePack != 0 {
		pkg += "sp" + strconv.Itoa(fixSpec.ServicePack)
	}
}

func packageString() (s string) {
	s = fmt.Sprintf("package %v\n", pkg)
	return
}

func writeField(field *datadictionary.FieldDef, componentName string) (s string) {
	if field.IsComponent() {
		imports[fmt.Sprintf("github.com/quickfixgo/quickfix/%v/%v", pkg, strings.ToLower(field.Component.Name))] = true

		s += fmt.Sprintf("//%v Component\n", field.Component.Name)
		s += fmt.Sprintf("%v %v.Component\n", field.Component.Name, strings.ToLower(field.Component.Name))
		return
	}

	if field.Required {
		s += fmt.Sprintf("//%v is a required field for %v.\n", field.Name, componentName)
	} else {
		s += fmt.Sprintf("//%v is a non-required field for %v.\n", field.Name, componentName)
	}

	if field.IsGroup() {
		if field.Required {
			s += fmt.Sprintf("%v []%v `fix:\"%v\"`\n", field.Name, field.Name, field.Tag)
		} else {
			s += fmt.Sprintf("%v []%v `fix:\"%v,omitempty\"`\n", field.Name, field.Name, field.Tag)
		}
		return
	}

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
	case "UTCDATEONLY":
		fallthrough
	case "UTCDATE":
		fallthrough
	case "TZTIMEONLY":
		fallthrough
	case "TZTIMESTAMP":
		fallthrough
	case "STRING":
		goType = "string"

	case "BOOLEAN":
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
		goType = "int"

	case "TIME":
		fallthrough
	case "UTCTIMESTAMP":
		imports["time"] = true
		goType = "time.Time"

	case "QTY":
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
		goType = "float64"

	default:
		fmt.Printf("Unknown type '%v' for tag '%v'\n", field.Type, field.Tag)
	}

	fixTags := strconv.Itoa(field.Tag)
	if field.Tag == 8 {
		if fixSpec.Major == 4 {
			fixTags = fmt.Sprintf("%v,default=FIX.%v.%v", fixTags, fixSpec.Major, fixSpec.Minor)
		} else {
			fixTags = fixTags + ",default=FIXT.1.1"
		}
	}

	if field.Required {
		s += fmt.Sprintf("%v %v `fix:\"%v\"`\n", field.Name, goType, fixTags)
	} else {
		s += fmt.Sprintf("%v *%v `fix:\"%v\"`\n", field.Name, goType, fixTags)
	}
	return
}

func genComponentImports() (fileOut string) {

	if len(imports) == 0 {
		return
	}

	fileOut += "import(\n"
	for i, _ := range imports {
		fileOut += fmt.Sprintf("\"%v\"\n", i)
	}
	fileOut += ")\n"

	return
}

type group struct {
	parent string
	field  *datadictionary.FieldDef
}

func collectGroups(parent string, field *datadictionary.FieldDef, groups []group) []group {
	if !field.IsGroup() {
		return groups
	}

	groups = append(groups, group{parent, field})
	for _, childField := range field.ChildFields {
		groups = collectGroups(field.Name, childField, groups)
	}

	return groups
}

func genGroupDeclaration(field *datadictionary.FieldDef, parent string) (fileOut string) {
	fileOut += fmt.Sprintf("//%v is a repeating group in %v\n", field.Name, parent)
	fileOut += fmt.Sprintf("type %v struct {\n", field.Name)
	for _, groupField := range field.ChildFields {
		fileOut += writeField(groupField, field.Name)
	}

	fileOut += "}\n"

	return
}

func genGroupDeclarations(name string, fields []*datadictionary.FieldDef) (fileOut string) {
	groups := []group{}
	for _, field := range fields {
		groups = collectGroups(name, field, groups)
	}

	for _, group := range groups {
		fileOut += genGroupDeclaration(group.field, group.parent)
	}

	return
}

func genHeader(header *datadictionary.MessageDef) {
	imports = make(map[string]bool)

	//delay field output to determine imports
	delayOut := genGroupDeclarations("Header", header.FieldsInDeclarationOrder)
	delayOut += fmt.Sprintf("//Header is the %v Header type\n", pkg)
	delayOut += "type Header struct {\n"
	for _, field := range header.FieldsInDeclarationOrder {
		delayOut += writeField(field, "Header")
	}
	delayOut += "}\n"

	fileOut := packageString()
	fileOut += genComponentImports()
	fileOut += delayOut

	gen.WriteFile(path.Join(pkg, "header.go"), fileOut)
}

func genTrailer(trailer *datadictionary.MessageDef) {
	imports = make(map[string]bool)
	fileOut := packageString()
	fileOut += fmt.Sprintf("//Trailer is the %v Trailer type\n", pkg)
	fileOut += "type Trailer struct {\n"
	for _, field := range trailer.FieldsInDeclarationOrder {
		fileOut += writeField(field, "Trailer")
	}
	fileOut += "}\n"

	gen.WriteFile(path.Join(pkg, "trailer.go"), fileOut)
}

func genComponent(name string, component *datadictionary.Component) {
	imports = make(map[string]bool)

	//delay output to determine imports
	delayOut := genGroupDeclarations(name, component.Fields)
	delayOut += fmt.Sprintf("//Component is a %v %v Component\n", pkg, name)
	delayOut += "type Component struct {\n"
	for _, field := range component.Fields {
		delayOut += writeField(field, name)
	}
	delayOut += "}\n"

	fileOut := fmt.Sprintf("package %v\n", strings.ToLower(name))
	fileOut += genComponentImports()
	fileOut += delayOut
	fileOut += "func New() *Component { return new(Component)}\n"

	gen.WriteFile(path.Join(pkg, strings.ToLower(name), name+".go"), fileOut)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
	}

	dataDict := flag.Arg(0)

	if spec, err := datadictionary.Parse(dataDict); err != nil {
		panic(err)
	} else {
		fixSpec = spec
	}

	initPackage()
	if fi, err := os.Stat(pkg); os.IsNotExist(err) {
		if err := os.Mkdir(pkg, os.ModePerm); err != nil {
			panic(err)
		}
	} else if !fi.IsDir() {
		panic(pkg + "/ is not a directory")
	}

	switch pkg {
	//uses fixt11 header/trailer
	case "fix50", "fix50sp1", "fix50sp2":
	default:
		genHeader(fixSpec.Header)
		genTrailer(fixSpec.Trailer)
	}

	for name, component := range fixSpec.Components {
		genComponent(name, component)
	}
}
