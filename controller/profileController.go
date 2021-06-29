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

func FindProfiles(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// req.Header.Set("Content-Type", "application/json")
	var conected = "mongodb://localhost:27017/"
	var database = "user"
	var collection = "profile"
	var IProfile []*model.Profiles

	sucess, message, Cursor := mongoserver.FindAll(conected, database, collection)
	fmt.Printf("\n")
	fmt.Printf("url : %s\n", strconv.FormatBool(sucess))
	fmt.Printf("database  : %s\n", message)
	fmt.Printf("collection: %s\n", Cursor)
	for Cursor.Next(context.TODO()) {
		var data model.Profiles
		err := Cursor.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		IProfile = append(IProfile, &data)

	}
	Cursor.Close(context.TODO())
	_respone := model.BaseResponse{
		Status:  sucess,
		Message: message,
		Value:   IProfile,
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(_respone)
}
func FindByIdProfiles(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// req.Header.Set("Content-Type", "application/json")
	var conected = "mongodb://localhost:27017/"
	var database = "user"
	var collection = "profile"
	var Id = "5f0c11c74cf27aa3366d7f5e"
	var IProfile []*model.Profiles

	sucess, message, Cursor := mongoserver.FindById(conected, database, collection, Id)
	fmt.Printf("\n")
	fmt.Printf("url : %s\n", strconv.FormatBool(sucess))
	fmt.Printf("databasess: %s\n", message)
	// fmt.Printf("collection: %s\n", Cursor)
	for Cursor.Next(context.TODO()) {
		var data model.Profiles
		err := Cursor.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
		IProfile = append(IProfile, &data)

	}
	Cursor.Close(context.TODO())
	_respone := model.BaseResponse{
		Status:  sucess,
		Message: message,
		Value:   IProfile,
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(_respone)
}

func InsertedProfiles(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	revcive := model.Customize{
		Id:          primitive.NewObjectID(),
		KartIndex:   90,
		OutfitIndex: 90,
	}
	var conected = "mongodb://localhost:27017/"
	var database = "user"
	var collection = "profile"
	mongoserver.InsertOne(conected, database, collection, revcive)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(revcive)
}
func UpdatedProfiles(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	objectId, err := primitive.ObjectIDFromHex("5ee1effd53c90fc7add066a9")
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
	var collection = "profile"
	mongoserver.UpdateOne(conected, database, collection, filter, update)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(filter)

}
func RemovedProfiles(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	objectId, err := primitive.ObjectIDFromHex("5ee1ef1405f138fdeff1d686")
	if err != nil {
		log.Println("Invalid id")
	}
	var conected = "mongodb://localhost:27017/"
	var database = "user"
	var collection = "profile"
	var filter = bson.M{"_id": objectId}
	mongoserver.RemoveOne(conected, database, collection, filter)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(filter)

}
