package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Like expression for where clause
func Like(m1 Modifier, m2 Modifier) Modifier {
	return &Decorate{
		func(q *query.Query) *query.Query {
			return q
		},
	}
}
