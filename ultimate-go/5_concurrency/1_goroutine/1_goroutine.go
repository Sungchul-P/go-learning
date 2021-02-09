package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 스케줄러에게 하나의 논리 프로세서만 할당할 것을 명시한다.
func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	// wg는 동시성을 관리하는데 사용된다.
	// 제로값으로 설정되며 제로값 상태에서 사용할 수 있는 Go의 매우 특별한 타입이다.
	// 비동기 계산 세마포어(Asynchronous Counting Semaphore)로 불린다.
	// 세 개의 메소드를 갖는다. n개의 고루틴은 이 메소드를 동시에 호출할 수 있고 모두 직렬화 되어 있다.
	// - Add: 얼마나 많은 고루틴이 있는지 계산한다.
	// - Done: 일부 고루틴이 종료될 예정이므로 값을 감소시킨다.
	// - Wait: 해당 카운트가 0이 될 때까지 프로그램을 유지한다.
	var wg sync.WaitGroup

	// 2개의 고루틴을 생성한다.
	// 만약 얼마나 많은 고루틴이 생성될지 모른다면, 이로 인해 데드락이 발생할 수 있다.
	wg.Add(2)
	fmt.Println("Start Goroutines")

	// 익명함수를 사용해 lowercase 함수에서 고루틴을 생성한다.
	// Go 스케줄러는 해당 함수를 G로 예약한다. 이를 G1 이라고 하자.
	// 그리고 P(프로세서)에 대해 일부 LRQ(Local Run Queue)를 로드한다.
	// 모든 G에 대해서 runnable state라면, 동시에 실행할 수 있다.
	// 싱글 프로세서(P) 이거나 싱글 스레드일지라도 전혀 상관없다.
	// 2개의 고루틴(main, G1)을 동시에 실행한다.
	go func() {
		lowercase()
		wg.Done()
	}()

	// lowercase 이후에 고루틴을 하나 더 생성한다.
	// 따라서, 이제는 3개의 고루틴이 동시에 실행된다.
	go func() {
		uppercase()
		wg.Done()
	}()

	// 고루틴이 종료되는지 모른다면 고루틴을 만들 수도 없다.
	// Wait는 두 개의 고루틴이 Done이 될 때까지 대기한다.
	// 2에서 0이 될 때까지 카운트하고, 0에 도달했을 때 스케줄러는 메인 고루틴을 마저 실행하고 종료한다.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

// lowercase 함수는 알파벳 소문자를 세번 반복 출력한다.
func lowercase() {
	for count := 0; count < 3; count++ {
		for r := 'a'; r <= 'z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
}

// uppercase 함수는 알파벳 대문자를 세번 반복 출력한다.
func uppercase() {
	for count := 0; count < 3; count++ {
		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("%c ", r)
		}
	}
}

// 순서 (Sequence)
// --------------
// lowercase 함수를 호출하고, uppercase 함수를 호출했지만, Go 스케줄러는 uppercase 함수를 먼저 호출했다.
// 싱글 스레드에서 동작하기에 순간에 1개의 고루틴만 실행할 수 있는 점을 기억하자.
// 동시성을 지키며 실행되는지 알 수 없는데, uppercase 함수가 lowercase 함수보다 먼저 실행되는지 알 수 없다.

// Wait을 하지 않으면 어떻게 될까?
// ---------------------------
// uppercase 함수와 lowercase 함수의 결과를 볼 수 없다.
// Go 스케줄러가 프로그램 종료를 막고 새로운 고루틴을 만들어 작업을 할당하기 전에 프로그램이 종료되는 일종의 경쟁으로 보인다.
// 기다리지 않기 때문에 고루틴은 실행할 기회가 전혀 없다.

// Done을 호출하지 않으면 어떻게 될까?
// ------------------------------
// 교착상태(Deadlock)가 발생한다.
// Go의 특별한 부분이며, 런타임에서 고루틴이 존재하지만 더 이상 진행할 수 없을 때 패닉(panic)상태가 된다.
