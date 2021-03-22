// --------------
// Race Detection
// --------------

// 멀티 스레드 소프트웨어를 작성할 때는 사실상 두 가지 선택지가 있다.
// - WaitGroup 에서 Add 와 Done, Wait 으로 제어하는 것 처럼, 공유 자원에 대한 접근 상태를 동기화
// - 고루틴을 예측 가능하고, 합리적으로 실행이 되도록 설계
// 채널이 없었을 때는, 아토믹 함수나 Mutex 를 사용하여 직접 구현했다.
// 채널은 간단한 제어 방법을 제공하지만, 대부분의 경우 아토믹 함수와 Mutex 를 사용하여 공유 자원에 대한 액세스 동기화를 사용하는 것이 가장 좋은 방법이다.
// 아토믹 연산은 Go 에서 가장 빠른 방법이다. 이는 Go는 메모리에서 한 번에 4~8 바이트씩 동기화를 하기 때문이다.
// 채널은 Mutex 뿐만 아니라 모든 데이터 구조와 로직이 함께 있기 때문에 매우 느린 편이다.

// 여러 개의 고루틴이 같은 자원에 접근하려 할 때 자원 경쟁이 발생한다.
// 자원 경쟁 상태를 만들어서 확인해 보자.

// go run -race 5_race_detection.go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// counter 는 모든 고루틴에 의해 증가되는 변수이다.
var counter int

func main() {
	// 사용 할 고루틴의 수
	const grs = 2

	// wg는 동시성을 관리하는 데 사용된다.
	var wg sync.WaitGroup
	wg.Add(grs)

	// 2개의 고루틴을 만들어 준다.
	for i := 0; i < grs; i++ {
		// 고루틴당 counter 를 2씩 증가 시킴
		// 프로그램을 실행할 때마다 출력은 4가 되어야 한다.
		go func() {
			for count := 0; count < 2; count++ {
				// 이 때의 counter 값을 저장해 둔다.
				// 여기서 발생하는 자원 경쟁 : 주어진 시간 동안 두 개의 고루틴은 동시에 읽고 쓸 수 있다.
				value := counter

				// 다른 고루틴에게 스레드를 양보하고, 다시 대기열에 들어간다. (테스트에만 사용할 것 !!)
				// 이 때, 공유 자원을 읽게 되면, 강제로 Context Switch 가 일어나게 되고, 자원 경쟁이 발생할 수 있다.
				runtime.Gosched()

				value++
				counter = value
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}
