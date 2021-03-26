//ch28/ex28.3/ex28_3_test.go
package main

import (
	"testing"
)

func TestAtoi1(t *testing.T) {
	n, err := Atoi("0")
	if err != nil {
		t.Fail()
	}
	if n != 0 {
		t.Fail()
	}
}

func TestAtoi2(t *testing.T) {
	n, err := Atoi("1")
	if err != nil {
		t.Fail()
	}
	if n != 1 {
		t.Fail()
	}

	n, err = Atoi("5")
	if err != nil {
		t.Fail()
	}
	if n != 5 {
		t.Fail()
	}
}

func TestAtoi3(t *testing.T) {
	n, err := Atoi("12")
	if err != nil {
		t.Fail()
	}
	if n != 12 {
		t.Fail()
	}

	n, err = Atoi("3523")
	if err != nil {
		t.Fail()
	}
	if n != 3523 {
		t.Fail()
	}
}

func TestAtoi4(t *testing.T) {
	n, err := Atoi("ab34c")
	if err == nil {
		t.Fail()
	}

	n, err = Atoi("  3523")
	if err != nil {
		t.Fail()
	}
	if n != 3523 {
		t.Fail()
	}
}

func TestAtoi5(t *testing.T) {
	n, err := Atoi("-12")
	if err != nil {
		t.Fail()
	}
	if n != -12 {
		t.Fail()
	}
}
