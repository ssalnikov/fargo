package parser

import (
	"github.com/gigovich/fargo/orm/model"
	"go/ast"
	"go/parser"
	"go/token"
)

// Parser golang file
type Parser struct {
	filePath    string
	models      map[string]*model.Meta
	modelParser *modelParser
	fileSet     *token.FileSet
}

// New parser instance
func New(filePath string) *Parser {
	models := make(map[string]*model.Meta)
	return &Parser{
		models:      models,
		modelParser: newModelParser(models),
		filePath:    filePath,
	}
}

// Parse go source file
func (p *Parser) Parse() (map[string]*model.Meta, error) {
	p.fileSet = token.NewFileSet()

	ctx := &Context{}
	ctx.Model.Package.Name = "model"
	ctx.Field.Package.Name = "field"

	// parse file
	parsed, err := parser.ParseFile(p.fileSet, p.filePath, nil, 0)
	if err != nil {
		return nil, err
	}

	ast.Inspect(parsed, p.InspectTypes(ctx))
	ast.Inspect(parsed, p.InspectVars(ctx))

	return p.models, nil
}

// InspectTypes declarations
func (p *Parser) InspectTypes(ctx *Context) func(ast.Node) bool {
	return func(n ast.Node) bool {
		gd, ok := n.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			return true
		}

		return p.modelParser.InspectTypeSpecs(ctx, gd.Specs)
	}
}

// InspectVars declarations
func (p *Parser) InspectVars(ctx *Context) func(ast.Node) bool {
	return func(n ast.Node) bool {
		gd, ok := n.(*ast.GenDecl)
		if !ok || gd.Tok != token.VAR {
			return true
		}
		return p.modelParser.InspectVarsSpecs(ctx, gd.Specs)
	}
}
