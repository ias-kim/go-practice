package main

import "fmt"

func main() {
	str := "Hello\t'World'\n"
	
	str2 := `Go is "awesome!\nGo is simple and\t'powerful'`

	fmt.Println(str)
	fmt.Println(str2)
}