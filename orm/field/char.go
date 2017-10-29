package field

// CharField database type
type CharField struct {
	Meta
}

// Char field constructor
func Char(name string, options ...Option) Mapper {
	f := &CharField{}
	f.Meta.Name = name
	Configure(&f.Meta, options...)
	return f
}

// GetMeta data
func (i *CharField) GetMeta() Meta {
	return i.Meta
}
