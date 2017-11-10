package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Count of elements
func Count(m Modifier) Modifier {
	return &Decorate{
		func(q *query.Query) *query.Query {
			return q
		},
	}
}
