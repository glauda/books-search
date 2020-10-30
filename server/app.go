package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"html/template"
)

// wiki tutorial (official go website)
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

/*
// create a handler struct
type HttpHandler struct {}

// implement 'ServeHTTP' method on 'HttpHandler' struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// create response binary data
	data := []byte("Hello World")

	// write 'data' to response
	res.Write(data)
}
*/

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		fmt.Fprintf(w, "<h1>error: this page does not exist</h1><div>%s</div>", err)
	} else{
		t, _ := template.ParseFiles("view.html")
		t.Execute(w, p)
	}
}


// TODO: render template 

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: save function")
	/* TODO
	title := r.URL.Path[len("/save/"):]
	*/
}

func main() {

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	/*
	// create a new handler
	handler := HttpHandler{}

	// listen and serve
	http.ListenAndServe(":8080", handler)
	*/
}

func Hello(name string) string{
	return "Hello, " + name
}