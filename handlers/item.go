package handlers

import (
	"encoding/json"
	"finanzas/misc/interfaces"
	"finanzas/misc/models"
	"finanzas/misc/repos"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
)

// ItemHandler is used to control all
// service interactions with the Item entity
type ItemHandler struct {
	db       *gorm.DB
	itemRepo interfaces.ItemRepo
}

// NewItemHandler returns a usable reference to an ItemHandler
func NewItemHandler(db *gorm.DB) *ItemHandler {
	return &ItemHandler{
		db:       db,
		itemRepo: repos.NewItemRepo(db),
	}
}

// FindByID is a GET handler that receives the ID of an item and call the
// respective method to get the corresponding item.
func (i *ItemHandler) FindByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		// json.NewEncoder(w).Encode(events)
	}

	item, err := i.itemRepo.FindByID(id)
	if err != nil {
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		// json.NewEncoder(w).Encode(events)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(item)
}

// Find is a POST handler that receives a JSON object that can be used to
// query items that match the supplied parameters.
func (i *ItemHandler) Find(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		// json.NewEncoder(w).Encode(events)
	}

	query := &models.Item{}
	err = json.Unmarshal(body, &query)
	if err != nil {
		w.WriteHeader(400)
	}

	items, err := i.itemRepo.Find(query)
	if err != nil {
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		// json.NewEncoder(w).Encode(events)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(items)
}

// Create is a POST handler that receives a JSON object that will be used
// to create a new Item record in the database.
func (i *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		// json.NewEncoder(w).Encode(events)
	}

	item := &models.Item{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		w.WriteHeader(400)
	}

	id, err := i.itemRepo.Create(item)
	if err != nil {
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		// json.NewEncoder(w).Encode(events)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(id)
}

// Update is a PUT handler that receies a JSON object that will be used
// to update an existing Item record in the database.
func (i *ItemHandler) Update(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		// json.NewEncoder(w).Encode(events)
	}

	item := &models.Item{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		w.WriteHeader(400)
	}

	err = i.itemRepo.Update(item)
	if err != nil {
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		// json.NewEncoder(w).Encode(events)
	}

	w.WriteHeader(202)
}
