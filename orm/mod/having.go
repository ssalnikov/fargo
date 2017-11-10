package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Having operation
func Having(mods ...Modifier) Modifier {
	return func(q *query.Query) *query.Query {
		return q
	}
}
