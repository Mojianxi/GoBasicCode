package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello beifeng!"))
	})
	http.ListenAndServe("127.0.0.1:8030", nil)
}
