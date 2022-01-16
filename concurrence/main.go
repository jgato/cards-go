package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
// thanks to https://stackoverflow.com/a/18479916/15383067
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://stackoverflowed.com",
		"http://redhat.com",
	}

	// dont really make good checks of params, but.. just playing
	// main.go [file_with_urls] [--parallel NUMCPUS]
	if len(os.Args) > 1 {
		linksFromFile, err := readLines(os.Args[1])
		if err != nil {
			fmt.Println("error opening file")
			os.Exit(1)
		}
		links = linksFromFile
	}
	fmt.Println("ok lets go")
	if len(os.Args) > 2 && os.Args[2] == "--parallel" { // I trust you also will give the num of cpus
		c := make(chan string)
		NUMCPUS, _ := strconv.Atoi(os.Args[3])
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
