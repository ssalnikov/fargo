package parser

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/model"
)

const newModelFuncName = "New"

// modelParser definition parser
type modelParser struct {
	fieldParser *fieldParser
	models      map[string]*model.Meta
}

func newModelParser(models map[string]*model.Meta) *modelParser {
	return &modelParser{
		fieldParser: newFieldParser(),
		models:      models,
	}
}

// InspectTypeSpecs collect all model declarations
func (p *modelParser) InspectTypeSpecs(ctx *Context, specs []ast.Spec) bool {
	for _, sp := range specs {
		ts, ok := sp.(*ast.TypeSpec)
		if !ok {
			continue
		}

		name := ts.Name.String()
		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			continue
		}

		for _, fieldItem := range st.Fields.List {
			se, ok := fieldItem.Type.(*ast.SelectorExpr)
			if !ok {
				continue
			}

			id, ok := se.X.(*ast.Ident)
			if !ok {
				continue
			}

			if id.Name == ctx.Model.Package.Name {
				p.models[name] = &model.Meta{
					Table: name,
				}
			}
		}
	}
	return true
}

// InspectVarsSpecs collect all model variables
func (p *modelParser) InspectVarsSpecs(ctx *Context, specs []ast.Spec) bool {
	for _, sp := range specs {
		vs, ok := sp.(*ast.ValueSpec)
		if !ok {
			continue
		}

		// we don't support slice of model definitions
		if len(vs.Values) > 1 || len(vs.Values) < 1 {
			continue
		}

		unop, ok := vs.Values[0].(*ast.UnaryExpr)
		if !ok || unop.Op != token.AND {
			// model variable description should be an address
			continue
		}

		cl, ok := unop.X.(*ast.CompositeLit)
		if !ok {
			continue
		}

		clt, ok := cl.Type.(*ast.Ident)
		if !ok {
			continue
		}

		modelType := clt.Name

		modelMeta, ok := p.models[modelType]
		if !ok {
			fmt.Printf("Struct '%v' not defined\n", modelType)
			continue
		}
		ctx.Model.Meta = modelMeta

		for _, elt := range cl.Elts {
			ce, ok := elt.(*ast.CallExpr)
			if !ok {
				continue
			}

			se, ok := ce.Fun.(*ast.SelectorExpr)
			if !ok {
				continue
			}

			if se.X.(*ast.Ident).Name == ctx.Model.Package.Name && se.Sel.Name == newModelFuncName {
				return p.inspectModelExpr(ctx, ce.Args)
			}
		}
	}
	return true
}

// inspectModelExpr parse orm.Model creation
func (p *modelParser) inspectModelExpr(ctx *Context, args []ast.Expr) bool {
	for _, arg := range args {
		ce, ok := arg.(*ast.CallExpr)
		if !ok {
			continue
		}

		se, ok := ce.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}

		id, ok := se.X.(*ast.Ident)
		if !ok {
			continue
		}

		if id.Name != ctx.Model.Package.Name {
			continue
		}

		switch se.Sel.Name {
		case "OptTable":
			if err := p.CallOptTable(ctx, ce.Args); err != nil {
				continue
			}
		case "OptFields":
			if err := p.CallOptFields(ctx, ce.Args); err != nil {
				continue
			}
		}
	}
	return true
}

// CallOptTable parses model.OptTable call for model defenition value
func (p *modelParser) CallOptTable(ctx *Context, args []ast.Expr) error {
	if len(args) != 1 {
		return fmt.Errorf("function model.OptTable requires one argument")
	}

	bl, ok := args[0].(*ast.BasicLit)
	if !ok || bl.Kind != token.STRING {
		return fmt.Errorf("function model.OptTable argument should be string type")
	}

	model.OptTable(bl.Value)(ctx.Model.Meta)
	return nil
}

// CallOptFields parses model.OptFields call which defines fields for model defenition value
func (p *modelParser) CallOptFields(ctx *Context, args []ast.Expr) error {
	var fieldMappers []field.Mapper
	for _, arg := range args {
		if m := p.fieldParser.inpsectFields(ctx, arg); m != nil {
			fieldMappers = append(fieldMappers, m)
		}
	}
	model.OptFields(fieldMappers...)(ctx.Model.Meta)
	return nil
}
