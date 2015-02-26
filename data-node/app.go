package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	file := r.Path("/files/{id}").Subrouter()
	file.Methods("GET").HandlerFunc(FileGetHandler)
	file.Methods("POST").HandlerFunc(FileCreateHandler)

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}

func FileCreateHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	path := "/Users/sebastian/" + id
	body, _ := ioutil.ReadAll(r.Body)

	file, _ := os.Create(path)
	defer file.Close()

	file.Write(body)
	file.Sync()

	fmt.Printf("Created file with id: %s", id)
	rw.WriteHeader(http.StatusOK)
}

func FileGetHandler(rw http.ResponseWriter, r *http.Request) {
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
