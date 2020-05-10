package mgorm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// InsertOne data from database
type InsertOne struct {
	Db
	ID   interface{}
	Data interface{}
}

// Exec insertone
func (i *InsertOne) Exec(mgo *mongo.Client) error {

	collection := mgo.Database(i.Name).Collection(i.Collection)
	insertResult, err := collection.InsertOne(context.TODO(), i.Data)
	if err != nil {
		return err
	}

	i.ID = insertResult.InsertedID
	return nil

}
