// -------------------------
// Buffered channel: Fan Out
// -------------------------

// 아이디어는 다음과 같다. 고루틴이 할일을 하다가 많은 데이터베이스 작업을 실행하기로 결정했다고 해보자.
// 이 고루틴은 그 일을 하기 위해 새로운 고루틴을 여러개, 예컨대 10개를 만들어낼 것이다.
// 각 고루틴은 데이터베이스 작업 2개를 수행할 것이다. 결국 데이터베이스 작업 20개가 고루틴 10개로 쪼개진다.
// 다시 말하면 원래 있던 고루틴이 고루틴 10개를 팬 아웃(fan out)하고 생성한 모든 고루틴이 작업을 끝내고 보고하기를 기다리는 것이다.
//
// 여기서는 버퍼 있는 채널이 딱인데 이는 우리가 미리 20개 작업을 수행하는 고루틴이 10개 있을 것이라는 점을 알기 때문이다.
// 따라서 버퍼 크기는 20이다. 우리는 결국 마지막에는 작업 신호를 받아야 한다는 점을 알기 때문에 어떤 작업 신호도 막힐 염려가 없다.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type result struct {
	id  int
	op  string
	err error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 고루틴 수와 연산 수를 설정한다.
	const routines = 10
	const inserts = routines * 2

	// 어떤 입력에 대해서든 정보를 받기 위해 버퍼 있는 채널을 연다.
	ch := make(chan result, inserts)

	// 우리가 처리해야 할 응답의 개수이다.
	// 이 고루틴은 자신의 스택 공간(stack space)를 관리할 수 있기 때문에
	// WaitGroup 을 쓰는 대신 지역 변수(local variable)을 WaitGroup 처럼 쓸 것이다.
	// 그러므로 시작하자마자 이 지역 변수를 입력 20개로 설정한다.
	waitInserts := inserts

	// 모든 입력을 수행한다. 이것이 바로 팬 아웃이다. 이제 우리에게는 고루틴이 10개 있다.
	// 각 고루틴은 입력 두 개를 수행한다. 입력의 결과는 ch 채널에서 쓰인다.
	// 이 채널은 버퍼 있는 채널이기 때문에 어떤 신호 보내기도 막히지(block) 않는다.
	for i := 0; i < routines; i++ {
		go func(id int) {
			ch <- insertUser(id)
			// 버퍼 있는 채널을 쓴 덕에 두번째 입력을 시작하기 위해 기다릴 필요가 없다. 첫번째 신호 보내기는 즉시 끝난다.
			ch <- insertTrans(id)
		}(i)
	}

	// 입력이 끝나면 입력 결과를 처리한다.
	for waitInserts > 0 {
		// 고루틴에서 오는 응답을 기다린다. 이건 신호 받기다. 한번에 결과 하나씩을 받고 waitInserts을 0이 될 때까지 줄일 것이다.
		r := <-ch
		log.Printf("N: %d ID: %d OP: %s ERR: %v", waitInserts, r.id, r.op, r.err)
		waitInserts--
	}
	log.Println("Inserts Complete")
}

func insertUser(id int) result {
	r := result{
		id: id,
		op: fmt.Sprintf("insert USERS value (%d)", id),
	}

	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("Unable to insert %d into USER table", id)
	}
	return r
}

func insertTrans(id int) result {
	r := result{
		id: id,
		op: fmt.Sprintf("insert TRANS value (%d)", id),
	}

	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("Unable to insert %d into USER table", id)
	}
	return r
}
