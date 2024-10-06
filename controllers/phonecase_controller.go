
package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"go-phonecase-crud/models"
)


var phoneCases []models.PhoneCase
var currentID int = 1

// CreatePhoneCase handles the creation of a new phone case
func CreatePhoneCase(w http.ResponseWriter, r *http.Request) {
	var newCase models.PhoneCase
	json.NewDecoder(r.Body).Decode(&newCase)
	newCase.ID = currentID
	currentID++
	newCase.Status = "available"
	phoneCases = append(phoneCases, newCase)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCase)
}

// GetAllPhoneCases returns all phone cases
func GetAllPhoneCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phoneCases)
}

// GetPhoneCaseByID returns a phone case by its ID
func GetPhoneCaseByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, phoneCase := range phoneCases {
		if phoneCase.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(phoneCase)
			return
		}
	}

	http.Error(w, "Phone case not found", http.StatusNotFound)
}

// UpdatePhoneCase updates an existing phone case by its ID
func UpdatePhoneCase(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, phoneCase := range phoneCases {
		if phoneCase.ID == id {
			phoneCases = append(phoneCases[:i], phoneCases[i+1:]...)

			var updatedCase models.PhoneCase
			json.NewDecoder(r.Body).Decode(&updatedCase)
			updatedCase.ID = id
			phoneCases = append(phoneCases, updatedCase)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedCase)
			return
		}
	}

	http.Error(w, "Phone case not found", http.StatusNotFound)
}

// DeletePhoneCase deletes a phone case by ID
func DeletePhoneCase(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, phoneCase := range phoneCases {
		if phoneCase.ID == id {
			phoneCases = append(phoneCases[:i], phoneCases[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Phone case not found", http.StatusNotFound)
}
