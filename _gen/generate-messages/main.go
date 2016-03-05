package main

import (
	"bytes"
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
	imports        map[string]bool
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
  "github.com/quickfixgo/quickfix/enum"
`
	fileOut += fmt.Sprintf("\"github.com/quickfixgo/quickfix/%v\"\n", headerTrailerPkg())

	for i := range imports {
		fileOut += fmt.Sprintf("\"%v\"\n", i)
	}
	fileOut += ")\n"
	return fileOut
}

func writeFieldDeclaration(field *datadictionary.FieldDef, componentName string) string {
	switch {
	case field.IsComponent():
		imports[fmt.Sprintf("github.com/quickfixgo/quickfix/%v/%v", pkg, strings.ToLower(field.Component.Name))] = true
	case !field.IsGroup():
		goType := gen.FixFieldTypeToGoType(field.Type)
		if goType == "time.Time" {
			imports["time"] = true
		}
	}

	return gen.WriteFieldDeclaration(fixSpec.Major, fixSpec.Minor, field, componentName)
}

func genGroupDeclaration(field *datadictionary.FieldDef, parent string) (fileOut string) {
	fileOut += fmt.Sprintf("//%v is a repeating group in %v\n", field.Name, parent)
	fileOut += fmt.Sprintf("type %v struct {\n", field.Name)
	for _, groupField := range field.ChildFields {
		fileOut += writeFieldDeclaration(groupField, field.Name)
	}

	fileOut += "}\n"

	writer := new(bytes.Buffer)
	if err := gen.WriteFieldSetters(writer, field.Name, field.ChildFields); err != nil {
		panic(err)
	}
	fileOut += writer.String()
	fileOut += "\n"

	return
}

type group struct {
	parent string
	field  *datadictionary.FieldDef
}

func collectGroups(parent string, field *datadictionary.FieldDef, groups []group) []group {
	if !field.IsGroup() {
		return groups
	}

	groups = append(groups, group{parent, field})
	for _, childField := range field.ChildFields {
		groups = collectGroups(field.Name, childField, groups)
	}

	return groups
}

func genGroupDeclarations(msg *datadictionary.MessageDef) (fileOut string) {
	groups := []group{}
	for _, field := range msg.FieldsInDeclarationOrder {
		groups = collectGroups(msg.Name, field, groups)
	}

	for _, group := range groups {
		fileOut += genGroupDeclaration(group.field, group.parent)
	}

	return
}

func headerTrailerPkg() string {
	switch pkg {
	case "fix50", "fix50sp1", "fix50sp2":
		return "fixt11"
	}

	return pkg
}

func genMessage(msg *datadictionary.MessageDef, requiredFields []*datadictionary.FieldDef) string {
	fileOut := fmt.Sprintf("//Message is a %v FIX Message\n", msg.Name)
	fileOut += "type Message struct {\n"
	fileOut += fmt.Sprintf("FIXMsgType string   `fix:\"%v\"`\n", msg.MsgType)
	fileOut += fmt.Sprintf("%v.Header\n", headerTrailerPkg())
	for _, field := range msg.FieldsInDeclarationOrder {
		fileOut += writeFieldDeclaration(field, msg.Name)
	}
	fileOut += fmt.Sprintf("%v.Trailer\n", headerTrailerPkg())
	fileOut += "}\n"
	fileOut += "//Marshal converts Message to a quickfix.Message instance\n"
	fileOut += "func (m Message) Marshal() quickfix.Message {return quickfix.Marshal(m)}\n"

	return fileOut
}

func genMessageSetters(msg *datadictionary.MessageDef) string {
	writer := new(bytes.Buffer)
	if err := gen.WriteFieldSetters(writer, "Message", msg.FieldsInDeclarationOrder); err != nil {
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
			beginStringEnum = fmt.Sprintf("enum.BeginStringFIX%v%v", fixSpec.Major, fixSpec.Minor)
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

func genMessagePkg(msg *datadictionary.MessageDef) {
	requiredFields := make([]*datadictionary.FieldDef, 0, len(msg.FieldsInDeclarationOrder))
	for _, field := range msg.FieldsInDeclarationOrder {
		if field.Required {
			requiredFields = append(requiredFields, field)
		}
	}

	pkgName := strings.ToLower(msg.Name)
	imports = make(map[string]bool)

	fileOut := fmt.Sprintf("//Package %v msg type = %v.\n", pkgName, msg.MsgType)
	fileOut += fmt.Sprintf("package %v\n", pkgName)

	//run through group and message declarations to collect required imports first
	delayOut := ""
	delayOut += genGroupDeclarations(msg)
	delayOut += genMessage(msg, requiredFields)

	fileOut += genMessageImports()
	fileOut += delayOut
	fileOut += genMessageSetters(msg)
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
