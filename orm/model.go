package orm

import (
	"github.com/gigovich/fargo/orm/field"
)

// Model base struct
type Model struct {
	Table  string
	Fields Fields
}

// GetPrimaryKey returns this model primary key or nil if primary key not defined
func (m *Model) GetPrimaryKey() field.Mapper {
	for _, f := range m.Fields {
		if f.GetMeta().Primary {
			return f
		}
	}

	return nil
}
