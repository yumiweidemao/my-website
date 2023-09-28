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

	// Serve the images from the "images" folder
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	log.Fatal(http.ListenAndServe(":80", nil))
}
