package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var fetched = struct {
	urls map[string]bool
	urlErrors map[string]error
	sync.Mutex
}{
	urls : make(map[string]bool),
	urlErrors : make(map[string]error),
}

func Crawl(url string, depth int, fetcher Fetcher) {

	// TODO:並列化
	// TODO:二度同じURLを読み込まない
	var crawl func(string, int)
	crawl = func (url string, depth int) {
		if depth <= 0 {
			return
		}

		fetched.Lock()
		if _, ok := fetched.urls[url]; ok {

			fetched.Unlock()
			return
		}

		fetched.urls[url] = true
		fetched.Unlock()
		
		body, urls, err := fetcher.Fetch(url)
		fetched.urlErrors[url] = err
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
	
		for _, u := range urls {
			go crawl(u, depth-1)
		}
	} 
	go crawl( url, depth)

	time.Sleep(3 * time.Second)
	
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

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
