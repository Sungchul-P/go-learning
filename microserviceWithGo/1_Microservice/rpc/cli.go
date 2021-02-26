package main

import (
	"fmt"

	"github.com/Sungchul-P/go-learning/microserviceWithGo/1_Microservice/rpc/client"
	"github.com/Sungchul-P/go-learning/microserviceWithGo/1_Microservice/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
