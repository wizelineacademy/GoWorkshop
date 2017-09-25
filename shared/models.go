package shared

import (
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

// CreateUser func
func CreateUser(email string) (string, error) {
	c := DbCollection("users")
	repo := &UserRepository{c}
	return repo.Create(&User{
		Email: email,
	})
}

// CreateItem func
func CreateItem(message string, userID string) (string, error) {
	c := DbCollection("list")
	repo := &ListRepository{c}

	return repo.Create(&Item{
		Message: message,
		UserId:  userID,
	})
}

// DeleteItem func
func DeleteItem(itemID string) error {
	c := DbCollection("list")
	repo := &ListRepository{c}
	return repo.Delete(itemID)
}

// GetUserItems func
func GetUserItems(userID string) []Item {
	c := DbCollection("list")
	repo := &ListRepository{c}
	return repo.GetAll(userID)
}

// Create user
func (r *UserRepository) Create(user *User) (string, error) {
	objID := bson.NewObjectId()
	user.Id = objID
	err := r.C.Insert(&user)
	return user.Id.Hex(), err
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
