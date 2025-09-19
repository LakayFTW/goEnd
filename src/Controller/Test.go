package controller

import (
	"html/template"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/test.html")
	err := t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
