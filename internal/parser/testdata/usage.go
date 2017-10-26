package main

import (
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/mod"
	"github.com/gigovich/fargo/orm/model"
)

type UserModel struct {
	model.Mapper
}

type ProfileModel struct {
	model.Mapper
}

type RoleModel struct {
	model.Mapper
}

// User model
var User = &UserModel{
	model.New(
		model.OptTable("users"),
		model.OptFields(
			field.Int("id", field.OptPrimary(), field.OptTags(`json:"id"`)),
			field.Int("profile_id", field.OptReferenceModel(Profile)),
			field.Int("role_id", field.OptReferenceField(Role.ID())),
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

	query := User.Query(
		mod.LeftJoin(Profile, mod.Eq(Profile.ID(), User.ProfileID())),
		mod.LeftJoin(Role, mod.Eq(Role.ID(), User.RoleID())),
		mod.GroupBy(User.Email()),
		mod.OrderBy(User.Email(), mod.Asc),
		mod.Having(mod.Gt(mod.Count(Profile.Number()), 32)),
		mod.Value(&sums, mod.Sum(Profile.Number())),
		mod.Eq(User.ID(), "32"),
		mod.Eq(User.Name(), "asdfasdf"),
	)

	records, err := User.First(query.Extend(
		mod.Or(
			mod.Like(User.Email(), "blah@blah.com"),
			mod.Like(User.Status(), User.StatusNo),
		),
	)).One()
	if err != nil {
		panic(err)
	}
}
