package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for { // 무한 루프
		fmt.Println("입력하세요.")
		var number int
		_, err := fmt.Scanln(&number) // 한 줄 입력을 받습니다.
		if err != nil { // 숫자가 아닌 경우
			fmt.Println("숫자를 입력하세요") 

			stdin.ReadString('\n') // 키보드 버퍼를 지워줌
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
		if number%2 == 0 {
			break
		}
	}
	fmt.Println("for문이 종료했습니다.")
}