package generator

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/gigovich/fargo/core/parser"
	"github.com/sergi/go-diff/diffmatchpatch"
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
		if err := checkFormatError(t, tempDir, err); err != nil {
			t.Error(err)
		}
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

// checkFormatError and print generated code part where this error happen
func checkFormatError(t *testing.T, tempDir string, origErr error) error {
	split := strings.Split(origErr.Error(), ":")
	if len(split) < 2 {
		return origErr
	}

	lineNo, err := strconv.Atoi(split[0])
	if err != nil {
		return origErr
	}

	_, err = strconv.Atoi(split[1])
	if err != nil {
		return origErr
	}

	f, err := os.Open(filepath.Join(tempDir, "usage_gen.go"))
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for i := 1; true; i++ {
		l, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if i > lineNo-3 && i < lineNo+3 {
			if i == lineNo {
				t.Logf(">>%4d: %s", i, l)
			} else {
				t.Logf("  %4d: %s", i, l)
			}
		}
	}

	t.Errorf("format error: %v", origErr)
	return nil
}
