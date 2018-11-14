package main

import (
	"net/http"
	"fmt"
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
func loginHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		r.ParseForm()
		fmt.Println("method:", r.Method) //获取请求的方法
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		if len(r.Form["username"])>0{
			formatter.HTML(w, http.StatusOK, "index", struct {
				ID      string `json:"id"`
				Content string `json:"content"`
			}{ID: r.Form["username"][0], Content: r.Form["password"][0]})
		}
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
	router.HandleFunc("/login",loginHandler(formatter))
	addStaticFileServer(router)
	// router goes last
	n.UseHandler(router)
	http.ListenAndServe(":8080", n)
}
