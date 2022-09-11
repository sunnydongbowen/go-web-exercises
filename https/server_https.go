package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello Go https")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":443", "server.crt", "server.pem", nil)
}
