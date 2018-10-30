package main

import (
	"log"
	"net/http"
	"github.com/ns7381/go_learn/blog/handler"
)

func main() {
	apiHandler, err := handler.CreateHttpAPIHandler()
	if err != nil {
		log.Fatalf("create http api handler error.", err)
	}
	http.Handle("/api/", apiHandler)

	log.Printf("start listening on localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
