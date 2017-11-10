package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Extend expression for where clause
func Extend(q *query.Query, modifiers ...Modifier) *query.Query {
	return q
}
