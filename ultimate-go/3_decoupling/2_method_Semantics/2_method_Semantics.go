// ---------------------------
// Value and Pointer semantics
// ---------------------------

// 숫자, 문자열, 불과 같은 기본 타입을 사용하는 경우, 값에 의한 호출을 사용하는 것을 권장한다.
// 만약 정수 또는 불 타입 변수의 주소 값을 사용해야 한다면 주의해야 한다. 상황에 따라, 이 방법이 맞을 수도 있고, 아닐 수도 있기 때문이다.
// 하지만 일반적으로, 메모리 누수의 위험이 있는 힙 메모리 영역에 이 변수들을 만들 필요는 없다.
// 그래서 이 타입의 변수들은 stack에 생성하는 것을 더 권장한다.
// 모든 것에는 예외가 있을 수 있지만, 그 예외를 적용하는 것이 적합하다고 판단하기 전에는 규칙을 따를 필요가 있다.

// 슬라이스, 맵, 채널, 인터페이스와 같이 참조 타입의 변수들 역시 기본적으로 값에 의한 호출을 사용하는 것을 권장한다.
// 다만 변수의 주소 값을 파라미터로 받는 Unmarshal같은 함수를 사용하기 위한 경우라면, 이 타입들의 주소 값을 사용해야 한다.

// 아래 예제들은 실제 Go의 표준 라이브러리에서 사용하는 코드들이다.
// 이들을 공부해보면, 값에 의한 호출과 참조에 의한 호출(pointer semantic) 중 하나를 일관되게 사용하는 것이 얼마나 중요한지 알 수 있다.
// 따라서 변수의 타입을 정할 때, 다음의 질문에 스스로 답해보자.

// - 이 타입은 값에 의한 호출과 참조에 의한 호출 중 어느 것이 더 적합한가?
// - 만약 이 변수의 값을 변경해야 한다면, 새로운 값을 가지는 복사본을 만드는 것과,
// 	 다른 곳에서도 이 변경된 값을 확인할 수 있게 이 변수의 값을 직접 변경하는 것 중 어느 것이 더 적합한가?
// 가장 중요한 것은 일관성이다. 처음에 한 결정이 잘못되었다고 판단되면, 그때 이를 변경하면 된다.

package main

import (
	"sync/atomic"
	"syscall"
)

// --------------
// Value semantic
// --------------

// Go의 net 패키지는 IP와 IPMask 타입을 제공하는데, 이들은 바이트 형 슬라이스이다.
// 아래의 예제들은 이 참조 타입들을 값에 의한 호출을 통해 사용하는 것을 보여준다.
type IP []byte
type IPMask []byte

// Mask는 IP 타입의 값 리시버를 사용하며, IP 타입의 값을 반환한다.
// 이 메서드는 IP 타입에 대해서 값에 의한 호출을 사용하는 것이다.
func (ip IP) Mask(mask IPMask) IP {
	if len(mask) == IPv6len && len(ip) == IPv4len && allFF(mask[:12]) {
		mask = mask[12:]
	}
	if len(mask) == IPv4len && len(ip) == IPv6len && bytesEqual(ip[:12], v4InV6Prefix) {
		ip = ip[12:]
	}
	n := len(ip)
	if n != len(mask) {
		return nil
	}
	out := make(IP, n)
	for i := 0; i < n; i++ {
		out[i] = ip[i] & mask[i]
	}
	return out
}

// ipEmptyString 은 IP 타입의 값을 파라미터로 받고, 문자열 타입의 값을 반환한다.
// 이 함수는 IP 타입에 대해서 값에 의한 호출을 사용하는 것이다.
func ipEmptyString(ip IP) string {
	if len(ip) == 0 {
		return ""
	}
	return ip.String()
}

// ----------------
// Pointer semantic
// ----------------

// Time 타입은 값에 의한 호출과 참조에 의한 호출 중 어떤 것을 사용해야 할까?
// 만약 Time 타입 변수의 값을 변경해야 한다면, 이 값을 직접 변경해야 할까? 아니면 복사본을 만들어서 값을 변경하는 것이 좋을까?
type Time struct {
	sec  int64
	nsec int32
	loc  *Location
}

// 타입에 대해 어떤 호출을 사용할지를 결정하는 가장 좋은 방법은 타입의 생성 함수를 확인하는 것이다.
// 이 생성 함수는 어떤 호출을 사용해야 하는지 알려준다. 이 예제에서, Now 함수는 Time 타입의 값을 반환한다.
// Time 타입의 값은 한번 복사가 이루어지고, 이 복사된 값은 이 함수를 호출한 곳으로 반환된다.
// 즉, 이 Time 타입의 값은 스택(stack)에 저장된다. 따라서 값에 의한 호출을 사용하는 것이 좋다.
func Now() Time {
	sec, nsec := now()
	return Time{sec + unixToInternal, nsec, Local}
}

// Add는 기존의 Time 타입의 값과는 다른 값을 얻기 위한 메서드다. 만약 값을 변경할 때는 무조건 참조에 의한 호출을 사용하고,
// 그렇지 않을 때는 값에 의한 호출을 사용해야 한다면 이 Add의 구현은 잘못되었다고 생각할 수 있다.
// 하지만, 타입은 어떤 호출을 사용할지를 책임지는 것이지, 메서드의 구현을 책임지는 것은 아니다.
// 메서드는 반드시 선택된 호출을 따라야 하고, 그래서 Add의 구현은 틀리지 않았다.

// Add는 값 리시버를 사용하고, Time 타입의 값을 반환한다.
// 즉, 이 메서드는 실제로 Time 타입 변수의 복사본을 변경하며, 완전히 새로운 값을 반환하는 것이다.
func (t Time) Add(d Duration) Time {
	t.sec += int64(d / 1e9)
	nsec := int32(t.nsec) + int32(d%1e9)
	if nsec >= 1e9 {
		t.sec++
		nsec -= 1e9
	} else if nsec < 0 {
		t.sec--
		nsec += 1e9
	}
	t.nsec = nsec
	return t
}

// div는 Time 타입의 파라미터를 받고, 기본 타입의 값들을 반환한다. 이 함수는 Time 타입에 대해 값에 의한 호출을 사용하는 것이다.
// func div(t Time, d Duration) (qmod2 int, r Duration) {}

// Time 타입에 대한 참조에 의한 호출은, 오직 주어진 데이터를 Time 타입으로 변환해 이 메서드를 호출한 변수를 수정할 할 때만 사용한다:
// func (t *Time) UnmarshalBinary(data []byte) error {}
// func (t *Time) GobDecode(data []byte) error {}
// func (t *Time) UnmarshalJSON(data []byte) error {}
// func (t *Time) UnmarshalText(data []byte) error {}

// Observation:
// ------------
// 대부분의 구조체(struct) 타입들은 값에 의한 호출을 잘 사용하지 않는다.
// 이는 다른 코드에서 함께 공유하거나 또는 공유하면 좋은 데이터들이기 때문이다.
// User 타입이 대표적인 예이다. User 타입의 변수를 복사하는 것은 가능은 하지만, 이는 실제로 좋은 구현이 아니다.

// Other examples:
// 앞서 이야기했듯이, 생성 함수는 어떤 호출을 사용해야 하는지를 알려준다.
// Open 함수는 File 타입 데이터의 주소 값, 즉 File 타입 포인터를 반환한다.
// 이는 File 타입에 대해, 참조에 의한 호출을 사용해서 이 File 타입의 값을 공유할 수 있다는 것을 뜻한다.
func Open(name string) (file *File, err error) {
	return OpenFile(name, O_RDONLY, 0)
}

// Chdir은 File 타입의 포인터 리시버를 사용한다. 즉, 이 메서드는 File 타입에 대해 참조에 의한 호출을 사용하는 것이다.
func (f *File) Chdir() error {
	if f == nil {
		return ErrInvalid
	}
	if e := syscall.Fchdir(f.fd); e != nil {
		return &PathError{"chdir", f.name, e}
	}
	return nil
}

// epipecheck는 File 타입의 포인터를 파라미터로 받는다. 따라서 이 함수는 File 타입에 대해 참조에 의한 호출을 사용하는 것이다.
func epipecheck(file *File, e error) {
	if e == syscall.EPIPE {
		if atomic.AddInt32(&file.nepipe, 1) >= 10 {
			sigpipe()
		}
	} else {
		atomic.StoreInt32(&file.nepipe, 0)
	}
}
