package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}

	// slice2 := make([]int, 3, 10) // 3이 최대치 주의 
	slice2 := make([]int, len(slice1))
	slice3 := make([]int, 10)

	cont1 := copy(slice2, slice1) // slice1 -> slice2 복사
	cont2 := copy(slice3, slice1) // slice1 -> slice3에 복사

	fmt.Println(cont1, slice2)
	fmt.Println(cont2, slice3)
}