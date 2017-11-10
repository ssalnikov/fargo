package main

import (
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/mod"
	"github.com/gigovich/fargo/orm/model"
)

type ProfileModel struct {
	model.Base
}

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

var Profile = &ProfileModel{
	model.New(
		model.OptTable("profiles"),
		model.OptFields(
			field.Int("id", field.OptPrimary()),
			field.Char("address"),
		),
	),
}

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
		mod.LeftJoin(Profile, mod.Eq(mod.Field(Profile.ID()), mod.Field(User.ProfileID()))),
		mod.LeftJoin(Role, mod.Eq(Role.ID(), User.RoleID())),
		mod.GroupBy(User.Name()),
		mod.OrderBy(User.Name(), mod.Asc),
		mod.Having(mod.Gt(mod.Count(Profile.Number()), mod.Scalar(32))),
		mod.Value(&sums, mod.Sum(Profile.Number())),
		mod.Eq(User.ID(), mod.Scalar("32")),
		mod.Eq(User.Name(), mod.Scalar("asdfasdf")),
	)

	records, err := User.One(query.Extend(
		mod.Or(
			mod.Like(User.Email(), "blah@blah.com"),
			mod.Like(User.Status(), User.StatusNo),
		),
	))
	if err != nil {
		panic(err)
	}
}
