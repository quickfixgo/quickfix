package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/gen"
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

type group struct {
	parent string
	field  *datadictionary.FieldDef
}

func collectGroups(parent string, part datadictionary.MessagePart, groups []group) []group {
	switch field := part.(type) {
	case *datadictionary.FieldDef:
		if !field.IsGroup() {
			return groups
		}

		groups = append(groups, group{parent, field})
		for _, childField := range field.Parts {
			groups = collectGroups(field.Name(), childField, groups)
		}
	}

	return groups
}

func genGroupDeclarations(name string, fields []datadictionary.MessagePart) (fileOut string) {
	groups := []group{}
	for _, field := range fields {
		groups = collectGroups(name, field, groups)
	}

	for _, group := range groups {
		fileOut += gen.WriteGroupDeclaration(fixSpec.Major, fixSpec.Minor, group.field, group.parent)
	}

	return
}

func genHeader(header *datadictionary.MessageDef) error {
	writer := new(bytes.Buffer)
	if err := gen.WritePackage(writer, pkg); err != nil {
		return err
	}
	if err := gen.WriteComponentImports(writer, pkg, header.Parts); err != nil {
		return err
	}

	fileOut := writer.String()
	fileOut += genGroupDeclarations("Header", header.Parts)
	fileOut += fmt.Sprintf("//Header is the %v Header type\n", pkg)
	fileOut += "type Header struct {\n"
	fileOut += gen.WriteFieldDeclarations(fixSpec.Major, fixSpec.Minor, header.Parts, "Header")
	fileOut += "}\n"

	writer = new(bytes.Buffer)
	if err := gen.WriteFieldSetters(writer, "Header", header.Parts); err != nil {
		return err
	}
	fileOut += writer.String()

	return gen.WriteFile(path.Join(pkg, "header.go"), fileOut)
}

func genTrailer(trailer *datadictionary.MessageDef) error {
	writer := new(bytes.Buffer)
	if err := gen.WritePackage(writer, pkg); err != nil {
		return err
	}
	if err := gen.WriteComponentImports(writer, pkg, trailer.Parts); err != nil {
		return err
	}
	fileOut := writer.String()

	fileOut += fmt.Sprintf("//Trailer is the %v Trailer type\n", pkg)
	fileOut += "type Trailer struct {\n"
	fileOut += gen.WriteFieldDeclarations(fixSpec.Major, fixSpec.Minor, trailer.Parts, "Trailer")
	fileOut += "}\n"

	writer = new(bytes.Buffer)
	if err := gen.WriteFieldSetters(writer, "Trailer", trailer.Parts); err != nil {
		return err
	}
	fileOut += writer.String()

	return gen.WriteFile(path.Join(pkg, "trailer.go"), fileOut)
}

func genComponent(name string, component *datadictionary.ComponentType) error {
	writer := new(bytes.Buffer)
	if err := gen.WritePackage(writer, strings.ToLower(name)); err != nil {
		return err
	}
	if err := gen.WriteComponentImports(writer, pkg, component.Parts()); err != nil {
		return err
	}

	if err := gen.WriteNewComponent(writer, *component); err != nil {
		return err
	}

	fileOut := writer.String()
	fileOut += genGroupDeclarations(name, component.Parts())
	fileOut += fmt.Sprintf("//%v is a %v Component\n", name, pkg)
	fileOut += fmt.Sprintf("type %v struct {\n", name)
	fileOut += gen.WriteFieldDeclarations(fixSpec.Major, fixSpec.Minor, component.Parts(), name)
	fileOut += "}\n"
	fileOut += genComponentSetters(component)

	return gen.WriteFile(path.Join(pkg, strings.ToLower(name), name+".go"), fileOut)
}

func genComponentSetters(component *datadictionary.ComponentType) string {
	writer := new(bytes.Buffer)
	if err := gen.WriteFieldSetters(writer, component.Name(), component.Parts()); err != nil {
		panic(err)
	}

	return writer.String()
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

	var h gen.ErrorHandler
	switch pkg {
	//uses fixt11 header/trailer
	case "fix50", "fix50sp1", "fix50sp2":
	default:
		h.Handle(genHeader(fixSpec.Header))
		h.Handle(genTrailer(fixSpec.Trailer))
	}

	for name, component := range fixSpec.ComponentTypes {
		h.Handle(genComponent(name, component))
	}

	os.Exit(h.ReturnCode)
}
