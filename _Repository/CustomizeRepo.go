package repository

import (
	_context "Admin-Tools-Api/Context"
	model "Admin-Tools-Api/model"
	"context"
	"log"
)

type ICustomizeRepository interface {
	FindAllCustomize(conected string,database string,collection string,_respone chan interface{})
    FindByIdCustomize(conected string,database string,collection string,Id string,_respone chan interface{})

}
type _CustomizeRepository struct{}

func CustomizeRepository() ICustomizeRepository{
	return _CustomizeRepository{}
}


func(repo _CustomizeRepository)FindAllCustomize(conected string,database string,collection string,_respone chan interface{}) {
	var  ICustomize []* model.Customize
	Icontext :=_context.ContextMongo()
    sucess, message,Cursor:= Icontext.FindAll(conected,database,collection)
	for Cursor.Next(context.TODO()) {
		var data  model.Customize
		err := Cursor.Decode(&data)
		if err != nil { 
			log.Fatal(err) 
		}
		ICustomize =append(ICustomize,&data)
	}
	Cursor.Close(context.TODO())
	_respone <- model.BaseResponse{
		Status:  sucess,
		Message: message,
		Value: ICustomize,
	}
}
func(repo _CustomizeRepository) FindByIdCustomize(conected string,database string,collection string,Id string,_respone chan interface{}) {
	var  ICustomize []* model.Customize
	Icontext :=_context.ContextMongo()
    sucess, message,Cursor	:=Icontext.FindById(conected,database,collection,Id)
	for Cursor.Next(context.TODO()) {
	    var data  model.Customize
		err := Cursor.Decode(&data)
		if err != nil { 
			log.Fatal(err) 
		}
		ICustomize =append(ICustomize,&data)
	}
	Cursor.Close(context.TODO())
	_respone <- model.BaseResponse{
		Status:  sucess,
		Message: message,
		Value: ICustomize,
	}
}