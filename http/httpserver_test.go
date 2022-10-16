package http

import (
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello bowen"))
	})
	http.ListenAndServe(":8080", nil)
}
