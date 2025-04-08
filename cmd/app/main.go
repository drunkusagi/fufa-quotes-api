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
		log.Println("error decoding data")
		return
	}

	app := http.NewServeMux()

	app.Handle("GET /", http.RedirectHandler("/random", http.StatusPermanentRedirect))

	app.Handle("GET /random", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		rn := rand.IntN(len(data) - 1)
		b, err := json.Marshal(data[rn])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Add("Content-Type", "text/html")
			w.Write([]byte("oopss!"))
		}

		w.Write(b)
	}))

	log.Println("starting app at port: " + APP_PORT)
	http.ListenAndServe(fmt.Sprintf("%s:%s", APP_HOST, APP_PORT), app)
}

func decodeData(q *[]QuotesDataStruct) error {
	dec := json.NewDecoder(bytes.NewReader(fufa))
	err := dec.Decode(q)
	if err != nil {
		return err
	}

	return nil
}
