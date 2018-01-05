package main

import (
	"fmt"

	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/mod"
	"github.com/gigovich/fargo/orm/model"
)

// ProfileModel definition struct
type ProfileModel struct {
	model.Base
}

// RoleModel definition struct
type RoleModel struct {
	model.Base
}

// User model
var User = &UserModel{
	model.New(
		model.OptTable("users"),
		model.OptFields(
			field.Int("id", field.OptPrimary(), field.OptTags(`json:"id"`)),
			field.Int("profile_id", field.OptReference(Profile.ID())),
			field.Int("role_id", field.OptReference(Role.ID())),
			field.Char("name"),
		),
	),
}

// Profile model instance
var Profile = &ProfileModel{
	model.New(
		model.OptTable("profiles"),
		model.OptFields(
			field.Int("id", field.OptPrimary()),
			field.Char("address"),
		),
	),
}

// Role model instance
var Role = &RoleModel{
	model.New(
		model.OptTable("roles"),
		model.OptFields(
			field.Int("id", field.OptPrimary()),
			field.Char("permissions"),
		),
	),
}

func main() {
	var sums int

	err := User.Insert(
		UserRecord{Name: "Me"},
		UserRecord{Name: "You"},
		UserRecord{Name: "Someone"},
	)
	if err != nil {
		panic(err)
	}

	query := User.Query(
		mod.LeftJoin(Profile, mod.Eq(Profile.ID(), User.ProfileID())),
		mod.LeftJoin(Role, mod.Eq(Role.ID(), User.RoleID())),
		mod.GroupBy(User.Name()),
		mod.OrderBy(User.Name(), mod.Asc),
		mod.Having(mod.Gt(mod.Count(Profile.Address()), mod.Scalar(32))),
		mod.Value(&sums, mod.Sum(Profile.ID())),
		mod.Eq(User.ID(), mod.Scalar("32")),
		mod.Eq(User.Name(), mod.Scalar("asdfasdf")),
	)

	records, err := User.One(
		mod.Extend(
			query,
			mod.Or(
				mod.Like(User.Name(), mod.Scalar("blah@blah.com")),
				mod.Like(User.ProfileID(), Profile.ID()),
			),
		),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(records)
}
