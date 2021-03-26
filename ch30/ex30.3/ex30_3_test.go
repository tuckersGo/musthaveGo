//ch30/ex30.3/ex30_3_test.go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestJsonHandler2(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil) // ❶ id 1 학생

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("aaa", student.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/2", nil) // ❷ id 2 학생
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err = json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("bbb", student.Name)
}

func TestJsonHandler3(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students",
		strings.NewReader(`{"Id":0,"Name":"ccc","Age":15,"Score":78}`)) // ❶ 새로운 학생 데이터

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code) // ❷ 응답 코드 검사

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil) // ❸ 추가된 학생 데이터
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("ccc", student.Name)
}

func TestJsonHandler4(t *testing.T) {
	assert := assert.New(t)

	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/students/1", nil) // ❶ DELETE 요청

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code) // ❷ 응답 코드 검사

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students", nil) // ❸ /students 경로
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(1, len(list)) // ❹ 결과 확인
	assert.Equal("bbb", list[0].Name)
}
