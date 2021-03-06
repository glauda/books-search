package main

import (
	"net/http"
	"fmt"
	"log"
//	"io/ioutil"
//	"html/template"
//	"regexp"
//	"errors"
)

// global variables
//var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
//var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// wiki tutorial (official go website)
/*
type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile("./wikifiles/" + filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
	body, err := ioutil.ReadFile("./wikifiles/" + filename)
	// body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
    return &Page{Title: title, Body: body}, nil
}


func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, "./static/" + tmpl + ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil
}

// http handler
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)	
	
}


func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}


func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
		}
		fn(w, r, m[2])
	}
}
*/

func main() {

	// http.Handle("/", http.FileServer(http.Dir("./static")))

	// serve production react app
	// caution: go run app.go must be executed in server/
	http.Handle("/", http.FileServer(http.Dir("../client/build")))

	// http.HandleFunc("/view/", makeHandler(viewHandler))
	// http.HandleFunc("/edit/", makeHandler(editHandler))
	// http.HandleFunc("/save/", makeHandler(saveHandler))
	
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}

func Hello(name string) string{
	return "Hello, " + name
}