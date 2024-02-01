package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tmpl *template.Template

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.Execute(w,nil)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	// Assuming you have a form field named "name"
	name := r.Form.Get("name")

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {

	tmpl =template.Must(template.ParseFiles("form.html"))
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
