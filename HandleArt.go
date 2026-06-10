package main

import (
	"net/http"
	"html/template"
	"ascii-art-web/ascii"
)

func HandleArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
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
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	filePath := "banners/" + banner + ".txt"
	LoadBan, err := ascii.LoadBanner(filePath)
	if err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	GenArt := ascii.GenerateArt(text, LoadBan)

	data := PageData{
		Result: GenArt,
	}

	templ, err := template.ParseFiles("templates/output.html")
	if err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	err = templ.Execute(w, data)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}