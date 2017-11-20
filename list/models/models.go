package models

import (
	"log"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// Item type
	Item struct {
		Id      bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserId  string        `bson:"user_id" json:"user_id"`
		Message string        `json:"message"`
	}
	// ListRepository type
	ListRepository struct {
		C *mgo.Collection
	}
)

var (
	mongoSession *mgo.Session
)

// CreateItem func
func CreateItem(message string, userID string) (string, error) {
	c := getSession().Copy().DB("todo").C("list")

	repo := &ListRepository{c}

	return repo.Create(&Item{
		Message: message,
		UserId:  userID,
	})
}

// DeleteItem func
func DeleteItem(itemID string) error {
	c := getSession().Copy().DB("todo").C("list")

	repo := &ListRepository{c}
	return repo.Delete(itemID)
}

// GetUserItems func
func GetUserItems(userID string) []Item {
	c := getSession().Copy().DB("todo").C("list")

	repo := &ListRepository{c}
	return repo.GetAll(userID)
}

// Create item
func (r *ListRepository) Create(item *Item) (string, error) {
	objID := bson.NewObjectId()
	item.Id = objID
	err := r.C.Insert(&item)
	return item.Id.Hex(), err
}

// GetAll items
func (r *ListRepository) GetAll(userID string) (items []Item) {
	iter := r.C.Find(bson.M{"user_id": userID}).Iter()
	result := Item{}
	for iter.Next(&result) {
		items = append(items, result)
	}
	return
}

// Delete item
func (r *ListRepository) Delete(id string) error {
	return r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
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
			log.Fatalf("[getSession]: %v\n", err)
		}
	}
	return mongoSession
}
