package service

import (
	repository "Admin-Tools-Api/Repository"
)

type IFileAcessService interface {
	ReadFileAccess(path string) interface{}
	WriteFileAccess(path string,data interface{})interface{}
}
type _FileAcessService struct{}



func FileAcessService() IFileAcessService{
	return _FileAcessService{}
}

func(FileAcess _FileAcessService)ReadFileAccess(path string) interface{} {
	_respone := make(chan interface{})
	_storage := repository.FileAcessRepository()
	go 	_storage.ReadFileAccess(path, _respone)
	return<-_respone
}

func(FileAcess _FileAcessService)WriteFileAccess(path string,data interface{})interface{}  {

	_respone := make(chan interface{})
	_storage := repository.FileAcessRepository()
	go 	_storage.WriteFileAccess(path,data,_respone)
	return<-_respone
}
