package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/Login", login)
	http.HandleFunc("/Home", home)
	http.HandleFunc("/Refresh", refresh)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
