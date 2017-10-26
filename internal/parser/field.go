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

	if len(ce.Args) < 1 {
		return nil
	}

	bl, ok := ce.Args[0].(*ast.BasicLit)
	if !ok && bl.Kind != token.STRING {
		return nil
	}

	switch se.Sel.Name {
	case "Int":
		return field.Int(bl.Value, f.getOptions(modelMeta, ce.Args[1:])...)
	case "Char":
		return field.Char(bl.Value, f.getOptions(modelMeta, ce.Args[1:])...)
	}
	return nil
}

func (f *fieldParser) getOptions(modelMeta *model.Meta, args []ast.Expr) (options []field.Option) {
	for _, arg := range args {
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
		case "OptReferenceModel":
			options = append(options, f.getOptReferenceModel(modelMeta, ce.Args))
		case "OptReferenceField":
			options = append(options, f.getOptReferenceField(modelMeta, ce.Args))
		}
	}
	return
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

	_, ok := args[0].(*ast.Ident)
	if !ok {
		return nil
	}

	// TODO: remove this stub object
	return field.OptReferenceModel(modelMeta)
}

func (f *fieldParser) getOptReferenceField(modelMeta *model.Meta, args []ast.Expr) field.Option {
	if len(args) != 1 {
		return nil
	}

	ce, ok := args[0].(*ast.CallExpr)
	if !ok {
		return nil
	}

	se, ok := ce.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil
	}

	_, ok = se.X.(*ast.Ident)
	if !ok {
		return nil
	}

	// TODO: remove this stub object
	return field.OptReferenceField(field.Int("stub"))
}
