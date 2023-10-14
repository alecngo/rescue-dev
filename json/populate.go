package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/petfinder-com/petfinder-go-sdk/pfapi"
)

var client *pfapi.Client

const (
	prefix = ""
	indent = "    "
)

func Populate(client *pfapi.Client) {
	types, err := client.GetAllTypes()
	if err != nil {
		fmt.Println("Error fetching all types: ", err)
		return
	}

	// For each type, fetch and populate its details (including breeds)
	var detailedTypes []pfapi.AnimalType
	for _, animalType := range types {
		detailedType, err := client.GetType(animalType.Name)
		if err != nil {
			fmt.Printf("Error fetching details for type %s: %v\n", animalType.Name, err)
			continue // Skip to next type in case of error
		}
		detailedTypes = append(detailedTypes, detailedType)
	}

	// Write the detailed types data to a JSON file
	err = writeDataToFile("animalTypes.json", detailedTypes)
	if err != nil {
		fmt.Println("Error writing data to file:", err)
		return
	}

	fmt.Println("Data written successfully to animalTypes.json!")
}

func writeDataToFile(filename string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, prefix, indent) // Indented for better readability
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, jsonData, 0644)
}
