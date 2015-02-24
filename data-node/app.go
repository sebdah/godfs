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

	files := r.Path("/files/{id}").Subrouter()
	files.Methods("POST").HandlerFunc(FileCreateHandle)

	http.ListenAndServe(":8000", r)
}

func FileCreateHandle(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	path := "/Users/sebastian/" + id
	body, _ := ioutil.ReadAll(r.Body)

	f, _ := os.Create(path)
	defer f.Close()

	f.Write(body)
	f.Sync()

	fmt.Fprintf(rw, "Created file with id: %s", id)
}
