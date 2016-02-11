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

	for i, _ := range imports {
		fileOut += fmt.Sprintf("\"%v\"\n", i)
	}
	fileOut += ")\n"
	return fileOut
}

func genFieldDeclaration(field *datadictionary.FieldDef, parent string) (fileOut string) {
	if field.IsComponent() {
		imports[fmt.Sprintf("github.com/quickfixgo/quickfix/%v/%v", pkg, strings.ToLower(field.Component.Name))] = true
		fileOut += fmt.Sprintf("//%v Component\n", field.Component.Name)
		fileOut += fmt.Sprintf("%v %v.Component\n", field.Component.Name, strings.ToLower(field.Component.Name))
		return
	}

	if field.Required {
		fileOut += fmt.Sprintf("//%v is a required field for %v.\n", field.Name, parent)
	} else {
		fileOut += fmt.Sprintf("//%v is a non-required field for %v.\n", field.Name, parent)
	}
	if field.IsGroup() {
		if field.Required {
			fileOut += fmt.Sprintf("%v []%v `fix:\"%v\"`\n", field.Name, field.Name, field.Tag)
		} else {
			fileOut += fmt.Sprintf("%v []%v `fix:\"%v,omitempty\"`\n", field.Name, field.Name, field.Tag)
		}
		return
	}

	goType := ""
	switch field.Type {
	case "MULTIPLESTRINGVALUE", "MULTIPLEVALUESTRING":
		fallthrough
	case "MULTIPLECHARVALUE":
		fallthrough
	case "CHAR":
		fallthrough
	case "CURRENCY":
		fallthrough
	case "DATA":
		fallthrough
	case "MONTHYEAR":
		fallthrough
	case "LOCALMKTDATE":
		fallthrough
	case "DATE":
		fallthrough
	case "EXCHANGE":
		fallthrough
	case "LANGUAGE":
		fallthrough
	case "XMLDATA":
		fallthrough
	case "COUNTRY":
		fallthrough
	case "UTCTIMEONLY":
		fallthrough
	case "UTCDATE":
		fallthrough
	case "UTCDATEONLY":
		fallthrough
	case "TZTIMEONLY":
		fallthrough
	case "TZTIMESTAMP":
		fallthrough
	case "STRING":
		goType = "string"

	case "BOOLEAN":
		goType = "bool"

	case "LENGTH":
		fallthrough
	case "DAYOFMONTH":
		fallthrough
	case "NUMINGROUP":
		fallthrough
	case "SEQNUM":
		fallthrough
	case "INT":
		goType = "int"

	case "TIME":
		fallthrough
	case "UTCTIMESTAMP":
		imports["time"] = true
		goType = "time.Time"

	case "QTY":
		fallthrough
	case "QUANTITY":
		fallthrough
	case "AMT":
		fallthrough
	case "PRICE":
		fallthrough
	case "PRICEOFFSET":
		fallthrough
	case "PERCENTAGE":
		fallthrough
	case "FLOAT":
		goType = "float64"

	default:
		fmt.Printf("Unknown type '%v' for tag '%v'\n", field.Type, field.Tag)
	}

	if field.Required {
		fileOut += fmt.Sprintf("%v %v `fix:\"%v\"`\n", field.Name, goType, field.Tag)
	} else {
		fileOut += fmt.Sprintf("%v *%v `fix:\"%v\"`\n", field.Name, goType, field.Tag)
	}

	return
}

func genGroupDeclaration(field *datadictionary.FieldDef, parent string) (fileOut string) {
	fileOut += fmt.Sprintf("//%v is a repeating group in %v\n", field.Name, parent)
	fileOut += fmt.Sprintf("type %v struct {\n", field.Name)
	for _, groupField := range field.ChildFields {
		fileOut += genFieldDeclaration(groupField, field.Name)
	}

	fileOut += "}\n"

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
	fileOut += fmt.Sprintf("Header %v.Header\n", headerTrailerPkg())
	for _, field := range msg.FieldsInDeclarationOrder {
		fileOut += genFieldDeclaration(field, msg.Name)
	}
	fileOut += fmt.Sprintf("Trailer %v.Trailer\n", headerTrailerPkg())
	fileOut += "}\n"
	fileOut += "//Marshal converts Message to a quickfix.Message instance\n"
	fileOut += "func (m Message) Marshal() quickfix.Message {return quickfix.Marshal(m)}\n"

	return fileOut
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
