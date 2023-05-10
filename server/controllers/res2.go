package controllers

import (
	"io"
	"net/http"
)

func Res2(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "response 2")
}
