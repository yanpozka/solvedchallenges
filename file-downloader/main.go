package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"

	"gopkg.in/cheggaaa/pb.v1"
)

const url = `http://0.0.0.0:9999/gato.zip`

func main() {
	var conLen int64
	{
		res, err := http.Head(url)
		check(err)
		// explict is better than implicit
		n, err := strconv.ParseInt(res.Header.Get("Content-Length"), 10, 64)
		check(err)
		conLen = int64(n)
	}
	const partLen = 3 * 1024 * 1024 // 3M

	bar := pb.New64(conLen/partLen + 1)
	bar.ShowCounters = false

	file, err := os.Create("file.zip")
	check(err)

	numCPUs := runtime.NumCPU()
	jobs := make(chan int64, numCPUs+1)

	worker := func() {
		for offset := range jobs {
			req, err := http.NewRequest("GET", url, nil)
			check(err)
			max := offset + partLen
			if max > conLen {
				max = conLen
			}
			req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", offset, max))

			res, err := http.DefaultClient.Do(req)
			check(err)
			if res.StatusCode == http.StatusRequestedRangeNotSatisfiable {
				log.Panic(http.StatusText(http.StatusRequestedRangeNotSatisfiable))
			}

			data, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			check(err)

			_, err = file.WriteAt(data, offset)
			check(err)

			bar.Increment()
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(numCPUs)

	for ix := 0; ix < numCPUs; ix++ {
		go func() {
			worker()
			wg.Done()
		}()
	}
	bar.Start()

	for offset := int64(0); offset < conLen; offset += partLen {
		jobs <- offset
	}
	close(jobs)

	wg.Wait()

	bar.FinishPrint(fmt.Sprintf("Finished: %d", conLen))
}

func dumpFile(name string) {
	file, err := os.Open(name)
	check(err)
	b, err := ioutil.ReadAll(file)
	check(err)
	fmt.Println(hex.Dump(b))
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
