package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	path := flag.String("path", "", "path to the files")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*path)))
	fmt.Println("The server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
