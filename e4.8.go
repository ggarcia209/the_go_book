package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	lCounts := make(map[rune]int)
	pCounts := make(map[rune]int)
	nCounts := make(map[rune]int)
	sCounts := make(map[rune]int)   // Counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			lCounts[r]++
		} else if unicode.IsNumber(r) {
			nCounts[r]++
		} else if unicode.IsPunct(r) {
			pCounts[r]++
		} else if unicode.IsSymbol(r) {
			sCounts[r]++
		}

		utflen[n]++
	}
	fmt.Printf("letter\tcount\n")
	for c, n := range lCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("number\tcount\n")
	for c, n := range nCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("punctuation\tcount\n")
	for c, n := range pCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("symbol\tcount\n")
	for c, n := range sCounts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
