package main

import (
	"fmt"
	"time"

	"github.com/eapache/go-resiliency/breaker"
)

func main() {
	// 요청에 세 번 실패하면 차단기 열린 상태
	// 반 개방 상태에서 1회 성공하면 닫힌 상태로 변경
	// 열린 상태에서 5초 후에 반 개방 상태로 변경
	b := breaker.New(3, 1, 5*time.Second)

	for {
		result := b.Run(func() error {
			// Call some service
			time.Sleep(2 * time.Second)
			return fmt.Errorf("Timeout")
		})

		switch result {
		case nil:
			// 성공
		case breaker.ErrBreakerOpen:
			// 회로 차단기가 열려 있기 때문에 코드를 실행하지 않는다.
			fmt.Println("Breaker open")
		default:
			fmt.Println(result)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
