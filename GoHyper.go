//@github.com/Seeps

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// Define flags for the URL and the output file
	url := flag.String("url", "http://example.com", "the URL to parse")
	outputFile := flag.String("output", "", "the file to write the results to")

	// Parse the command-line arguments
	flag.Parse()

	// Make a GET request to the URL
	resp, err := http.Get(*url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Parse the response body as HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Use a recursive function to visit each node in the HTML document and extract the hyperlinks
	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	// Check if the output file flag is set
	if *outputFile != "" {
		// Create the output file
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// Write the list of hyperlinks to the output file
		for _, link := range links {
			file.WriteString(link + "\n")
		}
	} else {
		// Print the list of hyperlinks to standard output
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
