package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Option struct {
	length    int
	upperCase bool
	lowerCase bool
	numbers   bool
	special   bool
}

var source = rand.NewSource(time.Now().UnixNano())

const _charsetLowerCase = "abcdefghijklmnopqrstuvwxyz"
const _charsetUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const _charsetNumbers = "0123456789"
const _charsetSpecial = "!@#$%^&*()_+{}[]:;?/|"

func generatePassword(opt Option) (string, error) {
	x := make([]byte, opt.length)
	charset := ""

	if opt.upperCase {
		charset += _charsetUpperCase
	}
	if opt.lowerCase {
		charset += _charsetLowerCase
	}
	if opt.numbers {
		charset += _charsetNumbers
	}
	if opt.special {
		charset += _charsetSpecial
	}

	if charset == "" {
		return "", errors.New("No charset selected")
	}

	for i := range x {
		x[i] = charset[source.Int63()%int64(len(charset))]
	}

	return string(x), nil
}

func main() {
	password, err := generatePassword(Option{length: 8, upperCase: true, lowerCase: true, numbers: true, special: true})

	fmt.Println(password)
	if err != nil {
		fmt.Println(err.Error())
	}
}
