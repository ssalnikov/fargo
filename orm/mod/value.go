package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Value extractor from query
func Value(v interface{}, m Modifier) Modifier {
	return &Decorate{
		func(q *query.Query) *query.Query {
			return q
		},
	}
}
