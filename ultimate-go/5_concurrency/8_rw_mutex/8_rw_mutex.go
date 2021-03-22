// ----------------
// Read/Write Mutex
// ----------------

// 많은 고루틴이 읽기 원하는 공유 자원이 있다고 하자.
//
// 때때로, 하나의 고루틴이 들어와서 리소스를 바꿀 수 있다.
// 그렇게 되면, 모두 읽는 것을 중단해야한다.
// 아무런 이유 없이 소프트웨어에 대기시간을 추가하기 때문에 이러한 유형의 사나리오에서 읽기를 동기화하는 것은 의미가 없다.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// data 는 공유될 slice 이다.
	data []string
	// rwMutex 는 코드의 임계 구역을 정의하는데 사용된다.
	rwMutex sync.RWMutex
	// 조회하는 시간에 시도된 읽기 수 (int64 타입을 보자마자 아토믹 함수 사용에 대해 생각해야 한다)
	readCount int64
)

// init 은 main 보다 먼저 호출된다.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// 10개의 서로 다른 쓰기를 수행하는 쓰기용 고루틴을 만든다.
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			writer(i)
		}
		wg.Done()
	}()

	// 영원히 실행되는 8개의 읽기용 고루틴을 만든다.
	for i := 0; i < 8; i++ {
		go func(i int) {
			for {
				reader(i)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Program Complete")
}

// 쓰기용 고루틴은 임의의 간격으로 슬라이스에 새 문자열을 추가한다.
func writer(i int) {
	// 한번에 오직 하나의 고루틴만이 슬라이스에 읽기/쓰기를 하도록 허용된다.
	rwMutex.Lock()
	{
		// 다른 고루틴이 읽기를 수행하지 않는 것을 보장해야 한다.
		rc := atomic.LoadInt64(&readCount)
		fmt.Printf("****> : Performing Write : RCount[%d]\n", rc)
		data = append(data, fmt.Sprintf("String: %d", i))
	}
	rwMutex.Unlock()
}

func reader(id int) {
	// 모든 고루틴은 쓰기 작업이 일어나지 않을 때 읽을 수 있다.
	rwMutex.RLock()
	{
		rc := atomic.AddInt64(&readCount, 1)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Printf("%d : Performing Read : Length[%d] RCount[%d]\n", id, len(data), rc)
		atomic.AddInt64(&readCount, -1)
	}
	rwMutex.RUnlock()
}

// Lesson:

// 아토믹 함수와 뮤텍스는 사용자의 소프트웨어에 대기시간을 만든다.
// 대기시간은 여러 고루틴 사이에 자원 접근에 대한 조율이 필요할 때 유용하다. Read/Write 뮤텍스는 대기시간을 줄이는 데 유용하다.
//
// 뮤텍스를 사용하는 경우, 잠금 이후 최대한 빨리 잠금을 해제해야 한다.
// 다른 불필요한 행위는 하지 않는 것이 좋다.
// 때로는 공유 자원을 읽기 위해 로컬 변수만 사용하는 것으로도 충분하다.
// 뮤텍스를 적게 사용할수록 좋다. 이를 통해 대기시간을 최소한으로 줄일 수 있다.
