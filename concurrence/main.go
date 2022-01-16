package main

import (
	"bufio"
	"errors"
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

// return error when no more elements
func popFirstLink(links []string) (string, error) {
	if len(links) > 0 {
		link := links[0]
		links = links[1:]
		return link, nil
	}
	return "", errors.New("no more elements")
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
		// create a set of maximum goroutines depending on the num of cpus
		// or the number of urls available
		for i := 0; i < NUMCPUS; i++ {
			link, err := popFirstLink(links)
			if err != nil {
				go checkLinkConcurrent(link, c)
			} else {
				break
			}
		}
		count := 0
		// from now on, each time of one goroutine finishes we create a new one
		// until we empty the list of links
		// that we ensure we have only NUMCPUS goroutines working on parallel
		for {
			fmt.Println(<-c)
			count++
			link, err := popFirstLink(links)
			if err != nil {
				go checkLinkConcurrent(link, c)
			} else {
				break
			}
		}
		fmt.Printf("processed %d urls\n", count)

	} else {
		count := 0
		for _, link := range links {
			checkLink(link)
			count++
		}
		fmt.Printf("processed %d urls\n", count)
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
