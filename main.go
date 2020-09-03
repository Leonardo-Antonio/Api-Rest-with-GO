package main

import (
	"log"
	"net/http"

	"github.com/Leonardo-Antonio/rest/handler"
	"github.com/Leonardo-Antonio/rest/storage"
)

func main() {
	mysql := storage.NewConnection()
	alumn := storage.NewAlumn(mysql.Pool())

	mux := http.NewServeMux()

	handler.RouterAlumn(mux, alumn)

	log.Println("Servidor corriendo en http://localhost:8080/v1/alumns/create")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor -> %+v\n", err)
	}
}
