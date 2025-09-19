package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Serve() {
	// Statische CSS-Dateien bereitstellen
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	// Handler für die Login-Seite
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        templates := template.Must(template.ParseGlob("templates/*.html"))

        data := map[string]interface{}{
            "Title": "GoEnd",
        }

        err := templates.ExecuteTemplate(w, "layout", data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

	fmt.Println("Server läuft auf http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
