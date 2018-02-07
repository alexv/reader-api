package main

import (
	"encoding/json"
	"net/http"

	"github.com/mmcdole/gofeed"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL("http://www.producthunt.com/feed")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(feed)
	})
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}
