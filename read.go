package mgorm

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Read data from database
type Read struct {
	Db
	Get    Cursor
	Filter bson.M
	Result []interface{}
}

// Exec exeuate data from read struct
func (m *Read) Exec(mgo *mongo.Client) error {

	collection := mgo.Database(m.Name).Collection(m.Collection)
	cur, err := collection.Find(context.TODO(), m.Filter)
	if err != nil {
		return err
	}

	m.Result = []interface{}{}
	for cur.Next(context.TODO()) {
		getData, err := m.Get.CursorDecoder(cur)
		if err != nil {
			return err
		}
		m.Result = append(m.Result, getData)
	}

	return nil

}
