// 메서드는 특수한 기능이 아니라 문법적인 필요에 의해 만들어진 것이다.
// 메서드는 데이터와 관련된 일부 기능을 외부에서 사용할 수 있는 것처럼 믿게 만든다.
// 객체지향 프로그래밍에서도 이러한 설계와 기능을 권장한다. 하지만 Go가 객체지향 프로그래밍를 추구하는 것은 아니다.
// 다만 데이터와 동작이 필요하기 때문에 이들이 있는 것이다.

// 어떤 때는 데이터가 일부 기능을 노출할 수 있지만 특정 목적을 위해 API를 설계하는 것은 아니다. 메서드는 함수와 거의 동일하다.

package main

import "fmt"

// data is a struct to bind methods to.
type data struct {
	name string
	age  int
}

// displayName은 data타입 d 변수의 name을 포함한 문자열을 출력한다.
// 이 메서드는 data를 값 리시버로 사용한다.
func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

// setAge는 data타입 d의 age를 수정하고, 이 값을 name과 함께 출력한다.
// 이 메서드는 data를 포인터 리시버로 사용한다.
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

func main() {
	// --------------------------
	// Methods are just functions
	// --------------------------

	// data 타입의 변수를 선언해보자.
	d := data{
		name: "Devnori",
	}

	fmt.Println("Proper Calls to Methods:")

	// How we actually call methods in Go.
	d.displayName()
	d.setAge(21)

	fmt.Println("\nWhat the Compiler is Doing:")

	// 다음 예제를 통해 Go가 내부적으로 어떻게 동작하는지 알 수 있다.
	// d.displayName()을 호출하면, 컴파일러는 data.displayName을 먼저 호출해서 data 타입의 값 리시버를 사용하는 것을 보여주고,
	// d를 첫번째 파라미터로 전달한다. func (d data) displayName()을 다시 자세히 보면, 리시버가 정말 파라미터임을 알 수 있다.
	// 즉, 이 리시버는 displayName의 첫번째 파라미터이다.
	// d.setAge(21) 역시 이와 비슷하다. Go는 포인터 리시버를 사용하는 함수를 호출하고, d를 함수의 파라미터로 넘긴다.
	// 추가로 d의 주소 값을 이용하기 위하여 약간의 조정이 필요하다.
	data.displayName(d)
	(*data).setAge(&d, 21)

	// -----------------
	// Function variable
	// -----------------

	fmt.Println("\nCall Value Receiver Methods with Variable:")

	// 함수형 변수를 선언하고, 이 변수에 d 변수의 메서드를 설정해보자.
	// 이 메서드, displayName은 값 리시버를 사용하기 때문에, 함수형 변수 f1은 d의 독립적인 복사본을 가진다.
	// 함수형 변수 f1은 참조 타입으로 포인터, 즉 주소 값을 저장한다.
	// displayName의 뒤에 ()을 붙이지 않았으므로, 이 메서드의 반환 값을 저장한 것이 아니다.
	f1 := d.displayName

	// 이 변수를 이용해서 메서드를 호출해보자.
	// f1은 포인터로, 이는 2개의 워드를 가지는 특별한 자료 구조를 가리킨다.
	// 첫 번째 워드는 실행 대상인 메서드를 가리키는데, 이 예제에서는 displayName이다.
	// 이 displayName는 값 리시버를 사용하므로 실행하기 위해서는 data 타입의 값이 필요하다.
	// 따라서 두 번째 워드는 이 값의 복사본을 가리킨다. displayName을 f1에 저장을 하면, 자동으로 d의 복사본이 만들어진다.
	//  -----
	// |  *  | --> code
	//  -----
	// |  *  | --> copy of d
	//  -----
	f1()

	// 만약 d의 멤버 변수인 name의 값을 “Sungchul Park”으로 변경하더라도, f1에는 이 변경이 적용되지 않는다.
	d.name = "Sungchul Park"

	// f1을 통해서 메서드를 호출해보자. 앞서 이야기했듯이, 결과는 변하지 않았다.
	f1()

	// 하지만 setAge 메서드를 저장한 f2는 d의 값이 변하면 그 자신의 결과도 변한다.

	fmt.Println("\nCall Pointer Receiver Method with Variable:")

	// 함수형 변수를 선언하고, 이 변수에 d 변수의 메서드를 설정해보자.
	// 이 메서드, setAge는 포인터 리시버를 사용하므로, 이 함수형 변수는 d의 주소 값을 가지게 된다.
	f2 := d.setAge

	// 이 함수형 변수를 이용해서 메서드를 호출해보자.
	// f2는 역시 포인터이고, 2개의 워드를 가지는 자료 구조를 가리킨다.
	// 첫 번째 워드는 setAge 메서드를 가리키지만, 두 번째 워드는 더 이상 복사본이 아니라 원본 d를 가리킨다.
	//  -----
	// |  *  | --> code
	//  -----
	// |  *  | --> original d
	//  -----

	d.name = "Sungchul Park Dev"

	f2(22)
}
