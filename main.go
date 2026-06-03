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
	fmt.Println("Server listening at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error Listening: ", err)
	}
}
