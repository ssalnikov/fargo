package generator

import (
	"fmt"
	"github.com/gigovich/fargo/internal/parser"
	"os"
	"path/filepath"
	"strings"
)

// Generator for code based on meta
type Generator struct {
	filePath string
}

// New generator instance constructor
func New(filePath string) *Generator {
	return &Generator{
		filePath: filePath,
	}
}

// Generate model code
func (g *Generator) Generate(ctx *parser.Context) error {
	f, err := os.Create(g.getGenFile())
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString("package main\n\n"); err != nil {
		return err
	}

	for name, def := range ctx.DefList {
		if !def.TypeDefined {
			fmt.Fprintf(f, "type %v struct {\n\tmodel.Mapper\n}\n\n", name)
		}

		for i, field := range def.Model.Fields {
			fmt.Fprintf(
				f,
				"func (m *%v) %s() fiel.Mapper {\n\treturn m.Fields[%v]\n}\n\n",
				name,
				g.formatFieldName(field.GetMeta().Name),
				i,
			)
		}
	}

	return nil
}

func (g *Generator) formatFieldName(name string) (modified string) {
	for _, part := range strings.Split(name, "_") {
		if substitute, ok := reservedFieldNames[strings.ToLower(part)]; ok {
			modified += substitute
		} else {
			modified += strings.Title(part)
		}
	}
	return modified
}

// getGenFile path
func (g *Generator) getGenFile() string {
	fp := filepath.Dir(g.filePath)
	filebase := filepath.Base(g.filePath)
	filebase = strings.TrimSuffix(filebase, ".go")
	return filepath.Join(fp, filebase+"_gen.go")
}

var reservedFieldNames = map[string]string{"id": "ID", "pk": "PK", "url": "URL", "uri": "URI"}
