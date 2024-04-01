package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	Routes := GetRoutes()

	for path, handler := range Routes {
		fmt.Println("Adding route", path, "to server.")
		http.HandleFunc(path, handler)
	}

	if len(os.Args) > 1 {
		mode := os.Args[1]
		if mode == "-D" {
			fmt.Println("Server running on: http://localhost:3000")
			http.ListenAndServe(":3000", nil)
		} else if mode == "-P" {
			fmt.Println("Server running on: https://maxvanasten.dev")
			err := http.ListenAndServeTLS(":443", "domain.cert.pem", "private.key.pem", nil)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Unknown 'mode' option, try -D for local development or -P for production.")
		}
	} else {
		fmt.Println("Please provide a 'mode' to run the server on, try -D for local development or -P for production.")
	}
}
