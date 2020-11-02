package main

import (
	"fmt"
	"net/http"
)

//curl -si "http://localhost:80/hello?name=World"
func helloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello")
	for _, v := range req.URL.Query() {
		fmt.Fprintf(w, " %s\n", v)
	}

}

func main() {
	http.HandleFunc("/hello", helloWorld)
	http.ListenAndServe(":80", nil)
}
