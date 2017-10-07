package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Field definition
type Field struct {
	Name string
	Type string
}

// Model definition
type Model struct {
	StructName string
	Fields     []Field
}

// Parser golang file
type Parser struct {
	filePath string
	fileSet  *token.FileSet
	models   map[string]Model
}

// New parser instance
func New(filePath string) *Parser {
	return &Parser{
		filePath: filePath,
		models:   make(map[string]Model),
	}
}

// Parse go source file
func (p *Parser) Parse() ([]Model, error) {
	p.fileSet = token.NewFileSet()

	// parse file
	parsed, err := parser.ParseFile(p.fileSet, p.filePath, nil, 0)
	if err != nil {
		return nil, err
	}

	// modelTypes := make(map[string]struct{})
	for _, decl := range parsed.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range decl.Specs {
				switch spec := spec.(type) {
				case *ast.TypeSpec:
					if spec.Name.Obj.Kind == ast.Typ {
						if p.ProcessModelStruct(spec) {
							continue
						}
					}
				}
			}
		}
	}

	models := make([]Model, 0, len(p.models))
	for _, m := range p.models {
		models = append(models, m)
	}
	return models, nil
}

// ProcessModelStruct detect types which embeds orm.Model and inits internal models map
func (p *Parser) ProcessModelStruct(n ast.Node) bool {
	typeSpec, ok := n.(*ast.TypeSpec)
	if !ok {
		return false
	}

	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return false
	}

	for _, field := range structType.Fields.List {
		if field, ok := field.Type.(*ast.SelectorExpr); ok {
			if pac, ok := field.X.(*ast.Ident); ok && pac.Name == "orm" {
				if field.Sel.Name == "Model" {
					p.models[typeSpec.Name.Name] = Model{
						StructName: typeSpec.Name.Name,
					}
					return true
				}
			}
		}
	}

	return false
}
