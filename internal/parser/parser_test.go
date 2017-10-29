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

	if len(models) != 3 {
		t.Errorf("parsed models count mismatch: expected 3, got %v", len(models))
		return
	}

	for _, k := range []string{"UserModel", "ProfileModel", "RoleModel"} {
		if _, ok := models[k]; !ok {
			t.Errorf("'%v' model not parsed", k)
			return
		}
	}
}
