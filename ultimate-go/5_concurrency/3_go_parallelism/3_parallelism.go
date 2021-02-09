// --------------------------
// Goroutines and parallelism
// --------------------------

// 이 프로그램은 고루틴이 병렬처리 되는 것을 보여준다.
// 2개의 프로세서(P), 2개의 스레드(m) 그리고 2개의 고루틴이 각각의 스레드(m)에서 병렬 처리 된다.
// 이전 프로그램과 비슷하지만 lowercase 함수와 uppercase 함수를 없애고, 익명 함수로 처리한다.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// 스케줄러에게 2개의 논리 프로세서를 할당한다.
	runtime.GOMAXPROCS(2)
}

func main() {
	// wg는 프로그램이 종료될때까지 기다리는데 사용한다.
	// Add에 2를 추가함으로, 2개의 고루틴이 종료될 때까지 대기한다.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 소문자 알파벳을 3번 출력하는 익명 함수를 선언하고, 고루틴을 생성한다.
	go func() {
		for count := 0; count < 3; count++ {
			for r := 'a'; r <= 'z'; r++ {
				fmt.Printf("%c ", r)
			}
		}

		// 메인(main)에게 작업이 끝났음을 알린다.
		wg.Done()
	}()

	// 대문자 알파벳을 3번 출력하는 익명 함수를 선언하고, 고루틴을 생성한다.
	go func() {
		for count := 0; count < 3; count++ {
			for r := 'A'; r <= 'Z'; r++ {
				fmt.Printf("%c ", r)
			}
		}

		// 메인(main)에게 작업이 끝났음을 알린다.
		wg.Done()
	}()

	// 고루틴이 끝나기를 기다린다.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
