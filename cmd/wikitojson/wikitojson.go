package main

import (
	"encoding/json"
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

// WikiJsonConverter converts XML pages into line delimited JSON
func WikiJsonConverter(in chan *wikitools.Page, out chan *string, filter *regexp.Regexp, wg *sync.WaitGroup) {
	defer wg.Done()

	for page := range in {
		canonicalTitle := wikitools.CanonicalizeTitle(page.Title)
		m := filter.MatchString(canonicalTitle)
		if !m && page.Redir.Title == "" {
			b, err := json.Marshal(page)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			line := string(b)
			out <- &line
		}
	}
}

func main() {

	filter := flag.String("filter", "^file:.*|^talk:.*|^special:.*|^wikipedia:.*|^wiktionary:.*|^user:.*|^user_talk:.*", "regex for pages to skip")
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
		log.Fatalln("Usage: wikitojson [OPTIONS] WIKIPEDIA-DUMP-XML")
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
	results := make(chan *string)
	// done chan
	done := make(chan bool)

	var wg sync.WaitGroup

	// workers
	for i := 0; i < *numWorkers; i++ {
		wg.Add(1)
		go WikiJsonConverter(queue, results, filterRx, &wg)
	}

	// output writer
	go wikitools.FanInLineWriter(os.Stdout, results, done)

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
