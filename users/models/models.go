package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name  string        `json:"name"`
		Email string        `json:"email"`
	}
	UserRepository struct {
		C *mgo.Collection
	}
)

// Create user
func (r *UserRepository) Create(user *User) (string, error) {
	obj_id := bson.NewObjectId()
	user.Id = obj_id
	err := r.C.Insert(&user)
	return user.Id.Hex(), err
}

// GetAll users
func (r *UserRepository) GetAll() []User {
	var users []User
	iter := r.C.Find(nil).Iter()
	result := User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}
