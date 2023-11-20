package main

import (
	"backend/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8069...")

	log.Fatal(http.ListenAndServe(":8069", r))
}
