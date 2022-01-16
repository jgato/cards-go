package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://stackoverflowed.com",
		"http://redhat.com",
	}

	if len(os.Args) > 1 && os.Args[1] == "--parallel" {
		c := make(chan string)
		NUMCPUS := 2
		init := 0
		end := NUMCPUS
		iterations := 1
		for {
			if end > len(links) {
				end = len(links)
			}
			jobs := links[init:end]
			for _, link := range jobs {
				go checkLinkConcurrent(link, c)
			}
			for i := 0; i < len(jobs); i++ {
				fmt.Println(<-c)
			}
			fmt.Printf("set %d of jobs done\n", iterations)
			if len(links) == end {
				break
			}
			init += NUMCPUS
			end += NUMCPUS
			iterations++
		}

	} else {
		for _, link := range links {
			checkLink(link)
		}
	}
}

func checkLink(link string) bool {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " could be down")
		return false
	}
	fmt.Println(link, " is oka")
	return true
}

func checkLinkConcurrent(link string, c chan string) bool {
	_, err := http.Get(link)
	if err != nil {
		c <- string(link + " could be down")
		return false
	}
	c <- string(link + " is ok")
	return true
}
