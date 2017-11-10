package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Sum aggregation
func Sum(m Modifier) Modifier {
	return &Decorate{
		func(q *query.Query) *query.Query {
			return q
		},
	}
}
