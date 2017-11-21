package generator

import (
	"text/template"

	"github.com/gigovich/fargo/internal/util"
)

var moduleTemplate = template.Must(
	template.New("module").
		Funcs(template.FuncMap{
			"fieldname":  util.FormatFieldName,
			"recordname": util.FormatRecordName,
			"gettags":    util.GetTags,
		}).
		Parse(
			`package {{ .PkgName }}

import (
	"github.com/gigovich/fargo/orm/mod"
	"github.com/gigovich/fargo/orm/model"
	"github.com/gigovich/fargo/orm/query"
)

{{range $modelName, $modelDef := .DefList }}
{{if not $modelDef.TypeDefined}}
// {{$modelName}} embends model meta mapper
type {{$modelName}} struct {
	model.Base
}
{{end}}

// {{recordname $modelName}} data object
type {{recordname $modelName}} struct { {{range $fieldIndex, $fieldDef := $modelDef.Model.Fields}}
	// {{fieldname $fieldDef.GetField.Name}} field
	{{fieldname $fieldDef.GetField.Name}} string {{gettags $fieldDef.GetField}}
{{end}}
}

{{range $fieldIndex, $fieldDef := $modelDef.Model.Fields}}
// {{fieldname $fieldDef.GetField.Name}} returns field mapper for column '{{$fieldDef.GetField.Name}}'
func (m *{{$modelName}}) {{fieldname $fieldDef.GetField.Name}}() model.Field {
	return model.Field{Model: m, Field: m.Fields[{{$fieldIndex}}]}
}
{{if $fieldDef.GetField.Primary}}
// GetPrimaryKey field of {{$modelName}}
func (m *{{$modelName}}) GetPrimaryKey() model.Field {
	return model.Field{Model: m, Field: m.Fields[{{$fieldIndex}}]}
}
{{end}}
{{end}}

// Insert '{{recordname $modelName}}' to database
func (m *{{$modelName}}) Insert(records ...{{recordname $modelName}}) error {
	return nil
}

// Query records for '{{$modelName}}'
func (m *{{$modelName}}) Query(mods ...mod.Modifier) *query.Query {
	return &query.Query{}
}

// Find returns first element from executed query
func (m *{{$modelName}}) Find(query *query.Query) ([]{{recordname $modelName}}, error) {
	return nil, nil
}

// One returns first element from executed query
func (m *{{$modelName}}) One(query *query.Query) (*{{recordname $modelName}}, error) {
	return nil, nil
}
{{end}}`))
