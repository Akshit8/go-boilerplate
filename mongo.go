// Package contact help manage contacts.
package contact

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultDatabase = "contactstore"

const collectionName = "contact"

// MongoHandler client and database
type MongoHandler struct {
	client   *mongo.Client
	database string
}

// NewHandler establishes connection with mongodb
func NewHandler(address string) *MongoHandler {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cl, _ := mongo.Connect(ctx, options.Client().ApplyURI(address))
	mh := &MongoHandler{
		client:   cl,
		database: defaultDatabase,
	}
	return mh
}

// GetOne populate c with filtered document
func (mh *MongoHandler) GetOne(c *Contact, filter interface{}) error {
	//Will automatically create a collection if not available
	collection := mh.client.Database(mh.database).Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, filter).Decode(c)
	return err
}

// Get returns all documents
func (mh *MongoHandler) Get(filter interface{}) []*Contact {
	collection := mh.client.Database(mh.database).Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	var result []*Contact
	for cur.Next(ctx) {
		contact := &Contact{}
		er := cur.Decode(contact)
		if er != nil {
			log.Fatal(er)
		}
		result = append(result, contact)
	}
	return result
}

// AddOne add new document
func (mh *MongoHandler) AddOne(c *Contact) (*mongo.InsertOneResult, error) {
	collection := mh.client.Database(mh.database).Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, c)
	return result, err
}

// Update the filtered document
func (mh *MongoHandler) Update(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := mh.client.Database(mh.database).Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.UpdateMany(ctx, filter, update)
	return result, err
}

// RemoveOne removes the filtered document
func (mh *MongoHandler) RemoveOne(filter interface{}) (*mongo.DeleteResult, error) {
	collection := mh.client.Database(mh.database).Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.DeleteOne(ctx, filter)
	return result, err
}

