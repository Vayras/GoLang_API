package server

import (
	"encoding/json"
	"fmt"
	"gormlogin/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func HandlePeople(w http.ResponseWriter, r *http.Request) {
	people, err := model.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Convert people to JSON
	peopleJSON, err := json.Marshal(people)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(peopleJSON)
}

func HandleCreatePerson(w http.ResponseWriter, r *http.Request) {
	var p model.Person
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	if err := model.CreatePerson(p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the user ID from the URL
	idParam := mux.Vars(r)["id"]
	fmt.Println("Extracted ID parameter:", idParam)
	idParam = strings.Trim(idParam, "{}")

	num, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Converted int:", num)

	// Call the DeletePersonByID function with the retrieved user ID
	if err := model.DeletePersonByID(num); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}

func PersonFinder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	idParam := mux.Vars(r)["id"]
	fmt.Println("Extracted ID parameter:", idParam)
	idParam = strings.Trim(idParam, "{}")

	num, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Converted int:", num)

	person, err := model.FindPersonByID(num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	personJSON, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(personJSON)
}

func FinderByNameAndPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	nameParam := mux.Vars(r)["name"]
	fmt.Println("Extracted ID parameter:", nameParam)
	nameParam = strings.Trim(nameParam, "{}")

	passParam := mux.Vars(r)["password"]
	fmt.Println("Extracted ID parameter:", passParam)
	passParam = strings.Trim(passParam, "{}")

	person, err := model.FindPersonByNameAndPassword(nameParam, passParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	personJSON, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(personJSON)
}
