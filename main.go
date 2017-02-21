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

type Config struct {
	Muse   Muse
	Index  *template.Template
	Search *template.Template
}

// Global Config
var config Config

func main() {
	// Setup config.
	port := flag.Int("port", 3000, "Port to run the server on")
	apiUrl := flag.String("apiurl", "https://api-v2.themuse.com/jobs", "The Muse api url")
	apiKey := flag.String("apikey", "", "The Muse api key")
	indexTemplate := flag.String("index", "index.html", "The index.html file location")
	searchTemplate := flag.String("search", "search.html", "The search.html file location")
	flag.Parse()

	config = Config{
		Muse:   Muse{Url: *apiUrl, ApiKey: *apiKey},
		Index:  template.Must(template.ParseFiles(*indexTemplate)),
		Search: template.Must(template.ParseFiles(*searchTemplate)),
	}

	// Setup routes.
	http.HandleFunc("/", index)
	http.HandleFunc("/search", search)

	// Run server.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

type SearchPage struct {
	Form    url.Values
	Results []Result
}

func search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Ensure there is a page value.
	if _, ok := r.Form["page"]; !ok {
		r.Form.Set("page", "0")
	}

	// Get job results.
	results, err := config.Muse.GetJobs(r.Form)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Get the page value as an int for the previous and next buttons.
	pageVal, err := strconv.Atoi(r.Form["page"][0])
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Set the previous and next buttons.
	// TODO: Remove the next button when there are no more pages to view.
	r.Form.Set("next", strconv.Itoa(pageVal+1))
	if pageVal > 0 {
		r.Form.Set("prev", strconv.Itoa(pageVal-1))
	}

	// Display the search page.
	sp := SearchPage{Form: r.Form, Results: results}
	if err := config.Search.Execute(w, sp); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// Display the index page.
	if err := config.Index.Execute(w, nil); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}
