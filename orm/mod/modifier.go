package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Modifier of query
type Modifier interface {
	Modify(*query.Query) *query.Query
}

// Decorate helps decorate any modification function to struct which implements Modifier interface
type Decorate struct {
	DecoratedFunc func(q *query.Query) *query.Query
}

// Modify interface realization which calls decorated modificator
func (c *Decorate) Modify(q *query.Query) *query.Query {
	return c.DecoratedFunc(q)
}
