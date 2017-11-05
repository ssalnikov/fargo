package generator

import (
	"github.com/gigovich/fargo/orm/field"
	"strings"
	"text/template"
)

var moduleTemplate = template.Must(
	template.New("module").
		Funcs(template.FuncMap{
			"fieldname":  formatFieldName,
			"recordname": formatRecordName,
			"gettags":    getTags,
		}).
		Parse(
			`package {{ .PkgName }}

import (
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/mod"
	"github.com/gigovich/fargo/orm/model"
	"github.com/gigovich/fargo/orm/query"
)

{{range $modelName, $modelDef := .DefList }}
{{if not $modelDef.TypeDefined}}
// {{$modelName}} embeds meta mapper
type {{$modelName}} struct {
	model.Meta
}

// {{recordname $modelName}} data object
type {{recordname $modelName}} struct { {{range $fieldIndex, $fieldDef := $modelDef.Model.Fields}}
	// {{fieldname $fieldDef.GetMeta.Name}} field
	{{fieldname $fieldDef.GetMeta.Name}} string {{gettags $fieldDef.GetMeta}}
{{end}}
}
{{end}}

{{range $fieldIndex, $fieldDef := $modelDef.Model.Fields}}
// {{fieldname $fieldDef.GetMeta.Name}} returns field mapper for column '{{$fieldDef.GetMeta.Name}}'
func (m *{{$modelName}}) {{fieldname $fieldDef.GetMeta.Name}}() field.Mapper {
	return m.Fields[{{$fieldIndex}}]
}
{{end}}

// Query records for '{{$modelName}}'
func (m *{{$modelName}}) Query(mods ...mod.Modifier) *model.Query {
	return &query.Query{}
}

// Find returns first element from executed query
func (m *{{$modelName}}) Find(query model.Query) ([]{{recordname $modelName}}, error) {
	return nil, nil
}

// One returns first element from executed query
func (m *{{$modelName}}) One(query model.Query) (*{{recordname $modelName}}, error) {
	return nil, nil
}
{{end}}`))

func formatFieldName(name string) (modified string) {
	for _, part := range strings.Split(name, "_") {
		if substitute, ok := reservedFieldNames[strings.ToLower(part)]; ok {
			modified += substitute
		} else {
			modified += strings.Title(part)
		}
	}
	return modified
}

func formatRecordName(name string) (modifier string) {
	return strings.TrimSuffix(name, "Model") + "Record"
}

func getTags(f field.Meta) string {
	if f.Tags == "" {
		return "`json:\"" + f.Name + "\"`"
	}
	return "`" + f.Tags + "`"
}
