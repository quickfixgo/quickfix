package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/gen"
)

type fieldTypeMap map[string]*datadictionary.FieldType

var (
	headerTemplate   *template.Template
	trailerTemplate  *template.Template
	messageTemplate  *template.Template
	globalFieldTypes fieldTypeMap

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

func dief(format string, a ...interface{}) {
	die(fmt.Sprintf(format, a...))
}

func die(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(3)
}

type component struct {
	Package          string
	FIXPackage       string
	TransportPackage string
	FIXSpec          *datadictionary.DataDictionary
	Name             string
	*datadictionary.MessageDef
}

func getGlobalFieldType(f *datadictionary.FieldDef) (t *datadictionary.FieldType, err error) {
	var ok bool
	t, ok = globalFieldTypes[f.Name()]
	if !ok {
		err = fmt.Errorf("Unknown global type for %v", f.Name())
	}

	return
}

func checkFieldTimeRequired(f *datadictionary.FieldDef) (required bool, err error) {
	var globalType *datadictionary.FieldType
	if globalType, err = getGlobalFieldType(f); err != nil {
		return
	}

	var t string
	if t, err = quickfixType(globalType); err != nil {
		return
	}

	var vt string
	if vt, err = quickfixValueType(t); err != nil {
		return
	}

	if vt == "time.Time" {
		required = true
		return
	}

	for _, groupField := range f.Fields {
		if required, err = checkFieldTimeRequired(groupField); required || err != nil {
			return
		}
	}

	return
}

func collectExtraImports(m *datadictionary.MessageDef) (imports []string, err error) {
	//NOTE: the time package is the only extra import considered here
	var timeRequired bool
	for _, f := range m.Fields {
		if timeRequired, err = checkFieldTimeRequired(f); timeRequired {
			imports = []string{"time"}
			break
		} else if err != nil {
			return
		}
	}

	return
}

func quickfixValueType(quickfixType string) (goType string, err error) {
	switch quickfixType {
	case "FIXString":
		goType = "string"
	case "FIXBoolean":
		goType = "bool"
	case "FIXInt":
		goType = "int"
	case "FIXUTCTimestamp":
		goType = "time.Time"
	case "FIXFloat":
		goType = "float64"
	default:
		err = fmt.Errorf("Unknown QuickFIX Type: %v", quickfixType)
	}

	return
}

func quickfixType(field *datadictionary.FieldType) (quickfixType string, err error) {
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
	case "TIME":
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
		quickfixType = "FIXString"

	case "BOOLEAN":
		quickfixType = "FIXBoolean"

	case "LENGTH":
		fallthrough
	case "DAYOFMONTH":
		fallthrough
	case "NUMINGROUP":
		fallthrough
	case "SEQNUM":
		fallthrough
	case "INT":
		quickfixType = "FIXInt"

	case "UTCTIMESTAMP":
		quickfixType = "FIXUTCTimestamp"

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
		quickfixType = "FIXFloat"

	default:
		err = fmt.Errorf("Unknown type '%v' for tag '%v'\n", field.Type, field.Tag())
	}

	return
}

func requiredFields(m *datadictionary.MessageDef) (required []*datadictionary.FieldDef) {
	for _, part := range m.RequiredParts() {
		if part.Required() {
			switch pType := part.(type) {
			case *datadictionary.FieldDef:
				if !pType.IsGroup() {
					required = append(required, pType)
				}
			case *datadictionary.Component:
				for _, f := range pType.RequiredFields() {
					if !f.IsGroup() {
						required = append(required, f)
					}
				}
			}
		}
	}

	return
}

func routerBeginString(spec *datadictionary.DataDictionary) (routerBeginString string) {
	switch {
	case spec.FIXType == "FIXT":
		routerBeginString = "FIXT.1.1"
	case spec.Major != 5 && spec.ServicePack == 0:
		routerBeginString = fmt.Sprintf("FIX.%v.%v", spec.Major, spec.Minor)
		//ApplVerID enums
	case spec.Major == 2:
		routerBeginString = "0"
	case spec.Major == 3:
		routerBeginString = "1"
	case spec.Major == 4 && spec.Minor == 0:
		routerBeginString = "2"
	case spec.Major == 4 && spec.Minor == 1:
		routerBeginString = "3"
	case spec.Major == 4 && spec.Minor == 2:
		routerBeginString = "4"
	case spec.Major == 4 && spec.Minor == 3:
		routerBeginString = "5"
	case spec.Major == 4 && spec.Minor == 4:
		routerBeginString = "6"
	case spec.Major == 5 && spec.Minor == 0 && spec.ServicePack == 0:
		routerBeginString = "7"
	case spec.Major == 5 && spec.Minor == 0 && spec.ServicePack == 1:
		routerBeginString = "8"
	case spec.Major == 5 && spec.Minor == 0 && spec.ServicePack == 2:
		routerBeginString = "9"
	}

	return
}

func init() {
	tmplFuncs := template.FuncMap{
		"toLower":             strings.ToLower,
		"requiredFields":      requiredFields,
		"routerBeginString":   routerBeginString,
		"importRootPath":      gen.GetImportPathRoot,
		"quickfixType":        quickfixType,
		"quickfixValueType":   quickfixValueType,
		"getGlobalFieldType":  getGlobalFieldType,
		"collectExtraImports": collectExtraImports,
	}

	baseTemplate := template.Must(template.New("Base").Funcs(tmplFuncs).Parse(`
{{ define "receiver" }}RECEIVER{{ end }}

{{ define "fieldsetter" -}}
{{- $field_type := getGlobalFieldType . -}}
{{- $qfix_type := quickfixType $field_type -}}
Set{{ .Name }}(v {{ quickfixValueType $qfix_type }}) {
	{{ template "receiver" }}.Set(field.New{{ .Name }}(v))
}
{{- end }}

{{ define "groupsetter" -}}
Set{{ .Name }}(f {{ .Name }}RepeatingGroup){
	{{ template "receiver" }}.SetGroup(f)
}
{{- end }}

{{ define "setters" }}
{{ range .Fields }}
//Set{{ .Name }} sets {{ .Name }}, Tag {{ .Tag }}
func ({{ template "receiver" }} {{ $.Name }}) {{ if .IsGroup }}{{ template "groupsetter" . }}{{ else }}{{ template "fieldsetter" . }}{{ end }}
{{ end }}{{ end }}

{{ define "fieldgetter" -}}
Get{{ .Name }}() (f field.{{ .Name }}Field, err quickfix.MessageRejectError) {
	err = {{ template "receiver" }}.Get(&f)
	return
}
{{- end }}

{{ define "groupgetter" -}}
Get{{ .Name }}() ({{ .Name }}RepeatingGroup, quickfix.MessageRejectError) {
	f := New{{ .Name }}RepeatingGroup()
	err := {{ template "receiver" }}.GetGroup(f)
	return f, err
}
{{- end }}


{{ define "getters" }}
{{ range .Fields }}
//Get{{ .Name }} gets {{ .Name }}, Tag {{ .Tag }}
func ({{ template "receiver" }} {{ $.Name }}) {{if .IsGroup}}{{ template "groupgetter" . }}{{ else }}{{ template "fieldgetter" .}}{{ end }}
{{ end }}{{ end }}

{{ define "hasers" }}
{{range .Fields}}
//Has{{ .Name}} returns true if {{ .Name}} is present, Tag {{ .Tag}}
func ({{ template "receiver" }} {{ $.Name }}) Has{{ .Name}}() bool {
	return {{ template "receiver" }}.Has(tag.{{ .Name}})
}
{{end}}{{ end }}

{{ define "group_template" }}
quickfix.GroupTemplate{
{{- range $index, $field := . }}{{if $index}},{{end}}quickfix.GroupElement(tag.{{$field.Name}}){{ end }} }
{{- end }}

{{ define "field_args" }}
{{- range $index, $field := . }}{{if $index}},{{end}}{{toLower $field.Name}} field.{{ .Name }}Field{{ end }}
{{- end }}

{{ define "groups" }}
{{ range .Fields }}
{{ if .IsGroup }}
//{{ .Name }} is a repeating group element, Tag {{ .Tag }}
type {{ .Name }} struct {
	quickfix.Group
}

{{ template "setters" .}}
{{ template "getters" . }}
{{ template "hasers" . }}
{{ template "groups" . }}

//{{ .Name }}RepeatingGroup is a repeating group, Tag {{ .Tag }}
type {{ .Name }}RepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//New{{ .Name }}RepeatingGroup returns an initialized, {{ .Name }}RepeatingGroup
func New{{ .Name }}RepeatingGroup() {{ .Name }}RepeatingGroup {
	return {{ .Name }}RepeatingGroup{
		quickfix.NewRepeatingGroup(tag.{{ .Name }}, {{ template "group_template" .Fields }})}
}

//Add create and append a new {{ .Name }} to this group
func ({{ template "receiver" }} {{ .Name }}RepeatingGroup) Add() {{ .Name }} {
	g := {{ template "receiver" }}.RepeatingGroup.Add()
	return {{ .Name }}{g}
}

//Get returns the ith {{ .Name }} in the {{ .Name }}RepeatinGroup
func ({{ template "receiver" }} {{ .Name}}RepeatingGroup) Get(i int) {{ .Name }} {
	return {{ .Name }}{ {{ template "receiver" }}.RepeatingGroup.Get(i) }
}

{{ end }}{{ end }}{{ end }}
`))

	headerTemplate = template.Must(template.Must(baseTemplate.Clone()).Parse(`
{{ define "receiver" }}h{{ end }}
package {{ .Package }}

import(
	{{- range collectExtraImports .MessageDef }}
	"{{ . }}"
	{{- end }}


	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/tag"
)

//Header is the {{ .Package }} Header type
type Header struct {
	quickfix.Header
}

{{ template "setters" .}}
{{ template "getters" . }}
{{ template "hasers" . }}
{{ template "groups" . }}
	`))

	trailerTemplate = template.Must(template.Must(baseTemplate.Clone()).Parse(`
{{ define "receiver" }}t{{ end }}
package {{ .Package }}

import(
	{{- range collectExtraImports .MessageDef }}
	"{{ . }}"
	{{- end }}

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/tag"
)

//Trailer is the {{ .Package }} Trailer type
type Trailer struct {
	quickfix.Trailer
}

{{ template "setters" .}}
{{ template "getters" . }}
{{ template "hasers" . }}
{{ template "groups" . }}
`))

	messageTemplate = template.Must(baseTemplate.Parse(`
{{ define "receiver" }}m{{ end }}
package {{ .Package }}

import(
	"time"

	"github.com/quickfixgo/quickfix"
	"{{ importRootPath }}/field"
	"{{ importRootPath }}/{{ .TransportPackage }}"
	"{{ importRootPath }}/tag"
)

//{{ .Name }} is the {{ .FIXPackage }} {{ .Name }} type, MsgType = {{ .MsgType }}
type {{ .Name }} struct {
	{{ .TransportPackage }}.Header
	quickfix.Body
	{{ .TransportPackage }}.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a {{ .Name }} from a quickfix.Message instance
func FromMessage(m quickfix.Message) {{ .Name }} {
	return {{ .Name }}{
		Header: {{ .TransportPackage }}.Header{ Header: m.Header },
		Body: m.Body,
		Trailer: {{ .TransportPackage }}.Trailer{ Trailer: m.Trailer },
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m {{ .Name }}) ToMessage() quickfix.Message {
	return quickfix.Message {
		Header: m.Header.Header,
		Body: m.Body,
		Trailer: m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

{{ $required_fields := requiredFields .MessageDef -}}
//New returns a {{ .Name }} initialized with the required fields for {{ .Name }} 
func New({{template "field_args" $required_fields }}) (m {{ .Name }}) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("{{ .MessageDef.MsgType }}"))
	{{- range $required_fields }}
	m.Set({{ toLower .Name }})
	{{- end }}

	return 
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg {{ .Name }}, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r:=func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "{{ routerBeginString .FIXSpec }}", "{{ .MessageDef.MsgType }}", r
}

{{ template "setters" . }}
{{ template "getters" . }}
{{ template "hasers" . }}
{{ template "groups" . }}
	`))
}

func genHeader(pkg string, spec *datadictionary.DataDictionary) {
	defer waitGroup.Done()
	writer := new(bytes.Buffer)
	c := component{
		Package:    pkg,
		Name:       "Header",
		MessageDef: spec.Header,
	}
	if err := headerTemplate.Execute(writer, c); err != nil {
		errors <- err
		return
	}

	if err := gen.WriteFile(path.Join(pkg, "header.go"), writer.String()); err != nil {
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
	if err := trailerTemplate.Execute(writer, c); err != nil {
		errors <- err
		return
	}

	if err := gen.WriteFile(path.Join(pkg, "trailer.go"), writer.String()); err != nil {
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

	if err := messageTemplate.Execute(writer, c); err != nil {
		errors <- err
		return
	}

	if err := gen.WriteFile(path.Join(fixPkg, pkgName, msg.Name+".go"), writer.String()); err != nil {
		errors <- err
	}
}

func buildGlobalFieldMap(specs []*datadictionary.DataDictionary) {
	globalFieldTypes = make(fieldTypeMap)
	for _, spec := range specs {
		for _, field := range spec.FieldTypeByTag {
			globalFieldTypes[field.Name()] = field
		}
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
			dief("Error Parsing %v: %v", dataDictPath, err)
		}
	}

	buildGlobalFieldMap(specs)

	for _, spec := range specs {
		pkg := getPackageName(spec)

		if fi, err := os.Stat(pkg); os.IsNotExist(err) {
			if err := os.Mkdir(pkg, os.ModePerm); err != nil {
				die(err.Error())
			}
		} else if !fi.IsDir() {
			dief("%v/ is not a directory", pkg)
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

	var h gen.ErrorHandler
	for err := range errors {
		h.Handle(err)
	}

	os.Exit(h.ReturnCode)
}
