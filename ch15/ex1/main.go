package main

import (
	"math/rand"
	"time"
	"fmt"
)


func main() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(100)
	fmt.Println(n)
}
