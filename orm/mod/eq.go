package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Eq is equal '=' expression modificator for two fields
func Eq(m1 Modifier, m2 Modifier) Modifier {
	return &Decorate{
		func(q *query.Query) *query.Query {
			return &query.Query{
				queries: []*query.Query{
					m1.Modify(&query.Query{}),
					&query.Query{Expr: "="},
					m2.Modify(&query.Query{}),
				},
			}
		},
	}
}
