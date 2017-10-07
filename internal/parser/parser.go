package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type FieldDef struct {
	Name string
	Type string
}

type ModelDef struct {
	StructName string
	Fields     []FieldDef
}

// Parser golang structs
type Parser struct {
	filePath string
	fileSet  *token.FileSet
}

func New(filePath string) *Parser {
	return &Parser{
		filePath: filePath,
		fileSet:  token.NewFileSet(),
	}
}

func (p *Parser) Parse() ([]ModelDef, error) {
	parsed, err := parser.ParseFile(p.fileSet, p.filePath, nil, 0)
	if err != nil {
		return nil, err
	}

	for _, d := range parsed.Decls {
		switch d.(type) {
		case *ast.GenDecl:
			g := d.(*ast.GenDecl).Tok
			ast.Print(p.fileSet, g.String())
		}
	}

	return nil, nil
}
