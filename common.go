package wikitools

import (
	"io"
	"net/url"
	"regexp"
	"strings"
)

const Version = "0.1.0"

type Redirect struct {
	Title string `xml:"title,attr" json:"title"`
}

// Page as it occurs on Wikipedia
type Page struct {
	Title          string   `xml:"title" json:"title"`
	CanonicalTitle string   `xml:"ctitle" json:"ctitle"`
	Redir          Redirect `xml:"redirect" json:"redirect"`
	Text           string   `xml:"revision>text" json:"text"`
}

// WikidataPage as it occurs on Wikidata, content will be turned from a string
// into a substructure with -d switch
type WikidataPage struct {
	Title          string      `xml:"title" json:"title"`
	CanonicalTitle string      `xml:"ctitle" json:"ctitle"`
	Redir          Redirect    `xml:"redirect" json:"redirect"`
	Content        interface{} `json:"content"`
}

// CanonicalizeTitle unifies a wikipedia page title
func CanonicalizeTitle(title string) string {
	can := strings.ToLower(title)
	can = strings.Replace(can, " ", "_", -1)
	can = url.QueryEscape(can)
	return can
}

// ExtractPageCategory returns a slice of categories for a given page
func ExtractPageCategory(page *Page, filter, pattern *regexp.Regexp) []string {
	var cats []string
	m := filter.MatchString(CanonicalizeTitle(page.Title))
	if !m && page.Redir.Title == "" {

		// specific to category extraction
		result := pattern.FindAllStringSubmatch(page.Text, -1)
		for _, value := range result {
			// replace anything after a |
			category := strings.TrimSpace(value[1])
			firstIndex := strings.Index(category, "|")
			if firstIndex != -1 {
				category = category[0:firstIndex]
			}
			cats = append(cats, category)
		}
	}
	return cats
}

// ExtractPageAuthorityControl extract raw authority control data from a page
func ExtractAuthorityControl(page *Page, filter, pattern *regexp.Regexp) []string {
	m := filter.MatchString(CanonicalizeTitle(page.Title))
	var result []string
	if !m && page.Redir.Title == "" {
		match := pattern.FindString(page.Text)
		if match != "" {
			// https://cdn.mediacru.sh/JsdjtGoLZBcR.png
			result = append(result, page.Title, strings.Replace(match, "\t", "", -1))
		}
	}
	return result
}

// FanInWriter writes the string arrays as tab separated values to the writer
func FanInTabWriter(writer io.Writer, in chan *[]string, done chan bool) {
	for fields := range in {
		writer.Write([]byte(strings.Join(*fields, "\t")))
		writer.Write([]byte("\n"))
	}
	done <- true
}

// FanInLineWriter writes the strings to the writer
func FanInLineWriter(writer io.Writer, in chan *string, done chan bool) {
	for s := range in {
		writer.Write([]byte(*s))
		writer.Write([]byte("\n"))
	}
	done <- true
}
