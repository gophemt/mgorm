package mgorm

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadOne data from database
type ReadOne struct {
	Db
	Filter bson.M
	Result Decoder
	Opts   *options.FindOneOptions
}

// Exec exeuate data from read struct
func (m *ReadOne) Exec(mgo *mongo.Client) error {

	collection := mgo.Database(m.Name).Collection(m.Collection)
	var dr *mongo.SingleResult
	if m.Opts != nil {
		dr = collection.FindOne(context.TODO(), m.Filter, m.Opts)
	} else {
		dr = collection.FindOne(context.TODO(), m.Filter)
	}

	var err error
	m.Result, err = m.Result.Decoder(dr)
	if err != nil {
		return err
	}

	return nil

}
