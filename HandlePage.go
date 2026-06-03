package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func HandlePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/input.html")
	if err != nil {
		fmt.Println("Error Parsing file: ", err)
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	 }
}