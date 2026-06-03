package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", HandlePage)
	http.HandleFunc("/ascii-art", HandlerArt)
	http.ListenAndServe(":8080", nil)
}

func HandlePage(w http.ResponseWriter, r *http.Request) {

}

func HandleArt(p http.ResponseWriter, q *http.Request) {

}
