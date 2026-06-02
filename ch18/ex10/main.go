package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {

}

func (f *File) Read() {

}

func ReadFile(reader Reader) {
	// Reader 인터페이스 변수를 Closer 인터페이스 타입 변환
	// 런 타임 에러가 발생
	c := reader.(Closer)
	c.Close()
}

func main() {
	// File 포인터 인스턴스를 ReadFile() 함수의 인수로 사용한다.
	file := &File{}
	ReadFile(file)
}