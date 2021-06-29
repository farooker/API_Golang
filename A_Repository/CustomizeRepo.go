package repository
import (
	model "Admin-Tools-Api/model"
)

type CustomizeRepository interface {

	FindAll([]model.Customize)
	FindByID()
	Inserted()
	Updated()
	Removed()
}