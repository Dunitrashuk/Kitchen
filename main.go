package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Dunitrashuk/Kitchen/config"
	"github.com/Dunitrashuk/Kitchen/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func getKitchen(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Kitchen Server is Listening on port 8081")
	fmt.Fprintf(w, "Kitchen Server is Listening on port 8081")
}

func getDish(w http.ResponseWriter, r *http.Request) {
	var dish models.Dish
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Dish %d received. Name: %s\n", dish.Dish_id, dish.Name)
}

func sendDishes() {
	time.Sleep(7 * time.Second)
	for i := 6; i <= 10; i++ {
		sendDish(i)
		time.Sleep(1 * time.Second)
	}
}

func sendDish(index int) {

	data := config.GetDish(index)
	jsonData, errMarshall := json.Marshal(data)
	if errMarshall != nil {
		log.Fatal(errMarshall)
	}
	resp, err := http.Post("http://"+config.GetHallAddress()+"/distribution", "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Dish %d sent to hall. Status: %d\n", data.Dish_id ,resp.StatusCode)
}

func kitchenServer() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", getKitchen).Methods("GET")
	myRouter.HandleFunc("/order", getDish).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+config.GetKitchenPort(), myRouter))
}

func main() {
	go sendDishes()
	kitchenServer()
}
