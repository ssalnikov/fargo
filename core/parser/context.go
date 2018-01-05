package parser

import (
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/model"
)

type FieldDef struct {
	// Field instance
	Field field.Mapper

	// MethodName in model instance which returns this field
	MethodName string
}

// ModelDef options
type ModelDef struct {
	// Model instance
	Model model.Base

	// Fields map by their model instance getter methods names
	Fields map[string]*FieldDef

	// TypeDefined in parsed file, and no need generate it struct
	TypeDefined bool
}

// NewFieldDef instance constructo, name should be method name not name from table
func NewFieldDef(fieldName string) *FieldDef {
	return &FieldDef{
		// Field mapper instance
		Field: &field.Base{},

		// MethodName which return field from model
		MethodName: fieldName,
	}
}

// NewModelDef instance constructor
func NewModelDef(modelName string) *ModelDef {
	return &ModelDef{
		Model: model.Base{
			Table: modelName,
		},
		Fields: make(map[string]*FieldDef),
	}
}

// Context for parsing
type Context struct {
	// currentDef for parsing
	currentDef *ModelDef

	// DefList map
	DefList map[string]*ModelDef

	// PkgFile path for parse
	PkgFile string

	// PkgName contains name of parsed package
	PkgName string

	// ModelImport name, by default `model`
	ModelImport string

	// FieldImport name, by default `field`
	FieldImport string
}
