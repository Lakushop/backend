package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lakushop/backend/router"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8090...")

	log.Fatal(http.ListenAndServe(":8090", r))
}
