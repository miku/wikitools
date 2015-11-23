// Copyright 2014 by Leipzig University Library, http://ub.uni-leipzig.de
//                   The Finc Authors, http://finc.info
//                   Martin Czygan, <martin.czygan@uni-leipzig.de>
//
// This file is part of some open source application.
//
// Some open source application is free software: you can redistribute
// it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation, either
// version 3 of the License, or (at your option) any later version.
//
// Some open source application is distributed in the hope that it will
// be useful, but WITHOUT ANY WARRANTY; without even the implied warranty
// of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//
// @license GPL-3.0+ <http://spdx.org/licenses/GPL-3.0+>
//
//
// Package wikitools provides support for wikitools command line tools.
package wikitools

import (
	"io"
	"net/url"
	"regexp"
	"strings"
)

// Version of wikitools.
const Version = "0.1.3"

// Redirect title.
type Redirect struct {
	Title string `xml:"title,attr" json:"title"`
}

// Page as it occurs on Wikipedia.
type Page struct {
	Title          string   `xml:"title" json:"title"`
	CanonicalTitle string   `xml:"ctitle" json:"ctitle"`
	Redir          Redirect `xml:"redirect" json:"redirect"`
	Text           string   `xml:"revision>text" json:"text"`
}

// WikidataPage as it occurs on Wikidata.
type WikidataPage struct {
	Title          string      `xml:"title" json:"title"`
	CanonicalTitle string      `xml:"ctitle" json:"ctitle"`
	Redir          Redirect    `xml:"redirect" json:"redirect"`
	Content        interface{} `json:"content"`
}

// CanonicalizeTitle unifies a wikipedia page title.
func CanonicalizeTitle(title string) string {
	can := strings.ToLower(title)
	can = strings.Replace(can, " ", "_", -1)
	can = url.QueryEscape(can)
	return can
}

// ExtractPageCategory returns a slice of categories for a given page.
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

// ExtractPageAuthorityControl extract raw authority control data from a page.
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

// FanInWriter writes the string arrays as tab separated values to the writer.
func FanInTabWriter(writer io.Writer, in chan *[]string, done chan bool) {
	for fields := range in {
		writer.Write([]byte(strings.Join(*fields, "\t")))
		writer.Write([]byte("\n"))
	}
	done <- true
}

// FanInLineWriter writes the strings to the writer.
func FanInLineWriter(writer io.Writer, in chan *string, done chan bool) {
	for s := range in {
		writer.Write([]byte(*s))
		writer.Write([]byte("\n"))
	}
	done <- true
}
