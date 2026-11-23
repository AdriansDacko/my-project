package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type dati struct {
	Light       float32 `json:"light"`
	Humidity    float32 `json:"humidity"`
	Temperature float32 `json:"temperature"`
}

func processData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var d dati

	err = json.Unmarshal(body, &d)
	if err != nil {
		panic(err)
	}

	fmt.Println(d)

}

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	http.HandleFunc("/sensor/data", processData)

	log.Println("listening on 0.0.0.0:5000")
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}

