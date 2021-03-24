// ---------------------------------
// Unbuffered channel (Tennis match)
// ---------------------------------

// 고루틴 2개를 테니스 매치에 넣을 것이다.
// 공이 양쪽 편에서 쳐지거나 놓쳤다는 보장을 필요로 하기 때문에 버퍼 없는 채널을 쓴다.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 버퍼 없는 채널을 생성한다.
	court := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// 선수 둘을 입장시킨다. 둘 모두 받기 모드에서 시작할 것이다.
	// 어떤 선수가 먼저 공을 받게 될 지는 알 수 없다.
	// 메인 고루틴을 심판이라고 생각하자. 누가 공을 먼저 받는지는 심판에게 달려 있다.
	go func() {
		player("Sungchul", court)
		wg.Done()
	}()

	go func() {
		player("Sean", court)
		wg.Done()
	}()

	// 테니스 경기를 시작한다. 메인 고루틴이 신호 보내기를 한다.
	// 선수 둘 모두 신호 받기 모드이므로 어느 쪽이 먼저 공을 받게 될지 알 수 없다.
	court <- 1

	// 경기가 끝날 때까지 기다린다.
	wg.Wait()
}

// player 는 테니스 경기를 하는 사람을 흉내낸다.
func player(name string, court chan int) {
	for {
		// 공이 다시 넘어올 때까지 기다린다. 이게 또 다른 형태의 신호 받기를 점을 놓치지 말자.
		// 단순히 값을 받는 대신에 신호 받기가 어떻게 반환되었는지 나타내는 플래그를 받을 수 있다.
		// 만약 신호가 데이터 때문에 생겼다면 ok는 true 일 것이다. 채널이 닫혔다면 ok는 false다.
		// 이를 통해 누가 이겼는지 결정할 수 있다.
		ball, ok := <-court

		// 채널이 닫혔다면 이긴 것이다.
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 무작위 값을 하나 정해서 공을 놓쳤는지 알아본다. 경기에서 진다면 채널을 닫는다.
		// 반대편 player는 데이터 없이 신호를 받았다는 것을 알게될 것이다.
		// 채널은 닫히고 반대편 player가 이기게 된다.
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// 졌다는 신호를 보내기 위해 채널을 닫는다.
			close(court)
			return
		}

		// 공을 친 횟수를 출력하고, 하나 증가시킨다.
		// 반대편 player는 여전히 신호 받기 모드에 있다는 것을 알고 있다.
		// 그러므로 신호를 보내는 쪽과 받는 쪽은 결국 함께 모일 것이다.
		// 버퍼 없는 채널에서는 전달이 보장되기 때문에 신호 받기가 먼저 일어난다.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// 공을 쳐서 상대 선수에게 다시 보낸다.
		court <- ball
	}
}
