// e4.2.go prints the SHA256 checksum of it's standard input by default and
// supports command line flags to print the SHA384 or SHA512 checksums instead
// ("-t" for SHA384 and "-f" for SHA512)

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// create flags to call different SHA functions
var flag384 bool
var flag512 bool

func init() {
	flag.BoolVar(&flag384, "t", false, "SHA384")
	flag.BoolVar(&flag512, "f", false, "SHA256")
}

func main() {
	// get input as []byte and pass to SHA functions
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Print(err)
	}
	flag.Parse()
	// create conditions for flags
	if !flag384 && !flag512 {
		fmt.Printf("SHA256: %d\n", get256(in))
	} else if !flag512 {
		fmt.Printf("SHA384: %d\n", get384(in))
	} else {
		fmt.Printf("SHA512: %d\n", get512(in))
	}
}

// calculate SHA256 checksum
func get256(in []byte) [32]byte {
	sha := sha256.Sum256(in)
	return sha
}

// calculate SHA384 checksum
func get384(in []byte) [48]byte {
	sha := sha512.Sum384(in)
	return sha
}

// calculate SHA512 checksum
func get512(in []byte) [64]byte {
	sha := sha512.Sum512(in)
	return sha
}
