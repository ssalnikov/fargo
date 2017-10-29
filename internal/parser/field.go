package parser

import (
	"go/ast"
	"go/token"

	"github.com/gigovich/fargo/orm/field"
)

type fieldParser struct {
}

func newFieldParser() *fieldParser {
	return &fieldParser{}
}

func (f *fieldParser) inpsectFields(ctx *Context, ae ast.Expr) field.Mapper {
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

	if id.Name != ctx.FieldImport {
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
		return field.Int(bl.Value, f.getOptions(ctx, ce.Args[1:])...)
	case "Char":
		return field.Char(bl.Value, f.getOptions(ctx, ce.Args[1:])...)
	}
	return nil
}

func (f *fieldParser) getOptions(ctx *Context, args []ast.Expr) (options []field.Option) {
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
		if !ok && id.Name != ctx.FieldImport {
			return nil
		}

		switch se.Sel.Name {
		case "OptPrimary":
			options = append(options, field.OptPrimary())
		case "OptTags":
			options = append(options, f.getOptTags(ctx, ce.Args))
		case "OptReferenceModel":
			options = append(options, f.getOptReferenceModel(ctx, ce.Args))
		case "OptReferenceField":
			options = append(options, f.getOptReferenceField(ctx, ce.Args))
		}
	}
	return
}

func (f *fieldParser) getOptTags(ctx *Context, args []ast.Expr) field.Option {
	if len(args) != 1 {
		return nil
	}

	bl, ok := args[0].(*ast.BasicLit)
	if !ok || bl.Kind != token.STRING {
		return nil
	}

	return field.OptTags(bl.Value)
}

func (f *fieldParser) getOptReferenceModel(ctx *Context, args []ast.Expr) field.Option {
	if len(args) != 1 {
		return nil
	}

	_, ok := args[0].(*ast.Ident)
	if !ok {
		return nil
	}

	// TODO: remove this stub object
	return field.OptReferenceModel(&ctx.currentDef.Model)
}

func (f *fieldParser) getOptReferenceField(ctx *Context, args []ast.Expr) field.Option {
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
