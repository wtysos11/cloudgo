package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func addStaticFileServer(r *mux.Router) {
	dir := "./file"
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))
}

func apiTestHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		content := vars["content"]
		formatter.JSON(w, http.StatusOK, struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: id, Content: content})
	}
}

func renderTestHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		content := vars["content"]
		formatter.HTML(w, http.StatusOK, "template", struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: id, Content: content})
	}
}

func main() {
	n := negroni.Classic()
	router := mux.NewRouter()
	formatter := render.New(render.Options{
		IndentJSON: true,
		Directory:  "file",
		Extensions: []string{".html"},
	})
	router.HandleFunc("/api/{id}/{content}", apiTestHandler(formatter))
	router.HandleFunc("/render/{id}/{content}", renderTestHandler(formatter))
	addStaticFileServer(router)
	// router goes last
	n.UseHandler(router)
	http.ListenAndServe(":8080", n)
}
