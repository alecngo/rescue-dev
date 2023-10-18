package main

import (
	"fmt"
	"net/http"

	"github.com/rescue-dev/backend/api"
	"github.com/rescue-dev/backend/pkg/pfapi"
	"github.com/rs/cors"
)

var pfclient *pfapi.Client

// Main() is used as an example for accessing petfinder api
func main() {
	_, err := pfapi.GetClient() // Initialize the client once during startup
	if err != nil {
		fmt.Println("Could not create client:", err)
		return
	}
	router := api.NewRouter()

	handler := cors.Default().Handler(router)
	http.ListenAndServe(":8080", handler)
}
