package model

import "github.com/gigovich/fargo/orm/field"

// Field contains both mappers for model and field
type Field struct {
	// Model mapper
	Model Mapper

	// field maper for associated model
	Field field.Mapper
}

// GetModel from field and model link object
func (f Field) GetModel() Mapper {
	return f.Model
}

// GetField from model and field link object
func (f Field) GetField() field.Mapper {
	return f.Field
}
