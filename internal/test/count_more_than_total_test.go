package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ra1nz0r/cafe_list_with_testify/internal/globals"
	"github.com/Ra1nz0r/cafe_list_with_testify/internal/handlers"
	"github.com/Ra1nz0r/cafe_list_with_testify/internal/params"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := len(globals.CafeList["moscow"])

	req := httptest.NewRequest("GET", params.UrlParams("moscow", "20"), nil)

	respRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.MainHandle)
	handler.ServeHTTP(respRecorder, req)

	resp := respRecorder.Result()
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	resFromBody := strings.Split(string(body), ", ")

	require.NoError(t, err)
	assert.Len(t, resFromBody, totalCount)
	assert.Equal(t, globals.CafeList["moscow"], resFromBody)
}
