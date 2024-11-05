package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello, htmx!")
	if err != nil {
		return
	}
}

func main() {
	log.Println("Info:", "Starting server...")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	port := "19080"
	log.Println("Info:", "Server started on http://localhost:"+port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalln("Error", "Server failed to start", err)
		return
	}
}
