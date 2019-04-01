package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model Structure

type Model struct {
	ModelName string `json:"model_name"`
	MinWorker int8   `json:"workers", string`
}

// Declare size of models
var models []Model

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct{ Status string }{Status: "healthy"})
}
func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
func scaleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	min_worker := r.URL.Query().Get("min_worker")
	params := mux.Vars(r)
	for _, item := range models {
		if item.ModelName == params["model_name"] {
			tmp, _ := strconv.ParseInt(min_worker, 10, 64)
			item.MinWorker = int8(tmp)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func descriptionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params
	for _, item := range models {
		if item.ModelName == params["model_name"] {

			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Model{})
}
func main() {
	//Init Router
	router := mux.NewRouter()

	models = append(models, Model{ModelName: "resnet"})

	//Route handlers
	router.HandleFunc("/ping", pingHandler).Methods("GET")
	router.HandleFunc("/models", registerHandler).Methods("POST")
	router.HandleFunc("/models/{model_name}", scaleHandler).Methods("PUT")
	router.HandleFunc("/models/{model_name}", descriptionHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
