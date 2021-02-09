// ----------------------
// Goroutine time slicing
// ----------------------

// Go의 스케줄러는 선점 스케줄러(preemptive)가 아닌 협력 스케줄러(cooperating scheduler)에도
// 선점된 것처럼 생각되는 이유는 런타임 스케줄러가 프로그래머에게 인식되기 전에 모든 처리를 하기 때문이다.

// 아래의 코드는 문맥교환을 보여주고, 언제 문맥교환이 발생하는지 예상할 수 있도록 보여준다.
// 1_goroutine.go 코드와 같은 패턴이지만 printPrime 함수가 새로 추가된다.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Create Goroutines")

	go func() {
		printPrime("A")
		wg.Done()
	}()

	go func() {
		printPrime("B")
		wg.Done()
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime 함수는 5000보다 작은 소수(prime numbers)를 출력한다.
// 특별한 함수는 아니지만, 완료하기 위해 약간의 시간이 필요하다.
// 프로그램을 실행하면 특정 소수에서 문맥교환(Context Switching)이 일어나는 것을 볼 수 있다.
// 하지만 문맥교환이 언제 일어날 지는 예측할 수 없기에 Go의 스케줄러가 협력 스케줄러임에도
// 불구하고 선점 스케줄러처럼 보인다고 말하는 이유이다.
func printPrime(prefix string) {
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}

		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)
}
