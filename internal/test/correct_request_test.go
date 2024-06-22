package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Ra1nz0r/cafe_list_with_testify/internal/handlers"
	"github.com/Ra1nz0r/cafe_list_with_testify/internal/params"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCorrectRequest(t *testing.T) {
	for _, testCase := range TestTable {
		num := len(testCase.Cafe)
		if num != 1 {
			num /= 2
		}

		req := httptest.NewRequest("GET", params.UrlParams(testCase.City, strconv.Itoa(num)), nil)

		respRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.MainHandle)
		handler.ServeHTTP(respRecorder, req)

		resp := respRecorder.Result()
		body, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		require.NoError(t, err)
		assert.NotEmpty(t, body)
		assert.Equal(t, http.StatusOK, respRecorder.Code)
	}
}
