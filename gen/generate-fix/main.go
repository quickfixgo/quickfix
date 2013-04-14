package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/cbusbey/quickfixgo/gen"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strconv"
	"strings"
)

var (
	pkg  string
	spec gen.FixSpec
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: generate-fix [flags] <path to data dictionary>\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func initPackage() {
	pkg = strings.ToLower(spec.FixType) + strconv.Itoa(spec.Major) + strconv.Itoa(spec.Minor)

	if spec.ServicePack != 0 {
		pkg += "sp" + strconv.Itoa(spec.ServicePack)
	}
}

func genMessages() {
	for _, m := range spec.Messages {
		genMessage(m)
	}
}

func genCracker() {
	fileOut := fmt.Sprintf("package %v\n", pkg)
	fileOut += buildCrackerImports()
	fileOut += buildCrack()
	fileOut += buildMessageRouter()
	fileOut += buildMessageCracker()

	filePath := pkg + "/message_cracker.go"
	writeFile(filePath, fileOut)
}

func buildCrackerImports() string {
	return `
import(
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session"
	"github.com/cbusbey/quickfixgo/fix"
)
`
}

func buildCrack() (out string) {
	out += "func Crack(msg message.Message, sessionID session.ID, router MessageRouter) reject.MessageReject {\n"
	out += "switch msgType, _ := msg.Header().StringField(fix.MsgType); msgType.Value() {"

	for _, m := range spec.Messages {
		out += fmt.Sprintf("case \"%v\":\n", m.MsgType)
		out += fmt.Sprintf("return router.On%v%v(%v{msg},sessionID)\n", strings.ToUpper(pkg), m.Name, m.Name)
	}
	out += "}\n"
	out += "return reject.NewInvalidMessageType(msg)\n"
	out += "}\n"

	return
}

func buildMessageRouter() (out string) {
	out += "type MessageRouter interface {\n"

	for _, m := range spec.Messages {
		out += fmt.Sprintf("On%v%v(msg %v, sessionID session.ID) reject.MessageReject\n", strings.ToUpper(pkg), m.Name, m.Name)
	}

	out += "}\n"

	return
}

func buildMessageCracker() (out string) {
	out += fmt.Sprintf("type %vMessageCracker struct {}\n", strings.ToUpper(pkg))

	for _, m := range spec.Messages {
		out += fmt.Sprintf("func (c * %vMessageCracker) On%v%v(msg %v, sessionId session.ID) reject.MessageReject {\n", strings.ToUpper(pkg), strings.ToUpper(pkg), m.Name, m.Name)
		out += "return reject.NewUnsupportedMessageType(msg)\n}\n"
	}

	return
}

func genMessage(msg gen.Message) {
	fileOut := fmt.Sprintf("package %v\n", pkg)
	fileOut += "import(\"github.com/cbusbey/quickfixgo/message\")\n"
	fileOut += fmt.Sprintf("type %v struct {\n message.Message}", msg.Name)

	filePath := pkg + "/" + msg.Name + ".go"
	writeFile(filePath, fileOut)
}

func writeFile(filePath, fileOut string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", fileOut, 0)
	if err != nil {
		fmt.Println("Failed to parse:\n", fileOut)
		panic(err)
	}

	ast.SortImports(fset, f)

	if file, err := os.Create(filePath); err == nil {
		defer file.Close()
		printer.Fprint(file, fset, f)
	} else {
		panic(err)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
	}

	dataDict := flag.Arg(0)

	xmlFile, err := os.Open(dataDict)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	decoder.Decode(&spec)

	initPackage()

	if fi, err := os.Stat(pkg); os.IsNotExist(err) {
		if err := os.Mkdir(pkg, os.ModePerm); err != nil {
			panic(err)
		}
	} else if !fi.IsDir() {
		panic(pkg + "/ is not a directory")
	}

	genCracker()
	genMessages()
}
