package _context

import "go.mongodb.org/mongo-driver/mongo"

type IConectext interface {

	Connection(url string) (bool, string, *mongo.Client)

	FindAll(conected string,database string,collection string) (bool, string,*mongo.Cursor)

	FindById(conected string,database string,collection string,Id string) (bool, string,*mongo.Cursor)

	InsertOne(conected string, database string, collection string , pram interface{})(bool, string)

	UpdateOne(conected string, database string, collection string , filter interface{},update interface {})(bool, string)

	RemoveOne(conected string, database string, collection string , filter interface{})(bool, string)
}
type Context struct{}

func ContextMongo() IConectext {
	return Context{}
}