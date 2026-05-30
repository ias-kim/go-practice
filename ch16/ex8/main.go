package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1)) // slice1과 같은 길이의 슬라이스 생성
	
	for i, v := range slice1 { // slice1의 모든 요소값 복사
		slice2[i] = v
	}

	slice1[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
}