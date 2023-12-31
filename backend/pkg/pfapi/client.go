package pfapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// DefaultBaseURL contains url for petfinder API
const DefaultBaseURL = "https://api.petfinder.com/v2"

var existingClient *Client
var once sync.Once
var clientErr error

// Client struct is used to hold http.Client
type Client struct {
	*http.Client
}

type Breed struct {
	Name string `json:"name"`
}

// url is a private function to determine what url to use
// It will use first the environment variable "PF_BASE_URL" or the constant "DefaultBaseURL"
func url() string {
	url := os.Getenv("PF_BASE_URL")
	if url != "" {
		return url
	}
	return DefaultBaseURL
}

func (c Client) httpGet(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-sdk", "petfinder-go-sdk (https://github.com/petfinder-com/petfinder-go-sdk)")
	response, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// sendRequest is a private function accepting a path as a variable
// It combines url + path to create the request and sends the request
func (c Client) sendGetRequest(path string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", url(), path)

	body, err := c.httpGet(url)

	return body, err
}

// NewClient accepts client id and secret client id issued by Petfinder
// It returns a struct callled Client that contains a pointer to http.Client
func GetClient() (*Client, error) {
	once.Do(func() {
		// Pull Client ID key and Client Secret Key from environment variables
		clientID := os.Getenv("PF_CLIENT_ID")
		clientSecret := os.Getenv("PF_CLIENT_SECRET")

		// Create pfclient Object
		pfclient, err := newClient(clientID, clientSecret)
		if err != nil {
			fmt.Println("Could not create client")
			clientErr = err
			return
		}

		existingClient = &pfclient
	})
	return existingClient, clientErr
}

func newClient(accessToken string, secretAccessToken string) (Client, error) {
	url := url()

	conf := &clientcredentials.Config{
		ClientID:     accessToken,
		ClientSecret: secretAccessToken,
		Scopes:       []string{},
		TokenURL:     url + "/oauth2/token/",
	}

	client := conf.Client(oauth2.NoContext)

	return Client{client}, nil
}

// GetAllTypes function is a method of Client
// It returns a struct of animals types and error
func (c Client) GetAllTypes() ([]AnimalType, error) {
	body, err := c.sendGetRequest("/types")

	var animalTypes []AnimalType
	var message interface{}
	err = json.Unmarshal(body, &message)
	if err != nil {
		return nil, err
	}
	messageMap := message.(map[string]interface{})
	typesMap := messageMap["types"].([]interface{})

	err = mapstructure.Decode(typesMap, &animalTypes)
	if err != nil {
		return nil, err
	}

	return animalTypes, nil
}

// GetType takes a string of the type name (dog, cat, etc) and returns
// an AnimalType struct and error.
func (c Client) GetType(reqType string) (AnimalType, error) {
	body, err := c.sendGetRequest("/types/" + reqType)

	var animalType AnimalType
	var message interface{}
	err = json.Unmarshal(body, &message)
	if err != nil {
		return AnimalType{}, err
	}
	messageMap := message.(map[string]interface{})
	typesMap := messageMap["type"].(map[string]interface{})

	err = mapstructure.Decode(typesMap, &animalType)
	if err != nil {
		return AnimalType{}, err
	}

	// Fetch the breeds for the given type.
	breeds, err := c.getBreedsForType(reqType)
	if err != nil {
		return AnimalType{}, err
	}

	// Extract breed names and populate the Breeds field.
	for _, breed := range breeds {
		animalType.Breeds = append(animalType.Breeds, breed.Name)
	}

	return animalType, nil
}

func (c Client) getBreedsForType(redType string) ([]Breed, error) {
	body, err := c.sendGetRequest("/types/" + redType + "/breeds")

	// Handle error from the request.
	if err != nil {
		return nil, err
	}

	var breeds []Breed
	var message interface{}
	err = json.Unmarshal(body, &message)
	if err != nil {
		return nil, err
	}

	messageMap := message.(map[string]interface{})
	breedsMap := messageMap["breeds"].([]interface{})

	err = mapstructure.Decode(breedsMap, &breeds)
	if err != nil {
		return nil, err
	}

	return breeds, nil
}

// GetAnimal takes a string of the type id (1234134) and returns
// an Animal struct and error.
func (c Client) GetAnimalById(animalID string) (Animal, error) {
	body, err := c.sendGetRequest("/animals/" + animalID)
	if err != nil {
		return Animal{}, err
	}
	var animal Animal
	var message interface{}

	err = json.Unmarshal(body, &message)
	if err != nil {
		return Animal{}, err
	}
	messageMap := message.(map[string]interface{})
	animalMap := messageMap["animal"].(map[string]interface{})

	err = mapstructure.Decode(animalMap, &animal)
	if err != nil {
		return Animal{}, err
	}

	return animal, nil
}

// GetAnimals takes a key,value pair for query string parameters
// It returns a hash of animals or error
func (c Client) GetAnimals(params SearchParams) (AnimalResponse, error) {
	paramString := params.CreateQueryString()
	url := fmt.Sprintf("/animals%s", paramString)
	body, err := c.sendGetRequest(url)

	var animals AnimalResponse
	var message interface{}

	err = json.Unmarshal(body, &message)
	if err != nil {
		return AnimalResponse{}, err
	}
	messageMap := message.(map[string]interface{})
	err = mapstructure.Decode(messageMap, &animals)
	if err != nil {
		return AnimalResponse{}, err
	}
	return animals, nil
}

// GetOrganizations takes a key,value pair for query string parameters
// It returns a hash of organizations or error
func (c Client) GetOrganizations() (OrganizationResponse, error) {
	//paramString := params.CreateQueryString()
	paramString := ""
	url := fmt.Sprintf("/organizations%s", paramString)
	body, err := c.sendGetRequest(url)
	//fmt.Println(string(body))
	var orgs OrganizationResponse
	var message interface{}

	err = json.Unmarshal(body, &message)
	if err != nil {
		return OrganizationResponse{}, err
	}

	messageMap := message.(map[string]interface{})

	err = mapstructure.Decode(messageMap, &orgs)
	if err != nil {
		return OrganizationResponse{}, err
	}
	return orgs, nil
}

// GetOrganizationsByID takes a string ID
// It returns a hash of organizations or error
func (c Client) GetOrganizationsByID(organizationID string) (Organization, error) {
	body, err := c.sendGetRequest("/organizations/" + organizationID)
	if err != nil {
		return Organization{}, err
	}

	var org Organization
	var message interface{}

	err = json.Unmarshal(body, &message)
	if err != nil {
		return Organization{}, err
	}

	messageMap := message.(map[string]interface{})
	orgMap := messageMap["organization"].(map[string]interface{})

	err = mapstructure.Decode(orgMap, &org)
	if err != nil {
		return Organization{}, err
	}
	return org, nil
}
