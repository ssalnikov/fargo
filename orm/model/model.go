package model

import (
	"github.com/gigovich/fargo/orm/field"
)

// OptionFunc function uses to configure model properties object
type OptionFunc func(*Meta)

// New model instance
func New(optFuncs ...OptionFunc) *Meta {
	meta := &Meta{}
	for _, f := range optFuncs {
		f(meta)
	}
	return meta
}

// OptTable name property configure
func OptTable(name string) OptionFunc {
	return func(instance *Meta) {
		instance.Table = name
	}
}

// OptFields list property configure
func OptFields(fieldMappers ...field.Mapper) OptionFunc {
	return func(meta *Meta) {
		for _, mapper := range fieldMappers {
			meta.Fields = append(meta.Fields, mapper)
		}
	}
}
