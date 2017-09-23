package op

import (
	"github.com/gigovich/fargo/orm"
)

// Option configure operation on query
type Option func(*orm.Query)
