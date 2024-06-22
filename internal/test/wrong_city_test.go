package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ra1nz0r/cafe_list_with_testify/internal/handlers"
	"github.com/Ra1nz0r/cafe_list_with_testify/internal/params"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", params.UrlParams("tulula", "1"), nil)

	respRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.MainHandle)
	handler.ServeHTTP(respRecorder, req)

	resp := respRecorder.Result()
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, respRecorder.Code)
	assert.Equal(t, "wrong city value", string(body))
}
