/* Change the 'findlinks1' program to traverse the 'n.FirstChild' linked list 
using recursive calls to 'visit' instead of a 'for' loop */

// Findlinks1 prints the links in an HTML document read from the standard input
// run as (in $GOPATH/bin) './fetch <url> | ./findlinks1.5'
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: &v\n, err")
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to 'links' each link found in 'n' and returns the result.
// Uses recursion to traverse the tree
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links, n = nextNode(links, n.FirstChild)
	return links
}

// nextNode uses recursion to traverse the FirstChild/NextSibling linked list
func nextNode(links []string, current *html.Node) ([]string, *html.Node) {
	if current != nil {
		links = visit(links, current)
		next := current.NextSibling
		return nextNode(links, next)
	}
	return links, current
}

/* Original visit function as found in Ch5.2 of "The Go Programming Language" (Donovan, Kernighan)
Not used - for comparison only */

// visit appends to 'links' each link found in 'n' and returns the result.
// Uses for loop and recursion to traverse the tree
func visitOG(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visitOG(links, c)
	}
	return links
}
