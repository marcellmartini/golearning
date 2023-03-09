package main

import (
	"http/upper"
	"log"
	"net/http"
)

// Req: http://localhost:1234/upper?word=abc
// Res: ABC
func main() {
	http.HandleFunc("/upper", upper.UpperCaseHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))
}
