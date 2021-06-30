package context_test

import (
	context "Admin-Tools-Api/Context"
	service "Admin-Tools-Api/Service"
	model "Admin-Tools-Api/model"
	"encoding/json"
	"fmt"
	"testing"
)

func TestConnectMongo(t *testing.T) {

	var url = "mongodb://localhost:27017/"
     _context := context.ContextMongo()
	sucess, _, _ := _context.Connection(url)
	if !sucess  {
		t.Error("connect mongo DB should be status is 'TRUE'", sucess)
	}
}
func TestFindAllMongo(t *testing.T) {

	var  conected = "mongodb://localhost:27017/"
	var  database = "user"
	var  collection = "customize"
    var  transforms  model.BaseResponse

    _service := service.CustomizeService()
	data := _service.FindAllCustomize(conected,database,collection)
	bodyBytes, _ := json.Marshal(data)
	json.Unmarshal(bodyBytes, &transforms)
	fmt.Print(transforms)
	if !transforms.Status  {
		t.Error("Find all should be status is 'TRUE'", transforms.Status)
	}
}

func TestFindByIdMongo(t *testing.T) {

	var  conected = "mongodb://localhost:27017/"
	var  database = "user"
	var  collection = "customize"
	var  Id = "5ee1f93421605bfa1187b448"
    var  transforms  model.BaseResponse

    _service := service.CustomizeService()
	data := _service.FindByIdCustomize(conected,database,collection,Id)
	bodyBytes, _ := json.Marshal(data)
	json.Unmarshal(bodyBytes, &transforms)
	if !transforms.Status  {
		t.Error("Find all should be status is 'TRUE'", transforms.Status)
	}
}