package main

import (
	"fmt"
	"log"
	"net/http"
)

type Storage struct {
	storageMap map[string]string
}

func (store *Storage) setHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	if len(queryParams) == 0 {
		http.Error(w, "No key value pairs provided", http.StatusBadRequest)
		return
	}

	for key, values := range queryParams {
		value := values[0]
		store.storageMap[key] = value
		fmt.Fprintf(w, "Set key '%s' to value '%s'\n", key, value)
	}
}

func (store *Storage) getHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	key := queryParams.Get("key")
	value := store.storageMap[key]

	if value == "" {
		http.Error(w, "No stored value for the key provided", http.StatusBadRequest)
		return
	}

	w.Write([]byte(value))
}

func main() {
	store := Storage{
		make(map[string]string),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/set", store.setHandler)
	mux.HandleFunc("/get", store.getHandler)

	fmt.Println("Starting server at :4000")
	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal(err)
	}
}
