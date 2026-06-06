package main
import "fmt"

func main() {
	m := make(map[int]int) // 맵 생성
	m[1] = 0
	m[2] = 2
	m[3] = 3


	v, ok := m[1] 

	if ok {
		fmt.Println(v, "m1 값 존재")
	}

	delete(m, 3) // 요소 삭제
	delete(m, 4) // 없는 요소 삭제 시도

	fmt.Println(m[3])
	


	fmt.Println(m[1])
}