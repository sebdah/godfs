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

	f, _ := os.Create(path)
	defer f.Close()

	f.Write(body)
	f.Sync()

	fmt.Fprintf(rw, "Created file with id: %s", id)
}

func FileGetHandler(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	path := "/Users/sebastian/" + id

	data, _ := ioutil.ReadFile(path)

	fmt.Fprint(rw, string(data))
}
