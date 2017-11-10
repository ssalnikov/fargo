package model

import (
	"github.com/gigovich/fargo/orm/field"
)

// OptionFunc function uses to configure model properties object
type OptionFunc func(*Base)

// OptTable name property configure
func OptTable(name string) OptionFunc {
	return func(instance *Base) {
		instance.Table = name
	}
}

// OptFields list property configure
func OptFields(fieldMappers ...field.Mapper) OptionFunc {
	return func(meta *Base) {
		for _, mapper := range fieldMappers {
			meta.Fields = append(meta.Fields, mapper)
		}
	}
}
