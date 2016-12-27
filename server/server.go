package main

import (
	"log"
	"net/http"
	"os/exec"
)

type PlayHandler struct{}
type StreamHandler struct{}

var (
	playHandler   PlayHandler
	streamHandler StreamHandler
)

func main() {
	//	runHttp()
	runTLS()
}

func runTLS() {
	http.Handle("/stream/", http.StripPrefix("/stream/", streamHandler))
	log.Fatal(http.ListenAndServeTLS(":8080", "./cert.pem", "./key.pem", nil))
}

func runHttp() {
	http.Handle("/play/", http.StripPrefix("/play/", playHandler))
	http.Handle("/stream/", http.StripPrefix("/stream/", streamHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (handler PlayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	runCmd("http://192.168.0.107:8080/" + r.URL.Path)
}

func (handler StreamHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	runCmd(r.URL.Path)
}

func runCmd(path string) {
	log.Println(path)
	cmd := exec.Command("vlc", "http://youtube.com/watch?"+path)
	err := cmd.Run()
	check(err)
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
