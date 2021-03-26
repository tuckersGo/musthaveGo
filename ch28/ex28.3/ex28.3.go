//ch28/ex28.3/ex28.3.go
package main

import (
	"errors"
	"strings"
)

func Atoi(input string) (int, error) {
	rst := 0
	negative := false
	input = strings.TrimSpace(input)
	if input[0] == '-' {
		negative = true
		input = input[1:]
	}
	for _, c := range input {
		if c >= '0' && c <= '9' {
			rst *= 10
			rst += int(c - '0')
		} else {
			return 0, errors.New("cannot convert to int")
		}
	}
	if negative {
		rst *= -1
	}
	return rst, nil
}
