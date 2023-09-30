package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Serve the "index.html" file
		http.ServeFile(w, r, "html/index.html")
	})

	// Serve files
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))

	log.Fatal(http.ListenAndServe(":80", nil))
}
