package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
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

func getListener() (net.Listener, error) {
	socketPath := os.Getenv("SOCKET_PATH")
	if socketPath != "" {
		return net.Listen("unix", socketPath)
	}

	listenAddress := os.Getenv("LISTEN_ADDRESS")
	if listenAddress != "" {
		return net.Listen("tcp", listenAddress)
	}

	// I would remove this variable in favor of LISTEN_ADDRESS,
	// but that would break compatibility.
	port := getVar("PORT", "8080")
	return net.Listen("tcp", ":"+port)
}

func handleInterrupt(srv *http.Server, connsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	// We received an interrupt signal, shut down.
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server shutdown: %v", err)
	}

	close(connsClosed)
}

func main() {
	var srv http.Server

	listener, err := getListener()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server at " + listener.Addr().String())

	http.HandleFunc("/.well-known/mta-sts.txt", handleRequest)
	http.HandleFunc("/", http.NotFound)

	connsClosed := make(chan struct{})
	go handleInterrupt(&srv, connsClosed)

	if err = srv.Serve(listener); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-connsClosed
}
