package main

import (
	"net/http"
	"fmt"
	"html/template"
	"ascii-art-web/ascii"
)
func HandleArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	text := r.FormValue("text")
	if len(text) == 0 {
		fmt.Println("Incorrect Request: ", nil)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	banner := r.FormValue("banners")
	if len(banner) == 0 {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	fontmap, err := ascii.LoadBanner(banner)
	if err != nil {
		fmt.Println("Error Reading banner file: ", err)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	generate := ascii.GenerateArt(text, fontmap)
	data := PageData{Result: generate}
	tmpl, err := template.ParseFiles("templates/input.html")
	if err != nil {
		fmt.Println("Error Parsing: ", err)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	 err = tmpl.Execute(w, data)
	 if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	 }
}
