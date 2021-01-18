package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"golangfirebase/router"
	"net/http"
)

func main() {
	fmt.Println("application started successfully")

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})
	r.HandleFunc("/posts", router.GetPosts).Methods("GET")
	r.HandleFunc("/posts", router.AddPost).Methods("POST")
	http.ListenAndServe(":8080", r)
}
