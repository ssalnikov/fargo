package mod

import (
	"github.com/gigovich/fargo/orm/model"
	"github.com/gigovich/fargo/orm/query"
)

// GroupBy operation
func GroupBy(fields ...model.Field) Modifier {
	return &Decorate{
		func(q *query.Query) *query.Query {
			return q
		},
	}
}
