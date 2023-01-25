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

func getMtaSts(w http.ResponseWriter, r *http.Request) {
	// Headers
	w.Header().Set("Content-Type", "text/plain")

	// STS Version
	io.WriteString(w, "version: STSv1\n")

	// Mode
	mode := getVar("MTA_STS_MODE", "enforce")
	io.WriteString(w, "mode: "+mode+"\n")

	// Max age
	maxAge := getVar("MTA_STS_MAX_AGE", "604800")
	io.WriteString(w, "max_age: "+maxAge+"\n")

	// MX servers
	mxServers := strings.Split(getVar("MTA_STS_MX", "mx.change.me"), ",")
	for _, server := range mxServers {
		io.WriteString(w, "mx: "+server+"\n")
	}
}

func main() {
	port := getVar("PORT", "8080")
	log.Println("Starting server at localhost:" + port)
	http.HandleFunc("/.well-known/mta-sts.txt", getMtaSts)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
