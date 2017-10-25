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
	modelPackageName string
	fieldPackageName string
	fieldParser      *fieldParser
	typesNames       map[string]*model.Meta
}

func newModelParser(modelPackageName, fieldPackageName string) *modelParser {
	return &modelParser{
		modelPackageName: modelPackageName,
		fieldPackageName: fieldPackageName,
		fieldParser:      newFieldParser(fieldPackageName),
		typesNames:       make(map[string]*model.Meta),
	}
}

func (p *modelParser) InspectTypeSpecs(specs []ast.Spec) bool {
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

			if id.Name == p.modelPackageName {
				p.typesNames[name] = &model.Meta{
					Table: name,
				}
			}
		}
	}
	return true
}

func (p *modelParser) InspectVarsSpecs(specs []ast.Spec) bool {
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

		modelMeta, ok := p.typesNames[modelType]
		if !ok {
			fmt.Printf("Struct '%v' not defined\n", modelType)
			continue
		}

		for _, elt := range cl.Elts {
			ce, ok := elt.(*ast.CallExpr)
			if !ok {
				continue
			}

			se, ok := ce.Fun.(*ast.SelectorExpr)
			if !ok {
				continue
			}

			if se.X.(*ast.Ident).Name == p.modelPackageName && se.Sel.Name == newModelFuncName {
				return p.inspectModelExpr(modelMeta, ce.Args)
			}
		}
	}
	return true
}

// inspectModelExpr parse orm.Model creation
func (p *modelParser) inspectModelExpr(modelMeta *model.Meta, args []ast.Expr) bool {
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

		if id.Name != p.modelPackageName {
			continue
		}

		switch se.Sel.Name {
		case "OptTable":
			if err := p.CallOptTable(modelMeta, ce.Args); err != nil {
				continue
			}
		case "OptFields":
			if err := p.CallOptFields(modelMeta, ce.Args); err != nil {
				continue
			}
		}
	}
	return true
}

func (p *modelParser) CallOptTable(modelMeta *model.Meta, args []ast.Expr) error {
	if len(args) != 1 {
		return fmt.Errorf("function model.OptTable requires one argument")
	}

	bl, ok := args[0].(*ast.BasicLit)
	if !ok || bl.Kind != token.STRING {
		return fmt.Errorf("function model.OptTable argument should be string type")
	}

	model.OptTable(bl.Value)(modelMeta)
	return nil
}

func (p *modelParser) CallOptFields(modelMeta *model.Meta, args []ast.Expr) error {
	var fieldMappers []field.Mapper
	for _, arg := range args {
		if m := p.fieldParser.inpsectFields(modelMeta, arg); m != nil {
			fieldMappers = append(fieldMappers, m)
		}
	}
	model.OptFields(fieldMappers...)
	return nil
}
