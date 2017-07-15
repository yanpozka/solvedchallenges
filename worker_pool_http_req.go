//
// Inspired by https://www.careercup.com/question?id=5710657300201472
//			   https://gobyexample.com/worker-pools
//
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"
)

const (
	wikiURL = "http://en.wikipedia.org/wiki/Main_Page"
	total   = 100
)

func main() {
	solution(runtime.NumCPU() * 2)
}

func solution(cores int) {
	log.SetFlags(log.Lshortfile)

	jobs := make(chan string, cores)
	result := make(chan int, cores)

	for w := 0; w < cores; w++ {
		go worker(w, jobs, result)
	}

	go func() {
		for ix := 0; ix < total; ix++ {
			jobs <- wikiURL
		}
		close(jobs)
	}()

	for p := 1; p <= total; p++ { // total is 100 so p match with current percent
		ctlen := <-result
		if ctlen == 0 {
			log.Println("empty response :(")
		}
		log.Printf("percent %d \n", p)
	}
}

func main_old() {
	for ix := 0; ix < total; ix++ {
		hitWiki(wikiURL)
	}
}

func worker(id int, jobs <-chan string, result chan<- int) {
	for url := range jobs {
		start := time.Now()
		if cl, err := hitWiki(url); err != nil {
			log.Panic(err)
		} else {
			result <- cl
		}

		log.Println("worker #", id, time.Since(start))
	}
}

func hitWiki(url string) (int, error) {
	client := new(http.Client)

	res, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	return len(bs), nil
}
