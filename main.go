package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/aelpxy/avatar-generator/lib"
)

const (
	PORT string = ":5000"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b := make([]byte, 32)
		rand.Read(b)

		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write([]byte(lib.GenerateImage(hex.EncodeToString(b))))
	})

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		query := params.Get("query")

		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write([]byte(lib.GenerateImage(query)))
	})

	log.Printf("Listening on http://0.0.0.0%s \n", PORT)
	http.ListenAndServe(PORT, nil)
}
