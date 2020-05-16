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
	Type   string
}

// Exec update one entry
func (u *UpdateOne) Exec(mgo *mongo.Client) error {

	collection := mgo.Database(u.Name).Collection(u.Collection)
	data := bson.D{}
	switch u.Type {
	case "set":
		data = bson.D{{Key: "$set", Value: u.Data}}
	case "inc":
		data = bson.D{{Key: "$inc", Value: u.Data}}
	default:
		data = bson.D{{Key: "$set", Value: u.Data}}
	}

	updatedResult, err := collection.UpdateOne(context.TODO(), u.Filter, data, u.Opts)
	if err != nil {
		return err
	}

	u.Count = updatedResult.ModifiedCount

	return nil

}
