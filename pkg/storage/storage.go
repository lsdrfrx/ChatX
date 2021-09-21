package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client     *mongo.Client
	DB         *mongo.Database
	Collection *mongo.Collection

	URI            string
	DatabaseName   string
	CollectionName string
}

type Respond struct {
	ID string `bson:"_id"`
}

func NewDatabase() *Database {
	return &Database{
		URI:            "mongodb://localhost:27017",
		DatabaseName:   "test",
		CollectionName: "test",
	}
}

func (db *Database) Open() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(db.URI))
	if err != nil {
		return err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	db.Client = client
	db.DB = client.Database(db.DatabaseName)
	db.Collection = db.DB.Collection(db.CollectionName)

	return nil
}

func (db *Database) Close() error {
	err := db.Client.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) Create(filter bson.M) error {
	_, err := db.Collection.InsertOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

// func (db *Database) Modify(filter bson.M, newVal bson.D) {
// 	_, err := db.Collection.UpdateOne(context.TODO(), filter,

// 	)
// 	if err != nil {
// 		return err
// 	}
// }

func (db *Database) Delete(filter bson.M) error {
	_, err := db.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) Get(filter bson.M) error {
	var respond Respond

	err := db.Collection.FindOne(context.TODO(), filter).Decode(&respond)
	if err != nil {
		return err
	}

	return nil
}
