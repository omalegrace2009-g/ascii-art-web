package main

import (
	"fmt"
	"net/http"
)
type PageData struct {
	Result string
}
func main() {
	http.HandleFunc("/", HandlePage)
	http.HandleFunc("/ascii-art", HandleArt)
	fmt.Println("Server Listening at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}