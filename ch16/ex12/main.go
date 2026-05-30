package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 5, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s)
}