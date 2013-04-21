package main

import (
	"flag"
	"fmt"
	"github.com/cbusbey/quickfixgo/gen"
	"os"
	"strconv"
	"strings"
)

var (
	pkg     string
	fixSpec *gen.FixSpec
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: generate-messages [flags] <path to data dictionary>\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func initPackage() {
	pkg = strings.ToLower(fixSpec.FixType) + strconv.Itoa(fixSpec.Major) + strconv.Itoa(fixSpec.Minor)

	if fixSpec.ServicePack != 0 {
		pkg += "sp" + strconv.Itoa(fixSpec.ServicePack)
	}
}

func genMessages() {
	for _, m := range fixSpec.Messages {
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
	gen.WriteFile(filePath, fileOut)
}

func buildCrackerImports() string {
	return `
import(
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/session"
	"github.com/cbusbey/quickfixgo/tag"
)
`
}

func buildCrack() (out string) {
	out += "func Crack(msg message.Message, sessionID session.ID, router MessageRouter) reject.MessageReject {\n"
	out += "switch msgType, _ := msg.Header().StringValue(tag.MsgType); msgType {"

	for _, m := range fixSpec.Messages {
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

	for _, m := range fixSpec.Messages {
		out += fmt.Sprintf("On%v%v(msg %v, sessionID session.ID) reject.MessageReject\n", strings.ToUpper(pkg), m.Name, m.Name)
	}

	out += "}\n"

	return
}

func buildMessageCracker() (out string) {
	out += fmt.Sprintf("type %vMessageCracker struct {}\n", strings.ToUpper(pkg))

	for _, m := range fixSpec.Messages {
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
	gen.WriteFile(filePath, fileOut)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
	}

	dataDict := flag.Arg(0)

	if spec, err := gen.ParseFixSpec(dataDict); err != nil {
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

	genCracker()
	genMessages()
}
