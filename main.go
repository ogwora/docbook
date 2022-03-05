package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static Webserver
	log.Fatal(http.ListenAndServe(":8080",
		http.FileServer(http.Dir("/perm/tdg"))))
}
