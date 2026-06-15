package main

import (
	"ascii-art-web/ascii"
	"html/template"
	"net/http"
)

func HandleArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if text == "" || banner == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	switch banner {
	case "standard", "shadow", "thinkertoy":
	default:
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	lban := "banner/" + banner + ".txt"
	ban, err := ascii.LoadBanner(lban)
	if err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	genArt := ascii.GenerateArt(text, ban)
	data := PageData{
		Result: genArt,
	}

	templ, err := template.ParseFiles("templates/output.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, data)
}
