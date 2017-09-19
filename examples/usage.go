package main

import (
	"fmt"
	"github.com/gigovich/fargo/orm"
	"github.com/gigovich/fargo/orm/field"
	"github.com/gigovich/fargo/orm/q"
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
		&field.Int(field.Meta{Name: "id", Primary: true}),
		&field.Char(field.Meta{Name: "name"}),
	},
}

func main() {
	user, err := User.Query(
		q.And(
			q.Eq(User.ID(), "32"),
			q.Eq(User.Name(), "asdfasdf"),
			q.Or(
				q.Like(User.Email(), "blah@blah.com"),
				q.Like(User.Status(), User.StatusNo),
			),
		),
	).First()
	if err != nil {
		panic(err)
	}
}
