# Petfinder.com API Client sdk for Golang

Build status: [![CircleCI](https://circleci.com/gh/petfinder-com/petfinder-go-sdk/tree/master.svg?style=svg)](https://circleci.com/gh/petfinder-com/petfinder-go-sdk/tree/master)

[Uses Petfinder API v2.](https://www.petfinder.com/developers/v2/docs/)

Please see example usage in cmd/main.go

Example client usage:
```go
//Pull Client ID key and Client Secret Key from environment variables
clientID := os.Getenv("PF_CLIENT_ID")
clientSecret := os.Getenv("PF_CLIENT_SECRET")

//Create pfclient Object
pfclient, err := pfapi.NewClient(clientID, clientSecret)
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
```
