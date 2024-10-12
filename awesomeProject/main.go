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

func main() {
	str1 := "abc"
	str2 := "абв"

	fmt.Println(len(str1))
	fmt.Println(len(str2))
}
