package main

import "fmt"

func main() {
	var a int = 500
	var p *int

	var f *float32

	p = &a

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가르키는 메모리의 값: %d\n", *p)
	fmt.Printf("f %d\n", f)

	*p = 100
	fmt.Printf("a의 값: %d\n", a)
}