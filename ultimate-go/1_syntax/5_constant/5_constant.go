package main

import "fmt"

func main() {

	// 타입이 없는 상수
	const ui = 12345    // kind: integer
	const uf = 3.141592 // kind: floating-point

	// 타입이 있는 상수
	const ti int = 12345        // type: int
	const tf float64 = 3.141592 // type: float64

	// 타입 제한을 초과하여 오버플로우 발생
	// Error => constant 1000 overflows uint8
	//const myUint8 uint8 = 1000

	// 상수는 다른 kind간 연산을 지원한다. Kind 승급(kind promotion)을 이용해서 어떤 kind로 연산할지 결정한다. (암묵적으로 수행)
	var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)
	fmt.Println(answer)

	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)
	fmt.Println(third)

	const zero = 1 / 3 // KindInt(1) / KindInt(3)
	fmt.Println(zero)

	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)

	fmt.Println(one)
	fmt.Println(two)

	const maxInt = 9223372036854775807 // 64bit 아키텍처에서 가장 큰 정수
	fmt.Println(maxInt)

	// int64 타입보다 훨씬 큰 숫자이지만 타입이 없는 상수이기 때문에 컴파일에 문제가 없다.
	const bigger = 9223372036854775808543522345

	// 타입을 명시한 경우는 컴파일 시 에러가 발생한다.
	// Error => constant 9223372036854775808543522345 overflows int64
	//const biggerInt int64 = 9223372036854775808543522345

	// iota is Greek alphabet (incrementing numbers)
	const (
		A1 = iota // 0 : 0에서 시작한다
		B1 = iota // 1 : 1 증가한다
		C1 = iota // 2 : 1 증가한다
	)

	fmt.Println("iota - 1:", A1, B1, C1)

	const (
		A2 = iota // 0 : 0에서 시작한다
		B2        // 1 : 1 증가한다
		C2        // 2 : 1 증가한다
	)

	fmt.Println("iota - 2:", A2, B2, C2)

	const (
		A3 = iota + 1 // 1 : 1에서 시작한다
		B3            // 2 : 1 증가한다
		C3            // 3 : 1 증가한다
	)

	fmt.Println("iota - 3:", A3, B3, C3)

	const (
		Ldate         = 1 << iota //  1 : 오른쪽으로 0번 시프트 된다. 0000 0001
		Ltime                     //  2 : 오른쪽으로 1번 시프트 된다. 0000 0010
		Lmicroseconds             //  4 : 오른쪽으로 2번 시프트 된다. 0000 0100
		Llongfile                 //  8 : 오른쪽으로 3번 시프트 된다. 0000 1000
		Lshortfile                // 16 : 오른쪽으로 4번 시프트 된다. 0001 0000
		LUTC                      // 32 : 오른쪽으로 5번 시프트 된다. 0010 0000
	)
	fmt.Println("iota shift Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}
