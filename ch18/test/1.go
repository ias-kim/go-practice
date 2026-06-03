package main

type ReadWriter interface {
	Read()
	Write()
}

type File struct {

}

func (f *File) Read() {

}

// func (f *File) Write() {

// }

func ReadWrite(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

func main() {
	f := &File{}
	ReadWrite(f) // 컴파일 오류 발생
}