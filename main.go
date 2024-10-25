package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func getVar(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	io.WriteString(w, "version: STSv1\n")

	mode := getVar("MTA_STS_MODE", "enforce")
	io.WriteString(w, "mode: "+mode+"\n")

	maxAge := getVar("MTA_STS_MAX_AGE", "604800")
	io.WriteString(w, "max_age: "+maxAge+"\n")

	mxServers := strings.Split(getVar("MTA_STS_MX", "mx.change.me"), ",")
	for _, server := range mxServers {
		io.WriteString(w, "mx: "+server+"\n")
	}
}

func main() {
	port := getVar("PORT", "8080")
	log.Println("Starting server at localhost:" + port)
	http.HandleFunc("/.well-known/mta-sts.txt", handleRequest)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
