package parser

import (
	"go/ast"
	"go/token"

	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/model"
)

type fieldParser struct {
	fieldPackageName string
}

func newFieldParser(fieldPackageName string) *fieldParser {
	return &fieldParser{
		fieldPackageName: fieldPackageName,
	}
}

func (f *fieldParser) inpsectFields(modelMeta *model.Meta, ae ast.Expr) field.Mapper {
	ce, ok := ae.(*ast.CallExpr)
	if !ok {
		return nil
	}

	se, ok := ce.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil
	}

	id, ok := se.X.(*ast.Ident)
	if !ok {
		return nil
	}

	if id.Name != f.fieldPackageName {
		return nil
	}

	switch se.Sel.Name {
	case "Int":
		return f.getIntField(modelMeta, ce.Args)
	}
	return nil
}

func (f *fieldParser) getIntField(modelMeta *model.Meta, args []ast.Expr) field.Mapper {
	if len(args) < 1 {
		return nil
	}

	bl, ok := args[0].(*ast.BasicLit)
	if !ok && bl.Kind != token.STRING {
		return nil
	}

	var options []field.Option
	for _, arg := range args[1:] {
		ce, ok := arg.(*ast.CallExpr)
		if !ok {
			return nil
		}

		se, ok := ce.Fun.(*ast.SelectorExpr)
		if !ok {
			return nil
		}

		id, ok := se.X.(*ast.Ident)
		if !ok && id.Name != f.fieldPackageName {
			return nil
		}

		switch se.Sel.Name {
		case "OptPrimary":
			options = append(options, field.OptPrimary())
		case "OptTags":
			options = append(options, f.getOptTags(modelMeta, ce.Args))
		}
	}

	return field.Int(bl.Value, options...)
}

func (f *fieldParser) getOptTags(modelMeta *model.Meta, args []ast.Expr) field.Option {
	if len(args) != 1 {
		return nil
	}

	bl, ok := args[0].(*ast.BasicLit)
	if !ok || bl.Kind != token.STRING {
		return nil
	}

	return field.OptTags(bl.Value)
}

func (f *fieldParser) getOptReferenceModel(modelMeta *model.Meta, args []ast.Expr) field.Option {
	if len(args) != 1 {
		return nil
	}

	bl, ok := args[0].(*ast.BasicLit)
	if !ok || bl.Kind != token.STRING {
		return nil
	}

	// TODO: continue here
}
