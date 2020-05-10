package mgorm

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Counter set counter and get number from counter
// Sequence name is id of counter that set
type Counter struct {
	DatabaseName string
	Sequence     string
}

type counter struct {
	Sequence string `bson:"_id"`
	Number   uint64
}

// Exec method for counter to get counter number
func (c *Counter) Exec(mgo *mongo.Client) (uint64, error) {

	db := Db{Name: c.DatabaseName, Collection: "counter"}

	read := ReadOne{
		Db:     db,
		Filter: bson.M{"_id": c.Sequence},
		Result: counter{},
	}

	err := read.Exec(mgo)
	if err != nil && err.Error() == "mongo: no documents in result" {

		insert := InsertOne{
			Db: db,
			Data: struct {
				Sequence string `bson:"_id"`
				Number   uint64 `bson:"number"`
			}{
				Sequence: c.Sequence,
				Number:   1,
			},
		}
		err := insert.Exec(mgo)
		if err != nil {

			return 0, err

		}

		return 1, nil

	} else if err != nil {

		return 0, err

	}

	newNumber := read.Result.(counter).Number + 1

	u := UpdateOne{
		Db:     db,
		Data:   bson.M{"number": newNumber},
		Filter: bson.M{"_id": c.Sequence},
	}
	err = u.Exec(mgo)
	if err != nil {

		return 0, err

	}

	return newNumber, nil

}

func (c counter) Decoder(r *mongo.SingleResult) (Decoder, error) {
	err := r.Decode(&c)
	if err != nil {
		return c, err
	}
	return c, nil
}
