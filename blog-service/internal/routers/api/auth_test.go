package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"unsafe"
)

func TestGetAuth(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:9999/auth", nil)
	q := req.URL.Query()
	q.Add("app_key", "eddycjy")
	q.Add("app_secret", "go-programming-tour-book")
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Fatal Error", err.Error())
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	fmt.Println(*(*string)(unsafe.Pointer(&content)))
}
