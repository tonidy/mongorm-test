package model

import "git.yud1z.my.id/lib/mongorm/pkg/mongorm"

type User struct {
	mongorm.Model
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
}
