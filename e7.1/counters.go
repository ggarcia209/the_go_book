// Create interfaces and methods to count words, bytes, and lines
// Create interfaces and methods to count words, bytes, and lines
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ByteCounter represents a string's number of bytes
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

// WordCounter represents a string's number of words
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	w := func() int {
		ct := 0
		s := bufio.NewScanner(bytes.NewReader(p))
		s.Split(bufio.ScanWords)
		for s.Scan() {
			ct++
		}
		return ct
	}
	*c = WordCounter(w())
	return w(), nil
}

// LineCounter represents a string's number of lines
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	l := func() int {
		ct := 0
		s := bufio.NewScanner(bytes.NewReader(p))
		for s.Scan() {
			ct++
		}
		return ct
	}
	*c = LineCounter(l())
	return l(), nil
}

func main() {
	t := "hello" // initial input

	input, err := ioutil.ReadAll(os.Stdin) //
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// byte count
	var b ByteCounter
	b.Write([]byte(t)) // var t byte count
	fmt.Printf("bytes: %d\n", b)

	b = 0                          // reset counter
	fmt.Fprintf(&b, string(input)) // stdin byte count
	fmt.Printf("\tnew write: %v\n", b)

	// word count
	var w WordCounter
	w.Write([]byte(t)) // var t word count
	fmt.Printf("words: %d\n", w)

	w = 0                          // reset
	fmt.Fprintf(&w, string(input)) // stdin word count
	fmt.Printf("\tnew write: %v\n", w)

	// line count
	var l LineCounter
	l.Write([]byte(t)) // var t line count
	fmt.Printf("lines: %d\n", l)

	l = 0                          // reset
	fmt.Fprintf(&l, string(input)) // stdin line count
	fmt.Printf("\tnew write: %v\n", l)
}

