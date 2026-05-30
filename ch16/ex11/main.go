package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	slice = append(slice, 0) // 맨 뒤에 요소 추가

	// 위치
	idx := 2

	println("iteration", len(slice), slice)
	// 맨 뒤부터 추가혀려는 위치까지 값을 하나씩 옮겨준다.
	for i := len(slice)-2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}

	slice[idx] = 100

	fmt.Println(slice)
}