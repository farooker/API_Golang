package _context
import (
	"context"
	"fmt"
	"log"
	"strconv"

	// "go.mongodb.org/mongo-driver/bson"
)

func (connect Context)RemoveOne(conected string, database string, collection string , filter interface{})(bool, string) {
	sucess, msg, client := connect.Connection(conected)
	fmt.Printf("\n")
	fmt.Printf("connect: %s\n", strconv.FormatBool(sucess))
	fmt.Printf("massage: %s\n", msg)
	if !sucess {
		return false,"connect database is Fiald"
	}
	schema := client.Database(database).Collection(collection)
	result, err := schema.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return false,err.Error()
		
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return true,"DeleteOne removed is sucesss"
}