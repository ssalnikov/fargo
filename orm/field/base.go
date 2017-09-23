package field

// Option setup function for field
type Option func(*Base)

// Base field struct with common properties for every fields
type Base struct {
	// primary key flag
	primary bool

	// alias name for field to table column
	alias string
}

// SetPrimary for field
func SetPrimary() Option {
	return func(b *Base) {
		b.primary = true
	}
}
