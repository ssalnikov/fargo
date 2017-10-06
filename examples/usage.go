package main

import (
	"github.com/gigovich/fargo/orm"
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/mod"
)

type UserModel struct {
	orm.Model
}

type ProfileModel struct {
	orm.Model
}

type RoleModel struct {
	orm.Model
}

// User model
var User = &UserModel{
	orm.Model{
		Table: "users",
		Fields: orm.Fields{
			field.Int("id",
				field.OptionPrimary(),
				field.OptionTags(`json:"id"`)),
			field.Int("profile_id", field.OptionReferenceModel(Profile)),
			field.Int("role_id", field.OptionReferenceModel(Role)),
			field.Char("name"),
		},
	},
}

var Profile = &ProfileModel{
	orm.Model{
		Table: "profiles",
		Fields: orm.Fields{
			field.Int("id",
				field.OptionPrimary()),
			field.Char("address"),
		},
	},
}

var Role = &RoleModel{
	orm.Model{
		Table: "roles",
		Fields: orm.Fields{
			field.Int("id",
				field.OptionPrimary()),
			field.Char("permissions"),
		},
	},
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
