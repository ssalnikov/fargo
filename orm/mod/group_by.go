package mod

import (
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/query"
)

// GroupBy operation
func GroupBy(fields ...field.Mapper) Modifier {
	return func(q *query.Query) *query.Query {
		return q
	}
}
