package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age int32 // 8
	Score float64 // 8
}

func main() {
	user := User{ 23, 77.2 }
	fmt.Println(unsafe.Sizeof(user)) // 16
}