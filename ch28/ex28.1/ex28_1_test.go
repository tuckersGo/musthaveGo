//ch28/ex28.1/ex28_1_test.go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {
	assert := assert.New(t)                               // ❶ 테스트 객체 생성
	assert.Equal(81, square(9), "square(9) should be 81") // ❷ 테스트 함수 호출
}

func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) should be 9")
}
