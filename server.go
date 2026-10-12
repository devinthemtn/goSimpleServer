package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

const redX = "\033[1;31m✘\033[0m"

// statusRecorder captures the status code written by the wrapped handler.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func (s *statusRecorder) Unwrap() http.ResponseWriter {
	return s.ResponseWriter
}

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		h.ServeHTTP(rec, r)
		if rec.status >= 400 {
			log.Println(redX, r.Method, r.URL.Path, rec.status)
		} else {
			log.Println(r.Method, r.URL.Path)
		}
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
