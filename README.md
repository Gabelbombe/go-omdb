The Golang Omdb API
=======
[![Build Status](https://travis-ci.org/ehime/go-omdb.svg?branch=master)](https://travis-ci.org/ehime/go-omdb)
[![GoDoc](https://godoc.org/github.com/ehime/go-omdb?status.svg)](https://godoc.org/github.com/ehime/go-omdb)


Author: Jd Daniel (dodomeki AT gmail DOT com)

<iframe src="http://githubbadge.appspot.com/ehime" style="border: 0;height: 142px;width: 200px;overflow: hidden;" frameBorder="0"></iframe>

This API uses the [omdbapi.com](http://omdbapi.com/) API by Brian Fritz

***
### OMDBAPI.com
This is an excellent open database for movie and film content.

I *strongly* encourage you to check it out and contribute to keep it growing.

### http://www.omdbapi.com
***
Project Usage
-------------
The API usage is very simple.

```go
package main

import (
	"fmt"
	"github.com/ehime/go-omdb"
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
```


See the project documentation to see the Response Objects and stuff
