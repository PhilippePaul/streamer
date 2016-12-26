package main

import (
	"log"
	"net/http"
	"os/exec"
)

func main() {
	runHttp()
}

type PlayHandler struct{}

var handler PlayHandler

func runHttp() {
	log.Fatal(http.ListenAndServe(":8080", http.StripPrefix("/play/", handler)))
}

func (handler PlayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	runCmd(r.URL.Path)
}

func runCmd(filename string) {
	cmd := exec.Command("vlc", "http://192.168.0.107:8080/"+filename)
	err := cmd.Run()
	check(err)
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
