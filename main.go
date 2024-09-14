package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Patient data
type Patient struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// in-memory store for patients
var patients []Patient

// Function to check if a patient with a given ID already exists
func patientExists(id string) bool {
	for _, p := range patients {
		if p.ID == id {
			return true
		}
	}
	return false
}

func addPatient(w http.ResponseWriter, r *http.Request) {
	var patient Patient

	err := json.NewDecoder(r.Body).Decode(&patient)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if patientExists(patient.ID) {
		http.Error(w, "Patient with this ID already exists", http.StatusConflict)
		return
	}

	patients = append(patients, patient)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(patient)
}

func getPatients(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(patients)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Patient Service!")
}

func handlePantients(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getPatients(w, r)
	case "POST":
		addPatient(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/patients", handlePantients)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
