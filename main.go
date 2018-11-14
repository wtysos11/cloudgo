package main

import (
	"fmt"
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
func loginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println("method:", r.Method) //获取请求的方法
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		if r.Method == "GET" {
			fmt.Fprintf(w, "Can't directly get the page, you should go to index.html first")
		} else {
			if len(r.Form["username"][0]) != 0 {
				formatter.HTML(w, http.StatusOK, "form", struct {
					USER string `json:"user"`
					PASS string `json:"password"`
				}{USER: r.Form["username"][0], PASS: r.Form["password"][0]})
			}
		}
	}
}

func developingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Somthing bad happened!"))
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
	router.HandleFunc("/login", loginHandler(formatter))
	router.HandleFunc("/unknown", developingHandler(formatter))
	addStaticFileServer(router)
	// router goes last
	n.UseHandler(router)
	http.ListenAndServe(":8080", n)
}
