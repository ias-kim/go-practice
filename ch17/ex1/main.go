package main

import "fmt"

type account struct {
	balance int
}

func widthdrawFunc(a *account, amount int) { // 일반 함수 표현
	a.balance -= amount
}

func (a *account) widthdrawMethod(amount int) { // 메서드 표현
	a.balance -= amount
}

func main() {
	a := &account{ 100 } // balance가 100인 acount 포인터 변수 생성
	
	widthdrawFunc(a, 30) // 함수 형태 호출
	
	a.widthdrawMethod(30) // 메서드 형태 호출

	fmt.Printf("%d \n", a.balance)
}