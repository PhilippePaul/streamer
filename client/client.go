package main

import (
	"log"
	"net/http"
)

func main() {
	runHttp()
}

func runHttp() {
	fs := http.FileServer(http.Dir("/home/phil/Stream"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8080", nil)
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
