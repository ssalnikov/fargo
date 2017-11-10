package mod

import (
	"github.com/gigovich/fargo/orm/model"
	"github.com/gigovich/fargo/orm/query"
)

// SortingOrder SQL type
type SortingOrder string

const (
	// Asc asceding SQL order
	Asc SortingOrder = "ASC"

	// Desc desceding SQL order
	Desc SortingOrder = "DESC"
)

// OrderBy operation
func OrderBy(field model.Field, order SortingOrder) Modifier {
	return func(q *query.Query) *query.Query {
		return q
	}
}
