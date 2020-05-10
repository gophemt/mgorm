package mgorm

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// DropDatabase drop's database
func (d *Db) DropDatabase(mgo *mongo.Client) error {
	err := mgo.Database(d.Name).Drop(context.TODO())
	if err != nil {
		return err
	}
	return nil
}
