package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Scalar value bind
func Scalar(v interface{}) Modifier {
	return &Decorate{
		func(q *query.Query) *query.Query {
			return q
		},
	}
}
