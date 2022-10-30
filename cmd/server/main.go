package main

import (
	"fmt"
	"net/http"
)

const port = "9090"

func main() {
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), http.FileServer(http.Dir("../../assets")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}

	fmt.Printf("Starting server on port %s..", port)
}
