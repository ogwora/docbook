package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static Webserver
	log.Fatal(http.ListenAndServe(":80",
		http.FileServer(http.Dir("/perm/tdg"))))
}
