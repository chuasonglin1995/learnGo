package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/xml"
)

/*
This XML file does not appear to have any style information associated with it. The document tree is shown below.
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:news="http://www.google.com/schemas/sitemap-news/0.9" xmlns:video="http://www.google.com/schemas/sitemap-video/1.1">
	<url>
		<loc>https://www.washingtonpost.com/world/2024/02/26/navalny-russia-prisoner-swap-krasikov/</loc>
		<lastmod>2024-02-26T12:10:45.048Z</lastmod>
		<news:news>
			<news:publication>
				<news:name>Washington Post</news:name>
				<news:language>en</news:language>
			</news:publication>
			<news:publication_date>2024-02-26T12:10:45.048Z</news:publication_date>
			<news:title>
				<![CDATA[ Navalny aide says prisoner swap was in the works before his death ]]>
			</news:title>
		</news:news>
		<changefreq>hourly</changefreq>
	</url>
</urlset>
*/

type UrlSet struct {
	XMLName xml.Name `xml:"urlset"`
	Url    	[]Url   `xml:"url"`
}

type Url struct {
	Loc string `xml:"loc"`
	News News `xml:"news"`
	ChangeFreq string `xml:"changefreq"`
}

type News struct {
	Publication      Publication `xml:"publication"`
	PublicationDate  string      `xml:"publication_date"`
	Title            string      `xml:"title"`
}

type Publication struct {
	Name     string `xml:"name"`
	Language string `xml:"language"`
}

type NewsMap struct {
	Location string
	Language  string
}

func main() {

	var s UrlSet
	news_map := make(map[string]NewsMap)

	// because range returns 2 values (index, value)
	resp, _ := http.Get("https://www.washingtonpost.com/news-world-sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)


	for _, Url := range s.Url {
		news_map[Url.News.Title] = NewsMap{Url.Loc, Url.News.Publication.Language}
	}
	fmt.Println(news_map)
}
