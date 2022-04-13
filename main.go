package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "cool")
}

func InitialRouters() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	return router
}

func Serve(port string, router *mux.Router) {
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("serve failed :%v", err)
	}
}

func main() {
	// Route initialize
	router := InitialRouters()
	Serve("8080", router)
}
