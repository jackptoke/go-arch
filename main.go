package main

import (
	"encoding/json"
	"go-arch/security"
	"log"
	"math/rand"
	"net/http"
)

type person struct {
	First string
}

// base64.StdEncoding.EncodeToString([]byte("hello world"))

func main() {
	//http.HandleFunc("/encode", encode)
	//http.HandleFunc("/decode", decode)
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	return
	//}
	password := "password"
	var key []byte
	for i := 1; i <= 64; i++ {
		key = append(key, byte(rand.Int()))
	}

	hashedPassword, err := security.HashPassword(password)
	checkError(err)
	err = security.ComparePassword(password, hashedPassword)
	if err != nil {
		log.Fatalln("Not logged in. Error: ", err)
	}
	log.Println("Logged in")
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
