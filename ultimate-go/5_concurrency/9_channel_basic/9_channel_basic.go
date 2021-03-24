package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("\n=> Double signal\n")
	signalAck()

	closeRange()

	fmt.Printf("\n=> Select and receive\n")
	selectRecv()

	fmt.Printf("\n=> Select and send\n")
	selectSend()

	fmt.Printf("\n=> Select and drop\n")
	selectDrop()
}

// ---------------------------------
// Unbuffered channel: Double signal
// ---------------------------------

func signalAck() {
	ch := make(chan string)

	// 이 고루틴은 신호가 전달되기 전까지 멈춘(block)다. 즉 우리가 신호를 받기 전까지 이 고루틴은 움직일 수 없다.
	go func() {
		fmt.Println(<-ch)
		ch <- "ok done"
	}()

	ch <- "do this"
	fmt.Println(<-ch)
}

// ---------------------------------
// Buffered channel: Close and range
// ---------------------------------

func closeRange() {
	// This is a buffered channel of 5.
	ch := make(chan int, 5)

	// Populate with value
	for i := 0; i < 5; i++ {
		ch <- i
	}

	// Close the channel.
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

// --------------------------------------
// Unbuffered channel: select and receive
// --------------------------------------

func selectRecv() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "work"
	}()

	// select 문을 써서 값을 받을 때까지 특정 시간만큼을 기다리는 방식
	select {
	case v := <-ch:
		fmt.Println(v)
	// 특정 시간만큼을 기다리고 신호 보내기를 수행한다.
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}

	// 위와 같은 코드를 작성한 후 고루틴에게 종료될 기회를 주지 않으면 버그가 발생한다.
	// 고루틴이 시간 초과되고 다음으로 넘어가면 대응하는 신호 받기(ch <-)가 없으므로 고루틴 누수(leak)가 생긴다.
	// 즉, 이 고루틴은 결코 종료되지 않는다.
}

// -----------------------------------
// Unbuffered channel: select and send
// -----------------------------------

func selectSend() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		fmt.Println(<-ch)
	}()

	// 이 방식도 고루틴 누수가 발생할 수 있다. (버퍼 채널을 사용해야 함)
	select {
	case ch <- "work":
		fmt.Println("send work")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timed out")
	}
}

// ---------------------------------
// Buffered channel: Select and drop
// ---------------------------------

func selectDrop() {
	ch := make(chan int, 5)

	go func() {
		// 여기서 신호 받기(ch <-) 루프에서 작업할 데이터가 오기를 기다리고 있다.
		for v := range ch {
			fmt.Println("recv", v)
		}
	}()

	// 작업을 채널에 보낼 것이다.
	// 버퍼가 다 차면, 버퍼는 block 되고 default case 가 실행되며 작업을 버리게 된다.
	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
			fmt.Println("send work", i)
		default:
			fmt.Println("drop", i)
		}
	}

	close(ch)
}
