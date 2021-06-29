package _context

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type ConectextServer interface {
    contextMongo() (bool, string,*mongo.Client)
}

type ContextURL struct {
	URL string
}

func (connect ContextURL) contextMongo() (bool, string, *mongo.Client) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connect.URL))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return false, "connection error", nil
	}
	return true, "connection Sucess", client
}
