package main

import (
	. "godroplet/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/newdroplet", PostDropletHandler)
	mux.HandleFunc("/droplets", GetDropletsHandler)
	mux.HandleFunc("/droplets/{dropletID}", GetDropletHandler)
	mux.HandleFunc("/droplets/{dropletID}/actions", GetDropletActionsHandler)
	mux.HandleFunc("/delete/{dropletID}", DeleteDropletHandler)
	mux.HandleFunc("/backups/{dropletID}", GetBackupsHandler)
	mux.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":9000", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
