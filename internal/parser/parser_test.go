package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	parser := New("testdata/usage.go")
	if err := parser.Parse(); err != nil {
		t.Error(err)
	}
}
