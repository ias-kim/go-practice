package main

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	Data uintptr
	Len int
}

func main() {
	str1 := "Hello World!"
	str2 := str1

	stringHeader1 := (*StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}