package _context
import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func (connect Context)FindAll(conected string,database string,collection string) (bool, string,*mongo.Cursor){
	sucess,msg,client := connect.Connection(conected)
	if (!sucess) {
		return false,"connect database is Fiald",nil
	}
	schema := client.Database(database).Collection(collection)
	cursor, err := schema.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return false,err.Error(),nil
	}
	return true,msg,cursor
}

func (connect Context)FindById(conected string,database string,collection string,Id string) (bool, string,*mongo.Cursor){
	sucess,msg,client := connect.Connection(conected)
	if (!sucess) {
		return false,"connect database is Fiald",nil
	}
	objectId, err := primitive.ObjectIDFromHex(Id)
	if err != nil{
		return false,"Invalid id",nil
	}

	schema := client.Database(database).Collection(collection)
	cursor, err := schema.Find(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return false,err.Error(),nil
	}
	return true,msg,cursor
}