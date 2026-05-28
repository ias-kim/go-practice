package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

// 사용자 입력 받기
func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// 랜덤값 생성
	r := rand.Intn(100)
	count := 1


	for {
		fmt.Printf("숫자값을 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			// - 사용자 입력이 더크면 더 크다고, 작으면 작다고 출력
			if n > r {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < r {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				// - 숫자가 맞으면 메시지 출력후 시도횟수 출력
				fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수:", count)
				break
			}
			count++
		}
	}
}


