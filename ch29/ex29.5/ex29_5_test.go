//ch29/ex29.5/ex29_5_test.go
package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil) // ❶ / 경로 테스트

	mux := MakeWebHandler() // ❷
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code) // ❸ Code 확인
	data, _ := io.ReadAll(res.Body)       // ❹ 데이터를 읽어서 확인
	assert.Equal("Hello World", string(data))
}

func TestBarHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil) // ➎ /bar 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello Bar", string(data))
}
