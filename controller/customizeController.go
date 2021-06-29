package controller

import (
	model "Admin-Tools-Api/model"
	mongoserver "Admin-Tools-Api/mongoserver"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindCustomize(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// req.Header.Set("Content-Type", "application/json")
	var conected = "mongodb://localhost:27017/"
	var database = "user"
	var collection = "customize"
	var ICustomize []*model.Customize

	sucess, message, Cursor := mongoserver.FindAll(conected, database, collection)
	fmt.Printf("\n")
	fmt.Printf("url : %s\n", strconv.FormatBool(sucess))
	fmt.Printf("database  : %s\n", message)
	fmt.Printf("collection: %s\n", Cursor)
	for Cursor.Next(context.TODO()) {
		var data model.Customize
		err := Cursor.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		ICustomize = append(ICustomize, &data)

	}
	Cursor.Close(context.TODO())
	_respone := model.BaseResponse{
		Status:  sucess,
		Message: message,
		Value:   ICustomize,
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(_respone)
}
func FindByIdCustomize(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// req.Header.Set("Content-Type", "application/json")
	var conected = "mongodb://localhost:27017/"
	var database = "user"
	var collection = "customize"
	var Id = "5ee1effd53c90fc7add066a9"
	var ICustomize []*model.Customize

	sucess, message, Cursor := mongoserver.FindById(conected, database, collection, Id)
	fmt.Printf("\n")
	fmt.Printf("url : %s\n", strconv.FormatBool(sucess))
	fmt.Printf("database  : %s\n", message)
	// fmt.Printf("collection: %s\n", Cursor)
	for Cursor.Next(context.TODO()) {
		var data model.Customize
		err := Cursor.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		ICustomize = append(ICustomize, &data)

	}
	Cursor.Close(context.TODO())
	_respone := model.BaseResponse{
		Status:  sucess,
		Message: message,
		Value:   ICustomize,
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(_respone)
}

func InsertedCustomize(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	revcive := model.Customize{
		Id:          primitive.NewObjectID(),
		KartIndex:   110,
		OutfitIndex: 100,
	}
	var conected = "mongodb://localhost:27017/"
	var database = "user"
	var collection = "customize"
	sucess, message := mongoserver.InsertOne(conected, database, collection, revcive)
	res.Header().Set("Content-Type", "application/json")
	_respone := model.BaseResponse{
		Status:  sucess,
		Message: message,
		Value:   nil,
	}
	json.NewEncoder(res).Encode(_respone)
}
func UpdatedCustomize(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	objectId, err := primitive.ObjectIDFromHex("60d4307cbdc96c05b6b41757")
	if err != nil {
		log.Println("Invalid id")
	}
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{
			primitive.E{Key: "kartIndex", Value: 50},
			primitive.E{Key: "outfitIndex", Value: 70},
		}}}
	var conected = "mongodb://localhost:27017/"
	var database = "user"
	var collection = "customize"
	sucess, message := mongoserver.UpdateOne(conected, database, collection, filter, update)
	res.Header().Set("Content-Type", "application/json")
	_respone := model.BaseResponse{
		Status:  sucess,
		Message: message,
		Value:   nil,
	}
	json.NewEncoder(res).Encode(_respone)

}
func RemovedCustomize(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	objectId, err := primitive.ObjectIDFromHex("5ee1effd53c90fc7add066a9")
	if err != nil {
		log.Println("Invalid id")
	}
	var conected = "mongodb://localhost:27017/"
	var database = "user"
	var collection = "customize"
	var filter = bson.M{"_id": objectId}
	sucess, message := mongoserver.RemoveOne(conected, database, collection, filter)
	res.Header().Set("Content-Type", "application/json")
	_respone := model.BaseResponse{
		Status:  sucess,
		Message: message,
		Value:   nil,
	}
	json.NewEncoder(res).Encode(_respone)

}
