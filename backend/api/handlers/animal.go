package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rescue-dev/backend/pkg/pfapi"
)

func GetAnimalByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allows any domain. For prod, set it to trusted domain only
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allowed methods
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
	// Check if it's just a preflight request
	if r.Method == "OPTIONS" {
		return
	}
	// Extract the ID from the URL using the mux.Vars function
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the singleton client instance
	client, err := pfapi.GetClient()
	if err != nil {
		http.Error(w, "Failed to initialize Petfinder client", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Use the provided GetAnimalById function to fetch the animal details
	animal, err := client.GetAnimalById(id)
	if err != nil {
		http.Error(w, "Failed to retrieve animal details", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Convert the Animal struct to JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animal)
}
