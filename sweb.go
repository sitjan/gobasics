package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New try")
}
func contact_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "try contacting my mail")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/contact", contact_handler)
	http.ListenAndServe(":8000", nil)
}
