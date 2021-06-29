package mongoserver

import (
	"context"
	"fmt"
	"log"
	"strconv"
)


func InsertOne(conected string, database string, collection string , pram interface{})(bool, string) {
	sucess, msg, client := Connected(conected)
	fmt.Printf("\n")
	fmt.Printf("connect: %s\n", strconv.FormatBool(sucess))
	fmt.Printf("massage: %s\n", msg)
	if !sucess {
		return false,"connect database is Fiald"
	}
	schema := client.Database(database).Collection(collection)
	insertResult, err := schema.InsertOne(context.TODO(), pram)
	if err != nil {
		log.Fatal(err)
		return false, err.Error()
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return true, "Inserted a single document is sucess"
}