// func removeDups removes adjacent duplicate characters in []string slice in place
package main

import (
	"fmt"
	"strings"
)

var str = strings.Split("aabccddee", "")

func removeDups(str []string) []string {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] {
			str = append(str[:i], str[i+1:]...)
		}
	}
	return str
}

func main() {
	fmt.Printf("No duplicates: %s\n", removeDups(str))
}
