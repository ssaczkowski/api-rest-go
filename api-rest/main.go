package main

import(
	"fmt"
	"net/http"
	"log"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hola mundo desde mi servidor web con Go")
	})

	server := http.ListenAndServe(":4000",nil)

	log.Fatal(server)
}
