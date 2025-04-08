package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

const (
	APP_HOST = "127.0.0.1"
	APP_PORT = "3030"
)

//go:embed data/fufufafa.json
var fufa []byte

type QuotesDataStruct struct {
	Id   int64 `json:"id"`
	Data struct {
		Content  string `json:"content"`
		Datetime string `json:"datetime"`
		Doksli   string `json:"doksli"`
	} `json:"data"`
}

func main() {
	var data []QuotesDataStruct

	err := decodeData(&data)
	if err != nil {
		log.Printf("Failed to decode JSON data: %v", err)
		return
	}

	app := http.NewServeMux()

	app.Handle("GET /", http.RedirectHandler("/random", http.StatusPermanentRedirect))

	app.Handle("GET /random", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		defer func() {
			if m := recover(); m != nil {
				log.Printf("Panic woyy! %v\n", m)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Something went wrong"))
			}
		}()

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rn := rand.IntN(len(data) - 1)
		b, err := json.Marshal(data[rn])
		if err != nil {
			log.Printf("Failed to marshal, error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to process quote"))
			return
		}

		w.Write(b)
	}))

	serverAddr := fmt.Sprintf("%s:%s", APP_HOST, APP_PORT)
	log.Printf("Server starting on http://%s", serverAddr)

	err = http.ListenAndServe(serverAddr, app)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func decodeData(q *[]QuotesDataStruct) error {
	dec := json.NewDecoder(bytes.NewReader(fufa))
	err := dec.Decode(q)
	if err != nil {
		return fmt.Errorf("JSON decode error: %w", err)
	}
	return nil
}
