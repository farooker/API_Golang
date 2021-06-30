package fileacess_test

import (
	service "Admin-Tools-Api/Service"
	model "Admin-Tools-Api/model"
	"encoding/json"
	"testing"
	"fmt"
	"path/filepath"
	"os"
	"log"
)

func TestReadingFile(t *testing.T) {
	var transforms  model.BaseResponse
	
	var path = "_Static/test.json"

	_service := service.FileAcessService()
	data := _service.ReadFileAccess(path)
	
	bodyBytes, _ := json.Marshal(data)
	json.Unmarshal(bodyBytes,&transforms)
	fmt.Println(transforms)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
    fmt.Println(dir)
	// if !transforms.Status {
	// 	t.Error("Reading file json should be status is 'TRUE'", transforms.Status)
	// }

}
