package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mx := http.NewServeMux()
	mx.Handle("POST /CreateMember", http.HandlerFunc(CreateMember))

	fmt.Println("Server started at port :8081")
	log.Fatal(http.ListenAndServe(":8081", mx))
}
