package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// user 구조체는 user 정보를 담고 있다.
type user struct {
	ID   int
	Name string
}

// updateStats 구조체는 업데이트 정보를 담고 있다.
type updateStats struct {
	Modified int
	Duration float64
	Success  bool
	Message  string
}

func main() {
	// user 프로필을 가져온다.
	u, err := retrieveUser("Devnori")
	if err != nil {
		fmt.Println(err)
		return
	}

	// user 프로필을 보여준다. `u`는 주소값이기에 *를 사용하여 값을 얻어낸다.
	fmt.Printf("%+v\n", *u)

	// user 의 name 을 업데이트 한다.
	// _(blank identifier)를 사용하여 리턴된 updateStats는 무시하며
	// if 범위 밖에서 사용할 값은 없으니 간결한 문법을 사용하였다.
	if _, err := updateUser(u); err != nil {
		fmt.Println(err)
		return
	}

	// 업데이트가 성공했다고 출력한다.
	fmt.Println("Updated user record for ID", u.ID)
}

// 특정한 사용자의 문서를 가져온다.
func retrieveUser(name string) (*user, error) {
	// getUser 함수를 호출하여 JSON 형식의 user를 전달 받는다.
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}

	// JSON 값을 unmarshal하여 저장할 user 타입인 변수 u를 생성한다.
	var u user

	// 변수 u를 json.Unmarshal 함수에 전달하면, 함수는 r로부터 JSON 을 읽어서 변수 u에 넣어준다.
	err = json.Unmarshal([]byte(r), &u)

	// retrieveUser 함수를 호출한 함수에게 u값을 전달한다. 이처럼 retrieveUser 함수에서
	// 생성한 변수의 주소값을 호출한 함수에게 전달하기에 이 변수는 힙 메모리에 할당된다.
	return &u, err
}

// 특정한 사용자에 대한 JSON 응답
func getUser(name string) (string, error) {
	response := `{"ID":101, "Name":"Devnori"}`
	return response, nil
}

// 특정 사용자가 업데이트 되었다는 응답
func updateUser(u *user) (*updateStats, error) {
	// response 변수는 JSON 응답을 시뮬레이션 한 것이다.
	response := `{"Modified":1, "Duration":0.005, "Success" : true, "Message": "updated"}`

	// JSON 문서를 userStats 구조체 타입의 변수로 unmarshal한다.
	var us updateStats
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	// update 성공여부를 확인한다.
	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}
