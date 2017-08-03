package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^([[:alpha:]]+)@golang.org$")
	match := re.FindStringSubmatch(r.URL.Path)
	if len(match) == 1 {
		fmt.Fprintf(w, "hello, %s", match[1])
		return
	}
	fmt.Fprintln(w, "hello, stranger")
}
