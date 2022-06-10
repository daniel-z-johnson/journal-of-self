package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	fmt.Println("Application Start")
	application()
}

func application() error {
	r := chi.NewMux()
	r.HandleFunc("/test", test)
	return http.ListenAndServe(":1117", r)
}

func test(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, "Success")
}
