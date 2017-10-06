package field

// PrimaryKeyGetter interface for models
type PrimaryKeyGetter interface {
	GetPrimaryKey() Mapper
}

// OptionPrimary set this field primary key flag enabled
func OptionPrimary() Option {
	return func(b *Meta) {
		b.Primary = true
	}
}

// OptionTags set this field primary key flag enabled
func OptionTags(tags string) Option {
	return func(b *Meta) {
		b.Tags = tags
	}
}

// OptionReferenceModel set this field referenced to other model by their primary key
func OptionReferenceModel(model PrimaryKeyGetter) Option {
	return func(b *Meta) {
		b.Reference = model.GetPrimaryKey()
	}
}

// OptionReferenceField set this field referenced to any other field
func OptionReferenceField(field Mapper) Option {
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
