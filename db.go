package mgorm

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Db structure is parent of all struct in mgorm
type Db struct {
	Name       string
	Collection string
}

// Decoder decode on types
type Decoder interface {
	Decoder(*mongo.SingleResult) (Decoder, error)
}

// Cursor retrive data from mongo cursor
type Cursor interface {
	CursorDecoder(*mongo.Cursor) (interface{}, error)
}

// GetClient is getclient
func GetClient() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return client

}
