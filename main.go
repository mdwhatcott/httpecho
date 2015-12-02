package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			response.WriteHeader(500)
		} else {
			fmt.Fprint(response, string(dump))
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
