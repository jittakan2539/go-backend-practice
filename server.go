package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
    //endpoint interface ex. http://localhost:4242/create-payment-intent
    http.HandleFunc("/create-payment-internet", handleCreatePaymentIntent)
   
    log.Println("Listening in localhost:4242")

    //address
    var err error = http.ListenAndServe("localhost:4242", nil) //error to take error return value
    if err != nil {
        log.Fatal(err)
    }
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Called.")
}