package main

import "fmt"

func main() {
	fmt.Println(3 * 4 ^ 7 << 2 + 3 * 5 == 7)
	// 우선 순위 소괄호로 묶는걸 추천함.
}