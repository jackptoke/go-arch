package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func encode(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "John",
	}
	p2 := person{
		First: "Jane",
	}

	people := []person{p1, p2}
	err := json.NewEncoder(w).Encode(people)
	checkError(err)
}

func decode(w http.ResponseWriter, r *http.Request) {
	var people []person
	err := json.NewDecoder(r.Body).Decode(&people)
	checkError(err)
	log.Println(people)
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}
