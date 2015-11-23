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
// wikinorm extract authority data from wikipedia XML dump.
package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"runtime/pprof"
	"sync"

	"github.com/miku/wikitools"
)

// AuthorityDataExtractor takes a channel of pages and yields PageAuthorityControl structs
func AuthorityDataExtractor(in chan *wikitools.Page, out chan *[]string, filter, pattern *regexp.Regexp, wg *sync.WaitGroup) {
	defer wg.Done()
	for page := range in {
		result := wikitools.ExtractAuthorityControl(page, filter, pattern)
		if len(result) > 0 {
			out <- &result
		}
	}
}

func main() {

	filter := flag.String("filter", "^file:.*|^talk:.*|^special:.*|^wikipedia:.*|^wiktionary:.*|^user:.*|^user_talk:.*", "regex for pages to skip")
	pattern := flag.String("pattern", "Authority Control", "word for Authority Control, will differ accross languages, e.g. for German this would be 'Normdaten'")
	version := flag.Bool("v", false, "prints current program version")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	numWorkers := flag.Int("w", runtime.NumCPU(), "number of workers")

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	filterRx, err := regexp.Compile(*filter)
	if err != nil {
		log.Fatalln(err)
	}

	flag.Parse()

	runtime.GOMAXPROCS(*numWorkers)

	if *version {
		fmt.Println(wikitools.Version)
		os.Exit(0)
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if len(flag.Args()) != 1 {
		log.Fatalln("Usage: wikinorm [OPTIONS] WIKIPEDIA-DUMP-XML")
	}

	// the input XML
	filename := flag.Args()[0]
	handle, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer handle.Close()

	// the parsed XML pages channel
	queue := make(chan *wikitools.Page)
	// output channel
	results := make(chan *[]string)
	// done chan
	done := make(chan bool)

	var wg sync.WaitGroup
	patternRx := regexp.MustCompile(`(?mi){{` + *pattern + `[^}]*}}`)

	// workers
	for i := 0; i < *numWorkers; i++ {
		wg.Add(1)
		go AuthorityDataExtractor(queue, results, filterRx, patternRx, &wg)
	}

	// output writer
	go wikitools.FanInTabWriter(os.Stdout, results, done)

	// XML decoder
	decoder := xml.NewDecoder(handle)
	var inElement string

	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			// If we just read a StartElement token
			inElement = se.Name.Local
			// ...and its name is "page"
			if inElement == "page" {
				var p wikitools.Page
				// decode a whole chunk of following XML into the
				// variable p which is a Page (se above)
				decoder.DecodeElement(&p, &se)
				queue <- &p
			}
		default:
		}
	}

	close(queue)
	wg.Wait()
	close(results)
	<-done
}
