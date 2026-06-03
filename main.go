package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
}

func main() {
	http.HandleFunc("/", HandlePage)
	http.HandleFunc("/ascii-art", HandleArt)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error Listening: ", err)
	}
}

func HandlePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/input.html")
	if err != nil {
		fmt.Println("Error Parsing file: ", err)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
}

func HandleArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banners")
	fontmap, err := LoadBanner(banner)
	if err != nil {
		fmt.Println("Error Reading banner file: ", err)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	generate := GenerateArt(text, fontmap)
	data := PageData{Result: generate}
	tmpl, err := template.ParseFiles("templates/input.html")
	if err != nil {
		fmt.Println("Error Parsing: ", err)
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, data)
}
