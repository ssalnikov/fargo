package generator

import (
	"github.com/gigovich/fargo/internal/parser"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestGeneration(t *testing.T) {
	testFile := "../parser/testdata/usage.go"

	data, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Error(err)
		return
	}

	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal(err)
		return
	}

	testFile = filepath.Join(dir, "usage.go")
	if err := ioutil.WriteFile(testFile, data, 0644); err != nil {
		t.Error(err)
		return
	}

	p := parser.New(testFile)
	ctx, err := p.Parse()
	if err != nil {
		t.Error(err)
		return
	}

	gen := New(testFile)
	if err := gen.Generate(ctx); err != nil {
		t.Error(err)
	}

	t.Log(dir)
}
