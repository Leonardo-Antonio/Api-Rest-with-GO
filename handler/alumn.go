package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Leonardo-Antonio/rest/model"
)

type alumn struct {
	storage Storage
}

func newAlumn(storage Storage) *alumn {
	return &alumn{storage}
}

func (a *alumn) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newRespone(Error, "Ha realizado mal la petición", nil)
		response.responseJSON(w, http.StatusBadRequest)
		return
	}
	data := model.Alumn{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newRespone(Error, "La persona no tiene la estructura correcta", nil)
		response.responseJSON(w, http.StatusInternalServerError)
		return
	}

	err = a.storage.Create(data)
	if err != nil {
		response := newRespone(Error, "Hubo un problema al crear el alumno", nil)
		response.responseJSON(w, http.StatusInternalServerError)
		return
	}
	response := newRespone(Message, "OK", nil)
	response.responseJSON(w, http.StatusOK)
}

func (a *alumn) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newRespone(Error, "Ha realizado mal la petición", nil)
		response.responseJSON(w, http.StatusBadRequest)
		return
	}
	data, err := a.storage.GetAll()
	if err != nil {
		response := newRespone(Error, "Ha ocurrido un error en la bd", nil)
		response.responseJSON(w, http.StatusInternalServerError)
		return
	}
	response := newRespone(Message, "OK", data)
	response.responseJSON(w, http.StatusOK)
}

func (a *alumn) getByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newRespone(Error, "Ha realizado mal la petición", nil)
		response.responseJSON(w, http.StatusBadRequest)
		return
	}
	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newRespone(Error, "Hubo problemas con el parametro en la url", nil)
		response.responseJSON(w, http.StatusInternalServerError)
		return
	}

	alumn, err := a.storage.GetByID(ID)
	if err != nil {
		response := newRespone(Error, "Ha ocurrido un error en la bd", nil)
		response.responseJSON(w, http.StatusInternalServerError)
		return
	}
	response := newRespone(Message, "OK", alumn)
	response.responseJSON(w, http.StatusOK)
}
func (a *alumn) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response := newRespone(Error, "Ha realizado mal la petición", nil)
		response.responseJSON(w, http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newRespone(Error, "Hubo problemas con el parametro en la url", nil)
		response.responseJSON(w, http.StatusInternalServerError)
		return
	}
	err = a.storage.Delete(ID)
	if err != nil {
		response := newRespone(Error, "Hubo problemas al intentar eliminar", nil)
		response.responseJSON(w, http.StatusInternalServerError)
		return
	}
	response := newRespone(Message, "OK", nil)
	response.responseJSON(w, http.StatusOK)
	return

}

func (a *alumn) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		response := newRespone(Error, "Ha realizado mal la petición", nil)
		response.responseJSON(w, http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response := newRespone(Error, "Hubo problemas con el parametro en la url", nil)
		response.responseJSON(w, http.StatusInternalServerError)
		return
	}

	alumn := model.Alumn{}
	err = json.NewDecoder(r.Body).Decode(&alumn)
	if err != nil {
		response := newRespone(Error, "No se puedo hacer la converción a struct", nil)
		response.responseJSON(w, http.StatusInternalServerError)
		return
	}

	err = a.storage.Update(ID, alumn)
	if err != nil {
		response := newRespone(Error, "No se puedo actualizar el alumno", nil)
		response.responseJSON(w, http.StatusInternalServerError)
		return
	}

	response := newRespone(Message, "OK", nil)
	response.responseJSON(w, http.StatusOK)
}
