package main

import (
	"fmt"
	"github.com/ehime/go-omdb/omdb"
)

func main() {
	query := &omdb.QueryData{Title: "Macbeth", SearchType: omdb.MovieSearch}
	res, err := omdb.Search(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Search)

	query = &omdb.QueryData{Title: "Macbeth", Year: "2015"}
	res2, err := omdb.MovieByTitle(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res2)

	res3, err := omdb.MovieByImdbID("tt2884018")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res3)
}