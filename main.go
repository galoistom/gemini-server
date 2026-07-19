package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	router := NewRouter()
	router.Handle("/", homeHandler)
	router.Handle("/about", aboutHandler)
	router.Handle("/hello", helloHandler)

	server := &Server{
		Addr:         ":1965",
		Handler:      router,
		ReadTimeOut:  5 * time.Second,
		WriteTimeOut: 5 * time.Second,
		CertFile:     "cert/server.crt",
		KeyFile:      "cert/server.key",
	}
	log.Println("starting gemini server ...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
