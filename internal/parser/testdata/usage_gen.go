package main

import (
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/mod"
	"github.com/gigovich/fargo/orm/model"
	"github.com/gigovich/fargo/orm/query"
)

// ProfileModel members meta mapper
type ProfileModel struct {
	model.Meta
}

// ProfileRecord data object
type ProfileRecord struct {
	// ID field
	ID string `json:"id"`

	// Address field
	Address string `json:"address"`
}

// ID returns field mapper for column 'id'
func (m *ProfileModel) ID() field.Mapper {
	return m.Fields[0]
}

// Address returns field mapper for column 'address'
func (m *ProfileModel) Address() field.Mapper {
	return m.Fields[1]
}

// Query records for 'ProfileModel'
func (m *ProfileModel) Query(mods ...mod.Modifier) *model.Query {
	return &query.Query{}
}

// Find returns first element from executed query
func (m *ProfileModel) Find(query model.Query) ([]ProfileRecord, error) {
	return nil, nil
}

// One returns first element from executed query
func (m *ProfileModel) One(query model.Query) (*ProfileRecord, error) {
	return nil, nil
}

// RoleModel members meta mapper
type RoleModel struct {
	model.Meta
}

// RoleRecord data object
type RoleRecord struct {
	// ID field
	ID string `json:"id"`

	// Permissions field
	Permissions string `json:"permissions"`
}

// ID returns field mapper for column 'id'
func (m *RoleModel) ID() field.Mapper {
	return m.Fields[0]
}

// Permissions returns field mapper for column 'permissions'
func (m *RoleModel) Permissions() field.Mapper {
	return m.Fields[1]
}

// Query records for 'RoleModel'
func (m *RoleModel) Query(mods ...mod.Modifier) *model.Query {
	return &query.Query{}
}

// Find returns first element from executed query
func (m *RoleModel) Find(query model.Query) ([]RoleRecord, error) {
	return nil, nil
}

// One returns first element from executed query
func (m *RoleModel) One(query model.Query) (*RoleRecord, error) {
	return nil, nil
}
