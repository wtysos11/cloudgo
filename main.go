package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func addStaticFileServer(r *mux.Router) {
	dir := "file"
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))
}

func main() {
	n := negroni.Classic()
	router := mux.NewRouter()
	//router.HandleFunc("/", )
	addStaticFileServer(router)
	// router goes last
	n.UseHandler(router)
	http.ListenAndServe(":8080", n)
}
