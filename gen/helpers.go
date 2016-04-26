package gen

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"

	"github.com/quickfixgo/quickfix/datadictionary"
)

var genTemplate *template.Template

func init() {
	tmplFuncs := template.FuncMap{
		"fixFieldTypeToGoType": fixFieldTypeToGoType,
		"toLower":              strings.ToLower,
		"partAsGoType":         partAsGoType,
	}

	genTemplate = template.Must(template.New("Setters").Funcs(tmplFuncs).Parse(`
		{{/* template writes out the package */}}
		{{define "package"}}package {{.}}
		{{end}}

		{{/* template writes out appropriate imports for message/component */}}
		{{define "imports"}}
import ({{ range .}}
	"{{.}}"
{{- end }}
)
		{{ end}}


{{/* template writes out a constructor for component type */}}
{{define "newcomponent"}}
//New returns an initialized {{.Name}} instance
func New({{ template "parts_args" .RequiredParts}}) *{{.Name}} {
	var m {{.Name}}
	{{- template "set_parts" .RequiredParts}}
	return &m
}
{{end}}

{{/* template writes out a constructor for message type */}}
{{define "newmessage"}}
//New returns an initialized {{.Name}} instance
func New({{ template "parts_args" .RequiredParts}}) *Message {
	var m Message
	{{- template "set_parts" .RequiredParts}}
	return &m
}
{{end}}

{{/* template writes out a constructor for group */}}
{{define "newgroup"}}
//New{{.Name}} returns an initialized {{.Name}} instance
func New{{.Name}}({{ template "parts_args" .RequiredParts}}) *{{.Name}} {
	var m {{.Name}}
	{{- template "set_parts" .RequiredParts}}
	return &m
}
{{end}}

{{/* template writes out a comma delimited list of parts to be used as an argument list*/}}
{{define "parts_args"}}
{{- range $index, $field := . }}{{if $index}},{{end}}{{toLower $field.Name}} {{partAsGoType $field}}{{ end }}
{{- end }}

{{/* template sets  parts*/}}
{{define "set_parts"}}
{{- range .}}
m.Set{{.Name}}({{toLower .Name}})
{{- end}}
{{- end}}

{{define "fieldSetter"}}
func (m *{{.Receiver}}) Set{{.Name}}(v {{ if .IsGroup}}[]{{.Name}}{{else}}{{fixFieldTypeToGoType .Type}}{{end}}) {
{{- if .IsGroup -}}m.{{.Name}} = v
{{- else if .Required -}}m.{{.Name}} = v
{{- else -}}m.{{.Name}} = &v
{{- end}}}{{end}}

{{define "compSetter"}}
func (m *{{.Receiver}}) Set{{.Name}}(v {{toLower .Name}}.{{ .Name}}) {
{{- if .Required -}}m.{{.Name}} = v
{{- else -}}m.{{.Name}} = &v
{{- end}}}{{end}}`))
}

func WriteNewComponent(writer io.Writer, comp datadictionary.ComponentType) error {
	return genTemplate.ExecuteTemplate(writer, "newcomponent", comp)
}

func WriteNewMessage(writer io.Writer, msg datadictionary.MessageDef) error {
	return genTemplate.ExecuteTemplate(writer, "newmessage", msg)
}

//WriteFieldSetters generates setters appropriate for Messages, Components or Repeating Groups
func WriteFieldSetters(writer io.Writer, receiver string, parts []datadictionary.MessagePart) error {
	type setter struct {
		Receiver string
		*datadictionary.FieldDef
	}

	type componentSetter struct {
		Receiver string
		datadictionary.Component
	}

	for _, part := range parts {
		switch field := part.(type) {
		case datadictionary.Component:
			if err := genTemplate.ExecuteTemplate(writer, "compSetter", componentSetter{receiver, field}); err != nil {
				return err
			}
		case *datadictionary.FieldDef:
			if err := genTemplate.ExecuteTemplate(writer, "fieldSetter", setter{receiver, field}); err != nil {
				return err
			}
		}
	}

	return nil
}

func HeaderTrailerPkg(pkg string) string {
	switch pkg {
	case "fix50", "fix50sp1", "fix50sp2":
		return "fixt11"
	}

	return pkg
}

func WritePackage(writer io.Writer, pkg string) error {
	return genTemplate.ExecuteTemplate(writer, "package", pkg)
}

func WriteMessageImports(writer io.Writer, pkg string, parts []datadictionary.MessagePart) error {
	imports := []string{
		"github.com/quickfixgo/quickfix",
		fmt.Sprintf("%s/enum", GetImportPathRoot()),
		fmt.Sprintf("%s/%v", GetImportPathRoot(), HeaderTrailerPkg(pkg)),
	}
	return writeImports(writer, pkg, imports, parts)
}

func WriteComponentImports(writer io.Writer, pkg string, parts []datadictionary.MessagePart) error {
	return writeImports(writer, pkg, nil, parts)
}

func writeImports(writer io.Writer, pkg string, imports []string, parts []datadictionary.MessagePart) error {
	importMap := make(map[string]interface{})
	for _, part := range parts {
		collectRequiredImports(importMap, pkg, part)
	}

	if len(imports) == 0 && len(importMap) == 0 {
		return nil
	}

	for ip := range importMap {
		imports = append(imports, ip)
	}

	return genTemplate.ExecuteTemplate(writer, "imports", imports)
}

func collectRequiredImports(imports map[string]interface{}, pkg string, part datadictionary.MessagePart) {
	if comp, ok := part.(datadictionary.Component); ok {
		imports[fmt.Sprintf("%s/%v/%v", GetImportPathRoot(), pkg, strings.ToLower(comp.Name()))] = nil

		return
	}

	field, ok := part.(*datadictionary.FieldDef)
	if !ok {
		panic("Expected FieldDef")
	}

	if fieldType := fixFieldTypeToGoType(field.Type); fieldType == "time.Time" {
		imports["time"] = nil
	}

	for _, part := range field.Parts {
		collectRequiredImports(imports, pkg, part)
	}
}

func WriteFieldDeclarations(fixSpecMajor int, fixSpecMinor int, parts []datadictionary.MessagePart, componentName string) (s string) {
	for _, part := range parts {
		s += writeFieldDeclaration(fixSpecMajor, fixSpecMinor, part, componentName)
	}
	return
}

func WriteGroupDeclaration(fixSpecMajor, fixSpecMinor int, field *datadictionary.FieldDef, parent string) (fileOut string) {
	fileOut += fmt.Sprintf("//%v is a repeating group in %v\n", field.Name(), parent)
	fileOut += fmt.Sprintf("type %v struct {\n", field.Name())
	fileOut += WriteFieldDeclarations(fixSpecMajor, fixSpecMinor, field.Parts, field.Name())
	fileOut += "}\n"

	writer := new(bytes.Buffer)
	if err := genTemplate.ExecuteTemplate(writer, "newgroup", *field); err != nil {
		panic(err)
	}

	if err := WriteFieldSetters(writer, field.Name(), field.Parts); err != nil {
		panic(err)
	}
	fileOut += writer.String()
	fileOut += "\n"

	return
}

func writeFieldDeclaration(fixSpecMajor int, fixSpecMinor int, part datadictionary.MessagePart, componentName string) (s string) {
	switch field := part.(type) {
	case datadictionary.Component:
		if field.Required() {
			s += fmt.Sprintf("//%v is a required component for %v.\n", field.Name(), componentName)
			s += fmt.Sprintf("%v.%v\n", strings.ToLower(field.Name()), field.Name())
		} else {
			s += fmt.Sprintf("//%v is a non-required component for %v.\n", field.Name(), componentName)
			s += fmt.Sprintf("%v *%v.%v\n", field.Name(), strings.ToLower(field.Name()), field.Name())
		}
		return

	case *datadictionary.FieldDef:
		if field.Required() {
			s += fmt.Sprintf("//%v is a required field for %v.\n", field.Name(), componentName)
		} else {
			s += fmt.Sprintf("//%v is a non-required field for %v.\n", field.Name(), componentName)
		}

		if field.IsGroup() {
			if field.Required() {
				s += fmt.Sprintf("%v []%v `fix:\"%v\"`\n", field.Name(), field.Name(), field.Tag())
			} else {
				s += fmt.Sprintf("%v []%v `fix:\"%v,omitempty\"`\n", field.Name(), field.Name(), field.Tag())
			}
			return
		}

		goType := fixFieldTypeToGoType(field.Type)
		fixTags := strconv.Itoa(field.Tag())
		if field.Tag() == 8 {
			if fixSpecMajor == 4 {
				fixTags = fmt.Sprintf("%v,default=FIX.%v.%v", fixTags, fixSpecMajor, fixSpecMinor)
			} else {
				fixTags = fixTags + ",default=FIXT.1.1"
			}
		}

		if field.Required() {
			s += fmt.Sprintf("%v %v `fix:\"%v\"`\n", field.Name(), goType, fixTags)
		} else {
			s += fmt.Sprintf("%v *%v `fix:\"%v\"`\n", field.Name(), goType, fixTags)
		}
	}
	return
}

func partAsGoType(part datadictionary.MessagePart) string {
	if comp, ok := part.(datadictionary.Component); ok {
		return fmt.Sprintf("%v.%v", strings.ToLower(comp.Name()), comp.Name())
	}

	field, ok := part.(*datadictionary.FieldDef)
	if !ok {
		panic("unknown part")
	}

	if field.IsGroup() {
		return fmt.Sprintf("[]%v", field.Name())
	}

	return fixFieldTypeToGoType(field.Type)
}

func fixFieldTypeToGoType(fieldType string) string {
	switch fieldType {
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
		return "string"

	case "BOOLEAN":
		return "bool"

	case "LENGTH":
		fallthrough
	case "DAYOFMONTH":
		fallthrough
	case "NUMINGROUP":
		fallthrough
	case "SEQNUM":
		fallthrough
	case "INT":
		return "int"

	case "TIME":
		fallthrough
	case "UTCTIMESTAMP":
		return "time.Time"

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
		return "float64"

	default:
		panic(fmt.Sprintf("Unknown type '%v'\n", fieldType))
	}
}

// GetImportPathRoot returns the root path to use in import statements.
// The root path is determined by stripping "$GOPATH/src/" from the current
// working directory.  For example, when generating code within the QuickFIX/Go
// source tree, the returned root path will be "github.com/quickfixgo/quickfix".
func GetImportPathRoot() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	goSrcPath := path.Join(os.Getenv("GOPATH"), "src")
	importPathRoot := strings.Replace(pwd, goSrcPath, "", 1)
	return strings.TrimLeft(importPathRoot, "/")
}
