package repository
import (
	model "Admin-Tools-Api/model"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type IFileAcessRepository interface {
	ReadFileAccess(path string,_respone chan interface{})
	WriteFileAccess(path string,data interface{},_respone chan interface{})
}
type _FileAcessRepository struct{}

func FileAcessRepository() IFileAcessRepository{
	return _FileAcessRepository{}
}

func(FileAcess _FileAcessRepository)WriteFileAccess(path string,data interface{},_respone chan interface{})  {

	file, _ := json.MarshalIndent(data, "", "")
	err := ioutil.WriteFile(path, file, 0644)
	if err != nil {
		fmt.Println(err)
		_respone <- model.BaseResponse{
			Status:  false,
			Message: err.Error(),
			Value: nil,
		}
		return;
	}
	var result map[string]interface{}
    json.Unmarshal([]byte(file), &result)
	_respone <- model.BaseResponse{
		Status:  true,
		Message: "Writing file sucess",
		Value: result,
	}
}
func(FileAcess _FileAcessRepository)ReadFileAccess(path string,_respone chan interface{})  {
	data, err:= ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		_respone <- model.BaseResponse{
			Status:  false,
			Message: err.Error(),
			Value: nil,
		}
		return;
	}
	var result map[string]interface{}
    json.Unmarshal([]byte(data), &result)
	_respone <- model.BaseResponse{
		Status:  true,
		Message: "Reading file sucess",
		Value: result,
	}
}