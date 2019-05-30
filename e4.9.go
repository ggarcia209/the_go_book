package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%v\n", wordFreq())
}

func wordFreq() map[string]int {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	count := make(map[string]int)
	for in.Scan() {
		count[in.Text()]++
	}
	return count
}
