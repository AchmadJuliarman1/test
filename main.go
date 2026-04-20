package main

import (
	"fmt"
	"net/http"
)

// handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// route endpoint "/"
	http.HandleFunc("/", helloHandler)

	// jalankan server di port 8080
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
