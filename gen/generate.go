package gen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path"
)

var (
	tabWidth    = 8
	printerMode printer.Mode
)

func init() {
	initPrinterMode()
}

func initPrinterMode() {
	printerMode = printer.UseSpaces | printer.TabIndent
}

func WriteFile(filePath, fileOut string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", fileOut, parser.ParseComments)
	if err != nil {
		fmt.Println("Failed to parse:\n", fileOut)
		panic(err)
	}

	ast.SortImports(fset, f)

	//create parentdir if it doesn't exist
	if parentdir := path.Dir(filePath); parentdir != "." {
		if err := os.MkdirAll(parentdir, os.ModePerm); err != nil {
			panic(err)
		}
	}

	if file, err := os.Create(filePath); err == nil {
		defer file.Close()
		(&printer.Config{Mode: printerMode, Tabwidth: tabWidth}).Fprint(file, fset, f)
	} else {
		panic(err)
	}
}
