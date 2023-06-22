package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// //	type book struct {
// //		Id  string `json:"t1"`
// //		Title string `json:"t2"`
// //		Author string `json:"t3"`
// //		Quantity int  `json:"t4"`
// //	}
// func initializeRouter() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/users", GetUsers).Methods("GET")
// 	r.HandleFunc("/users{id}", GetUsers).Methods("GET")
// 	r.HandleFunc("/users", CreateUsers).Methods("POST")
// 	// r.HandleFunc("/users{id}", UpdateUsers).Methods("PUT")
// 	// r.HandleFunc("/users{id}", DeleteUsers).Methods("DELETE")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

// func main() {
// 	InitalizeMigration()
// 	initializeRouter()
// }
