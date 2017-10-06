package field

// Option setup function for field
type Option func(*Meta)

// Meta struct with common field properties
type Meta struct {
	// primary key flag
	Primary bool

	// Name of the field
	Name string

	// tags for this field in record struct
	Tags string

	// reference field (this value as usual should be set to ID field of other table)
	Reference Mapper
}
