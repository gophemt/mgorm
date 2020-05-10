package mgorm

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UpdateOne database entry with filter
type UpdateOne struct {
	Db
	Data   bson.M
	Filter bson.M
	Count  int64
}

// Exec update one entry
func (u *UpdateOne) Exec(mgo *mongo.Client) error {

	collection := mgo.Database(u.Name).Collection(u.Collection)
	atualizacao := bson.D{{Key: "$set", Value: u.Data}}
	updatedResult, err := collection.UpdateOne(context.TODO(), u.Filter, atualizacao)
	if err != nil {
		return err
	}

	u.Count = updatedResult.ModifiedCount

	return nil

}
