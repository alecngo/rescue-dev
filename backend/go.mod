module github.com/rescue-dev/backend

go 1.21.3

replace github.com/petfinder-com/petfinder-go-sdk => ./backend/pkg/pfapi

require (
	github.com/mitchellh/mapstructure v1.1.2
	golang.org/x/oauth2 v0.0.0-20190226205417-e64efc72b421
)

require github.com/rs/cors v1.10.1

require (
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0
	golang.org/x/net v0.0.0-20190311183353-d8887717615a // indirect
	google.golang.org/appengine v1.4.0 // indirect
)
