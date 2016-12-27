package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:         ":3000",
		Handler:      http.HandlerFunc(handler),
		WriteTimeout: 2 * time.Second,
	}
	err := s.ListenAndServeTLS("contrib/localhost.crt", "contrib/localhost.key")
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, 世界\n"))
}
