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
	router.HandleFunc("/contact", Contact)

	fmt.Println("El servidor está corriendo en http://localhost:3000.")
	server := http.ListenAndServe(":3000", router)

	log.Fatal(server)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con Go.")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ésta es la página de contacto.")
}
