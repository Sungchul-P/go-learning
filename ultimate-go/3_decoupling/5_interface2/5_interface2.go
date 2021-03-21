package main

import "fmt"

// notifier 는 인터페이스로 알림을 보내는 타입들의 동작을 정의한다.
type notifier interface {
	notify()
}

// printer 역시 인터페이스로, 어떤 정보를 출력하는 동작을 정의한다.
type printer interface {
	print()
}

// user 는 프로그램 내의 사용자 정보를 담을 타입을 정의한다.
type user struct {
	name  string
	email string
}

// print 는 user 타입의 name 과 email 정보를 출력한다.
func (u user) print() {
	fmt.Printf("My name is %s and my email is %s\n", u.name, u.email)
}

// ------------------------------
// Interface via pointer receiver
// ------------------------------

// notify 는 포인터 리시버를 사용해서 notifier 인터페이스를 구현한다.
func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// String 은 fmt.Stringer 인터페이스를 구현한다.
// fmt 는 지금까지 화면에 정보를 출력하기 위해서 사용한 패키지로,
// String 을 구현하는 데이터가 주어지면 기존 동작이 아닌 새롭게 구현된 동작을 수행한다.
// 참조에 의한 호출을 사용하기 때문에, 포인터만이 이 인터페이스를 만족한다.
func (u *user) String() string {
	return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
}

func main() {
	// user 타입의 데이터를 만들어보자.
	u := user{"Devnori", "devnori@email.com"}

	// 참조에 의한 호출로 u를 전달하는 다형성 함수를 호출하자: sendNotification(u).
	// 하지만 컴파일러는 이 호출을 허락하지 않고 다음의 에러 메시지를 보여준다:
	// "cannot use u (type user) as type notifier in argument to sendNotification:
	// user does not implement notifier (notify method has pointer receiver)"
	// 이는 무결성을 위반해서 발생한 문제이다.

	// 포인터 리시버를 이용해서 인터페이스를 구현한다면, 반드시 참조에 의한 호출을 사용해야 한다.
	// 만약 값 리시버를 사용해서 인터페이스를 구현한다면, 값에 의한 호출과 참조에 의한 호출 모두를 사용할 수 있다.
	// 하지만, 일관성을 유지하기 위해서 이 경우에도 값에 의한 호출을 사용하는 것을 권장한다.
	// 다만 Unmarshal 과 같은 함수를 위해서는 참조에 의한 호출을 사용해야만 할 수도 있다.
	//
	// 이 문제를 해결하기 위해서는, u 대신 u의 주소 값(&u)를 전달해야 한다.
	// user 의 값을 만들고 이 값의 주소 값을 전달하면, 인터페이스는 user 타입의 주소 값을 가지게 되고 원본 값을 가리킬 수 있게 된다.
	//  -------
	// | *User |
	//  -------
	// |   *   | --> original user value
	//  -------
	sendNotification(&u)

	// 이와 유사하게, u의 값을 Println 에게 전달하면, 기존 형태의 출력을 보게 된다.
	// 하지만 만약 u의 주소 값을 전달하면, 앞서 정의한 String 으로 기존 동작을 덮어써서 새로운 형태의 출력을 사용하게 된다.
	fmt.Println(u)  // {Devnori devnori@email.com}
	fmt.Println(&u) // My name is "Devnori" and my email is "devnori@email.com"

	// ------------------
	// Slice of interface
	// ------------------

	// 인터페이스를 저장할 수 있는 슬라이스를 만들자.
	// 이 슬라이스에는 printer 인터페이스를 구현하는 모든 타입의 값이나 포인터를 저장할 수 있다.

	//   index 0   index 1
	//  -------------------
	// |   User  |  *User  |
	//  -------------------
	// |    *    |    *    |
	//  -------------------
	//      A         A
	//      |         |
	//     copy    original

	entities := []printer{
		// 이 슬라이스에 값을 저장하면, 인터페이스의 값은 복사본을 가지게 된다.
		// 따라서 원본 데이터에 변경이 발생하더라도, 이를 확인할 수는 없다.
		u,

		// 슬라이스에 포인터를 저장하면, 인터페이스의 값은 원본 데이터를 가리키는 주소 값을 복사하고 이를 가지게 된다.
		// 따라서 원본 데이터의 변경을 확인할 수 있다.
		&u,
	}

	// u의 name 과 email 을 변경해보자.
	u.name = "Sungchul Park"
	u.email = "sungchul.dev@gmail.com"

	// 슬라이스를 순회하면서 복사된 인터페이스의 값을 통해 print 를 호출해보자.
	for _, e := range entities {
		e.print()
	}
}

// sendNotification 은 다형성 함수이다.
// 이 함수는 notifier 인터페이스를 구현하는 타입의 값을 받아 알림을 보낸다.
// 이 함수는 다음과 같이 말하는 것이다: 나는 notifier 인터페이스를 구현하는 타입의 값 또는 포인터를 받을 것이다.
// 그리고 나는 인터페이스를 통해서 동작을 호출할 것이다.
func sendNotification(n notifier) {
	n.notify()
}
