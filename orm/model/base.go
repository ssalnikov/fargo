package model

import (
	"github.com/gigovich/fargo/orm/field"
)

// Mapper interface to model properties
type Mapper interface {
	// GetModel base object
	GetModel() *Base
}

// Base model with table common properties and fields definitions
type Base struct {
	// Table name in database
	Table string

	// Fields definitions list
	Fields []field.Mapper
}

// New base model instance constructor
func New(optFuncs ...OptionFunc) Base {
	meta := Base{}
	for _, f := range optFuncs {
		f(&meta)
	}
	return meta
}

// GetPrimaryKey returns this model primary key or nil if primary key not defined
func (m Base) GetPrimaryKey() field.Mapper {
	for _, f := range m.Fields {
		if f.GetField().Primary {
			return f
		}
	}

	return nil
}

// GetModel base object
func (m *Base) GetModel() *Base {
	return m
}
