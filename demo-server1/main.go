package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ready", withLogging(withCORS(withJSON(handleReady))))
	mux.HandleFunc("/healthy", withLogging(withCORS(withJSON(handleHealthy))))
	mux.HandleFunc("/kill", handleKill)
	mux.HandleFunc("/", withLogging(withCORS(withJSON(handleRoot))))
	http.ListenAndServe(":3000", mux)
}

func handleReady(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"status": "READY"}`)
}

func handleHealthy(w http.ResponseWriter, r *http.Request) {
	if rand.Intn(2) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status": "HEALTHY"}`)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"status": "NOT HEALTHY"}`)
	}
}

func handleKill(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"status": "terminating"}`)
	defer os.Exit(1)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().UTC()
	type response = struct {
		Message   string
		Host      string
		IP        string
		Port      int
		Timestamp time.Time
	}
	addr, port := getOutboundIP()
	name, err := os.Hostname()
	if err != nil {
		name = "N/A"
	}
	body, _ := json.Marshal(response{
		"OK",
		name,
		net.IP.String(addr),
		port,
		timestamp,
	})
	fmt.Fprintln(w, string(body))
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timestamp := time.Now().UTC()
		next(w, r)
		diff := time.Now().Sub(timestamp)
		fmt.Println("[ACCESS]", timestamp, r.RemoteAddr, r.RequestURI, diff)
	}
}

func withJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(w, r)
	}
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		next(w, r)
	}
}

func getOutboundIP() (net.IP, int) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, localAddr.Port
}
