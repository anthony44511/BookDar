package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running at http://localhost:8081")
	http.ListenAndServe(":8081", nil)
}
