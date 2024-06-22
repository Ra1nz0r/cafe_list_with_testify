package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Ra1nz0r/cafe_list_with_testify/internal/globals"
	"github.com/Ra1nz0r/cafe_list_with_testify/internal/logerr"
)

func MainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		logerr.LogErrWrite(w.Write([]byte("count missing")))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logerr.LogErrWrite(w.Write([]byte("wrong count value")))

		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := globals.CafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		logerr.LogErrWrite(w.Write([]byte("wrong city value")))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ", ")

	w.WriteHeader(http.StatusOK)
	logerr.LogErrWrite(w.Write([]byte(answer)))

}
