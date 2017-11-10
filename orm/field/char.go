package field

// CharField database type
type CharField struct {
	Base
}

// Char field constructor
func Char(name string, options ...Option) Mapper {
	f := &CharField{}
	f.Base.Name = name
	Configure(&f.Base, options...)
	return f
}

// GetField data
func (i *CharField) GetField() Base {
	return i.Base
}
