package mod

import (
	"github.com/gigovich/fargo/orm"
	"github.com/gigovich/fargo/orm/query"
)

// LeftJoin modificator
func LeftJoin(mapper orm.ModelMapper, mods ...Modifier) Modifier {
	return func(q *query.Query) *query.Query {
		return q
	}
}
