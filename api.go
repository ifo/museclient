package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Muse struct {
	Url    string
	ApiKey string
	Client http.Client
}

type JobsResponse struct {
	Page      int64    `json:"page"`
	PageCount int64    `json:"page_count"`
	Results   []Result `json:"results"`
	TimedOut  bool     `json:"timed_out"`
	Took      int64    `json:"took"`
	Total     int64    `json:"total"`
}

type Result struct {
	Categories []struct {
		Name string `json:"name"`
	} `json:"categories"`
	Company struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		ShortName string `json:"short_name"`
	} `json:"company"`
	Contents string `json:"contents"`
	ID       int64  `json:"id"`
	Levels   []struct {
		Name      string `json:"name"`
		ShortName string `json:"short_name"`
	} `json:"levels"`
	Locations []struct {
		Name string `json:"name"`
	} `json:"locations"`
	ModelType       string `json:"model_type"`
	Name            string `json:"name"`
	PublicationDate string `json:"publication_date"`
	Refs            struct {
		LandingPage string `json:"landing_page"`
	} `json:"refs"`
	ShortName string        `json:"short_name"`
	Tags      []interface{} `json:"tags"`
	Type      string        `json:"type"`
}

// GetJobs retrieves all jobs matching the query provided by the url Values.
func (m *Muse) GetJobs(v url.Values) ([]Result, error) {
	// Set ApiKey if we have one.
	if m.ApiKey != "" {
		v.Set("api_key", m.ApiKey)
	}

	getUrl := fmt.Sprintf("%s?%s", m.Url, v.Encode())
	resp, err := m.Client.Get(getUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Successful requests might still not give results.
	if resp.StatusCode != 200 {
		// TODO: Give more specific status code errors
		return nil, fmt.Errorf("API request was not successful, status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jr JobsResponse
	err = json.Unmarshal(body, &jr)
	if err != nil {
		return nil, err
	}
	return jr.Results, nil
}
