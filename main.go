package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server Start")
	http.Handle("/", http.FileServer(http.Dir("./wasm/")))
	http.ListenAndServe(":8080", nil)
}
