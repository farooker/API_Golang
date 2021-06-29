package mongoserver

import (
	"context"
	//"fmt"
    "log"
	"time"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectServer interface {
    connectDatabase() (bool, string,*mongo.Client)
}

type ConnectString struct {
    URL string
}
func (connect ConnectString) connectDatabase() (bool, string,*mongo.Client) {
    client, err := mongo.NewClient(options.Client().ApplyURI(connect.URL))
	if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
        return false,"connection error" , nil
    }
    return true,"connection Sucess" ,client
}


func Connected(url string) (bool, string,*mongo.Client) {
    client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
        return false,"connection error" , nil
    }
    return true,"connection Sucess" ,client
}
