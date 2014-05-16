package main

import (
	"flag"
	"fmt"
	"github.com/quickfixgo/quickfix/_gen"
	"github.com/quickfixgo/quickfix/datadictionary"
	"os"
	"path"
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
		genMessagePkg(m)
	}
}

func genMessageImports() string {
	fileOut := `
import( 
  "github.com/quickfixgo/quickfix"
  "github.com/quickfixgo/quickfix/fix"
  "github.com/quickfixgo/quickfix/fix/field"
)
`

	if fixSpec.Major == 5 {
		fileOut += `
import( 
  "github.com/quickfixgo/quickfix/fix/enum"
)
`
	}
	return fileOut
}

func genMessage(msg *datadictionary.MessageDef, requiredFields []*datadictionary.FieldDef) string {
	fileOut := fmt.Sprintf("//Message is a %v wrapper for the generic Message type\n", msg.Name)
	fileOut += "type Message struct {\n quickfix.Message}\n"

	for _, field := range msg.FieldsInDeclarationOrder {
		if field.Required {
			fileOut += fmt.Sprintf("//%v is a required field for %v.\n", field.Name, msg.Name)
		} else {
			fileOut += fmt.Sprintf("//%v is a non-required field for %v.\n", field.Name, msg.Name)
		}
		fileOut += fmt.Sprintf("func (m Message) %v() (*field.%vField, quickfix.MessageRejectError) {\n", field.Name, field.Name)
		fileOut += fmt.Sprintf("f := &field.%vField{}\n", field.Name)
		fileOut += "err:=m.Body.Get(f)\n"
		fileOut += "return f, err\n}\n"

		fileOut += fmt.Sprintf("//Get%v reads a %v from %v.\n", field.Name, field.Name, msg.Name)
		fileOut += fmt.Sprintf("func (m Message) Get%v(f *field.%vField) quickfix.MessageRejectError {\n", field.Name, field.Name)
		fileOut += "return m.Body.Get(f)\n}\n"
	}

	return fileOut
}

func genMessageBuilder(msg *datadictionary.MessageDef, requiredFields []*datadictionary.FieldDef) string {
	fileOut := fmt.Sprintf("//MessageBuilder builds %v messages.\n", msg.Name)
	fileOut += "type MessageBuilder struct {\n quickfix.MessageBuilder}\n"

	fileOut += fmt.Sprintf("//Builder returns an initialized MessageBuilder with specified required fields for %v.\n", msg.Name)
	fileOut += "func Builder(\n"
	builderArgs := make([]string, len(requiredFields))
	for i, field := range requiredFields {
		builderArgs[i] = fmt.Sprintf("%v *field.%vField", strings.ToLower(field.Name), field.Name)
	}
	fileOut += strings.Join(builderArgs, ",\n")
	fileOut += ") MessageBuilder {\n"
	fileOut += "var builder MessageBuilder\n"
	fileOut += "builder.MessageBuilder = quickfix.NewMessageBuilder()\n"

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

	return fileOut
}

func genMessageRoute(msg *datadictionary.MessageDef) string {
	var beginStringEnum string
	if fixSpec.FIXType == "FIXT" {
		beginStringEnum = "fix.BeginString_FIXT11"
	} else {
		if fixSpec.ServicePack == 0 {
			beginStringEnum = fmt.Sprintf("fix.BeginString_FIX%v%v", fixSpec.Major, fixSpec.Minor)
		} else {
			beginStringEnum = fmt.Sprintf("enum.ApplVerID_FIX%v%vSP%v", fixSpec.Major, fixSpec.Minor, fixSpec.ServicePack)
		}
	}

	fileOut := `
//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string,string,quickfix.MessageRoute) {
	r:=func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	`
	fileOut += fmt.Sprintf("return %v,\"%v\", r\n", beginStringEnum, msg.MsgType)
	fileOut += "}\n"

	return fileOut
}

func genMessagePkg(msg *datadictionary.MessageDef) {
	requiredFields := make([]*datadictionary.FieldDef, 0, len(msg.FieldsInDeclarationOrder))
	for _, field := range msg.FieldsInDeclarationOrder {
		if field.Required {
			requiredFields = append(requiredFields, field)
		}
	}

	pkgName := strings.ToLower(msg.Name)

	fileOut := fmt.Sprintf("//Package %v msg type = %v.\n", pkgName, msg.MsgType)
	fileOut += fmt.Sprintf("package %v\n", pkgName)
	fileOut += genMessageImports()

	fileOut += genMessage(msg, requiredFields)
	fileOut += genMessageBuilder(msg, requiredFields)
	fileOut += genMessageRoute(msg)

	gen.WriteFile(path.Join(pkg, strings.ToLower(msg.Name), msg.Name+".go"), fileOut)
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

	genMessages()
}
