package main

import (
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/mod"
	"github.com/gigovich/fargo/orm/model"
	"github.com/gigovich/fargo/orm/query"
)

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

// GetPrimaryKey field of ProfileModel
func (m *ProfileModel) GetPrimaryKey() field.Mapper {
	return m.Fields[0]
}

// Address returns field mapper for column 'address'
func (m *ProfileModel) Address() field.Mapper {
	return m.Fields[1]
}

// Query records for 'ProfileModel'
func (m *ProfileModel) Query(mods ...mod.Modifier) *query.Query {
	return &query.Query{}
}

// Find returns first element from executed query
func (m *ProfileModel) Find(query *query.Query) ([]ProfileRecord, error) {
	return nil, nil
}

// One returns first element from executed query
func (m *ProfileModel) One(query *query.Query) (*ProfileRecord, error) {
	return nil, nil
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

// GetPrimaryKey field of RoleModel
func (m *RoleModel) GetPrimaryKey() field.Mapper {
	return m.Fields[0]
}

// Permissions returns field mapper for column 'permissions'
func (m *RoleModel) Permissions() field.Mapper {
	return m.Fields[1]
}

// Query records for 'RoleModel'
func (m *RoleModel) Query(mods ...mod.Modifier) *query.Query {
	return &query.Query{}
}

// Find returns first element from executed query
func (m *RoleModel) Find(query *query.Query) ([]RoleRecord, error) {
	return nil, nil
}

// One returns first element from executed query
func (m *RoleModel) One(query *query.Query) (*RoleRecord, error) {
	return nil, nil
}

// UserModel embends model meta mapper
type UserModel struct {
	model.Mapper
}

// UserRecord data object
type UserRecord struct {
	// ID field
	ID string `json:"id"`

	// ProfileID field
	ProfileID string `json:"profile_id"`

	// RoleID field
	RoleID string `json:"role_id"`

	// Name field
	Name string `json:"name"`
}

// ID returns field mapper for column 'id'
func (m *UserModel) ID() field.Mapper {
	return m.Fields[0]
}

// GetPrimaryKey field of UserModel
func (m *UserModel) GetPrimaryKey() field.Mapper {
	return m.Fields[0]
}

// ProfileID returns field mapper for column 'profile_id'
func (m *UserModel) ProfileID() field.Mapper {
	return m.Fields[1]
}

// RoleID returns field mapper for column 'role_id'
func (m *UserModel) RoleID() field.Mapper {
	return m.Fields[2]
}

// Name returns field mapper for column 'name'
func (m *UserModel) Name() field.Mapper {
	return m.Fields[3]
}

// Query records for 'UserModel'
func (m *UserModel) Query(mods ...mod.Modifier) *query.Query {
	return &query.Query{}
}

// Find returns first element from executed query
func (m *UserModel) Find(query *query.Query) ([]UserRecord, error) {
	return nil, nil
}

// One returns first element from executed query
func (m *UserModel) One(query *query.Query) (*UserRecord, error) {
	return nil, nil
}
