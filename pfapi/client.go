package pfapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

//DefaultBaseURL contains url for petfinder API
const DefaultBaseURL = "https://api.petfinder.com/v2"

//Client struct is used to hold http.Client
type Client struct {
	*http.Client
}

//url is a private function to determine what url to use
//It will use first the environment variable "PF_BASE_URL" or the constant "DefaultBaseURL"
func url() string {
	url := os.Getenv("PF_BASE_URL")
	if url != "" {
		return url
	}
	return DefaultBaseURL
}

//NewClient accepts client id and secret client id issued by Petfinder
//It returns a struct callled Client that contains a pointer to http.Client
func NewClient(accessToken string, secretAccessToken string) (Client, error) {
	//New attempt
	url := url()
	fmt.Println(url)
	conf := &clientcredentials.Config{
		ClientID:     accessToken,
		ClientSecret: secretAccessToken,
		Scopes:       []string{},
		TokenURL:     url + "/oauth2/token/",
	}

	client := conf.Client(oauth2.NoContext)

	return Client{client}, nil
}

//GetAllTypes function is a method of Client
//It returns a struct of animals types and error
func (c Client) GetAllTypes() ([]AnimalType, error) {
	url := fmt.Sprintf("%s/types", url())
	response, err := c.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	var animalTypes []AnimalType
	var message interface{}
	json.Unmarshal(body, &message)
	messageMap := message.(map[string]interface{})
	typesMap := messageMap["types"].([]interface{})

	err = mapstructure.Decode(typesMap, &animalTypes)
	if err != nil {
		fmt.Println("OH NO: ", err)
	}

	// for _, at := range typesMap {
	// 	var animal AnimalType
	// 	at := at.(map[string]interface{})
	// 	err := mapstructure.Decode(at, &animal)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	animalTypes = append(animalTypes, animal)
	// }

	// messageMap := message.(map[string]interface{})
	// typesMap := messageMap["types"].([]interface{})
	// //fmt.Println(typeField["types"])

	// for _, at := range typesMap {
	// 	a := AnimalType{}
	// 	at := at.(map[string]interface{})
	// 	a.Name = at["name"].(string)
	// 	a.Colors = at["colors"].([]string)
	// 	fmt.Println(a)
	// 	animalTypes = append(animalTypes, a)
	// }

	// var animalTypes []AnimalType
	// json.Unmarshal(messageMap["types"].([]byte), &animalTypes)
	// fmt.Println(animalTypes)

	return animalTypes, nil
}
