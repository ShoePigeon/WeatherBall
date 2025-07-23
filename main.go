package main

import (
	"fmt"
	"net/http"
	"text/template"
	"weatherball/web"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").Parse(web.IndexHTML)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Page")
}

func main() {

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/search", searchHandler)

	fmt.Println("Server starting on port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
