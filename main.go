package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/somen440/slim-go/app"
)

func main() {
	var addr = flag.String("addr", ":8080", "application address")
	flag.Parse()

	log.Println("Web ", *addr)
	srv := &http.Server{
		Handler:      app.Routes(),
		Addr:         *addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
