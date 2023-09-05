package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func getFilms() []Film {
	films := []Film{
		{Title: "The Godfather", Director: "Francis Ford Coppola"},
		{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
		{Title: "Schindler's List", Director: "Steven Spielberg"},
		{Title: "Raging Bull", Director: "Martin Scorsese"},
	}

	return films
}

func index(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": getFilms(),
	}

	tmpl.Execute(w, films)
}

func addFilm(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")

	htmlStr := fmt.Sprintf("<div class=\"flex mb-4 items-center\"><p class=\"w-full text-grey-darkest\">%s - %s</p></div>", title, director)
	tmpl, _ := template.New("add-film").Parse(htmlStr)

	tmpl.Execute(w, nil)
}

func main() {
	fmt.Println("Hello, World!")

	http.HandleFunc("/", index)
	http.HandleFunc("/add-film", addFilm)

	http.ListenAndServe(":3333", nil)
}
