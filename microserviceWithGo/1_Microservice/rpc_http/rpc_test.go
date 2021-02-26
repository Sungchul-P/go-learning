package main

import (
	"testing"

	"github.com/Sungchul-P/go-learning/microserviceWithGo/1_Microservice/rpc_http/client"
	"github.com/Sungchul-P/go-learning/microserviceWithGo/1_Microservice/rpc_http/server"
)

func BenchmarkDialHTTP(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c := client.CreateClient()
		c.Close()
	}
}

func BenchmarkHelloWorldHandlerHTTP(b *testing.B) {
	b.ResetTimer()

	c := client.CreateClient()

	for i := 0; i < b.N; i++ {
		_ = client.PerformRequest(c)
	}

	c.Close()
}

func init() {
	// start the server
	go server.StartServer()
}
