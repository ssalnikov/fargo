package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Or expression of two elements
func Or(m1 Modifier, m2 Modifier) Modifier {
	return &Decorate{
		func(q *query.Query) *query.Query {
			return q
		},
	}
}
