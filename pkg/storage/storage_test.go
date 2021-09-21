package storage

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestCreate(t *testing.T) {
	db := NewDatabase()
	err := db.Open()
	if err != nil {
		t.Fatalf("Unable to open database: %s", err.Error())
	}

	err = db.Create(bson.M{"_id": "Peter"})
	if err != nil {
		t.Fatalf("Unable to create new document: %s", err.Error())
	}
}
