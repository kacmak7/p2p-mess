package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}
/*
func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}
*/
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

/*
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
*/

	err := http.ListenAndServe(":8081", mux)
	panic(err)
}
