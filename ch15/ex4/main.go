package main

import (
	"os"
	"bufio"
	"fmt"
	"math/rand"
	"time"
)

const (
	// 초기 잔액
	Balance = 1000
	// 맞췄을 때 버는 양
	EarnPoint = 500
	// 틀렸을 때 잃는 양
	LosePoint = 100
	// 게임 승리 시 포인트
	VictoryPPoint = 5000
	GameoverPoint = 0
)


var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)

	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err

}


func main () {
	rand.Seed(time.Now().UnixNano())
	balance := Balance

	for {
		fmt.Print("1~5 사이의 값을 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else if n < 1 || n > 5 {
			fmt.Println("1~5 사이의 값만 출력하세요.")
		} else {
			r := rand.Intn(5) + 1
			if n == r {
				balance += EarnPoint
				fmt.Println("축하합니다. 맞추셨습니다. 남은 돈:", balance)
				if balance >= VictoryPPoint {
					fmt.Println("게임 승리")
					break
				}
			} else {
				balance -= LosePoint
				fmt.Println("꽝 아쉽지만 다음 기회를.. 남은 돈:", balance)
				if balance <= GameoverPoint {
					fmt.Println("게임 종료")
					break
				}
			}
		}
	}
	
}