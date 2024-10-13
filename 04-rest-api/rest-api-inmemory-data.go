package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// go get -u github.com/gorilla/mux
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

var (
	memoryStore = make(map[int]Item)
	idCounter   = 1
	mu          sync.Mutex
)

func getItems(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var items []Item
	for _, item := range memoryStore {
		items = append(items, item)
	}
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	item, exists := memoryStore[id]
	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = idCounter
	idCounter++
	memoryStore[item.ID] = item
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, exists := memoryStore[id]
	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	var updatedItem Item
	_ = json.NewDecoder(r.Body).Decode(&updatedItem)
	updatedItem.ID = id
	memoryStore[id] = updatedItem
	json.NewEncoder(w).Encode(updatedItem)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, exists := memoryStore[id]
	if !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	delete(memoryStore, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := mux.NewRouter()

	// Endpointy pro REST API
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

