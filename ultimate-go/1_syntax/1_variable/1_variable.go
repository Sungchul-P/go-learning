package main

import "fmt"

func main() {
	// var로 변수를 선언하면 타입의 제로값으로 초기화
	var a int
	var b string
	var c float64
	var d bool
	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// 짧은 변수 선언(short variable declaration) 연산자로 선언과 동시에 초기화
	aa := 10
	bb := "hello" // 첫 번째 워드는 문자들의 배열을 기리키는 포인터이고, 두 번째 워드는 5이다.
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// Conversion
	aaa := int32(10) // int -> int32
	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}
