package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Having operation
func Having(mods ...Modifier) Modifier {
	return &Decorate{
		func(q *query.Query) *query.Query {
			return q
		},
	}
}
