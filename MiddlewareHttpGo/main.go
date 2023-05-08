package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", myMiddleware(myStep))
	http.ListenAndServe(":8080", nil)
}
func myStep(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello Handler!")
	w.WriteHeader(http.StatusOK)

}

func myMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello from middleware!")
		next(w, r)
	}
}
