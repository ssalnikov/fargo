package field

// Getter field meta object
type Getter interface {
	GetField() Mapper
}

// OptPrimary set this field primary key flag enabled
func OptPrimary() Option {
	return func(b *Base) {
		b.Primary = true
	}
}

// OptTags set this field primary key flag enabled
func OptTags(tags string) Option {
	return func(b *Base) {
		b.Tags = tags
	}
}

// OptReference set this field referenced to any other field
func OptReference(field Getter) Option {
	return func(b *Base) {
		b.Reference = field.GetField()
	}
}

// Configure meta object by sequental call option cofigure functions
func Configure(meta *Base, options ...Option) {
	for _, o := range options {
		o(meta)
	}
}
