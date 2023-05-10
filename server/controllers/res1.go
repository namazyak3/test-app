package controllers

import (
	"io"
	"net/http"
)

func Handler1(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "response 1")
}
