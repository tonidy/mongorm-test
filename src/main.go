package main

import (
	"context"
	"fmt"

	"git.yud1z.my.id/lib/mongorm/pkg/mongorm"
	"github.com/tonidy/mongorm-test/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Replace the connection string with your own
	client, err := mongorm.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	db := client.Database("test_db")

	// Create a new user
	user := model.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	err = user.Create(
		context.Background(),
		db,
		"users",
		&user,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User created: %v\n", user)

	// Read a user by ID
	var readUser model.User
	err = readUser.ReadOne(
		context.Background(),
		db,
		"users",
		bson.M{"_id": user.ID},
		&readUser,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User read: %v\n", readUser)

	// Update a user's email
	update := bson.M{
		"$set": bson.M{
			"email":      "john.doe_updated@example.com",
			"updated_at": primitive.NewDateTimeFromTime(user.UpdatedAt),
		},
	}
	err = user.Update(
		context.Background(),
		db,
		"users",
		bson.M{
			"_id": user.ID,
		},
		update,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User updated: %v\n", user)

	// Delete a user by ID
	// err = user.Delete(context.Background(), db, "users", bson.M{"_id": user.ID})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User deleted")
}
