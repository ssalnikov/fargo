package model

// Mapper returns Model
type Mapper interface {
	// GetMeta data of model
	GetMeta() *Meta
}
