package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Gt is great '>' expression modificator for two expressions
func Gt(m1 Modifier, m2 Modifier) Modifier {
	return &Decorate{
		func(q *query.Query) *query.Query {
			return q
		},
	}
}
