package generator

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"strings"

	"github.com/gigovich/fargo/internal/parser"
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

	buff := bytes.Buffer{}
	if err := moduleTemplate.Execute(&buff, ctx); err != nil {
		return err
	}

	// check code validity
	formatted, err := format.Source(buff.Bytes())
	if err == nil {
		// save formated code
		if _, err := f.Write(formatted); err != nil {
			return err
		}
	} else {
		// even broken code we should save to see problem place
		if _, err := f.Write(buff.Bytes()); err != nil {
			return err
		}
	}
	return err
}

// getGenFile path
func (g *Generator) getGenFile() string {
	fp := filepath.Dir(g.filePath)
	filebase := filepath.Base(g.filePath)
	filebase = strings.TrimSuffix(filebase, ".go")
	return filepath.Join(fp, filebase+"_gen.go")
}
