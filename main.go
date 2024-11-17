package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	LoadConfig()
	ConnectDB()
	
	mx := http.NewServeMux()
	mx.Handle("POST /CreateMember", http.HandlerFunc(CreateMember))
	mx.Handle("POST /AddNewBook", http.HandlerFunc(AddNewBook))
	mx.Handle("POST /IssueBook", http.HandlerFunc(GetBook))
	mx.Handle("POST /ReturnBook", http.HandlerFunc(BackBook))

	fmt.Println("Server started at port :8081")
	log.Fatal(http.ListenAndServe(":8081", mx))
}
