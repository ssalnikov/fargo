package main

import (
	"fmt"
	"github.com/gigovich/fargo/orm"
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/op"
)

type UserModel struct {
	orm.Model
	orm.Fields
}

// User model
var User = &UserModel{
	orm.Model{
		Table: "users",
	},
	orm.Fields{
		field.Int("id", field.SetPrimary()),
		field.Char("name"),
	},
}

func main() {
	query := User.Query(
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

	records, err := User.First(query.Extend()).One()
	if err != nil {
		panic(err)
	}
}
