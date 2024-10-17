package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Title   string
	Message string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/submit", formHandler)

	port:= "8080"

	if len(os.Args) > 1{
		port = os.Args[1]
	 
	}


	// Static file handling (for serving CSS, JS, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Server starting at http://localhost:",port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {

	        fmt.Println("Error Starting Server at  http://localhost:",port)
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Home",
		Message: "Welcome to the Go Web App!",
	}
	renderTemplate(w, "home", data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "About",
		Message: "This is a simple Go web application.",
	}
	renderTemplate(w, "about", data)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		message := fmt.Sprintf("Hello, %s! Thanks for submitting the form.", name)
		data := PageData{
			Title:   "Form Submitted",
			Message: message,
		}
		renderTemplate(w, "submit", data)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	tmplPath := fmt.Sprintf("templates/%s.html", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}

