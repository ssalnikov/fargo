package parser

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/gigovich/fargo/internal/util"
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/model"
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

	bl.Value = strings.Trim(bl.Value, "\"")
	var fieldInstance field.Mapper
	switch se.Sel.Name {
	case "Int":
		fieldInstance = field.Int(bl.Value, f.getOptions(ctx, ce.Args[1:])...)
	case "Char":
		fieldInstance = field.Char(bl.Value, f.getOptions(ctx, ce.Args[1:])...)
	}

	fieldMethodName := util.FormatFieldName(bl.Value)
	fieldDef := NewFieldDef(fieldMethodName)
	fieldDef.Field = fieldInstance
	ctx.currentDef.Fields[fieldMethodName] = fieldDef

	return fieldInstance
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
		case "OptReference":
			options = append(options, f.getOptReference(ctx, ce.Args))
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

func (f *fieldParser) getOptReference(ctx *Context, args []ast.Expr) field.Option {
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

	name, ok := se.X.(*ast.Ident)
	if !ok {
		return nil
	}
	modelName := name.Name
	fieldName := se.Sel.Name

	modelDef, ok := ctx.DefList[modelName]
	if !ok {
		modelDef = NewModelDef(modelName)
		ctx.DefList[modelName] = modelDef
	}

	fieldDef, ok := modelDef.Fields[fieldName]
	if !ok {
		fieldDef = NewFieldDef(fieldName)
	}

	return field.OptReference(&model.Field{Model: modelDef.Model, Field: fieldDef.Field})
}
