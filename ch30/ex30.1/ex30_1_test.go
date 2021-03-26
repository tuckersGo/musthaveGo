//ch30/ex30.1/ex30_1_test.go
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
	req := httptest.NewRequest("GET", "/students", nil) // ❶ /students 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list) // ❷ 결과 변환
	assert.Nil(err)                                // ❸ 결과 확인
	assert.Equal(2, len(list))
	assert.Equal("aaa", list[0].Name)
	assert.Equal("bbb", list[1].Name)
}
