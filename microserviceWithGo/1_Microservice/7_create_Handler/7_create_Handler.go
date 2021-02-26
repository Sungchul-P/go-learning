package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080

	handler := newValidationHandler(newHelloWorldHandler())

	http.Handle("/helloworld", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// Handler 인터페이스의 메서드를 구현할 구조체 필드를 정의
type validationHandler struct {
	next http.Handler
}

// 새 핸들러를 리턴하는 함수 추가
func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	// 요청을 디코딩하는 과정에서 에러가 리턴되면 응답에 500 에러가 기록되며,
	// 핸들러 체인이 여기에서 중단된다.
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	// 체인의 다음 핸들러 호출
	h.next.ServeHTTP(rw, r)
}

// 응답을 작성하는 helloWorldHandler 타입
type helloWorldHandler struct{}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello"}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
