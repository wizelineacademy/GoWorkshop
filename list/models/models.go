package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	Item struct {
		Id      bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserId  string        `json:"user_id"`
		Message string        `json:"message"`
	}
	ListRepository struct {
		C *mgo.Collection
	}
)

// Create item
func (r *ListRepository) Create(item *Item) error {
	obj_id := bson.NewObjectId()
	item.Id = obj_id
	err := r.C.Insert(&item)
	return err
}

// GetAll items
func (r *ListRepository) GetAll() []Item {
	var items []Item
	iter := r.C.Find(nil).Iter()
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
