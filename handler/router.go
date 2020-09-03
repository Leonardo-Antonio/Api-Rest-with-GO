package handler

import (
	"net/http"
)

// RouterAlumn .
func RouterAlumn(mux *http.ServeMux, storage Storage) {
	hlr := newAlumn(storage)
	mux.HandleFunc("/v1/alumns/create", hlr.create)
	mux.HandleFunc("/v1/alumns/get-all", hlr.getAll)
	mux.HandleFunc("/v1/alumns/search", hlr.getByID)
	mux.HandleFunc("/v1/alumns/delete", hlr.delete)
	mux.HandleFunc("/v1/alumns/update", hlr.update)
}
