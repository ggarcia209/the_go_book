/* Write a function to populate a mapping from element names (p, div, span)
to the number of elements with that name in an HTML document tree */

// mapHtmlNode prints the count of each element in an HTML document read from the standard input
// run as (in $GOPATH/bin) './fetch <url> | ./mapHtmlNode'
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	var eleMap = make(map[string]int)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elements: &v\n, err")
		os.Exit(1)
	}
	for k, v := range visit(eleMap, doc) {
		fmt.Println(k, v)
	}
}

// visit maps the count of each element found in n and returns the result.
func visit(eleMap map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		eleMap[n.Data]++
	}
	eleMap, n = nextNode(eleMap, n.FirstChild)
	return eleMap
}

// nextNode uses recursion to traverse the FirstChild/NextSiblink linked list
func nextNode(eleMap map[string]int, current *html.Node) (map[string]int, *html.Node) {
	if current != nil {
		eleMap = visit(eleMap, current)
		next := current.NextSibling
		return nextNode(eleMap, next)
	}
	return eleMap, current
}
