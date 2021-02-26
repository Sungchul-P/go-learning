package main

import (
	"fmt"

	"github.com/Sungchul-P/go-learning/microserviceWithGo/1_Microservice/rpc_http/client"
	"github.com/Sungchul-P/go-learning/microserviceWithGo/1_Microservice/rpc_http/server"
)

func main() {
	server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Println(reply.Message)
}
