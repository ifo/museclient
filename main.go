package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	// Setup config.
	port := flag.Int("port", 3000, "Port to run the server on")
	flag.Parse()

	// Setup routes.
	http.HandleFunc("/", index)
	http.HandleFunc("/search", search)

	// Run server.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

type SearchPage struct {
	Form    url.Values
	Results map[string]string
}

func search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Ensure there is a page value.
	if _, ok := r.Form["page"]; !ok {
		r.Form.Set("page", "0")
	}

	// Get the page value as an int for the previous and next buttons.
	pageVal, err := strconv.Atoi(r.Form["page"][0])
	if err != nil {
		log.Println(err)
		// return 500
		return
	}

	// Set the previous and next buttons.
	r.Form.Set("next", strconv.Itoa(pageVal+1))
	if pageVal > 0 {
		r.Form.Set("prev", strconv.Itoa(pageVal-1))
	}

	search, err := template.ParseFiles("./search.html")
	if err != nil {
		log.Println(err)
		return
	}

	// Display the search page.
	sp := SearchPage{Form: r.Form, Results: map[string]string{"example": "result"}}
	if err := search.Execute(w, sp); err != nil {
		log.Println(err)
		// return 500
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	idx, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Println(err)
		return
	}

	// Display the index page.
	if err := idx.Execute(w, nil); err != nil {
		log.Println(err)
		// return 500
	}
}
