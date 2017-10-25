package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Parser golang file
type Parser struct {
	filePath    string
	modelParser *modelParser
	fileSet     *token.FileSet
}

// New parser instance
func New(filePath string) *Parser {
	return &Parser{
		modelParser: newModelParser("model", "field"),
		filePath:    filePath,
	}
}

// Parse go source file
func (p *Parser) Parse() error {
	p.fileSet = token.NewFileSet()

	// parse file
	parsed, err := parser.ParseFile(p.fileSet, p.filePath, nil, 0)
	if err != nil {
		return err
	}

	ast.Inspect(parsed, p.InspectTypes)
	ast.Inspect(parsed, p.InspectVars)

	return nil
}

// InspectTypes declarations
func (p *Parser) InspectTypes(n ast.Node) bool {
	gd, ok := n.(*ast.GenDecl)
	if !ok || gd.Tok != token.TYPE {
		return true
	}

	return p.modelParser.InspectTypeSpecs(gd.Specs)
}

// InspectVars declarations
func (p *Parser) InspectVars(n ast.Node) bool {
	gd, ok := n.(*ast.GenDecl)
	if !ok || gd.Tok != token.VAR {
		return true
	}
	return p.modelParser.InspectVarsSpecs(gd.Specs)
}
