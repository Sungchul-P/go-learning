package main

import "fmt"

type user struct {
	name  string
	email string
}

// notify 는 값 리시버(value receiver)를 가지는 메서드(method)이다.
// u는 user 타입으로, Go 에서는 함수가 리시버와 함께 선언된다면 이를 메서드라고 한다.
// 리시버는 파라미터와 비슷하게 보이지만, 이는 자신만의 역할이 있다.
// 값 리시버를 사용하면, 메서드는 자신을 호출한 변수를 복사하고, 그 복사본을 가지고 동작한다.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail은 포인터 리시버를 가지는 메서드이다: u는 user의 포인터 타입으로,
// 포인터 리시버를 이용하면 메서드를 호출한 변수를 공유하면서 바로 접근이 가능하다.
func (u *user) changeEmail(email string) {
	u.email = email
	fmt.Printf("Changed User Email To %s\n", email)
}

// 위의 두 메서드들은 값 리시버와 포인터 리시버의 차이를 이해하기 위해서 같이 사용되었다.
// 하지만 실제 개발에서는 하나의 리시버를 사용하는 것을 권장한다.

func main() {
	// -------------------------------
	// Value and pointer receiver call
	// -------------------------------

	// user 타입의 변수는 값 리시버와 포인터 리시버를 사용하는 모든 메서드를 호출할 수 있다.
	psc := user{"PSC", "psc@email.com"}
	psc.notify()
	psc.changeEmail("psc@hotmail.com")

	// user 의 포인터 타입 변수 역시 값 리시버와 포인터 리시버를 사용하는 모든 메서드를 호출할 수 있다.
	devnori := &user{"Devnori", "devnori@email.com"}
	devnori.notify()
	devnori.changeEmail("devnori101@gmail.com")

	// 이 예제에서 devnori 는 user 타입을 가리키는 포인터 변수이다. 하지만 값 리시버로 구현된 notify를 호출할 수 있다.
	// user 타입의 변수로 메서드를 호출하지만, Go는 내부적으로 이를 (*devnori).notify()로 호출한다.
	// Go는 devnori 가 가리키는 값을 찾고, 이 값을 복사하여 notify를 값에 의한 호출(value semantic)이 가능하도록 한다.

	// 이와 유사하게 psc는 user 타입의 변수이지만, 포인터 리시버로 구현된 changeEmail을 호출할 수 있다.
	// Go는 psc의 주소를 찾고, 내부적으로 (&psc).changeEmail()을 호출한다.

	// 두 개의 user 타입 원소를 가진 슬라이스를 만들자.
	users := []user{
		{"psc", "psc@email.com"},
		{"devnori", "devnori@email.com"},
	}

	// 이 슬라이스에 for ... range를 사용하면, 각 원소의 복사본을 만들고, notify 호출을 위해 또 다른 복사본을 만들게 된다.
	for _, u := range users {
		u.notify()
	}

	// 포인터 리시버를 사용하는 changeEmail을 for ... range 안에서 사용해보자. 이는 복사본의 값을 변경하는 것으로 이렇게 사용하면 안 된다.
	for _, u := range users {
		u.changeEmail("sungchul-p@github.com")
	}
}
