package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"github.com/Dunitrashuk/Kitchen/config"
	"github.com/Dunitrashuk/Kitchen/structs"
	_ "github.com/Dunitrashuk/Kitchen/structs"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "time"
)

var orderList []structs.Order

func getKitchen(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Kitchen Server is Listening on port 8081")
	fmt.Fprintf(w, "Kitchen Server is Listening on port 8081")
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	var order structs.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	orderList = append(orderList, order)
	fmt.Printf("Order %s received. Max_wait: %d\n", order.Order_Id, order.Max_Wait)
}

//func sendDish(index int) {
//
//	data := config.GetDish(index)
//	jsonData, errMarshall := json.Marshal(data)
//	if errMarshall != nil {
//		log.Fatal(errMarshall)
//	}
//	resp, err := http.Post("http://"+config.GetHallAddress()+"/distribution", "application/json",
//		bytes.NewBuffer(jsonData))
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Dish %d sent to hall. Status: %d\n", data.Dish_id, resp.StatusCode)
//}

func createCooks() {
	for i := 1; i <= config.NrOfCooks(); i++ {
		go cook(i)
	}
}

func cook(cookId int) {
	fmt.Printf("Cook %d, started\n", cookId)

}

func kitchenServer() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", getKitchen).Methods("GET")
	myRouter.HandleFunc("/order", getOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+config.GetKitchenPort(), myRouter))
}

func main() {
	createCooks()
	kitchenServer()
}
