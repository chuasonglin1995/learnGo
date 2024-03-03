package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Url     []Url    `xml:"url"`
}

type Url struct {
	Loc        string `xml:"loc"`
	News       News   `xml:"news"`
	ChangeFreq string `xml:"changefreq"`
}

type News struct {
	Publication     Publication `xml:"publication"`
	PublicationDate string      `xml:"publication_date"`
	Title           string      `xml:"title"`
}

type Publication struct {
	Name     string `xml:"name"`
	Language string `xml:"language"`
}

type NewsForDisplay struct {
	Title					  string
	Location        string
	PublicationDate string `xml:"publication_date"`
}

type NewsAggPage struct {
	Title string
	News  map[string]NewsForDisplay
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

func newsRoutine(c chan NewsForDisplay, url Url) {
	defer wg.Done()
	var n NewsForDisplay

	// I cant actually get anything from new url, beause its blocked by paywall. But its just for practise purposes.
	resp, _ := http.Get(url.Loc)
	ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	n = NewsForDisplay{url.News.Title, url.Loc, url.News.PublicationDate}
	fmt.Println(n)
	c <- n
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s UrlSet
	news_map := make(map[string]NewsForDisplay)

	resp, _ := http.Get("https://www.washingtonpost.com/news-world-sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)
	queue := make(chan NewsForDisplay, 30)

	for _, Url := range s.Url {
		wg.Add(1)
		go newsRoutine(queue, Url)
	}

	wg.Wait()
	close(queue)

	for elem := range queue {
		news_map[elem.Title] = elem
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
	t, _ := template.ParseFiles("newsaggtemplate.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
