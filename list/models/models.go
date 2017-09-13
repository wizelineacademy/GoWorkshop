package models

import (
	"gopkg.in/mgo.v2"
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

// Create item
func (r *ListRepository) Create(item *Item) (string, error) {
	objID := bson.NewObjectId()
	item.Id = objID
	err := r.C.Insert(&item)
	return item.Id.Hex(), err
}

// GetAll items
func (r *ListRepository) GetAll(userID string) []Item {
	var items []Item
	iter := r.C.Find(bson.M{"user_id": userID}).Iter()
	result := Item{}
	for iter.Next(&result) {
		items = append(items, result)
	}
	return items
}

// Delete item
func (r *ListRepository) Delete(id string) error {
	return r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
