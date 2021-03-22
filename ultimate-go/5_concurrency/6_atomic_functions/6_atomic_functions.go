// ----------------
// Atomic Functions
// ----------------

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// counter 는 모든 고루틴에 의해 증가되는 변수이다.
// 해당 변수가 int 가 아닌 int64 타입이라는 것에 유의해야한다.
// 아토믹 함수의 경우 정확성을 요구하기 때문에 구체적인 타입을 명시해야한다.
var counter int64

func main() {
	// 고루틴의 수 지정
	const grs = 2

	// 동시성을 관리하기 위한 wg 생성
	var wg sync.WaitGroup
	wg.Add(grs)

	// 2개의 고루틴들을 생성한다.
	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				// counter 에 안전하게 1을 더해준다.
				// 아토믹 연산 함수는 첫 번째 매개변수로 동기화 보장을 원하는 대상의 주소를 받는다.
				// 같은 주소에 대해 이러한 함수들을 사용하면 직렬화 된다.
				atomic.AddInt64(&counter, 1)

				// 이 호출은 AddInt64 함수 호출이 완료됐을 때 counter 가 이미 증가했으므로 큰 의미가 없다.
				runtime.Gosched()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
