package field

// PrimaryKeyGetter interface for models
type PrimaryKeyGetter interface {
	GetPrimaryKey() Mapper
}

// OptPrimary set this field primary key flag enabled
func OptPrimary() Option {
	return func(b *Meta) {
		b.Primary = true
	}
}

// OptTags set this field primary key flag enabled
func OptTags(tags string) Option {
	return func(b *Meta) {
		b.Tags = tags
	}
}

// OptReferenceModel set this field referenced to other model by their primary key
func OptReferenceModel(model PrimaryKeyGetter) Option {
	return func(b *Meta) {
		b.Reference = model.GetPrimaryKey()
	}
}

// OptReferenceField set this field referenced to any other field
func OptReferenceField(field Mapper) Option {
	return func(b *Meta) {
		b.Reference = field
	}
}

// Configure meta object by sequental call option cofigure functions
func Configure(meta *Meta, options ...Option) {
	for _, o := range options {
		o(meta)
	}
}
