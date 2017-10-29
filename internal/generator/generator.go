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
			fmt.Fprintf(f, "type %v struct {\n\tmodel.Mapper\n}\n", name)
		}
	}

	return nil
}

// getGenFile path
func (g *Generator) getGenFile() string {
	fp := filepath.Dir(g.filePath)
	filebase := filepath.Base(g.filePath)
	filebase = strings.TrimSuffix(filebase, ".go")
	return filepath.Join(fp, filebase+"_gen.go")
}
