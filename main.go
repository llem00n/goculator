package main

import (
	"fmt"
	"goculator/handlers"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()

	r.PathPrefix("/assets/").
    Handler(http.StripPrefix("/assets/",
        http.FileServer(http.Dir("./assets/"))))

	r.Handle("/", handlers.NewPageHandler()).Methods("GET")
	r.Handle("/{operation}", handlers.NewOperationHandler()).Methods("POST")

	fmt.Println("Listening on http://localhost:8008")
	http.ListenAndServe(":8008", r)
}
