package main

import (
	"fmt"
	"homecook/sorting/mapsort"  // correct import path as needed
	"sort"
)

// Maps to be sorted can either be initialized as type MsFmtMap (MapSort Format Map -
// map[interface{}]interface{}), preformatted for use with the MapSort function
var testMap1 = make(mapsort.MsFmtMap)

// Or they can be initialized as a standard map format, and formatted later
// Alternatively, they do not need to be formatted at all (see func alt below)
var testMap2 = map[string]float64{
	"a": 9.0,
	"f": 7.5,
	"k": 12.2,
	"b": 10.4,
	"z": 2.9,
}

func main() {
	var fmtTM2 = make(mapsort.MsFmtMap) // corresponds to and used to format testMap2
	// populate testMap1
	for i := 0; i < 10; i++ {
		testMap1[i] = "flame emoji"
	}
	// format testMap2 by copying k/v pairs to fmtM2
	for k, v := range testMap2 {
		fmtTM2[k] = v
	}

	// initialize KeySort and ValSort instances
	var ks mapsort.KeySort
	var vs mapsort.ValSort

	// Sort the formatted maps
	// keySorted, valSorted are sorted slices of structs
	// Each struct in slice represents a key/value pair
	keySorted := ks.MapSort(testMap1)
	valSorted := vs.MapSort(fmtTM2)

	// Print the maps' data by accessing the struct fields of each k/v pair
	fmt.Println("***** Sorted by Key *****")
	for _, e := range keySorted {
		fmt.Printf("Key: %v, Value: %v\n", e.Key, e.Val)
	}
	fmt.Println()
	fmt.Println("***** Sorted by Value *****")
	for _, e := range valSorted {
		fmt.Printf("Key: %v, Value: %v\n", e.Key, e.Val)
	}
}

// alt demonstrates how to sort a map without formatting it
func alt() {
	// initialize empty slice of []Entry (k/v pair) structs
	var ks mapsort.KeySort

	// range over map and create new Entry for each k/v pair
	for k, v := range testMap2 {
		ks = append(ks, mapsort.Entry{Key: k, Val: v})
	}

	// sort
	sort.Sort(ks)

	// print results by ranging over es []Entry slice
	for _, e := range ks {
		fmt.Printf("%v: %v\n", e.Key, e.Val)
	}
	return
}
