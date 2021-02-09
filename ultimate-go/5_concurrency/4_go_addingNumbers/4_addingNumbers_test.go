// 가비지 컬렉터를 끈 상태에서 천 만개의 숫자 컬렉션으로 벤치마크 진행

// 모든 고루틴에 단일 OS/하드웨어 스레드만 사용할 수 있는 경우, Sequantial이 더 빠르다.
// 이는 Concurrent는 단일 OS 스레드에서 컨텍스트 스위치의 오버헤드와 고루틴 관리 때문이다.
// $ GOGC=off go test -cpu 1 -run none -bench . -benchtime 3s
// Processing 10000000 numbers using 8 goroutines
// goos: darwin
// goarch: amd64
// pkg: github.com/Sungchul-P/go-learning/ultimate-go/5_concurrency/4_go_addingNumbers
// BenchmarkSequential                  625           5577831 ns/op
// BenchmarkConcurrent                  573           6226007 ns/op
// BenchmarkSequentialAgain             540           5936244 ns/op
// BenchmarkConcurrentAgain             559           6876999 ns/op
// PASS
// ok      github.com/Sungchul-P/go-learning/ultimate-go/5_concurrency/4_go_addingNumbers  17.386s

// -----------------------------------------------------------------------------
// [Concurrency WITH Parallelism]
// $ GOGC=off go test -cpu 8 -run none -bench . -benchtime 3s
// Processing 10000000 numbers using 8 goroutines
// goos: darwin
// goarch: amd64
// pkg: github.com/Sungchul-P/go-learning/ultimate-go/5_concurrency/4_go_addingNumbers
// BenchmarkSequential-8                602           5851905 ns/op
// BenchmarkConcurrent-8               1694           1867099 ns/op
// BenchmarkSequentialAgain-8           606           5977516 ns/op
// BenchmarkConcurrentAgain-8          1791           1901268 ns/op
// PASS
// ok      github.com/Sungchul-P/go-learning/ultimate-go/5_concurrency/4_go_addingNumbers  16.055s

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

var numbers []int

func init() {
	rand.Seed(time.Now().UnixNano())
	numbers = generateList(1e7)
	fmt.Printf("Processing %d numbers using %d goroutines\n", len(numbers), runtime.NumCPU())
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), numbers)
	}
}

func BenchmarkSequentialAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkConcurrentAgain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addConcurrent(runtime.NumCPU(), numbers)
	}
}
