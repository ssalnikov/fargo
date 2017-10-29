package field

// IntField database type
type IntField struct {
	Meta
}

// Int field constructor
func Int(name string, options ...Option) Mapper {
	f := &IntField{}
	f.Meta.Name = name
	Configure(&f.Meta, options...)
	return f
}

// GetMeta data
func (i *IntField) GetMeta() Meta {
	return i.Meta
}
