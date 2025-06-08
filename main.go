package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const port = "8080"

func main() {

	server := &http.Server{
		Addr:    ":" + port,
		Handler: routes(),
	}

	log.Fatal(server.ListenAndServe())
}

func routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", handler)

	return r
}

type responsejson struct {
	Error   bool        `json:"error"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := responsejson{
		Error:   false,
		Message: "Hi Request recieved",
	}
	output, _ := json.MarshalIndent(response, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(output)

}
