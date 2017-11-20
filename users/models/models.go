package models

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// User type
	User struct {
		Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Email string        `json:"email"`
	}
	// UserRepository type
	UserRepository struct {
		C *mgo.Collection
	}
)

var (
	mongoSession *mgo.Session
)

// CreateUser func
func CreateUser(email string) (string, error) {
	c := getSession().Copy().DB("todo").C("users")

	repo := &UserRepository{c}
	return repo.Create(&User{
		Email: email,
	})
}

// Create user
func (r *UserRepository) Create(user *User) (string, error) {
	objID := bson.NewObjectId()
	user.Id = objID
	err := r.C.Insert(&user)
	return user.Id.Hex(), err
}

func getSession() *mgo.Session {
	if mongoSession == nil {
		var err error
		mongoSession, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{os.Getenv("MONGO_HOST")},
			Username: "",
			Password: "",
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.WithError(err).Fatal("could not connect to mongo")
		}
	}
	return mongoSession
}
