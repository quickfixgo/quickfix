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
		"toLower":                               strings.ToLower,
		"requiredFields":                        requiredFields,
		"beginString":                           beginString,
		"routerBeginString":                     routerBeginString,
		"importRootPath":                        getImportPathRoot,
		"quickfixType":                          quickfixType,
		"quickfixValueType":                     quickfixValueType,
		"getGlobalFieldType":                    getGlobalFieldType,
		"collectExtraImports":                   collectExtraImports,
		"checkIfDecimalImportRequiredForFields": checkIfDecimalImportRequiredForFields,
		"checkIfEnumImportRequired":             checkIfEnumImportRequired,
	}

	baseTemplate := template.Must(template.New("Base").Funcs(tmplFuncs).Parse(`
{{ define "receiver" }}RECEIVER{{ end }}

{{ define "fieldsetter" -}}
{{- $field_type := getGlobalFieldType . -}}
{{- $qfix_type := quickfixType $field_type -}}
{{- if and ($field_type.Enums) (ne $qfix_type "FIXBoolean") -}}
Set{{ .Name }}(v enum.{{ .Name }}) {
	{{ template "receiver" }}.Set(field.New{{ .Name }}(v))
}
{{- else if eq $qfix_type "FIXDecimal" -}}
Set{{ .Name }}(value decimal.Decimal, scale int32) {
	{{ template "receiver" }}.Set(field.New{{ .Name }}(value, scale))
}
{{- else -}}
Set{{ .Name }}(v {{ quickfixValueType $qfix_type }}) {
	{{ template "receiver" }}.Set(field.New{{ .Name }}(v))
}
{{- end }}{{ end }}

{{ define "groupsetter" -}}
Set{{ .Name }}(f {{ .Name }}RepeatingGroup){
	{{ template "receiver" }}.SetGroup(f)
}
{{- end }}

{{ define "setters" }}
{{ range .Fields }}
// Set{{ .Name }} sets {{ .Name }}, Tag {{ .Tag }}.
func ({{ template "receiver" }} {{ $.Name }}) {{ if .IsGroup }}{{ template "groupsetter" . }}{{ else }}{{ template "fieldsetter" . }}{{ end }}
{{ end }}{{ end }}

{{ define "fieldgetter" -}}
Get{{ .Name }}() (f field.{{ .Name }}Field, err quickfix.MessageRejectError) {
	err = {{ template "receiver" }}.Get(&f)
	return
}
{{- end }}

{{ define "fieldvaluegetter" -}}
{{- $ft := getGlobalFieldType . -}}
{{- $bt := quickfixType $ft -}}
{{- if and $ft.Enums (ne $bt "FIXBoolean") -}}
Get{{ .Name }}() (v enum.{{ .Name }}, err quickfix.MessageRejectError) {
{{- else if eq $bt "FIXDecimal" -}}
Get{{ .Name }}() (v decimal.Decimal, err quickfix.MessageRejectError) {
{{- else -}}
Get{{ .Name }}() (v {{ quickfixValueType $bt }}, err quickfix.MessageRejectError) {
{{- end }}
	var f field.{{ .Name }}Field
	if err = {{ template "receiver" }}.Get(&f); err == nil {
		v = f.Value()
	}
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
// Get{{ .Name }} gets {{ .Name }}, Tag {{ .Tag }}.
func ({{ template "receiver" }} {{ $.Name }}) {{if .IsGroup}}{{ template "groupgetter" . }}{{ else }}{{ template "fieldvaluegetter" .}}{{ end }}
{{ end }}{{ end }}

{{ define "hasers" }}
{{range .Fields}}
// Has{{ .Name}} returns true if {{ .Name}} is present, Tag {{ .Tag}}.
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
// {{ .Name }} is a repeating group element, Tag {{ .Tag }}.
type {{ .Name }} struct {
	*quickfix.Group
}

{{ template "setters" .}}
{{ template "getters" . }}
{{ template "hasers" . }}
{{ template "groups" . }}

// {{ .Name }}RepeatingGroup is a repeating group, Tag {{ .Tag }}.
type {{ .Name }}RepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// New{{ .Name }}RepeatingGroup returns an initialized, {{ .Name }}RepeatingGroup.
func New{{ .Name }}RepeatingGroup() {{ .Name }}RepeatingGroup {
	return {{ .Name }}RepeatingGroup{
		quickfix.NewRepeatingGroup(tag.{{ .Name }}, {{ template "group_template" .Fields }})}
}

// Add create and append a new {{ .Name }} to this group.
func ({{ template "receiver" }} {{ .Name }}RepeatingGroup) Add() {{ .Name }} {
	g := {{ template "receiver" }}.RepeatingGroup.Add()
	return {{ .Name }}{g}
}

// Get returns the ith {{ .Name }} in the {{ .Name }}RepeatinGroup.
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
	{{- if checkIfEnumImportRequired .MessageDef}}
	"{{ importRootPath }}/enum"
	{{- end }}
	"{{ importRootPath }}/field"
	"{{ importRootPath }}/tag"
)

// Header is the {{ .Package }} Header type.
type Header struct {
	*quickfix.Header
}

// NewHeader returns a new, initialized Header instance.
func NewHeader(header *quickfix.Header) (h Header) {
	h.Header = header
	h.SetBeginString("{{ beginString .FIXSpec }}")
	return
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
	{{- if checkIfEnumImportRequired .MessageDef}}
	"{{ importRootPath }}/enum"
	{{- end }}
	"{{ importRootPath }}/field"
	"{{ importRootPath }}/tag"
)

// Trailer is the {{ .Package }} Trailer type.
type Trailer struct {
	*quickfix.Trailer
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
	{{- range collectExtraImports .MessageDef }}
	"{{ . }}"
	{{- end }}

	"github.com/quickfixgo/quickfix"
	{{- if checkIfEnumImportRequired .MessageDef}}
	"{{ importRootPath }}/enum"
	{{- end }}
	"{{ importRootPath }}/field"
	"{{ importRootPath }}/{{ .TransportPackage }}"
	"{{ importRootPath }}/tag"
)

// {{ .Name }} is the {{ .FIXPackage }} {{ .Name }} type, MsgType = {{ .MsgType }}.
type {{ .Name }} struct {
	{{ .TransportPackage }}.Header
	*quickfix.Body
	{{ .TransportPackage }}.Trailer
	Message *quickfix.Message
}

// FromMessage creates a {{ .Name }} from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) {{ .Name }} {
	return {{ .Name }}{
		Header: {{ .TransportPackage}}.Header{&m.Header},
		Body: &m.Body,
		Trailer: {{ .TransportPackage}}.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m {{ .Name }}) ToMessage() *quickfix.Message {
	return m.Message
}

{{ $required_fields := requiredFields .MessageDef -}}
// New returns a {{ .Name }} initialized with the required fields for {{ .Name }}.
func New({{template "field_args" $required_fields }}) (m {{ .Name }}) {
	m.Message = quickfix.NewMessage()
	m.Header = {{ .TransportPackage }}.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("{{ .MessageDef.MsgType }}"))
	{{- range $required_fields }}
	m.Set({{ toLower .Name }})
	{{- end }}

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg {{ .Name }}, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r:=func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
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
	"{{ importRootPath }}/enum"
	"{{ importRootPath }}/tag"
{{ if checkIfDecimalImportRequiredForFields . }} "github.com/shopspring/decimal" {{ end }}
	"time"
)

{{ range . }}
{{- $base_type := quickfixType . -}}

{{ if and .Enums (ne $base_type "FIXBoolean") }}
// {{ .Name }}Field is a enum.{{ .Name }} field.
type {{ .Name }}Field struct { quickfix.FIXString }
{{ else }}
// {{ .Name }}Field is a {{ .Type }} field.
type {{ .Name }}Field struct { quickfix.{{ $base_type }} }
{{ end }}

// Tag returns tag.{{ .Name }} ({{ .Tag }}).
func (f {{ .Name }}Field) Tag() quickfix.Tag { return tag.{{ .Name }} }

{{ if eq $base_type "FIXUTCTimestamp" }}
// New{{ .Name }} returns a new {{ .Name }}Field initialized with val.
func New{{ .Name }}(val time.Time) {{ .Name }}Field {
	return New{{ .Name }}WithPrecision(val, quickfix.Millis)
}

// New{{ .Name }}NoMillis returns a new {{ .Name }}Field initialized with val without millisecs.
func New{{ .Name }}NoMillis(val time.Time) {{ .Name }}Field {
	return New{{ .Name }}WithPrecision(val, quickfix.Seconds)
}

// New{{ .Name }}WithPrecision returns a new {{ .Name }}Field initialized with val of specified precision.
func New{{ .Name }}WithPrecision(val time.Time, precision quickfix.TimestampPrecision) {{ .Name }}Field {
	return {{ .Name }}Field{ quickfix.FIXUTCTimestamp{ Time: val, Precision: precision } }
}

{{ else if and  .Enums (ne $base_type "FIXBoolean") }}
func New{{ .Name }}(val enum.{{ .Name }}) {{ .Name }}Field {
	return {{ .Name }}Field{ quickfix.FIXString(val) }
}
{{ else if eq $base_type "FIXDecimal" }}
// New{{ .Name }} returns a new {{ .Name }}Field initialized with val and scale.
func New{{ .Name }}(val decimal.Decimal, scale int32) {{ .Name }}Field {
	return {{ .Name }}Field{ quickfix.FIXDecimal{ Decimal: val, Scale: scale} }
}
{{ else }}
// New{{ .Name }} returns a new {{ .Name }}Field initialized with val.
func New{{ .Name }}(val {{ quickfixValueType $base_type }}) {{ .Name }}Field {
	return {{ .Name }}Field{ quickfix.{{ $base_type }}(val) }
}
{{ end }}

{{ if and  .Enums (ne $base_type "FIXBoolean") }}
func (f {{ .Name }}Field) Value() enum.{{ .Name }} { return enum.{{ .Name }}(f.String()) }
{{ else if eq $base_type "FIXDecimal" }}
func (f {{ .Name }}Field) Value() (val decimal.Decimal) { return f.Decimal }
{{ else }}
func (f {{ .Name }}Field) Value() ({{ quickfixValueType $base_type }}) {
{{- if eq $base_type "FIXString" -}}
 return f.String() }
{{- else if eq $base_type "FIXBoolean" -}}
 return f.Bool() }
{{- else if eq $base_type "FIXInt" -}}
 return f.Int() }
{{- else if eq $base_type "FIXUTCTimestamp" -}}
 return f.Time }
{{- else if eq $base_type "FIXFloat" -}}
 return f.Float() }
{{- else -}}
 TEMPLATE ERROR: Value() for {{ $base_type }}
{{ end }}{{ end }}{{ end }}
`))

	EnumTemplate = template.Must(template.New("Enum").Parse(`
package enum
{{ range $ft := . }}
{{ if $ft.Enums }}
// {{ $ft.Name }} field enumeration values.
type {{ $ft.Name }} string
const(
{{ range $ft.Enums }}
{{ $ft.Name }}_{{ .Description }} {{ $ft.Name }} = "{{ .Value }}"
{{- end }}
)
{{ end }}{{ end }}
	`))
}
