package chapter

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var Priorty int = 0

type Chapter struct {
	Priorty int
	Url     string
	Title   string
	Date    string
}

// extrackDataFromLi extracts chapter data from an HTML node and returns a Chapter struct.
// It iterates through the child nodes of the given HTML node and populates the Chapter struct
// with the URL, title, and date information.
// The title is cleaned by removing newline and tab characters.
func extrackDataFromLi(n *html.Node) Chapter {
	var chapter Chapter
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "a" {
			for _, a := range c.Attr {
				if a.Key == "href" {
					chapter.Url = a.Val
				}
			}
			chapter.Title = c.FirstChild.Data
		}
		if c.Type == html.ElementNode && c.Data == "span" {
			for _, a := range c.Attr {
				if a.Key == "class" && a.Val == "chapter-release-date" {
					chapter.Date = c.FirstChild.NextSibling.FirstChild.Data
				}
			}
		}
	}
	// remove \n and \t from title
	chapter.Title = strings.ReplaceAll(chapter.Title, "\n", "")
	chapter.Title = strings.ReplaceAll(chapter.Title, "\t", "")
	return chapter
}

// extrackChapterUrlsFromHtml extracts chapter URLs from an HTML node and returns a slice of Chapter.
// It recursively traverses the HTML node and extracts chapter information from each "li" element.
// The extracted chapters are assigned a priority based on their order in the HTML structure.
// The function returns a slice of Chapter containing the extracted chapter information.
func extrackChapterUrlsFromHtml(n *html.Node) []Chapter {

	var chapters []Chapter
	if n.Type == html.ElementNode && n.Data == "li" {
		var tmpChapter Chapter
		Priorty++
		tmpChapter = extrackDataFromLi(n)
		tmpChapter.Priorty = Priorty
		chapters = append(chapters, tmpChapter)

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		chapters = append(chapters, extrackChapterUrlsFromHtml(c)...)
	}
	return chapters
}

// GetChapters retrieves the chapters from the specified AJAX URL.
// It sends a POST request to the URL and parses the response body to extract the chapter URLs.
// Returns a slice of Chapter URLs.
func GetChapters(CHAPTERS_AJAX_URL string) []Chapter {
	// Get chapters from ajax url with post request
	res, err := http.PostForm(CHAPTERS_AJAX_URL, nil)
	if err != nil {
		fmt.Println("Error while getting chapters")
	}
	// read response body HTML
	body, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println("Error while parsing response body")
	}
	// extract chapter urls from html
	urls := extrackChapterUrlsFromHtml(body)
	return urls
}
