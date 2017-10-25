package mod

import (
	"github.com/gigovich/fargo/orm/model"
	"github.com/gigovich/fargo/orm/query"
)

// LeftJoin modificator
func LeftJoin(mapper model.Mapper, mods ...Modifier) Modifier {
	return func(q *query.Query) *query.Query {
		return q
	}
}
