package main

import (
	"fmt"

	"github.com/petfinder-com/petfinder-go-sdk/json"
	"github.com/petfinder-com/petfinder-go-sdk/pfapi"
)

// Main() is used as an example for accessing petfinder api
func main() {

	//Pull Client ID key and Client Secret Key from environment variables
	//Create pfclient Object
	pfclient, err := pfapi.GetClient()
	if err != nil {
		fmt.Println("Could not create client")
	}

	//Create variable based on AnimalType struct
	var types []pfapi.AnimalType

	//Retreive all animal types, put into struct
	types, _ = pfclient.GetAllTypes()

	//Iterate through animal types using struct data
	for _, t := range types {
		fmt.Println("Name: ", t.Name)
		fmt.Println("Colors: ", t.Colors)
		fmt.Println("Self Link: ", t.Links.Self.Href)
	}

	//Get a specific type
	myType, err := pfclient.GetType("dog")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myType.Name)
	fmt.Println(myType.Coats)
	fmt.Println(myType.Genders)
	fmt.Println(myType.Colors)
	fmt.Println(myType.Breeds)

	json.Populate(pfclient)
}
