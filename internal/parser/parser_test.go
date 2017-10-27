package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	parser := New("testdata/usage.go")
	models, err := parser.Parse()
	if err != nil {
		t.Error(err)
	}

	t.Log(models["UserModel"].Fields[0])
}
