package server

import (
	"net/http"

	"github.com/Ra1nz0r/cafe_list_with_testify/internal/handlers"
	"github.com/Ra1nz0r/cafe_list_with_testify/internal/logerr"
)

func StartServe() {
	http.HandleFunc(`/cafe`, handlers.MainHandle)
	err := http.ListenAndServe(":8080", nil)
	logerr.LogErr(err)
}
