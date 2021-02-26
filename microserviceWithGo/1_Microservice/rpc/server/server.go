package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/Sungchul-P/go-learning/microserviceWithGo/1_Microservice/rpc/contract"
)

const port = 1234

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}

func StartServer() {
	// 핸들러의 새 인스턴스를 만든 다음, 기본 RPC 서버에 등록한다.
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)

	// func Listen(network, address string) (Listener, error)
	// Listener 인터페이스 구현
	// Accept() : 리스너의 다음 연결을 기다리고 있다가 그것을 리턴한다.
	// Close() : 리스너를 닫는다. 연결을 기다리고 있던 Accept 동작에서 빠져나와 에러를 리턴한다.
	// Addr() : 리스너의 네트워크 주소를 리턴한다.
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
	defer l.Close()

	for { // 무한루프
		conn, _ := l.Accept()  // 연결이 이루어 질 때까지 대기
		go rpc.ServeConn(conn) // 클라이언트가 완료될 때까지 대기(고루틴을 쓰기 때문에 즉시 다음 연결 처리 가능)
	}
}

type HelloWorldHandler struct{}

// 인터페이스를 준수할 필요 없이, 구조체 필드 및 관련 메서드를 가지고 있다면 이것을 RPC 서버에 등록할 수 있다.
func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
