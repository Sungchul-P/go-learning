package main

import "fmt"

func main() {

	// 인덱스 0의 문자열은 문자들을 실제로 저장하고 있는 배열에 대한 포인터와 길이 정보 5를 가지게 된다.
	var strings [5]string

	//  -----         -------------------
	// |  *  |  ---> | A | p | p | l | e | (1)
	//  -----         -------------------
	// |  5  |                  A
	//  -----                   |
	//                          |
	//                          |
	//     ---------------------
	//    |
	//  -----------------------------
	// |  *  | nil | nil | nil | nil |
	//  -----------------------------
	// |  5  |  0  |  0  |  0  |  0  |
	//  -----------------------------
	strings[0] = "Apple"
	strings[1] = "Orange"
	strings[2] = "Banana"
	strings[3] = "Grape"
	strings[4] = "Plum"

	// ---------------------------------
	// 문자열 배열 반복 (Iterate over the array of strings)
	// ---------------------------------

	// 문자열의 길이를 알고 있으니 스택에 둘 수 있고, 그 덕분에 힙에 할당하여 GC를 해야하는 부담을 덜게 된다.
	fmt.Printf("\n=> Iterate over array\n")
	for i, fruit := range strings {
		fmt.Println(i, fruit)
	}

	// 리터럴 표기법(literal syntax)으로 값 초기화
	numbers := [4]int{10, 20, 30, 40}

	fmt.Printf("\n=> Iterate over array using traditional style\n")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	// ---------------------
	// 다른 타입의 배열들 (Different type arrays)
	// ---------------------

	var five [5]int

	four := [4]int{10, 20, 30, 40}

	fmt.Printf("\n=> Different type arrays\n")
	fmt.Println(five)
	fmt.Println(four)

	// -----------------------------
	// 연속 메모리 할당 (Contiguous memory allocations)
	// -----------------------------

	six := [6]string{"Annie", "Betty", "Charley", "Doug", "Edward", "Hoanh"}

	fmt.Printf("\n=> Contiguous memory allocations\n")
	for i, v := range six {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &six[i])
	}
	// => Contiguous memory allocations
	// Value[Annie]    Address[0xc00008e230] IndexAddr[0xc0000b0120]
	// Value[Betty]    Address[0xc00008e230] IndexAddr[0xc0000b0130]
	// Value[Charley]  Address[0xc00008e230] IndexAddr[0xc0000b0140]
	// Value[Doug]     Address[0xc00008e230] IndexAddr[0xc0000b0150]
	// Value[Edward]   Address[0xc00008e230] IndexAddr[0xc0000b0160]
	// Value[Hoanh]    Address[0xc00008e230] IndexAddr[0xc0000b0170]

}
