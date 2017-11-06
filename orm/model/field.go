package model

import "github.com/gigovich/fargo/orm/field"

// Field contains both mappers for model and field
type Field struct {
	// Model mapper
	Model Mapper

	// Field maper for associated model
	Field field.Mapper
}
