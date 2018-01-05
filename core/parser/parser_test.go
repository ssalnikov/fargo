package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	parser := New("testdata/usage.go")
	context, err := parser.Parse()
	if err != nil {
		t.Error(err)
	}

	if len(context.DefList) != 3 {
		t.Errorf("parsed models count mismatch: expected 3, got %v", len(context.DefList))
		return
	}

	for _, k := range []string{"UserModel", "ProfileModel", "RoleModel"} {
		if _, ok := context.DefList[k]; !ok {
			t.Errorf("'%v' model not parsed", k)
			return
		}
	}
}
