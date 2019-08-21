// clockwall streams time data from multiple instances of clock2 servers
// and prints each server's time at every second, until the connection is closed.
// Run clockwall binary when finished executing clock2 binary instances
// Set locations/timezones and corresponding port numbers as command line arguments
// Ex: $ TZ=Asia/Tokyo clock2/clock2 -port 8030 &
//     $ clockwall/clockwall LosAngeles=8010 NewYork=8020 Tokyo=8030
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sort"
	"sync"
	"time"

	"homecook/sorting/mapsort" // correct this import path as needed - mapsort package available in e8.1 directory
)

var clocks = make(map[string][]byte)

var mu sync.Mutex

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		tz, port := getVar(arg)
		go getTime(tz, port)  // goroutine retrieves time values from each server concurrently
	}
	for {
		showTimes()
		time.Sleep(1 * time.Second)
	}
}

// getVar parses command line argument and derives
// var name ('tz' - timezone) and value ('port' - tcp port)
func getVar(dec string) (string, string) {
	for i, s := range dec {
		if s == strToRune("=") {
			tz := dec[:i]     // timezone - left side of "="
			port := dec[i+1:] // port number - right side of "="
			return tz, port
		}
	}
	return "", ""
}

// strToRune converts string character to rune
// in order to test equivalence in func getVar
func strToRune(s string) rune {
	rs := []rune(s)
	r := rs[0]
	return r
}

// getTime retrieves current time at every second
// from the clock2 server at the specified port
func getTime(tz, port string) {
	buf := make([]byte, 9) // exact size of timestamp in bytes
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		mu.Lock()
		clocks[tz] = buf[:n] // update map to reflect current time
		mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}

// showTimes prints the current time received from each clock2 server instance
// and uses mapsort.ValSort interface to sort time values from least to greatest
func showTimes() {
	mu.Lock()
	var es mapsort.ValSort
	for k, v := range clocks {
		es = append(es, mapsort.Entry{Key: k, Val: string(v)})
	}
	sort.Sort(es)
	for _, e := range es {
		fmt.Printf("%s local time: %s", e.Key, e.Val)
	}
	fmt.Println("---------------")
	mu.Unlock()
}
