The Golang Omdb API
=======
[![Build Status](https://travis-ci.org/ehime/gomdb.svg?branch=master)](https://travis-ci.org/ehime/go-omdb)
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
	"github.com/ehime/gomdb"
)

func main() {
	query := &gomdb.QueryData{Title: "Macbeth", SearchType: gomdb.MovieSearch}
	res, err := gomdb.Search(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Search)

	query = &gomdb.QueryData{Title: "Macbeth", Year: "2015"}
	res2, err := gomdb.MovieByTitle(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res2)

	res3, err := gomdb.MovieByImdbID("tt2884018")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res3)
}
```


See the project documentation to see the Response Objects and stuff
