package main

import (
	"fmt"
	"net/http"
)

// Go并发编程实战汪明，书上的内容

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello Go https")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":443", "server.crt", "server.pem", nil)
}
