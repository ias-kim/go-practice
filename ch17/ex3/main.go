package main

import "fmt"

type account struct {
	balance int
	firstName string
	lastName string
}

// 포인터 메서드
func (a1 *account) widthdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 타입 메서드
func (a2 account) widthdrawValue(amount int) account {
	a2.balance -= amount
	return a2
}

// 변경된 값을 변환하는 값 타입 메서드
func (a3 account) widthdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{ 100, "Joe", "Park" }
	mainA.widthdrawPointer(30)	// 포인터 메서드 호출
	fmt.Println(mainA.balance)

	mainA.widthdrawValue(20)	// 값 타입 메서드 호출
	fmt.Println(mainA.balance)

	var mainB account = mainA.widthdrawReturnValue(20)
	fmt.Println(mainB.balance)

	mainB.widthdrawPointer(30) // 포인터 메서드 호출
	fmt.Println(mainB.balance)
}