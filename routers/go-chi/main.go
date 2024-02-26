package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
)

// PageData структура для передачі даних у шаблон
type PageData struct {
	Title string
	Name  string
}

func main() {
	router := chi.NewRouter()

	// Головна сторінка
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", PageData{Title: "Home", Name: ""})
	})

	// Вітання з ім'ям
	router.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		renderTemplate(w, "hello", PageData{Title: "Hello", Name: name})
	})

	// Інша сторінка
	router.Get("/another", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "another", PageData{Title: "Another Page", Name: ""})
	})

	http.ListenAndServe(":8080", router)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	// Зчитування шаблону з файлу
	tmplPath := "templates/" + tmpl + ".html"
	tmplContent, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Використання шаблону для відображення сторінки
	err = tmplContent.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
