package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"strings"
	"sync"
	"sync/atomic"
)

type channel struct {
	Items []channelItem `xml:"items,omitempty"`
}

type channelItem struct {
	Title       string `xml:"title,omitempty"`
	Description string `xml:"description,omitempty"`
}

type document struct {
	Channel channel `xml:"channel,omitempty"`
}

func main() {

	// start cpu profile with stdlib
	// pprof.StartCPUProfile(os.Stdout)
	// defer pprof.StopCPUProfile()

	trace.Start(os.Stdout)
	defer trace.Stop()

	docs := make([]string, 10000)
	for i := range docs {
		docs[i] = "newfeed.xml"
	}

	topic := "president"

	n := find3(topic, docs)

	log.Printf("found %s %d times", topic, n)
	// createDoc()
}

func createDoc() {
	d := document{
		Channel: channel{
			Items: []channelItem{
				channelItem{
					Title:       "president hello",
					Description: "world",
				},
			},
		},
	}

	b, _ := xml.Marshal(d)
	ioutil.WriteFile("newfeed.xml", b, 0600)
}

func find3(topic string, docs []string) int {
	var found int32

	g := runtime.NumCPU()

	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, len(docs))
	for _, doc := range docs {
		ch <- doc
	}
	close(ch)

	// create N go routine to process the buffered channel
	// we now make local variable lFound much longer, the cache line will only be thrashing N times where N equals to cpu number
	for i := 0; i < g; i++ {
		go func() {
			var lFound int32
			defer func() {
				atomic.AddInt32(&found, lFound)
				wg.Done()
			}()
			for doc := range ch {
				f, err := os.OpenFile(doc, os.O_RDONLY, 0)
				if err != nil {
					return
				}

				data, err := ioutil.ReadAll(f)
				if err != nil {
					f.Close()
					return
				}
				f.Close()

				var d document
				if err := xml.Unmarshal(data, &d); err != nil {
					return
				}

				for _, item := range d.Channel.Items {
					if strings.Contains(item.Title, topic) {
						lFound++
						continue
					}
					if strings.Contains(item.Description, topic) {
						lFound++
					}
				}
			}
		}()
	}
	wg.Wait()
	return int(found)
}

func find2(topic string, docs []string) int {
	var found int32

	g := len(docs)

	var wg sync.WaitGroup
	wg.Add(g)

	// first, we use goroutine to make full use of cpu
	// then we use local variable lFound to count for found, because lFound will be allocate on stack
	// that will reduce cpu cache threshing, we sync global counter found per doc using atomic
	for _, doc := range docs {
		go func(doc string) {
			var lFound int32
			defer func() {
				atomic.AddInt32(&found, lFound)
				wg.Done()
			}()
			f, err := os.OpenFile(doc, os.O_RDONLY, 0)
			if err != nil {
				return
			}

			data, err := ioutil.ReadAll(f)
			if err != nil {
				f.Close()
				return
			}
			f.Close()

			var d document
			if err := xml.Unmarshal(data, &d); err != nil {
				return
			}

			for _, item := range d.Channel.Items {
				if strings.Contains(item.Title, topic) {
					lFound++
					continue
				}
				if strings.Contains(item.Description, topic) {
					lFound++
				}
			}

			atomic.AddInt32(&found, lFound)
		}(doc)
	}
	wg.Wait()
	return int(found)
}

func find(topic string, docs []string) int {
	var found int

	for _, doc := range docs {
		f, err := os.OpenFile(doc, os.O_RDONLY, 0)
		if err != nil {
			return 0
		}

		data, err := ioutil.ReadAll(f)
		if err != nil {
			f.Close()
			return 0
		}
		f.Close()

		var d document
		if err := xml.Unmarshal(data, &d); err != nil {
			return 0
		}

		for _, item := range d.Channel.Items {
			if strings.Contains(item.Title, topic) {
				found++
				continue
			}
			if strings.Contains(item.Description, topic) {
				found++
			}
		}
	}
	return found
}
