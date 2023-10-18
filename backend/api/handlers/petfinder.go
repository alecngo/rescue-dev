package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rescue-dev/backend/pkg/pfapi"
)

func FindPetsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allows any domain. For prod, set it to trusted domain only
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allowed methods
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")

	// Check if it's just a preflight request
	if r.Method == "OPTIONS" {
		return
	}
	zip := r.URL.Query().Get("zip")
	distance := r.URL.Query().Get("distance")

	// Get the singleton client instance
	client, err := pfapi.GetClient()
	if err != nil {
		http.Error(w, "Failed to initialize Petfinder client", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	myParams := pfapi.NewPetSearchParams()
	myParams.AddParam("location", zip)
	myParams.AddParam("distance", distance)
	myParams.AddParam("sort", "distance")

	myAnimals, err := client.GetAnimals(myParams)
	if err != nil {
		http.Error(w, "Failed to fetch animals from Petfinder", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Convert the Animals struct to JSON and send it as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myAnimals.Animals)
}
