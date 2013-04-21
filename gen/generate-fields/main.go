package main

import (
	"flag"
	"fmt"
	"github.com/cbusbey/quickfixgo/gen"
	"os"
	"sort"
)

var (
	fieldMap map[string]int
	pkg      = "fix"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: generate-fields [flags] <path to data dictionary> ... \n")
	flag.PrintDefaults()
	os.Exit(2)
}

func genTags() {
	fileOut := "package tag\n"

	fileOut += "const (\n"
	sortedTags := make([]string, len(fieldMap))
	i := 0
	for f, _ := range fieldMap {
		sortedTags[i] = f
		i++
	}
	sort.Strings(sortedTags)

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

	for _, dataDict := range flag.Args() {
		spec, err := gen.ParseFixSpec(dataDict)

		if err != nil {
			panic(err)
		}

		for _, field := range spec.FieldTypes {
			fieldMap[field.Name] = field.Number
		}
	}

	genTags()
}
