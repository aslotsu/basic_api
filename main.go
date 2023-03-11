package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	result, err := w.Write([]byte("This silly api is working I am shocked"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func sayAnotherHello(w http.ResponseWriter, r *http.Request) {
	result, err := w.Write([]byte("I really want to get that EC2 instance working then I will be fine"))
	if err != nil {
		fmt.Println(result)
	}
}

func main() {

	http.HandleFunc("/", sayHello)
	http.HandleFunc("/hi", sayAnotherHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
	}
}
