package service

import (
	repository "Admin-Tools-Api/Repository"
)

type ICustomizeService interface {

	FindAllCustomize(conected string,database string,collection string) interface{}
	FindByIdCustomize(conected string,database string,collection string,Id string) interface{}
}
type _CustomizeService struct{}

func CustomizeService() ICustomizeService{
	return _CustomizeService{}
}

func(service _CustomizeService)FindAllCustomize(conected string,database string,collection string) interface{} {

	_respone := make(chan interface{})
	_customize := repository.CustomizeRepository()
	go 	_customize.FindAllCustomize(conected,database,collection,_respone)
	return<-_respone
}

func(service _CustomizeService)FindByIdCustomize(conected string,database string,collection string,Id string) interface{} {

	_respone := make(chan interface{})
	_customize := repository.CustomizeRepository()
	go 	_customize.FindByIdCustomize(conected,database,collection,Id,_respone)
	return<-_respone
}