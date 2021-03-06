package controller

import (
	service "Admin-Tools-Api/Service"
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func FindAllCustomize(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	// req.Header.Set("Content-Type", "application/json")
    var  conected = "mongodb://localhost:27017/"
	var  database = "user"
	var  collection = "customize"

    IService := service.CustomizeService()
	respone :=IService.FindAllCustomize(conected,database,collection)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(respone)
}

func FindByIdCustomize(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    var  conected = "mongodb://localhost:27017/"
	var  database = "user"
	var  collection = "customize"
	var  Id = "5ee1f93421605bfa1187b448"

	IService := service.CustomizeService()
	respone :=IService.FindByIdCustomize(conected,database,collection,Id)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(respone)

}

// func InsertedCustomize(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
// 	revcive := model.Customize {
// 		Id: primitive.NewObjectID(),
// 		KartIndex : 110,
// 		OutfitIndex :100,
// 	}
// 	var  conected = "mongodb://localhost:27017/"
// 	var  database = "user"
// 	var  collection = "customize"
// 	sucess, message:=mongoserver.InsertOne(conected, database,collection ,revcive)
// 	res.Header().Set("Content-Type", "application/json")
// 	_respone := model.BaseResponse{
// 		Status:  sucess,
// 		Message: message,
// 		Value: nil,
// 	}
// 	json.NewEncoder(res).Encode(_respone)
// }
// func UpdatedCustomize(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
// objectId, err := primitive.ObjectIDFromHex("60d4307cbdc96c05b6b41757")
// if err != nil{
// 		log.Println("Invalid id")
// }
// filter :=bson.D{primitive.E{Key: "_id", Value: objectId}}
// update :=  bson.D{
// 	primitive.E{Key: "$set", Value: bson.D{
// 	primitive.E{Key: "kartIndex", Value:50},
// 	primitive.E{Key: "outfitIndex", Value:70},
// }}}
// var  conected = "mongodb://localhost:27017/"
// var  database = "user"
// var  collection = "customize"
// sucess, message :=mongoserver.UpdateOne(conected, database,collection ,filter,update)
// res.Header().Set("Content-Type", "application/json")
// _respone := model.BaseResponse{
// 	Status:  sucess,
// 	Message: message,
// 	Value: nil,
// }
// json.NewEncoder(res).Encode(_respone)

// }
// func RemovedCustomize(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
// 	objectId, err := primitive.ObjectIDFromHex("5ee1effd53c90fc7add066a9")
// 	if err != nil{
// 		log.Println("Invalid id")
// 	}
// 	var  conected = "mongodb://localhost:27017/"
// 	var  database = "user"
// 	var  collection = "customize"
// 	var  filter = bson.M{"_id": objectId}
// 	sucess, message :=mongoserver.RemoveOne(conected, database,collection ,filter)
// 	res.Header().Set("Content-Type", "application/json")
// 	_respone := model.BaseResponse{
// 		Status:  sucess,
// 		Message: message,
// 		Value: nil,
// 	}
// 	json.NewEncoder(res).Encode(_respone)

// }


