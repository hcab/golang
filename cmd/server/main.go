package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yourusername/my-go-webapp/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HelloHandler)
	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}