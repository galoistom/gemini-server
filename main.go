package main

import "log"

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	router := NewRouter()
	router.Handle("/", homeHandler)
	router.Handle("/about", aboutHandler)
	router.Handle("/hello", helloHandler)

	server := &Server{
		Addr:         ":1965",
		Handler:      router,
		ReadTimeOut:  5,
		WriteTimeOut: 5,
		CertFile:     "cert/server.crt",
		KeyFile:      "cert/server.key",
	}
	log.Println("starting gemini server ...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
