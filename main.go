package main

import (
	"net/http"
)
func main() {
	http.HandleFunc("/", HandleThings)
	http.ListenAndServe(":8080", nil)
}
func HandleThings(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/input.html")
}