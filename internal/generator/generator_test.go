package generator

import (
	"bytes"
	"github.com/gigovich/fargo/internal/parser"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestGeneration(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal(err)
		return
	}

	origDir := "../parser/testdata/"

	usageOrigData, err := ioutil.ReadFile(filepath.Join(origDir, "usage.go"))
	if err != nil {
		t.Error(err)
		return
	}

	genOrigData, err := ioutil.ReadFile(filepath.Join(origDir, "usage_gen.go"))
	if err != nil {
		t.Error(err)
		return
	}

	// and copy there `usage.go` content
	usageTempFile := filepath.Join(tempDir, "usage.go")
	if err := ioutil.WriteFile(usageTempFile, usageOrigData, 0644); err != nil {
		t.Error(err)
		return
	}

	ctx, err := parser.New(usageTempFile).Parse()
	if err != nil {
		t.Error(err)
		return
	}

	if err := New(usageTempFile).Generate(ctx); err != nil {
		t.Error(err)
		return
	}

	genTempData, err := ioutil.ReadFile(filepath.Join(tempDir, "usage_gen.go"))
	if err != nil {
		t.Error(err)
		return
	}

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(string(genOrigData), string(genTempData), false)
	if !bytes.Equal(genOrigData, genTempData) {
		t.Log(dmp.DiffPrettyText(diffs))
		t.Error("oiriginal and generated content mismatch")
		return
	}
}
