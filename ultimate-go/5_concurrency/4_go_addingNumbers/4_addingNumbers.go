package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	numbers := generateList(1e7)

	fmt.Println(add(numbers))
	fmt.Println(addConcurrent(runtime.NumCPU(), numbers))
}

func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)
	for i := 0; i < totalNumbers; i++ {
		numbers[i] = rand.Intn(totalNumbers)
	}
	return numbers
}

// add()는 비순차(out of order) 실행에 적합한 워크로드인가? (YES)
// --------------------------------------------------------
// []int는 더 작은 리스트로 분할될 수 있고 동시에 처리될 수 있다.
// 분할된 리스트의 합계를 더하면 순차처리와 동일한 결과를 낼 수 있다.

// 최상의 처리량을 얻으려면 몇 개의 작은 리스트를(smaller lists) 개별적으로 만들고 처리해야 하는가?
// --------------------------------------------------------
// 이 질문에 답하려면 add 작업이 어떤 작업량을 수행하는지 알아야 한다.
// add()는 알고리즘이 단순 계산을 수행하고 고루틴이 대기 상태로 들어가는 원인이 되지 않기 때문에
// CPU 바운드 워크로드를 수행한다.
// 이는 OS / 하드웨어 스레드 당 하나의 고루틴을 사용하는 것만으로도 우수한 처리량을 얻을 수 있음을 의미한다.
func add(numbers []int) int {
	var v int
	for _, n := range numbers {
		v += n
	}
	return v
}

func addConcurrent(goroutines int, numbers []int) int {
	var v int64
	totalNumbers := len(numbers)
	lastGoroutine := goroutines - 1

	// 각 고루틴은 고유하지만 더 작은 숫자 리스트를 갖게 된다.
	// 리스트의 크기는 컬렉션의 크기를 고루틴 수로 나누어 계산한다.
	stride := totalNumbers / goroutines

	var wg sync.WaitGroup
	wg.Add(goroutines)

	// adding 작업을 수행하기 위해 고루틴 풀을 생성한다.
	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride

			// 마지막 고루틴은 다른 고루틴보다 클 수 있는 나머지 숫자 목록을 추가한다.
			if g == lastGoroutine {
				end = totalNumbers
			}

			var lv int
			for _, n := range numbers[start:end] {
				lv += n
			}

			// 작은 리스트의 합이 모두 합쳐져 최종 합이됩니다.
			atomic.AddInt64(&v, int64(lv))
			wg.Done()
		}(g)
	}

	wg.Wait()

	return int(v)
}

// addConcurrent()는 add()에 비해 확실히 더 복잡하지만 그 복잡성이 그만한 가치가 있는가?
// 이를 확인하기 위한 가장 좋은 방법은 벤치마크를 만드는 것입니다.
