package http

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func httpGet() {
	resp, err := http.Get("http://127.0.0.1:8080/?key=bowen")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// ioutil包部分函数被io和os
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("http://127.0.0.1:8080/", "application/x-www-form-urlencoded", strings.NewReader("name=Go"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}
func main() {
	httpGet()
	httpPost()
}
