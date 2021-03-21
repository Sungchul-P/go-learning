// --------------------
// Grouping By Behavior
// --------------------

// 이번 파트는 구성과 인터페이스를 활용한 예시이며 이는 Go 에서 사용되면 좋은 방식이다.

// 공통된 상태가 아닌 공통된 행동으로 그룹화하는 이 패턴은 Go 프로그램에서 좋은 설계 원칙이다.
// Go 의 뛰어난 특성중 하나는 미리 구성할 필요가 없다는 점이다.
// 컴파일러는 컴파일 타임에 인터페이스와 행동을 자동으로 식별한다.
// 이는 현재 또는 미래에 작성한 인터페이스와 호환될 수 있는 코드를 지금 작성할 수 있다는 의미이다.
// 또한 컴파일러가 즉석에서 행동을 식별하기 때문에 선언된 위치도 중요하지 않다.
// 대신 어떤 행동을 해야 하는지를 고려해야 한다.

package main

import "fmt"

// Speaker 는 그룹이 일원이 되기 위해 따라야 할 공통된 행동을 정의하고 있다.
// Speaker 는 구체적인 타입에 대한 계약이다. Animal 타입은 제거한다.
type Speaker interface {
	Speak()
}

// Dog 는 Dog 가 필요한 모든 것을 정의하고 있다.
type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

// Speak 는 개가 말하는 방식을 의미한다.
// 말하는 방식이 정의된 Dog 는 구체적인 타입인 Speaker 그룹의 일원이 된다.
func (d Dog) Speak() {
	fmt.Println("Woof!",
		"My name is", d.Name,
		", it is", d.IsMammal,
		"I am a mammal with a pack factor of", d.PackFactor)
}

// Cat 는 Cat 이 필요한 모든 것을 정의하고 있다.
// 복사 붙여넣기를 하면 약간의 시간이 걸릴지도 모르지만, 대부분의 경우 디커플링은 코드 재사용보다 더 나은 방식이다.
type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

// Speak 는 고양이가 말하는 방식을 의미한다.
// 말하는 방식이 정의된 Cat 는 구체적인 타입인 Speaker 그룹의 일원이 된다.
func (c Cat) Speak() {
	fmt.Println("Meow!",
		"My name is", c.Name,
		", it is", c.IsMammal,
		"I am a mammal with a climb factor of", c.ClimbFactor)
}

func main() {
	// 말하는 방식이 정의된 동물들을 생성해보자.
	speakers := []Speaker{
		// Dog 의 속성을 초기화해서 Dog 를 생성한다.
		Dog{
			Name:       "Fido",
			IsMammal:   true,
			PackFactor: 5,
		},

		// Cat 의 속성을 초기화해서 Cat 을 생성한다.
		Cat{
			Name:        "Milo",
			IsMammal:    true,
			ClimbFactor: 4,
		},
	}

	// Speaker 들이 Speak()을 호출하도록 한다.
	for _, spkr := range speakers {
		spkr.Speak()
	}
}

// ---------------------------------
// Guidelines around declaring types
// ---------------------------------

// - 새롭거나 유일한 것을 대표하는 타입을 선언한다. 가독성을 위해 별칭(alias)을 생성하지 않는다.
// - 모든 타입의 값이 자체적으로 생성되거나 사용되었는지 확인한다.
// - 상태가 아니라 행동을 위한 타입을 임베드 해야 한다. 행동에 대해 고려하지 않는다면 미래에 유지보수하기 어려운 설계가 될 수 있다.
// - 특정 타입이 이미 있는 타입을 위한 별칭이거나 추상화하고 있다면 의문을 가져야 한다.
// - 특정 타입이 유일한 목적이 공통 상태를 공유하는 것이라면 의문을 가져야 한다.
