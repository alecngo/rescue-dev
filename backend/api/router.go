package api

import (
	"github.com/gorilla/mux"
	"github.com/rescue-dev/backend/api/handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/findPets", handlers.FindPetsHandler).Methods("GET")
	r.HandleFunc("/animal/{id}", handlers.GetAnimalByIdHandler).Methods("GET")
	return r
}
