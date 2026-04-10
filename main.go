package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Text   string
	Result string
	Banner string
	Error  string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	log.Println("Server listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	if r.Method != http.MethodGet {
		handleError(w, http.StatusMethodNotAllowed, "405 Method Not Allowed")
		return
	}

	renderTemplate(w, PageData{})
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(w, http.StatusMethodNotAllowed, "405 Method Not Allowed")
		return
	}

	err := r.ParseForm()
	if err != nil {
		handleError(w, http.StatusBadRequest, "400 Bad Request")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		handleError(w, http.StatusBadRequest, "400 Bad Request: Missing text or banner")
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		handleError(w, http.StatusBadRequest, "400 Bad Request: Invalid banner")
		return
	}

	// Check if banner file exists
	if _, err := os.Stat(banner + ".txt"); os.IsNotExist(err) {
		handleError(w, http.StatusNotFound, "404 Not Found: Banner not found")
		return
	}

	result, err := GenerateAsciiArt(text, banner)
	if err != nil {
		if err == ErrInvalidASCII {
			handleError(w, http.StatusBadRequest, "400 Bad Request: Invalid character in text")
			return
		}
		handleError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	data := PageData{
		Text:   text,
		Banner: banner,
		Result: result,
	}

	renderTemplate(w, data)
}

func renderTemplate(w http.ResponseWriter, data PageData) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError, "500 Internal Server Error: Template not found")
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "500 Internal Server Error")
	}
}

func handleError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	errTmpl, err := template.ParseFiles("templates/error.html")
	if err == nil {
		errTmpl.Execute(w, struct{ Message string }{Message: message})
	} else {
		w.Write([]byte(message))
	}
}
