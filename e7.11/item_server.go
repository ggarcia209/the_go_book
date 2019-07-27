/* Create server with handlers to enable clients to
Create, Read, Update and Delete inventory database entries.
Ex ("http://localhost:8000/update?item=shirts&price=15") */

package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/boltdb/bolt"
	// "homecook/conv"  // imported functions' source code at bottom
)

// Lock when creating/updating/deleting db values
var mutex = &sync.Mutex{}

func main() {
	// create "db" directory if not exists
	if _, err := os.Stat("db"); os.IsNotExist(err) {
		os.Mkdir("db", 0755)
	}

	dbMap := make(database) // in memory store

	// load db into in-memory map
	db, err := bolt.Open("db/inventory.db", 0755, nil)
	if err != nil {
		log.Fatal(err)
	}

	// read/write transaction
	if err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("inventory"))
		if err != nil {
			return fmt.Errorf("could not load database\n%v", err)
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			dbMap[string(k)] = dollars(BytesToFloat64(v)) // decode bytes to type dollars
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	db.Close()

	http.HandleFunc("/list", dbMap.list)
	http.HandleFunc("/price", dbMap.price)
	http.HandleFunc("/update", dbMap.update)
	http.HandleFunc("/delete", dbMap.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

/* dollars interface */
type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

/* database interface */
type database map[string]dollars

// List (Read) all items in database
func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price.String())
	}
}

// Read price for specified item
func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price.String())
}

// Create new item or Update existing entry
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := strings.ToLower(req.URL.Query().Get("item"))
	price := req.URL.Query().Get("price")
	if price == "" {
		fmt.Fprintf(w, "error: price not set")
		return
	}

	// convert string from URL to float64 and check value
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Fprintf(w, "error: price must be numerical value")
		return
	}

	// convert to dollars and verify price is >= 0
	dv := dollars(p)
	if dv < 0 {
		fmt.Fprintf(w, "error: price must be greater than or equal to 0")
		return
	}

	mutex.Lock()
	oldb, err := bolt.Open("db/inventory.db", 0755, nil) // offline database
	if err != nil {
		fmt.Fprintf(w, "error: offline database could not be opened; try again\n%v", err)
		return
	}

	// Create/Update transaction
	if err := oldb.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("inventory"))
		if err != nil {
			return fmt.Errorf("offline database could not be opened; try again\n%v", err)
		}
		if err := b.Put([]byte(item), Float64ToBytes((p))); err != nil { // serialize k,v
			return fmt.Errorf("could not update; try again\n%v", err)
		}
		return nil
	}); err != nil {
		fmt.Fprintf(w, "error: data store unsuccessful\n%v", err)
		return
	}

	oldb.Close()
	db[item] = dv // store in memory after successful disk storage
	mutex.Unlock()

	fmt.Fprintf(w, "stored in database: %s: %s\n", item, db[item].String())
}

// Delete specified entry
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := strings.ToLower(req.URL.Query().Get("item"))
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	mutex.Lock()
	oldb, err := bolt.Open("db/inventory.db", 0755, nil) // offline database
	if err != nil {
		fmt.Fprintf(w, "error: offline database could not be opened; try again\n%v", err)
		return
	}

	// Delete transaction
	if err := oldb.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("inventory"))
		if err != nil {
			return fmt.Errorf("offline database could not be opened; try again\n%v", err)
		}
		if err := b.Delete([]byte(item)); err != nil {
			return fmt.Errorf("could not delete; try again\n%v", err)
		}
		return nil
	}); err != nil {
		fmt.Fprintf(w, "error: deletion unsuccessful\n%v", err)
		return
	}

	oldb.Close()
	delete(db, item) // update in memory after successful deletion from disk
	mutex.Unlock()

	fmt.Fprintf(w, "item deleted: %s\n", item)
}

/* "imported" functions from 'homecook/conv' (home-made utilities packages) */

// BytesToUint64 decodes a byte slice representing sinlge int value to type uint64
func BytesToUint64(bs []byte) uint64 {
	return binary.BigEndian.Uint64(bs)
}

// Float64ToBytes encodes a uint64 value to a byte slice
func Float64ToBytes(fl float64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, math.Float64bits(fl))
	return b
}

// BytesToFloat64 decodes a byte slice representing sinlge int value to type float64
func BytesToFloat64(bs []byte) float64 {
	return math.Float64frombits(BytesToUint64(bs))
}
