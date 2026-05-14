package main

import "fmt"

var g int = 10 // 패키지 전역 변수 선언

func main() {
	var m int = 20
	
	{	
		var s int = 50 // 지역변수선언
		fmt.Println(m, s, g)
	} // s 지역변수 없어짐
	m = s + 20 // Error

}