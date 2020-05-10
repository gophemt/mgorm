package mgorm

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ReadOne data from database
type ReadOne struct {
	Db
	Filter bson.M
	Result Decoder
}

// Exec exeuate data from read struct
func (m *ReadOne) Exec(mgo *mongo.Client) error {

	collection := mgo.Database(m.Name).Collection(m.Collection)
	documentReturned := collection.FindOne(context.TODO(), m.Filter)

	var err error
	m.Result, err = m.Result.Decoder(documentReturned)
	if err != nil {
		return err
	}

	return nil

}
