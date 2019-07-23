/* Many GUI's provide a table widget with a stateful multi-tier sort:
the primary sort key is th emost recently clicked column head,
the secondary sortkey is the second-most recently clicked column head,
and so on. Define an implementation of 'sort.Interface' for use by
such a table. */

package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// Track represents a song and its data
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// Entry stores column header name, its corresponding sort key function,
// and it's position in the sort order.
type Entry struct {
	Key  string
	Func func(x, y *Track) (bool, int)
	Val  int
}

// Custom Sort interface
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// Sort Entries interface
type entries []Entry

func (x entries) Len() int           { return len(x) }
func (x entries) Less(i, j int) bool { return x[i].Val < x[j].Val }
func (x entries) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// sortOrder records the most recently "clicked" column headers
// and is initialized with a default order.
var sortOrder = map[string]int{
	"Title":  0,
	"Artist": 1,
	"Album":  2,
	"Year":   3,
	"Length": 4,
}

// sortFuncs maps the sort key functions and is used in conjunction 
// with sortOrder to generate primary, secondary sort keys and so on...
var sortFuncs = map[string]func(x, y *Track) (bool, int){
	"Title":  sorttitle,
	"Artist": sortartist,
	"Album":  sortalbum,
	"Year":   sortyear,
	"Length": sortlength,
}

// List of songs to be sorted
var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"No One", "Alicia Keys", "As I Am", 2007, length("4m14s")},
}

// Converts time string value in Track to type time.Duration
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {
	selectKey("Artist")
	// selectKey("Length")  // try one at a time to simulate clicking column head in browser
	// selectKey("Album")
	// selectKey("Title")
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		var b bool // records whether x.(attr) < y.(attr)
		var p int  // records whether values are inequal, pass to next key if == 0 (equal)
		es := sortFuncKeys()
		for _, e := range es {
			b, p = e.Func(x, y)
			if p == 1 { // values are not equal
				return b
			}
		}
		return false
	}})
	printTracks(tracks)
}

// Select primary sort key and maintain order of recency
// by incrementing sortOrder values of the other keys
func selectKey(head string) map[string]int {
	for k := range sortOrder {
		if k == head {
			sortOrder[k] = 0
		} else {
			sortOrder[k]++
		}
	}
	return sortOrder
}

// Create ordered list of sort keys (by sortOrder values) using 'entries' interface
func sortFuncKeys() entries {
	var es entries
	for k, v := range sortOrder {
		es = append(es, Entry{Key: k, Val: v, Func: sortFuncs[k]})
	}
	sort.Sort(es)
	return es
}

// Print sorted Tracks as table
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

// Each of the following functions acts as a sort key. If 'false, 0' is returned,
// values are equivalent and Tracks are passed to next function (key) in sort order
func sorttitle(x, y *Track) (bool, int) {
	if x.Title != y.Title {
		return x.Title < y.Title, 1
	}
	return false, 0
}

func sortartist(x, y *Track) (bool, int) {
	if x.Artist != y.Artist {
		return x.Artist < y.Artist, 1
	}
	return false, 0
}

func sortalbum(x, y *Track) (bool, int) {
	if x.Album != y.Album {
		return x.Album < y.Album, 1
	}
	return false, 0
}

func sortyear(x, y *Track) (bool, int) {
	if x.Year != y.Year {
		return x.Year < y.Year, 1
	}
	return false, 0
}

func sortlength(x, y *Track) (bool, int) {
	if x.Length != y.Length {
		return x.Length < y.Length, 1
	}
	return false, 0
}
