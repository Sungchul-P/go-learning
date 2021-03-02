package main

import "fmt"

// 프로그램에서 사용할 user를 정의한다.
type user struct {
	name    string
	surname string
}

func main() {
	// ----------------------
	// Declare and initialize
	// ----------------------

	// string 타입을 키로, user 타입을 값으로 갖는 맵을 선언하고 만든다.
	users1 := make(map[string]user)

	// 맵에 키/값 쌍을 추가한다.
	users1["Roy"] = user{"Rob", "Roy"}
	users1["Ford"] = user{"Henry", "Ford"}
	users1["Mouse"] = user{"Mickey", "Mouse"}
	users1["Jackson"] = user{"Michael", "Jackson"}

	// ----------------
	// Iterate over map
	// ----------------

	// map을 순회한다.
	fmt.Printf("\n=> Iterate over map\n")
	for key, value := range users1 {
		fmt.Println(key, value)
	}

	// ------------
	// Map literals
	// ------------

	// 초기값을 갖는 맵을 선언하고 초기화한다.
	users2 := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// map을 순회한다.
	fmt.Printf("\n=> Map literals\n")
	for key, value := range users2 {
		fmt.Println(key, value)
	}

	// ----------
	// Delete key
	// ----------

	delete(users2, "Roy")

	// --------
	// Find key
	// --------

	// 키 Roy를 찾아보자.
	// 만약 키 중에 Roy가 존재한다면, 그에 해당하는 값을 가져와 할당한다.
	// 그렇지 않다면 u는 여전히 user 타입의 값을 가지겠지만, 그 값은 제로 값으로 설정된다.
	u1, found1 := users2["Roy"]
	u2, found2 := users2["Ford"]

	// 값과 키의 존재 여부를 나타낸다.
	fmt.Printf("\n=> Find key\n")
	fmt.Println("Roy", found1, u1)
	fmt.Println("Ford", found2, u2)

	// --------------------
	// Map key restrictions
	// --------------------

	type users []user
	// 이 구문을 사용하여 users 를 새로 정의할 수 있으며, 이는 users를 정의하는 두 번째 방법이다.
	// 이처럼 이미 존재하는 타입을 통해, 다른 타입의 타입으로 사용할 수 있다. 이 때 두 타입은 서로 연관성이 없다.
	// 하지만 다음의 코드와 같이 키로서 사용하고자 할 때,
	// u := make(map[users]int)
	// 컴파일러는 다음의 오류를 발생시킨다.
	// Invalid map key type: the comparison operators == and != must be fully defined for key type
	//
	// 그 이유는, 키로 어떤 것을 사용하던지 그 값은 반드시 비교 가능해야 하기 때문이다. 맵이 키의 해시 값을 만들 수 있는 지 보여주는 일종의 불리언 표현식을 사용해야한다.
}
