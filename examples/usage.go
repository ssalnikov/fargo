package main

import (
	"fmt"
	"github.com/gigovich/fargo/orm"
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/op"
)

type UserModel struct {
	orm.Meta
	orm.Fields
}

// User model
var User = &UserModel{
	orm.Meta{
		Table: "users",
	},
	field.Fields{
		field.Int("id", field.Primary()),
		field.Char("name"),
	},
}

func main() {
	user, err := User.One(
		op.Value(&sums, op.Sum(Profile.Number())),
		op.Eq(User.ID(), "32"),
		op.Eq(User.Name(), "asdfasdf"),
		op.Or(
			op.Like(User.Email(), "blah@blah.com"),
			op.Like(User.Status(), User.StatusNo),
		),
		op.GroupBy(User.Email()),
		op.OrderBy(op.Asc, User.Email()),
		op.LeftJoin(op.Eq(Profile.ID(), User.ProfileID())),
		op.LeftJoin(op.Eq(Profile.ID(), User.ProfileID())),
		op.Having(op.Gt(op.Count(Profile.Number()), 32)),
	)
	if err != nil {
		panic(err)
	}
}
