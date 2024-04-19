package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
	errCode := http.StatusOK
	req := httptest.NewRequest("GET", "/cafe?count=1&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	statusCode := responseRecorder.Code

	require.Equal(t, errCode, statusCode)
	assert.NotEmpty(t, body)

}

func TestMainHandlerWhenCityDoesntSupport(t *testing.T) {
	errText := "wrong city value"
	errCode := http.StatusBadRequest
	req := httptest.NewRequest("GET", "/cafe?count=10&city=tula", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	statusCode := responseRecorder.Code

	// здесь нужно добавить необходимые проверки
	assert.Equal(t, errCode, statusCode)
	assert.Equal(t, errText, body)

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	// здесь нужно добавить необходимые проверки
	assert.Len(t, list, totalCount)
}
