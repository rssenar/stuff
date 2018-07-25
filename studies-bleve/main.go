package main

import (
	"fmt"

	"github.com/blevesearch/bleve"
)

func main() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("./BleveIndex/CustomerDB", mapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name    string
		Address string
		City    string
		State   string
		Zip     string
	}{
		Name:    "Richard Senar",
		Address: "3219 Hannover St",
		City:    "Corona",
		State:   "CA",
		Zip:     "92882",
	}

	data2 := struct {
		Name    string
		Address string
		City    string
		State   string
		Zip     string
	}{
		Name:    "Gabie Senar",
		Address: "3220 Hannover St",
		City:    "Irvine",
		State:   "CA",
		Zip:     "92618",
	}

	data3 := struct {
		Name    string
		Address string
		City    string
		State   string
		Zip     string
	}{
		Name:    "Noah Senar",
		Address: "125 Anywhere St",
		City:    "Ney York",
		State:   "NY",
		Zip:     "10125",
	}

	// index some data
	index.Index("A12345", data)
	index.Index("A12346", data2)
	index.Index("A12347", data2)
	index.Index("A12348", data3)
	index.Index("A12349", data)
	index.Index("A12350", data)
	index.Index("A12351", data)
	index.Index("A12352", data)

	// search for some text
	query := bleve.NewMatchQuery("York")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)
}
