package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	pathVersionInfo   = "/{module:.+}/@v/{version}.info"
	pathVersionModule = "/{module:.+}/@v/{version}.mod"
	pathVersionZip    = "/{module:.+}/@v/{version}.zip"
	pathList          = "/{module:.+}/@v/list"
	pathLatest        = "/{module:.+}/@latest"
)

func main() {
	proxyURL := os.Getenv("GOPROXY")
	if proxyURL == "" {
		log.Fatalf("no GOPROXY set!")
	}
	r := mux.NewRouter()
	r.Handle(pathVersionInfo, redirTo(proxyURL)).Methods("GET")
	r.Handle(pathVersionModule, redirTo(proxyURL)).Methods("GET")
	r.Handle(pathVersionZip, redirTo(proxyURL)).Methods("GET")

	stg := NewStorage()
	r.Handle(pathList, list(proxyURL, stg)).Methods("GET")
	r.Handle(pathLatest, latest(proxyURL, stg)).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Print("Serving on port 8081")
	if err := http.ListenAndServe(":8081", loggedRouter); err != nil {
		log.Printf("The server crashed! (%s)", err)
		os.Exit(1)
	}
}
