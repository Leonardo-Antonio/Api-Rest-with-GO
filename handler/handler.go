package handler

import "github.com/Leonardo-Antonio/rest/model"

// Storage interface what implemet class
type Storage interface {
	Create(alumn model.Alumn) error
	GetAll() ([]model.Alumn, error)
	GetByID(id int) (model.Alumn, error)
	Delete(id int) error
	Update(int, model.Alumn) error
}
