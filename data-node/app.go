package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

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
	body, _ := ioutil.ReadAll(r.Body)

	//f, err := os.Create("/home/sebastian/test/")

	fmt.Fprintf(rw, "Created file with id: %s and data: %s\n", id, body)
}
