package main

import (
	"fmt"
	"errors"
	"bufio"
	"strings"
	"strconv"
)

func MultipleFormString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) // 스캐너 생성
	scanner.Split(bufio.ScanWords)	// 한 단어씩 끊어 내기

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
		// 에러 감싸기
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	return a * b, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() { // 단어 읽기
		return 0, 0, fmt.Errorf("Failed to scan") 
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word) // 문자열 숫자로 변환
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int word%s err:%w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFormString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			fmt.Println("number error:", numError)
		}// 감싸진 에러가 NumError인지 확인
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}