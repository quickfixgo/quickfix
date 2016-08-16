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
	defer waitGroup.Done()
	writer := new(bytes.Buffer)
	c := component{
		Package:    pkg,
		Name:       "Header",
		MessageDef: spec.Header,
		FIXSpec:    spec,
	}
	if err := internal.HeaderTemplate.Execute(writer, c); err != nil {
		errors <- err
		return
	}

	if err := internal.WriteFile(path.Join(pkg, "header.generated.go"), writer.String()); err != nil {
		errors <- err
	}
}

func genTrailer(pkg string, spec *datadictionary.DataDictionary) {
	defer waitGroup.Done()
	writer := new(bytes.Buffer)
	c := component{
		Package:    pkg,
		Name:       "Trailer",
		MessageDef: spec.Trailer,
	}
	if err := internal.TrailerTemplate.Execute(writer, c); err != nil {
		errors <- err
		return
	}

	if err := internal.WriteFile(path.Join(pkg, "trailer.generated.go"), writer.String()); err != nil {
		errors <- err
	}
}

func genMessage(fixPkg string, spec *datadictionary.DataDictionary, msg *datadictionary.MessageDef) {
	defer waitGroup.Done()
	pkgName := strings.ToLower(msg.Name)
	transportPkg := getTransportPackageName(spec)

	writer := new(bytes.Buffer)
	c := component{
		Package:          pkgName,
		FIXPackage:       fixPkg,
		TransportPackage: transportPkg,
		FIXSpec:          spec,
		Name:             msg.Name,
		MessageDef:       msg,
	}

	if err := internal.MessageTemplate.Execute(writer, c); err != nil {
		errors <- err
		return
	}

	if err := internal.WriteFile(path.Join(fixPkg, pkgName, msg.Name+".generated.go"), writer.String()); err != nil {
		errors <- err
	}
}

func genTags() {
	defer waitGroup.Done()
	writer := new(bytes.Buffer)

	if err := internal.TagTemplate.Execute(writer, internal.GlobalFieldTypes); err != nil {
		errors <- err
		return
	}

	if err := internal.WriteFile("tag/tag_numbers.generated.go", writer.String()); err != nil {
		errors <- err
	}
}

func genFields() {
	defer waitGroup.Done()
	writer := new(bytes.Buffer)

	if err := internal.FieldTemplate.Execute(writer, internal.GlobalFieldTypes); err != nil {
		errors <- err
		return
	}

	if err := internal.WriteFile("field/fields.generated.go", writer.String()); err != nil {
		errors <- err
	}
}

func genEnums() {
	defer waitGroup.Done()
	writer := new(bytes.Buffer)

	if err := internal.EnumTemplate.Execute(writer, internal.GlobalFieldTypes); err != nil {
		errors <- err
		return
	}

	if err := internal.WriteFile("enum/enums.generated.go", writer.String()); err != nil {
		errors <- err
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
	}

	specs := make([]*datadictionary.DataDictionary, flag.NArg())
	for i, dataDictPath := range flag.Args() {
		var err error
		specs[i], err = datadictionary.Parse(dataDictPath)
		if err != nil {
			log.Fatalf("Error Parsing %v: %v", dataDictPath, err)
		}
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
		//uses fixt11 header/trailer
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
