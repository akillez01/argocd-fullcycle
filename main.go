package main

import "net/http"

func main() {

	// Set up a simple HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello, Achilles!</h1>"))
	})

	http.ListenAndServe(":8080", nil)
}
