// -------------------------------------
// Decoupling With Interface Composition
// -------------------------------------

//        System                                ps
//  --------------------                      ---------
// |  _________         |-pull               |         |-pull
// | |         |        |-store              | *System |-store
// | | *Xenia  |-pull   |                    |         |
// | |         |        | <------------------ ---------
// |  ---------         |      p             |         |
// | |         |        |    -----           |    *    |
// | |    *    |------- |-> |     |-pull     |         |
// | |         |        |    -----            ---------
// |  ---------         |
// |
// |  __________        |
// | |          |       |
// | | * Pillar |-store |
// | |          |       |
// |  ----------        |      s
// | |          |       |    -----                         p                   s
// | |    *     |------ |-> |     |-store               ---------           ---------
// | |          |       |    -----                     |         |-pull    |         |-store
// |  ----------        |                              | *System |         | *System |
//  --------------------                               |         |         |         |
//          A                                           ---------           ---------
//          |                                          |         |         |         |
//           ------------------------------------------|    *    | ------- |    *    |
//                                                     |         |         |         |
//                                                      ---------           ---------

package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Data 는 복사 할 데이터의 구조체이다.
type Data struct {
	Line string
}

// Puller 는 데이터 추출 동작을 선언한다.
type Puller interface {
	Pull(d *Data) error
}

// Storer 는 데이터 저장 동작을 선언한다.
type Storer interface {
	Store(d *Data) error
}

// PullStorer 는 추출 및 저장에 대한 동작을 선언한다.
type PullStorer interface {
	Puller
	Storer
}

// Xenia 는 우리가 데이터를 빼내야(pull data) 하는 시스템이다.
type Xenia struct {
	Host    string
	Timeout time.Duration
}

// Pull 은 Xenia 에서 데이터를 가져오는 방법을 알고 있다.
// 메서드의 파라미터로 (d *Data) 를 선언하면 Data 의 타입과 크기를 미리 알 수 있다. 따라서 스택에 저장할 수 있게 된다.
func (*Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF

	case 5:
		return errors.New("Error reading data from Xenia")

	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

// Pillar 는 우리가 데이터를 저장 해야 하는 시스템이다.
type Pillar struct {
	Host    string
	Timeout time.Duration
}

// Store 은 Pillar 에 데이터를 저장하는 방법을 알고 있다.
func (*Pillar) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// System 은 Pullers 와 Stores 를 하나의 시스템으로 결합한다.
type System struct {
	Puller
	Storer
}

// pull 은 Puller 로 부터 많은 양의 데이터를 추출하는 방법을 알고 있다.
func pull(p Puller, data []Data) (int, error) {
	// 데이터 슬라이스를 순회하면서 각각의 원소를 Puller 의 Pull 메서드에 전달한다.
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// store 은 Storer 로 부터 많은 양의 데이터를 저장하는 방법을 알고 있다.
func store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Copy 는 System 에서 데이터를 추출하고 저장하는 방법을 알고 있다.
// Puller 에서 Storer 로 데이터를 전달한다.
func Copy(ps PullStorer, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(ps, data)
		if i > 0 {
			if _, err := store(ps, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

func main() {
	sys := System{
		Puller: &Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Storer: &Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}
