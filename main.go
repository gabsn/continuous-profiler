package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type App struct {
	Name  string
	Hosts []string
}

type Index struct {
	Apps []App
}

func newMockIndex() Index {
	return Index{
		Apps: []App{
			App{
				Name:  "destro",
				Hosts: []string{"i-0383f57df06a90be1", "i-09cdd8858d79497fe", "i-0b66d9d3f18af8514"},
			},
			App{
				Name:  "spidly",
				Hosts: []string{"i-09cdd8858d79497fe", "i-0b66d9d3f18af8514"},
			},
			App{
				Name:  "propjoe",
				Hosts: []string{"i-0c98f18a8d1d221cf", "i-0dcb66a8bd9dae274", "i-0c930e3e64c4d682a", "i-09cdd8858d79497fe", "i-0b66d9d3f18af8514"},
			},
		},
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index := newMockIndex()

	js, err := json.Marshal(index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func flamegraphHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./flamegraph.svg")
}

func main() {
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/flamegraph", flamegraphHandler)
	log.Fatal(http.ListenAndServe(":8765", nil))
}
