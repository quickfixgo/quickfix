package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/gen"
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

func genGroupDeclarations(msg *datadictionary.MessageDef) (fileOut string) {
	groups := []group{}
	for _, field := range msg.Parts {
		groups = collectGroups(msg.Name, field, groups)
	}

	for _, group := range groups {
		fileOut += gen.WriteGroupDeclaration(fixSpec.Major, fixSpec.Minor, group.field, group.parent)
	}

	return
}

func genMessage(msg *datadictionary.MessageDef) string {
	fileOut := fmt.Sprintf("//Message is a %v FIX Message\n", msg.Name)
	fileOut += "type Message struct {\n"
	fileOut += fmt.Sprintf("FIXMsgType string   `fix:\"%v\"`\n", msg.MsgType)
	fileOut += fmt.Sprintf("%v.Header\n", gen.HeaderTrailerPkg(pkg))
	fileOut += gen.WriteFieldDeclarations(fixSpec.Major, fixSpec.Minor, msg.Parts, msg.Name)
	fileOut += fmt.Sprintf("%v.Trailer\n", gen.HeaderTrailerPkg(pkg))
	fileOut += "}\n"
	fileOut += "//Marshal converts Message to a quickfix.Message instance\n"
	fileOut += "func (m Message) Marshal() quickfix.Message {return quickfix.Marshal(m)}\n"

	return fileOut
}

func genMessageNew(msg *datadictionary.MessageDef) string {
	writer := new(bytes.Buffer)
	gen.WriteNewMessage(writer, *msg)
	return writer.String()
}

func genMessageSetters(msg *datadictionary.MessageDef) string {
	writer := new(bytes.Buffer)
	if err := gen.WriteFieldSetters(writer, "Message", msg.Parts); err != nil {
		panic(err)
	}
	return writer.String()
}

func genMessageRoute(msg *datadictionary.MessageDef) string {
	var beginStringEnum string
	if fixSpec.FIXType == "FIXT" {
		beginStringEnum = "enum.BeginStringFIXT11"
	} else {
		if fixSpec.ServicePack == 0 {
			if fixSpec.Major != 5 {
				beginStringEnum = fmt.Sprintf("enum.BeginStringFIX%v%v", fixSpec.Major, fixSpec.Minor)
			} else {
				beginStringEnum = fmt.Sprintf("enum.ApplVerID_FIX%v%v", fixSpec.Major, fixSpec.Minor)
			}
		} else {
			beginStringEnum = fmt.Sprintf("enum.ApplVerID_FIX%v%vSP%v", fixSpec.Major, fixSpec.Minor, fixSpec.ServicePack)
		}
	}

	fileOut := `
//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string,string,quickfix.MessageRoute) {
	r:=func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m:=new(Message)
		if err:=quickfix.Unmarshal(msg, m); err!=nil	{
			return err
		}
		return router(*m, sessionID)
	}
	`
	fileOut += fmt.Sprintf("return %v,\"%v\", r\n", beginStringEnum, msg.MsgType)
	fileOut += "}\n"

	return fileOut
}

func genMessagePkg(msg *datadictionary.MessageDef) error {
	pkgName := strings.ToLower(msg.Name)
	fileOut := fmt.Sprintf("//Package %v msg type = %v.\n", pkgName, msg.MsgType)

	writer := new(bytes.Buffer)
	if err := gen.WritePackage(writer, pkgName); err != nil {
		return err
	}
	if err := gen.WriteMessageImports(writer, pkg, msg.Parts); err != nil {
		return err
	}
	fileOut += writer.String()
	fileOut += genGroupDeclarations(msg)
	fileOut += genMessage(msg)
	fileOut += genMessageNew(msg)
	fileOut += genMessageSetters(msg)
	fileOut += genMessageRoute(msg)

	return gen.WriteFile(path.Join(pkg, strings.ToLower(msg.Name), msg.Name+".go"), fileOut)
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

	var h gen.ErrorHandler
	for _, m := range fixSpec.Messages {
		h.Handle(genMessagePkg(m))
	}
	os.Exit(h.ReturnCode)
}
