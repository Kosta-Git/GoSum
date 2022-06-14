package main

import (
	"fmt"
	"math"
)

type VerifiedChar struct {
	Char   string
	Status int
}
type Reverser []VerifiedChar

func (this Reverser) Reverse() {
	length := len(this)

	for i := 0; i < length/2; i++ {
		this[i], this[length-(i+1)] = this[length-(i+1)], this[i]
	}
}

func Verify(hash string, toVerify string, isReverseVerification bool) []VerifiedChar {
	chars := Reverser{}
	min := int(math.Min(float64(len(hash)), float64(len(toVerify))))

	if isReverseVerification {
		hash = reverse(hash)
		toVerify = reverse(toVerify)
	}

	for i := 0; i < min; i++ {
		status := -1
		if hash[i] == toVerify[i] {
			status = 1
		}

		chars = append(chars, VerifiedChar{
			Char:   string(hash[i]),
			Status: status,
		})
	}

	if len(hash) > len(toVerify) {
		for i := min; i < len(hash); i++ {
			chars = append(chars, VerifiedChar{
				Char:   string(hash[i]),
				Status: 0,
			})
		}
	}

	if isReverseVerification {
		chars.Reverse()
	}

	return chars
}

func PrintVerified(chars []VerifiedChar) {
	for _, c := range chars {
		switch c.Status {
		case 1:
			fmt.Print("\033[32m", c.Char, "\033[0m")
		case -1:
			fmt.Print("\033[31m", c.Char, "\033[0m")
		case 0:
			fmt.Print(c.Char)
		}
	}
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
