package main

import (
	"flag"
	"fmt"
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/gen"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	pkg            string
	fixSpec        *datadictionary.DataDictionary
	sortedMsgTypes []string
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: generate-messages [flags] <path to data dictionary>\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func initPackage() {
	pkg = strings.ToLower(fixSpec.FIXType) + strconv.Itoa(fixSpec.Major) + strconv.Itoa(fixSpec.Minor)

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
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)
`
}

func buildCrack() (out string) {
	out += "func Crack(msg message.Message, sessionID quickfix.SessionID, router MessageRouter) errors.MessageRejectError {\n"
	out += `
  msgType:=&field.MsgTypeField{}
switch msg.Header.Get(msgType); msgType.Value {
`

	for _, msgType := range sortedMsgTypes {
		m, _ := fixSpec.Messages[msgType]

		out += fmt.Sprintf("case \"%v\":\n", msgType)
		out += fmt.Sprintf("return router.On%v%v(%v{msg},sessionID)\n", strings.ToUpper(pkg), m.Name, m.Name)
	}
	out += "}\n"
	out += "return errors.InvalidMessageType()\n"
	out += "}\n"

	return
}

func buildMessageRouter() (out string) {
	out += "type MessageRouter interface {\n"

	for _, msgType := range sortedMsgTypes {
		m, _ := fixSpec.Messages[msgType]
		out += fmt.Sprintf("On%v%v(msg %v, sessionID quickfix.SessionID) errors.MessageRejectError\n", strings.ToUpper(pkg), m.Name, m.Name)
	}

	out += "}\n"

	return
}

func buildMessageCracker() (out string) {
	out += fmt.Sprintf("type %vMessageCracker struct {}\n", strings.ToUpper(pkg))

	for _, msgType := range sortedMsgTypes {
		m, _ := fixSpec.Messages[msgType]
		out += fmt.Sprintf("//On%v%v is a Callback for %v messages.\n", strings.ToUpper(pkg), m.Name, m.Name)
		out += fmt.Sprintf("func (c * %vMessageCracker) On%v%v(msg %v, sessionID quickfix.SessionID) errors.MessageRejectError {\n", strings.ToUpper(pkg), strings.ToUpper(pkg), m.Name, m.Name)
		out += "return errors.UnsupportedMessageType()\n}\n"
	}

	return
}

func genMessage(msg *datadictionary.MessageDef) {
	fileOut := fmt.Sprintf("package %v\n", pkg)
	fileOut += `
import( 
  "github.com/quickfixgo/quickfix/errors"
  "github.com/quickfixgo/quickfix/fix"
  "github.com/quickfixgo/quickfix/fix/field"
  "github.com/quickfixgo/quickfix/message"
)
`

	if fixSpec.Major == 5 {
		fileOut += `
import( 
  "github.com/quickfixgo/quickfix/fix/enum"
)
`
	}

	fileOut += fmt.Sprintf("//%v msg type = %v.\n", msg.Name, msg.MsgType)
	fileOut += fmt.Sprintf("type %v struct {\n message.Message}\n", msg.Name)

	fileOut += fmt.Sprintf("//%vBuilder builds %v messages.\n", msg.Name, msg.Name)
	fileOut += fmt.Sprintf("type %vBuilder struct {\n message.MessageBuilder}\n", msg.Name)

	requiredFields := make([]*datadictionary.FieldDef, 0, len(msg.FieldsInDeclarationOrder))
	for _, field := range msg.FieldsInDeclarationOrder {
		if field.Required {
			requiredFields = append(requiredFields, field)
		}
	}

	fileOut += fmt.Sprintf("//Create%vBuilder returns an initialized %vBuilder with specified required fields.\n", msg.Name, msg.Name)
	fileOut += fmt.Sprintf("func Create%vBuilder(\n", msg.Name)
	builderArgs := make([]string, len(requiredFields))
	for i, field := range requiredFields {
		builderArgs[i] = fmt.Sprintf("%v *field.%vField", strings.ToLower(field.Name), field.Name)
	}
	fileOut += strings.Join(builderArgs, ",\n")
	fileOut += fmt.Sprintf(") %vBuilder {\n", msg.Name)
	fileOut += fmt.Sprintf("var builder %vBuilder\n", msg.Name)
	fileOut += "builder.MessageBuilder = message.Builder()\n"

	if fixSpec.FIXType == "FIXT" {
		fileOut += fmt.Sprintf("builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))\n")
	} else {
		if fixSpec.Major == 5 {
			fileOut += fmt.Sprintf("builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))\n")
			switch fixSpec.ServicePack {
			case 0:
				fileOut += fmt.Sprintf("builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))\n")
			default:
				fileOut += fmt.Sprintf("builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP%v))\n", fixSpec.ServicePack)
			}
		} else {
			fileOut += fmt.Sprintf("builder.Header().Set(field.NewBeginString(fix.BeginString_FIX%v%v))\n", fixSpec.Major, fixSpec.Minor)
		}
	}

	fileOut += fmt.Sprintf("builder.Header().Set(field.NewMsgType(\"%v\"))\n", msg.MsgType)

	for _, field := range requiredFields {
		fileOut += fmt.Sprintf("builder.Body().Set(%v)\n", strings.ToLower(field.Name))
	}

	fileOut += "return builder\n"
	fileOut += "}\n"

	for _, field := range msg.FieldsInDeclarationOrder {
		if field.Required {
			fileOut += fmt.Sprintf("//%v is a required field for %v.\n", field.Name, msg.Name)
		} else {
			fileOut += fmt.Sprintf("//%v is a non-required field for %v.\n", field.Name, msg.Name)
		}
		fileOut += fmt.Sprintf("func (m %v) %v() (*field.%vField, errors.MessageRejectError) {\n", msg.Name, field.Name, field.Name)
		fileOut += fmt.Sprintf("f := &field.%vField{}\n", field.Name)
		fileOut += "err:=m.Body.Get(f)\n"
		fileOut += "return f, err\n}\n"

		fileOut += fmt.Sprintf("//Get%v reads a %v from %v.\n", field.Name, field.Name, msg.Name)
		fileOut += fmt.Sprintf("func (m %v) Get%v(f *field.%vField) errors.MessageRejectError {\n", msg.Name, field.Name, field.Name)
		fileOut += "return m.Body.Get(f)\n}\n"
	}

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

	if spec, err := datadictionary.Parse(dataDict); err != nil {
		panic(err)
	} else {
		fixSpec = spec
	}

	sortedMsgTypes = make([]string, len(fixSpec.Messages))
	i := 0
	for f := range fixSpec.Messages {
		sortedMsgTypes[i] = f
		i++
	}
	sort.Strings(sortedMsgTypes)

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
