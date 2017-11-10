package field

// IntField database type
type IntField struct {
	Base
}

// Int field constructor
func Int(name string, options ...Option) Mapper {
	f := &IntField{}
	f.Base.Name = name
	Configure(&f.Base, options...)
	return f
}

// GetField data
func (i *IntField) GetField() Base {
	return i.Base
}
