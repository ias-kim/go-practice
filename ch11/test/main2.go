package main

import (
	"fmt"
	"unsafe"
)

type Padding struct {
	A int8
	G int8
	D uint16
	F float32
	B int
	C float64
	E int
}

func main() {
	user := Padding{ 1, 2, 3, 4, 5, 6, 7 }
	fmt.Println(unsafe.Sizeof(user))
}