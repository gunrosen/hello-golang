package main

import (
	"fmt"
	"sync"
)

/*
	Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.
*/

type SafeCrawledResult struct {
	mu sync.Mutex
	v  map[string]bool
}

func (c *SafeCrawledResult) SetCrawledUrl(url string) {
	c.mu.Lock()
	c.v[url] = true
	c.mu.Unlock()
}

func (c *SafeCrawledResult) IsCrawled(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[url]
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
var wg sync.WaitGroup

func Crawl(url string, depth int, fetcher Fetcher, scr *SafeCrawledResult) {
	// Fetch URLs in parallel.
	// Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	// Check if url already crawled then ignore
	if scr.IsCrawled(url) {
		// Do nothing
		//fmt.Printf("visit already crawl: %s \n", url)
	} else {
		scr.SetCrawledUrl(url)
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			wg.Add(1)
			u := u
			go func() {
				defer wg.Done()
				Crawl(u, depth-1, fetcher, scr)
			}()
		}
	}
	return
}

func main() {
	scr := SafeCrawledResult{v: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, &scr)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
