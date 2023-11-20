package main

import (
	"fmt"
	"net/http"
)


func defaultServerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "something")
	fmt.Fprintf(w, "Anotherone")
}


func main() {
	http.HandleFunc("/", defaultServerFunc)
	http.ListenAndServe(":8000", nil)
}
