package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Email string        `json:"email"`
	}
	UserRepository struct {
		C *mgo.Collection
	}
)

// Create user
func (r *UserRepository) Create(user *User) (string, error) {
	objID := bson.NewObjectId()
	user.Id = objID
	err := r.C.Insert(&user)
	return user.Id.Hex(), err
}
