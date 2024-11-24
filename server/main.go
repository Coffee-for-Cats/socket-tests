package main

import "net/http"

func main() {
	http.HandleFunc("/connect", Connect)
	http.ListenAndServe(":8080", nil)
}
