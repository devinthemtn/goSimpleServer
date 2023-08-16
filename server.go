package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func main() {
	curWD, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	// log.Printf("type of wd: %T", curWD)
	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", curWD, "the directory of the static files to serve")
	flag.Parse()

	http.Handle("/", loggingHandler(http.FileServer(http.Dir(*directory))))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
