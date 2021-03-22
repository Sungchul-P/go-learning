// -------
// Mutexes
// -------

// 일반적으로 데이터 공유를 하기 위해 매번 4-8바이트를 할당할 만큼 메모리가 여유롭지 않다. 이럴 때 뮤텍스를 사용하면 좋다.
// 뮤텍스를 사용하면 모든 고루틴이 한번에 하나씩 실행할 수 있는 WaitGroup(Add, Done and Wait)과 같은 API 를 사용할 수 있다.

package main

import (
	"fmt"
	"sync"
)

var (
	counter int

	// mutex 는 코드의 임계 구역을 정의하는데 사용된다.
	// 한번에 하나의 고루틴만이 이동할 수 있으며, 순서는 스케줄러가 조절한다. (예측불가)
	// 모든 고루틴들은 다른 고루틴이 들어오도록 나갈 때, 잠금과 해제를 요청한다.
	// 두 개의 함수가 동일한 mutex 를 사용할 수 있으므로 한 번에 하나의 고루틴만 주어진 함수를 실행할 수 있다.
	mutex sync.Mutex
)

func main() {
	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				mutex.Lock()
				{
					value := counter
					value++
					counter = value
				}
				mutex.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter: %d\n", counter)
}
