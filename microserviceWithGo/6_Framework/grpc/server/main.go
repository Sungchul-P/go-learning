package main

import (
	"fmt"
	"log"
	"net"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	proto "github.com/Sungchul-P/go-learning/microserviceWithGo/6_Framework/grpc/proto"
)

type kittenServer struct{} // 핸들러 정의

func (k *kittenServer) Hello(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	response := &proto.Response{} // 응답 객체 생성
	response.Msg = fmt.Sprintf("Hello %v", request.Name)

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000)) // 리스너 생성
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterKittensServer(grpcServer, &kittenServer{}) // 서버 인스턴스 생성
	grpcServer.Serve(lis)
}
