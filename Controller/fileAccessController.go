package controller

import (
	service "Admin-Tools-Api/Service"
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)
 
type Salary struct {

    Basic, HRA, TA float64
}
 
type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func WriteFileAccess(res http.ResponseWriter, req *http.Request, _ httprouter.Params)  {
	data := Employee{
        FirstName: "Mark",
        LastName:  "Jones",
        Email:     "mark@gmail.com",
        Age:       25,
        MonthlySalary: []Salary{
            Salary{
                Basic: 10.00,
                HRA:   5000.00,
                TA:    2000.00,
            },
            Salary{
                Basic: 16000.00,
                HRA:   5000.00,
                TA:    2100.00,
            },
            Salary{
                Basic: 170,
                HRA:   500,
                TA:    2200.00,
            },
        },
    }
	var path = "_Static/test.json"

	_service := service.FileAcessService()
	respone := _service.WriteFileAccess(path ,data)

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(respone)
}


func ReadFileAccess(res http.ResponseWriter, req *http.Request, _ httprouter.Params)  {
	var path = "_Static/test.json"

	_service := service.FileAcessService()
	respone := _service.ReadFileAccess(path)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(respone)
}