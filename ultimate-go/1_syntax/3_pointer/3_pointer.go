package main

import "fmt"

func main() {

	count := 10 // 이 변수는 스택에 저장 된다.
	// 변수의 주소를 얻기 위해 &를 사용한다.
	fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// count의 값을 전달한다.
	increment1(count)

	// increment1 를 실행한 다음의 count 값을 출력한다. 바뀐 것이 없다.
	fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// count의 주소를 전달한다. 이것 역시 "pass by value", 즉, 값을 전달하는 것이다.
	// "pass by reference" 가 아니다. 주소 역시 값인 것이다.
	increment2(&count)

	// increment2 를 실행한 다음 count 값을 출력한다. 값이 변경되었다.
	fmt.Println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// ---------------
	// Escape analysis : 무엇을 스택 또는 힙에 할당할 것인지 결정한다.
	// ---------------

	stayOnStack()  // 값의 복사본을 전달하므로 스택 프레임에 위치
	escapeToHeap() // 힙을 가리키는 포인터를 반환
}

func increment1(inc int) {
	// inc 의 값을 증가 시킨다.
	inc++
	fmt.Println("inc1:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

// increment2 는 inc를 포인터 변수로 선언했다. 이 변수는 주소값을 가지며, int 타입의 값을 가리킨다.
// *는 연산자가 아니라 타입 이름의 일부이다. 이미 선언된 타입이건, 당신이 선언한 타입이건
// 모든 타입은 선언이 되면 포인터 타입도 가지게 된다.
func increment2(inc *int) {
	// inc 포인터 변수가 가리키고 있는 int 변수의 값을 증가시킨다.
	// 여기서 *는 연산자이며 포인터 변수가 가리키고 있는 값을 의미한다.
	*inc++
	fmt.Println("inc2:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}

// user는 시스템의 user를 의미한다.
type user struct {
	name  string
	email string
}

func stayOnStack() user {
	// 스택 프레임에 변수를 생성하고 초기화한다.
	u := user{
		name:  "Devnori",
		email: "sungchul.dev@gmail.com",
	}

	// 값을 리턴하여 main의 스택 프레임으로 전달한다.
	return u
}

func escapeToHeap() *user {
	u := user{
		name:  "Devnori",
		email: "sungchul.dev@gmail.com",
	}

	// 값의 주소를 리턴한다. (주소 값을 콜 스택으로 전달)
	return &u
}
