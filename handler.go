package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, %q", html.EscapeString(r.URL.Path))
 }

 func TodoIndex(w http.ResponseWriter, r *http.Request) {
 		todos := Todos{
 			Todo{Name: "Write Presentation"},
 			Todo{Name: "Host Meetup"},
 		}

 		json.NewEncoder(w).Encode(todos)
 }

 func TodoShow(w http.ResponseWriter, r *http.Request) {
 		vars := mux.Vars(r)
 		todoID := vars["todoID"]
 		fmt.Fprintln(w, "Todo Show:", todoID)
 }