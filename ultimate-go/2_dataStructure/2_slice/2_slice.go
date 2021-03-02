package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// ----------------------
	// Declare and initialize
	// ----------------------

	// 5개의 요소를 갖는 슬라이스(slice)를 생성해보자.
	// make 함수는 슬라이스(slice)와 맵(map) 그리고 채널(channel) 타입에서 사용하는, 특별한 내장 함수이다.
	// make 함수를 사용하여 5개의 문자열 배열을 갖는 슬라이스를 생성하면, 3개의 워드(word) 데이터 구조가 만들어진다.
	// 첫 번째 워드는 배열을 위치를 가리키고 두 번째 워드는 길이를, 세 번째 워드는 용량을 나타낸다.
	//  -----
	// |  *  | --> | nil | nil | nil | nil | nil |
	//  -----      |  0  |  0  |  0  |  0  |  0  |
	// |  5  | (Length)
	//  -----
	// |  5  | (Capacity)
	//  -----

	// ------------------
	// Length vs Capacity
	// ------------------

	// 길이(Length)는, 포인터의 위치에서부터 접근해서 읽고 쓸 수 있는 요소의 수를 의미하며,
	// 용량(Capacity)은 포인터의 위치에서부터 배열에 존재할 수 있는 요소의 총량을 뜻한다.
	// 슬라이스는 언뜻 배열처럼 보인다. 비용도 배열과 동일하게 발생한다.
	// 하지만, 한 가지 다른 점은 make 함수의 []string의 대괄호 안에 값이 없다는 것이다. 이것으로 배열과 슬라이스를 구분할 수 있다.

	slice1 := make([]string, 5)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Plum"

	// 슬라이스의 길이를 넘는 인덱스에는 접근할 수 없다.
	// panic: runtime error: index out of range [5] with length 5
	//slice1[5] = "Runtime error"

	// 슬라이스의 주소가 아닌 값(value)을 전달한다. 따라서, Println함수는 슬라이스의 복사본을 갖게 된다.
	fmt.Printf("\n=> Printing a slice\n")
	fmt.Println(slice1)

	// --------------
	// Reference type
	// --------------

	// 5개의 요소를 갖고 용량이 8개인 슬라이스를 만들기 위해 make 키워드를 이용할 수 있으며,
	// 이를 통해 초기화 시점에 직접 용량을 정할 수 있다.
	// 결국, 차례대로 8개의 요소를 갖는 "배열을 가르키는 포인터"와 "길이"는 5, "용량"은 8을 갖는 3개의 워드(word)형 자료 구조를 갖게된다.
	// 이는 첫 5개의 요소에 대해 읽고 쓸 수 있으며, 필요시 이용 가능한 3개의 용량을 갖는 것을 뜻한다.
	//  -----
	// |  *  | --> | nil | nil | nil | nil | nil | nil | nil | nil |
	//  -----      |  0  |  0  |  0  |  0  |  0  |  0  |  0  |  0  |
	// |  5  |
	//  -----
	// |  8  |
	//  -----

	slice2 := make([]string, 5, 8)
	slice2[0] = "Apple"
	slice2[1] = "Orange"
	slice2[2] = "Banana"
	slice2[3] = "Grape"
	slice2[4] = "Plum"

	fmt.Printf("\n=> Length vs Capacity\n")
	inspectSlice(slice2)

	// --------------------------------------------------------
	// Idea of appending: making slice a dynamic data structure
	// --------------------------------------------------------
	fmt.Printf("\n=> Idea of appending\n")

	// 문자열로 구성될 nil 슬라이스를 선언하고 그 값을 제로 값(zero value)으로 설정한다.
	// 이 때, 첫 번째 워드(word)는 nil을 가르키는 포인터로, 두 번째와 세 번째 값은 0을 나타내는 3개의 워드(word) 자료구조를 갖는다.
	var data []string

	// 만약, data := string{}을 하게되면, 이 둘은 서로 같을까?
	// 그렇지 않다. 왜냐하면 이 경우, data는 값이 제로 값(zero value)으로 설정되지 않기 때문이다.
	// 빈 리터럴로 생성되는 모든 타입이 제로 값(zero value)을 반환하지는 않기 때문에, 제로 값(zero value)을 위해 var을 사용하는 이유기도하다.
	// 위의 경우에, 반환 되는 슬라이스는 nil 슬라이스가 아닌 포인터를 갖고 있는 빈 슬라이스가 된다.
	// nil 슬라이스와 빈 슬라이스 간에는 각기 다른 의미가 있는데, 제로 값(zero value)으로 설정된 참조 타입은 nil로 여길 수 있다는 점이다.
	// marshal 함수에 nil 슬라이스를 넘긴다면 null을 반환하고, 빈 슬라이스를 넘긴다면 빈 JSON을 반환하게 된다.
	// 그렇다면 이 때 포인터는 어떤 것을 가르키게 될까? 바로, 나중에 살펴볼 빈 구조체를 가르킨다.

	// 슬라이스의 용량(capacity)을 가져오자.
	lastCap := cap(data)

	// 슬라이스에 약 10만개의 문자열을 덧붙인다.
	for record := 1; record <= 102400; record++ {

		// 내장 함수인 append를 사용해서 슬라이스를 덧붙일 수 있다.
		// 이 함수를 통해 슬라이스에 값을 추가할 수 있으며 자료 구조를 동적으로 만들 수 있으면서도,
		// 기계적 동정심(mechanical sympathy)을 통해 예측 가능한 접근 패턴을 제공함으로써 여전히 인접한 메모리 블럭을 이용할 수 있게 된다.
		// append 함수는 값 개념(value semantic)으로 동작한다. 슬라이스 자체를 공유하는 것이 아니라, 슬라이스에 값을 덧붙이고 그 복사본을 반환하는 식이다.
		// 따라서 슬라이스는 힙 메모리가 아닌 스택에 위치하게 된다.
		data = append(data, fmt.Sprintf("Rec: %d", record))

		// append가 동작할 때 마다, 매번 길이와 용량을 확인한다. 만약 두 값이 동일하다면, 더 이상 남은 공간이 없다는 것을 뜻한다.
		// 이 때, append 함수는 기존보다 2배를 늘린 크기를 갖는 새로운 배열을 만들어서 예전 값을 복사한 뒤, 새 값을 추가하게 된다.
		// 그리고 스택 프레임에 존재하는 값을 변경시킨 뒤, 그 복사본을 반환한다. 그렇게 기존의 슬라이스가 새로운 복사본으로 치환된다.
		// 만약 길이와 용량이 같지 않다면, 슬라이스 안에 아직 사용할 수 있는 공간이 남아있다는 것을 뜻하므로, 새 복사본을 만드는 일 없이 값을 추가 할 수 있다.
		// 이것은 굉장히 효율적이다. 출력의 마지막 열을 확인해보자. 배열의 요소가 1000개 혹은 그 이하일 때, 배열의 크기는 2배로 늘어난다.
		// 요소의 개수가 1000개를 넘고 나면, 용량의 변화율은 25%로 변한다. 슬라이스의 용량이 변경될 때, 그 변화를 나타낸다.
		if lastCap != cap(data) {
			// 변화율을 계산한다.
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// lastCap에 새 용량을 저장한다.
			lastCap = cap(data)

			// 결과를 표시한다.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n", &data[0], record, cap(data), capChg)
		}
	}

	// --------------
	// Slice of slice
	// --------------

	// slice2의 인덱스 2, 인덱스 3의 값을 갖는 slice3을 생성하자. slice3의 길이는 2이고 용량은 6이다.
	// 매개변수는 [시작 인덱스:(시작 인덱스 + 길이)] 형태이다.
	// 결과를 통해 두 슬라이스는 같은 배열을 공유하고 있는 것을 알 수 있다. 슬라이스의 헤더는 값의 개념으로 사용 될 때 스택에 존재한다. 오직 공유되는 배열만이 힙에 위치한다.
	slice3 := slice2[2:4]

	fmt.Printf("\n=> Slice of slice (before)\n")
	inspectSlice(slice2)
	inspectSlice(slice3)

	// slice3의 인덱스 0의 값을 바꾸면, 어떤 슬라이스가 변경 될까?
	slice3[0] = "CHANGED"

	// 두 슬라이스 모두 변한다. 생성되어 있는 슬라이스를 변경한다는 것을 잊지 말아야 한다. 어디서 이 슬라이스를 사용하는지, 또 배열을 공유하고 있는지를 주의깊게 살펴야 한다.
	fmt.Printf("\n=> Slice of slice (after)\n")
	inspectSlice(slice2)
	inspectSlice(slice3)

	// slice3 := append(slice3, "CHANGED")는 어떨까? 슬라이스의 길이와 용량이 다르면, append를 사용 할 때 비슷한 문제가 발생한다.
	// slice3의 0번째 인덱스 값을 변경하는 대신에, append를 호출 해보자. slice3의 길이는 2이고 용량은 6이라서 수정을 위한 여유 공간을 가지고 있다.
	// slice2의 4번째 인덱스와 같은 주소 값을 가지는 slice3의 3번째 인덱스의 원소부터 변경 된다. 이런 상황은 굉장히 위험하다.
	// 그러면 슬라이스의 길이와 용량이 서로 같으면 어떨까? 슬라이싱 구문(slicing syntax)의 또 다른 매개변수를 추가하여, slice3의 용량을 6 대신 2로 만들어보자: slice3 := slice2[2:4:4]
	//
	// 길이와 용량이 같은 슬라이스에 대해 append가 호출 되면, slice2의 4번째 원소를 가지고 오지 않는다. 이것은 분리되어 있다.
	// slice3는 길이가 2이고 용량이 2이면서 여전히 slice2와 같은 배열을 공유하고 있다. append가 호출 되면, 길이와 용량이 달라지게 된다.
	// 주소 또한 달라지게 된다. 길이는 3인 새로운 슬라이스가 된다. 새로운 슬라이스 소유의 배열을 갖게 되고 더 이상 원본 슬라이스의 영향을 받지 않는다.

	// ------------
	// Copy a slice
	// ------------

	// 복사는 문자열과 슬라이스 타입에서만 동작 한다. 원본의 요소들을 담을 수 있을 만큼의 크기로 새로운 슬라이스를 만들고 내장 함수인 copy를 사용해서 값을 복사한다.
	slice4 := make([]string, len(slice2))
	copy(slice4, slice2)

	fmt.Printf("\n=> Copy a slice\n")
	inspectSlice(slice4)

	// -------------------
	// Slice and reference
	// -------------------

	// 길이가 7인 정수형 슬라이스를 선언하자.
	x := make([]int, 7)

	// 임의의 값을 넣어준다.
	for i := 0; i < 7; i++ {
		x[i] = i * 100
	}

	// 슬라이스의 두 번째 원소의 포인터를 변수에 할당한다.
	twohundred := &x[1]

	// 슬라이스에 새로운 값을 추가해보자. 이 코드는 위험 하다. x 슬라이스는 길이가 7이고 용량 7이다.
	// 길이와 용량이 같기 때문에 용량이 두 배로 늘어나고 값들이 복사된다. 이제 x 슬라이스는 길이가 8이고 용량이 14이며 다른 메모리 블록을 가르킨다.
	x = append(x, 800)

	// 슬라이스의 두 번째 원소의 값을 변경 할 때, twohundred는 변경되지 않는다. 이전의 슬라이스를 가리키기 때문이다. 이 변수를 읽을 때 마다, 잘못된 값을 얻는다.
	x[1]++

	fmt.Printf("\n=> Slice and reference\n")
	fmt.Println("twohundred:", *twohundred, "x[1]:", x[1])

	// -----
	// UTF-8
	// -----

	fmt.Printf("\n=> UTF-8\n")

	// Go의 모든 것은 UTF-8 문자 집합(chrarcter sets)을 근간으로 한다. 만약 다른 인코딩 구조를 사용한다면 문제가 발생한다.
	//
	// 중국어와 영어로 문자열을 선언하자. 중국 문자는 각각 3 바이트를 사용한다. UTF-8는 바이트, 코드 포인트(code point) 그리고 문자로 3 계층을 이루고 있다.
	// Go 관점에서 문자열은 단지 저장되는 바이트일 뿐이다.
	//
	// 아래 예제에서 첫 번째 3 바이트는 하나의 코드 포인트를 표현한다. 하나의 코드 포인트는 하나의 문자를 표현한다.
	// 1 바이트부터 4 바이트를 가지고 코드 포인트를 표현 할 수 있고(코드 포인트는 32 비트 값이다.) 1 부터 대다수의 코드 포인트는 문자를 표현 할 수 있다.
	// 간단하게, 3 바이트로 1 코드 포인트로 1 문자를 표현한다. 그래서 s 문자열을 3 바이트, 3 바이트, 1 바이트, 1 바이트.. 로 읽는다(앞 부분에 중국 문자가 2개 있고 나머지는 영어이기 때문에)
	s := "世界 means world"

	// UTFMax 상수 4다. – 인코딩 된 rune 당 최대 4 바이트 -> 모든 코드 포인트를 표현하기 위해 필요한 최대 바이트 수는 4 바이트다.
	// Rune은 자체 타입이다. 이것은 int32 타입의 별칭이다. 우리가 사용하는 byte 타입은 uint8의 별칭이다.
	var buf [utf8.UTFMax]byte

	// 위의 문자열을 순회 할 때, 바이트에서 바이트, 코드 포인트에서 코드 포인트, 문자에서 문자로 중 어느 방식으로 순회할까?
	// 정답은 코드 포인트에서 코드 포인트이다. 첫 번째 순회에서 i는 0이다. 그 다음 i는 다음 코드 포인트로 이동되기 때문에 3이다. 그 다음은 6이다.
	for i, r := range s {
		// 룬/코드 포인트의 바이트 수를 출력해보자.
		rl := utf8.RuneLen(r)

		// 룬을 표현하는 바이트들의 차이를 계산하자.
		si := i + rl

		// 문자열로부터 룬을 버퍼에 복사하자. 모든 코드 포인트를 순회하며 배열 버퍼에 복사하고 화면에 출력하려는 것이다.
		// Go 에서 “배열은 언제든 슬라이스가 될 준비가 되어있다.” 슬라이싱 구문을 사용하여 buf 배열을 가리키는 슬라이스 헤더를 만든다. 헤더는 스택에 생성되기에 힙에 할당하지 않는다.
		copy(buf[:], s[i:si])
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])

	}

}

// inspectSlice는 리뷰를 위해 슬라이스 헤더를 보여주는 함수이다.
// 파라미터 : 다시 말하지만, []string의 대괄호 속에 값이 없으므로 슬라이스를 사용함을 알 수 있다.
// 배열에서 했던 것과 마찬가지로, 슬라이스를 순회한다.
// `len`이 슬라이스의 길이를 알려주며, `cap`은 슬라이스의 용량을 알려준다.
// 결과를 보면, 예상대로 슬라이스의 주소 값들이 정렬되어 표시되는 것을 볼 수 있다.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], slice[i])
	}
}
