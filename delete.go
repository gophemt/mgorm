package mgorm

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Remove data from mongo database
type Remove struct {
	Db
	Filter bson.M
	Count  int64
}

// Exec remove type
func (r *Remove) Exec(mgo *mongo.Client) error {

	collection := mgo.Database(r.Name).Collection(r.Collection)
	deleteResult, err := collection.DeleteOne(context.TODO(), r.Filter)
	if err != nil {
		return err
	}

	r.Count = deleteResult.DeletedCount

	return nil

}
