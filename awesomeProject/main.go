package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}

func myPrint(n int) {
	fmt.Println("Это выполняется в потоке ", n)
}

func main() {

	go myPrint(1)

	fmt.Println("Это выполняется в основном потоке ")

	a := 0
	for i := 0; i < 10000000; i++ {
		a++
	}

	go myPrint(2)

	b := 0
	for i := 0; i < 10000000; i++ {
		b++
	}
}
