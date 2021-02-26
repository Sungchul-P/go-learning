package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/Sungchul-P/go-learning/microserviceWithGo/1_Microservice/rpc/contract"
)

const port = 1234

func CreateClient() *rpc.Client {
	// Dial() 을 사용해 클라이언트 자체를 생성
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return client
}

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse

	// 클라이언트의 Call() 메서드를 사용해 서버의 이름 붙여진 함수를 호출한다.
	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
