package mgorm

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateOne database entry with filter
type UpdateOne struct {
	Db
	Data   interface{}
	Filter bson.M
	Count  int64
	Opts   *options.UpdateOptions
}

// Exec update one entry
func (u *UpdateOne) Exec(mgo *mongo.Client) error {

	collection := mgo.Database(u.Name).Collection(u.Collection)
	atualizacao := bson.D{{Key: "$set", Value: u.Data}}
	updatedResult, err := collection.UpdateOne(context.TODO(), u.Filter, atualizacao, u.Opts)
	if err != nil {
		return err
	}

	u.Count = updatedResult.ModifiedCount

	return nil

}
