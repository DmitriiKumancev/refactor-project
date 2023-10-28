package main

import (
	"log"
	"net/http"

	"github.com/DmitriiKumancev/refactor-project/api"
	"github.com/DmitriiKumancev/refactor-project/storage"
)

func main() {
	storage.InitStore()

	r := api.NewRouter()
	http.Handle("/", r)

	addr := ":3333"
	log.Printf("Server is running on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
