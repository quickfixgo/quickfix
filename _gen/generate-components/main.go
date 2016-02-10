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
	if field.Required {
		s += fmt.Sprintf("//%v is a required field for %v.\n", field.Name, componentName)
	} else {
		s += fmt.Sprintf("//%v is a non-required field for %v.\n", field.Name, componentName)
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

func genComponentImports() string {
	fileOut := "import \"time\"\n"
	return fileOut
}

func genHeader(header *datadictionary.MessageDef) {
	fileOut := packageString()
	fileOut += genComponentImports()

	fileOut += fmt.Sprintf("//Header is the %v Header type\n", pkg)
	fileOut += "type Header struct {\n"
	for _, field := range header.FieldsInDeclarationOrder {
		fileOut += writeField(field, "Header")
	}
	fileOut += "}\n"

	gen.WriteFile(path.Join(pkg, "header.go"), fileOut)
}

func genTrailer(trailer *datadictionary.MessageDef) {
	fileOut := packageString()
	fileOut += fmt.Sprintf("//Trailer is the %v Trailer type\n", pkg)
	fileOut += "type Trailer struct {\n"
	for _, field := range trailer.FieldsInDeclarationOrder {
		fileOut += writeField(field, "Trailer")
	}
	fileOut += "}\n"

	gen.WriteFile(path.Join(pkg, "trailer.go"), fileOut)
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
}
