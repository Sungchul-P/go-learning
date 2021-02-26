package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

type validationContextKey string

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

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	// 유효한 요청이 있을 때 이 요청에 대한 새 컨텍스트를 만든 다음,
	// 요청의 Name 필드 값을 컨텍스트에 설정한다.
	// 컨텍스트를 사용하는 또 다른 패키지와의 충돌을 피하기 위해 명시적으로 선언된 문자열로 사용한다. (validationContextKey)
	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

type helloWorldHandler struct {
}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// 값을 가져오려면 컨텍스트를 가져온 다음 Value 메서드를 호출해 문자열로 변환하면 된다.
	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloWorldResponse{Message: "Hello" + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func fetchGoogle(t *testing.T) {
	r, _ := http.NewRequest("GET", "https://google.com", nil)

	// 시간 초과(timeout) 컨텍스트를 만든다.
	// 컨텍스트가 자동으로 취소되는 인바운드 요청과 달리 아웃바운드 요청에서는 취소 단계를 수동으로 수행해야 한다.
	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)
	defer cancelFunc()

	// func (r *Request) WithContext(ctx context.Context) *Request
	// ctx 컨텍스트로 원본 요청의 컨텍스트를 변경한 얕은 복사본(shalow copy)을 리턴
	// 이 함수를 실행하면 1밀리초 후 에러가 발생하면서 요청이 끝난다.
	r = r.WithContext(timeoutRequest)

	// 컨텍스트는 요청이 완료되기 전에 시간 초과되며, do 메서드는 즉시 리턴된다.
	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
