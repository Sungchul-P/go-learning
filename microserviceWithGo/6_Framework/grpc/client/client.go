package main

import (
	"fmt"
	"log"

	proto "github.com/Sungchul-P/go-learning/microserviceWithGo/6_Framework/grpc/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// 서버에 연결한 다음 요청을 초기화
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure()) //전송 보안 비활성화
	if err != nil {
		log.Fatal("Unable to create connection to server: ", err)
	}

	client := proto.NewKittensClient(conn) // 클라이언트 생성
	response, err := client.Hello(context.Background(), &proto.Request{Name: "Nic"})

	if err != nil {
		log.Fatal("Error calling service: ", err)
	}

	fmt.Println(response.Msg)
}
