package main
import (
	"fmt"
	"container/list"
)

type Queue struct { // Queue 구조체 정의
	v *list.List
}

func NewQueue() *Queue {
	return &Queue{ list.New() }
}

func (q *Queue) Push(val interface{}) { // 요소 추가
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func main() {
	queue := NewQueue() // 새로운 큐 생성

	for i := 1; i < 5; i++ {
		queue.Push(i) // 요소 입력
	}

	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v ->" , v)
		v = queue.Pop()
	}
}