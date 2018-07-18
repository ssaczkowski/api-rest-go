package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/peliculas", MovieList)
	router.HandleFunc("/peliculas/{id}", MovieShow)

	fmt.Println("El servidor está corriendo en http://localhost:3000.")
	server := http.ListenAndServe(":3000", router)

	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from my web server with Go.")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List Movies.")
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    movie_id := params["id"]
    fmt.Fprintf(w,"You have uploaded the movie n°%s", movie_id)
}
