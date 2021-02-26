package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Sungchul-P/go-learning/microserviceWithGo/4_Test/data"
	"github.com/Sungchul-P/go-learning/microserviceWithGo/4_Test/handlers"
)

func main() {
	serverURI := "localhost"
	if os.Getenv("DOCKER_IP") != "" {
		serverURI = os.Getenv("DOCKER_IP")
	}

	store, err := data.NewMongoStore(serverURI)
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.Search{DataStore: store}
	err = http.ListenAndServe(":8323", &handler)
	if err != nil {
		log.Fatal(err)
	}
}
