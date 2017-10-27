package parser

import (
	"github.com/gigovich/fargo/orm/model"
)

// Context for parsing
type Context struct {
	Model struct {
		Meta    *model.Meta
		Package struct {
			Name string
		}
	}

	Field struct {
		Package struct {
			Name string
		}
	}
}
