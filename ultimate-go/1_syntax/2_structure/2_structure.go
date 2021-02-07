package main

import "fmt"

// 구조체 필드의 순서는 메모리 할당 크기가 큰 순서대로 위치시키는 것이 좋다.
type example struct {
	counter int64   // 8바이트
	pi      float32 // 4바이트
	flag    bool    // 1바이트
}

func main() {
	// 구조체 타입의 변수를 선언하면, 구조체의 필드들은 제로값으로 초기화된다.
	var e1 example

	fmt.Printf("%+v\n", e1)
	fmt.Println("")

	// 구조체 리터럴로 초기화
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
	fmt.Println("")

	// 익명 타입 변수를 선언하고, 구조체 리터럴로 초기화. 익명 타입은 재사용 할 수 없다.
	e3 := struct {
		counter int64   // 8바이트
		pi      float32 // 4바이트
		flag    bool    // 1바이트
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}
	fmt.Println("Flag", e3.flag)
	fmt.Println("Counter", e3.counter)
	fmt.Println("Pi", e3.pi)
	fmt.Println("")

	// 같은 구조체 타입이라도 각기 다른 변수끼리 대입은 할 수 없다.
	// e1 = e2 (X)
	// e1 = example(e2) (O) 명시적 변환이 필요
	// 하지만, 동일한 구조의 익명 구조테 타입은 대입이 가능하다.
	var e4 example
	e4 = e3
	fmt.Printf("%+v\n", e4)
}
