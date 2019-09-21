package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

//HomeHandler will handle get request to get the current shared clipboard
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile("clipboard.txt")
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(dat)

}

//HomePostHandler will handle post request to save the current shared clipboard
func HomePostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	err = ioutil.WriteFile("clipboard.txt", body, 0644)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/", HomePostHandler).Methods("POST")
	http.ListenAndServe(":1234", r)
}
