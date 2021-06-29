package _context
import (
	"context"
	"fmt"
	"log"
	"strconv"

	// "go.mongodb.org/mongo-driver/bson"
)

func (connect Context)UpdateOne(conected string, database string, collection string , filter interface{},update interface {})(bool, string) {
	sucess, msg, client := connect.Connection(conected)
	fmt.Printf("\n")
	fmt.Printf("connect: %s\n", strconv.FormatBool(sucess))
	fmt.Printf("massage: %s\n", msg)
	if !sucess {
		//panic(err)

		return false,"connect database is Fiald"
	}
	schema := client.Database(database).Collection(collection)
	updateResult, err := schema.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return false,err.Error()
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", 
	updateResult.MatchedCount, updateResult.ModifiedCount)
	return true,"Update documents is sucess"
}