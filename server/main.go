package main

import (
	"backend-api/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", router.Router()))
	fmt.Println("Server running on port 8080")
}
