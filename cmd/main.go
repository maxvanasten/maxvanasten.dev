package main

import (
	"fmt"
	"net/http"
)

func main() {
	Routes := GetRoutes()

	for path, handler := range Routes {
		fmt.Println("Adding route", path, "to server.")
		http.HandleFunc(path, handler)
	}

	fmt.Println("http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
