package parser

import (
	"github.com/gigovich/fargo/orm/model"
)

// ModelDef options
type ModelDef struct {
	// Model instance
	Model model.Meta

	// TypeDefined in parsed file, and no need generate it struct
	TypeDefined bool
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
