package parser

import (
	"fmt"
	"go/ast"
	"go/token"
)

// modelParser definition parser
type modelParser struct {
	typesNames     map[string]modelDef
	ormPackageName string
}

func newModelParser(ormPackageName string) *modelParser {
	return &modelParser{
		ormPackageName: ormPackageName,
		typesNames:     make(map[string]modelDef),
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

			if id.Name == p.ormPackageName {
				p.typesNames[name] = modelDef{
					StructName: name,
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

		if _, ok := p.typesNames[modelType]; !ok {
			fmt.Printf("Struct '%v' not defined\n", modelType)
			continue
		}

		for _, elt := range cl.Elts {
			cl, ok := elt.(*ast.CompositeLit)
			if !ok {
				continue
			}

			se, ok := cl.Type.(*ast.SelectorExpr)
			if !ok {
				continue
			}

			if se.X.(*ast.Ident).Name == p.ormPackageName && se.Sel.Name == "Model" {
				return p.inspectModelExpr(cl.Elts)
			}
		}
	}
	return true
}

// inspectModelExpr parse orm.Model creation
func (p *modelParser) inspectModelExpr(elts []ast.Expr) bool {
	for _, elt := range elts {

	}
	return true
}

// fieldDef definition
type fieldDef struct {
	Name string
	Type string
}

// modelDef definition
type modelDef struct {
	StructName string
	Fields     []fieldDef
}
