package main

import (
	"fmt"
	"github.com/Dunitrashuk/Kitchen/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getKitchen(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Kitchen Server is Listening on port 8081")
	fmt.Fprintf(w, "Kitchen Server is Listening on port 8081")
}

func kitchenServer() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", getKitchen).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+config.GetKitchenPort(), myRouter))
}

func main() {
	kitchenServer();
}
