package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	file := r.Path("/files/{id}").Subrouter()
	file.Methods("GET").HandlerFunc(GetHandler)
	file.Methods("POST").HandlerFunc(CreateHandler)
	file.Methods("PUT").HandlerFunc(UpdateHandler)
	file.Methods("DELETE").HandlerFunc(DeleteHandler)

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}

func CreateHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	path := "/Users/sebastian/" + id
	body, _ := ioutil.ReadAll(r.Body)

	// Check if the file exists
	if _, err := os.Stat(path); err == nil {
		rw.WriteHeader(http.StatusConflict)
		rw.Write([]byte("File exists"))
		return
	}

	file, err := os.Create(path)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error: " + err.Error()))
		return
	}
	defer file.Close()

	file.Write(body)
	file.Sync()

	rw.WriteHeader(http.StatusOK)
	log.Printf("Created file with id: %s\n", id)
}

func DeleteHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	path := "/Users/sebastian/" + id

	err := os.Remove(path)
	if err != nil {
		// Better error handling would be nice..
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)
	log.Printf("Deleted file with id: %s\n", id)
}

func GetHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	path := "/Users/sebastian/" + id

	data, err := ioutil.ReadFile(path)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}

func UpdateHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	path := "/Users/sebastian/" + id
	body, _ := ioutil.ReadAll(r.Body)

	// Ensure that the file exists
	if _, err := os.Stat(path); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("Error: File not found"))
		return
	}

	// Delete the file
	err := os.Remove(path)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error: " + err.Error()))
		return
	}

	// Create the file
	file, err := os.Create(path)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Error: " + err.Error()))
		return
	}
	defer file.Close()

	file.Write(body)
	file.Sync()

	rw.WriteHeader(http.StatusOK)
	log.Printf("Updated file with id: %s\n", id)
}
