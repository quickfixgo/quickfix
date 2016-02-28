package gen

import (
	"fmt"
	"github.com/quickfixgo/quickfix/datadictionary"
	"io"
	"text/template"
)

var fieldSetterTemplate *template.Template

func init() {
	tmplFuncs := make(template.FuncMap)
	tmplFuncs["fixFieldTypeToGoType"] = fixFieldTypeToGoType

	fieldSetterTemplate = template.Must(template.New("Setters").Funcs(tmplFuncs).Parse(`
func (m *{{.Receiver}}) Set{{.Name}}(v {{ if .IsGroup}}[]{{.Name}}{{else}}{{fixFieldTypeToGoType .Type}}{{end}}) {
{{- if .IsGroup -}}m.{{.Name}} = v
{{- else if .Required -}}m.{{.Name}} = v
{{- else -}}m.{{.Name}} = &v
{{- end}}}`))
}

//WriteFieldSetters generates setters appropriate for Messages, Components and Repeating Groups
func WriteFieldSetters(writer io.Writer, receiver string, fields []*datadictionary.FieldDef) error {
	type setter struct {
		Receiver string
		*datadictionary.FieldDef
	}

	for _, field := range fields {
		if field.IsComponent() {
			continue
		}

		if err := fieldSetterTemplate.Execute(writer, setter{receiver, field}); err != nil {
			return err
		}
	}

	return nil
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
