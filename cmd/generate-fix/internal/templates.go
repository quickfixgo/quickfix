package internal

import (
	"strings"
	"text/template"
)

var (
	HeaderTemplate  *template.Template
	TrailerTemplate *template.Template
	MessageTemplate *template.Template
	TagTemplate     *template.Template
	FieldTemplate   *template.Template
	EnumTemplate    *template.Template
)

func init() {
	tmplFuncs := template.FuncMap{
		"toLower":             strings.ToLower,
		"requiredFields":      requiredFields,
		"routerBeginString":   routerBeginString,
		"importRootPath":      getImportPathRoot,
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
{{- range $index, $field := . }}{{if $index}},{{end}}{{if $field.IsGroup }}New{{ $field.Name }}RepeatingGroup(){{else}}quickfix.GroupElement(tag.{{$field.Name}}){{ end }}{{ end }} }
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

	HeaderTemplate = template.Must(template.Must(baseTemplate.Clone()).Parse(`
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

	TrailerTemplate = template.Must(template.Must(baseTemplate.Clone()).Parse(`
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

	MessageTemplate = template.Must(baseTemplate.Parse(`
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
	m.Header.Set(field.NewBeginString("{{ routerBeginString .FIXSpec }}"))
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

	TagTemplate = template.Must(template.New("Tag").Parse(`
package tag
import("github.com/quickfixgo/quickfix")

const (
{{- range .}}
{{ .Name }} quickfix.Tag =  {{ .Tag }}
{{- end }}
)
	`))

	FieldTemplate = template.Must(template.New("Field").Funcs(tmplFuncs).Parse(`
package field
import(
	"github.com/quickfixgo/quickfix"
	"{{ importRootPath }}/tag"

	"time"
)

{{ range . }}
{{- $base_type := quickfixType . -}}
//{{ .Name }}Field is a {{ .Type }} field
type {{ .Name }}Field struct { quickfix.{{ $base_type }} }

//Tag returns tag.{{ .Name }} ({{ .Tag }})
func (f {{ .Name }}Field) Tag() quickfix.Tag { return tag.{{ .Name }} }

{{ if eq $base_type "FIXUTCTimestamp" }} 
//New{{ .Name }} returns a new {{ .Name }}Field initialized with val
func New{{ .Name }}(val time.Time) {{ .Name }}Field {
	return {{ .Name }}Field{ quickfix.FIXUTCTimestamp{ Time: val } } 
}

//New{{ .Name }}NoMillis returns a new {{ .Name }}Field initialized with val without millisecs
func New{{ .Name }}NoMillis(val time.Time) {{ .Name }}Field {
	return {{ .Name }}Field{ quickfix.FIXUTCTimestamp{ Time: val, NoMillis: true } } 
}

{{ else }}
//New{{ .Name }} returns a new {{ .Name }}Field initialized with val
func New{{ .Name }}(val {{ quickfixValueType $base_type }}) {{ .Name }}Field {
	return {{ .Name }}Field{ quickfix.{{ $base_type }}(val) }
}
{{ end }}{{ end }}
`))

	EnumTemplate = template.Must(template.New("Enum").Parse(`
package enum
{{ range $ft := . }}
{{ if $ft.Enums }} 
//Enum values for {{ $ft.Name }}
const(
{{ range $ft.Enums }}
{{ $ft.Name }}_{{ .Description }} = "{{ .Value }}"
{{- end }}
)
{{ end }}{{ end }}
	`))
}
