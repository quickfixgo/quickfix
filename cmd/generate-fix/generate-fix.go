package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"github.com/quickfixgo/quickfix/cmd/generate-fix/internal"
	"github.com/quickfixgo/quickfix/datadictionary"
)

var (
	waitGroup sync.WaitGroup
	errors    = make(chan error)
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %v [flags] <path to data dictionary> ... \n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func getPackageName(fixSpec *datadictionary.DataDictionary) string {
	pkg := strings.ToLower(fixSpec.FIXType) + strconv.Itoa(fixSpec.Major) + strconv.Itoa(fixSpec.Minor)

	if fixSpec.ServicePack != 0 {
		pkg += "sp" + strconv.Itoa(fixSpec.ServicePack)
	}

	return pkg
}

func getTransportPackageName(fixSpec *datadictionary.DataDictionary) string {
	if fixSpec.Major >= 5 {
		return "fixt11"
	}
	return getPackageName(fixSpec)
}

type component struct {
	Package          string
	FIXPackage       string
	TransportPackage string
	FIXSpec          *datadictionary.DataDictionary
	Name             string
	*datadictionary.MessageDef
}

func genHeader(pkg string, spec *datadictionary.DataDictionary) {
	c := component{
		Package:    pkg,
		Name:       "Header",
		MessageDef: spec.Header,
		FIXSpec:    spec,
	}
	gen(internal.HeaderTemplate, path.Join(pkg, "header.generated.go"), c)
}

func genTrailer(pkg string, spec *datadictionary.DataDictionary) {
	c := component{
		Package:    pkg,
		Name:       "Trailer",
		MessageDef: spec.Trailer,
	}
	gen(internal.TrailerTemplate, path.Join(pkg, "trailer.generated.go"), c)
}

func genMessage(fixPkg string, spec *datadictionary.DataDictionary, msg *datadictionary.MessageDef) {
	pkgName := strings.ToLower(msg.Name)
	transportPkg := getTransportPackageName(spec)

	c := component{
		Package:          pkgName,
		FIXPackage:       fixPkg,
		TransportPackage: transportPkg,
		FIXSpec:          spec,
		Name:             msg.Name,
		MessageDef:       msg,
	}

	gen(internal.MessageTemplate, path.Join(fixPkg, pkgName, msg.Name+".generated.go"), c)
}

func genTags() {
	gen(internal.TagTemplate, "tag/tag_numbers.generated.go", internal.GlobalFieldTypes)
}

func genFields() {
	gen(internal.FieldTemplate, "field/fields.generated.go", internal.GlobalFieldTypes)
}

func genEnums() {
	gen(internal.EnumTemplate, "enum/enums.generated.go", internal.GlobalFieldTypes)
}

func gen(t *template.Template, fileOut string, data interface{}) {
	defer waitGroup.Done()
	writer := new(bytes.Buffer)

	if err := t.Execute(writer, data); err != nil {
		errors <- err
		return
	}

	if err := internal.WriteFile(fileOut, writer.String()); err != nil {
		errors <- err
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
	}

	args := flag.Args()
	if len(args) == 1 {
		dictpath := args[0]
		if strings.Contains(dictpath, "FIX50SP1") {
			args = append(args, strings.Replace(dictpath, "FIX50SP1", "FIXT11", -1))
		} else if strings.Contains(dictpath, "FIX50SP2") {
			args = append(args, strings.Replace(dictpath, "FIX50SP2", "FIXT11", -1))
		} else if strings.Contains(dictpath, "FIX50") {
			args = append(args, strings.Replace(dictpath, "FIX50", "FIXT11", -1))
		}
	}
	specs := []*datadictionary.DataDictionary{}

	for _, dataDictPath := range args {
		spec, err := datadictionary.Parse(dataDictPath)
		if err != nil {
			log.Fatalf("Error Parsing %v: %v", dataDictPath, err)
		}
		specs = append(specs, spec)
	}

	internal.BuildGlobalFieldTypes(specs)

	waitGroup.Add(1)
	go genTags()
	waitGroup.Add(1)
	go genFields()
	waitGroup.Add(1)
	go genEnums()

	for _, spec := range specs {
		pkg := getPackageName(spec)

		if fi, err := os.Stat(pkg); os.IsNotExist(err) {
			if err := os.Mkdir(pkg, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		} else if !fi.IsDir() {
			log.Fatalf("%v/ is not a directory", pkg)
		}

		switch pkg {
		// Uses fixt11 header/trailer.
		case "fix50", "fix50sp1", "fix50sp2":
		default:
			waitGroup.Add(1)
			go genHeader(pkg, spec)

			waitGroup.Add(1)
			go genTrailer(pkg, spec)
		}

		for _, m := range spec.Messages {
			waitGroup.Add(1)
			go genMessage(pkg, spec, m)
		}
	}

	go func() {
		waitGroup.Wait()
		close(errors)
	}()

	var h internal.ErrorHandler
	for err := range errors {
		h.Handle(err)
	}

	os.Exit(h.ReturnCode)
}
