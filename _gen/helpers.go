package gen

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/template"

	"github.com/quickfixgo/quickfix/datadictionary"
)

var fieldSetterTemplate *template.Template

func init() {
	tmplFuncs := template.FuncMap{
		"fixFieldTypeToGoType": FixFieldTypeToGoType,
		"toLower":              strings.ToLower,
	}

	fieldSetterTemplate = template.Must(template.New("Setters").Funcs(tmplFuncs).Parse(`
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
			if err := fieldSetterTemplate.ExecuteTemplate(writer, "compSetter", componentSetter{receiver, field}); err != nil {
				return err
			}
		case *datadictionary.FieldDef:
			if err := fieldSetterTemplate.ExecuteTemplate(writer, "fieldSetter", setter{receiver, field}); err != nil {
				return err
			}
		}
	}

	return nil
}

func WriteFieldDeclaration(fixSpecMajor int, fixSpecMinor int, part datadictionary.MessagePart, componentName string) (s string) {
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
				s += fmt.Sprintf("%v []%v `fix:\"%v\"`\n", field.Name(), field.Name(), field.Tag)
			} else {
				s += fmt.Sprintf("%v []%v `fix:\"%v,omitempty\"`\n", field.Name(), field.Name(), field.Tag)
			}
			return
		}

		goType := FixFieldTypeToGoType(field.Type)
		fixTags := strconv.Itoa(field.Tag)
		if field.Tag == 8 {
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

func FixFieldTypeToGoType(fieldType string) string {
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
