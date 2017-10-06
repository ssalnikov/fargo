package mod

import (
	"github.com/gigovich/fargo/orm/query"
)

// Modifier of query
type Modifier func(*query.Query) *query.Query
