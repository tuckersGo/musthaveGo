//ch29/ex29.6/ex26_6_test.go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil) // ❶ /student 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	student := new(Student)
	err := json.NewDecoder(res.Body).Decode(student) // ❷ 결과 변환
	assert.Nil(err)                                  // ❸ 결과 확인
	assert.Equal("aaa", student.Name)
	assert.Equal(16, student.Age)
	assert.Equal(87, student.Score)
}
