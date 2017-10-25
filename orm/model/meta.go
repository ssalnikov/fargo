package model

import (
	"github.com/gigovich/fargo/orm/field"
)

// Meta
type Meta struct {
	Table  string
	Fields []field.Mapper
}

// GetPrimaryKey returns this model primary key or nil if primary key not defined
func (m *Meta) GetPrimaryKey() field.Mapper {
	for _, f := range m.Fields {
		if f.GetMeta().Primary {
			return f
		}
	}

	return nil
}
