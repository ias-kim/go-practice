package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)

	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5)) // 인수 5개 사용
	fmt.Println(sum(10, 20))		// 인수 2개 사용
	fmt.Println(sum())				// 인수 0개 사용
}