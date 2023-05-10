package controllers

import (
	"io"
	"net/http"
)

func Res1(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "response 1")
}
