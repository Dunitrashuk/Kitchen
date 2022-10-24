package main

import (
	"bytes"
	_ "bytes"
	"encoding/json"
	"fmt"
	"github.com/Dunitrashuk/Kitchen/config"
	"github.com/Dunitrashuk/Kitchen/structs"
	_ "github.com/Dunitrashuk/Kitchen/structs"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
	_ "time"
)

var mutex sync.Mutex
var orderList []structs.Order
var finishedOrders []structs.FinishedOrder
var ovens []structs.Apparatus
var stoves []structs.Apparatus


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
	fmt.Printf("Order %d received. Max_wait: %d\n", order.Order_Id, order.Max_Wait)
	go addOrder(order)
}

func addOrder(order structs.Order) {
	mutex.Lock()
	orderList = append(orderList, order)
	finishedOrders = append(finishedOrders, structs.FinishedOrder{
		Order_id:        order.Order_Id,
		Table_id:        order.Table_Id,
		Waiter_id:       order.Waiter_Id,
		Items:           order.Items,
		Priority:        order.Priority,
		Max_wait:        order.Max_Wait,
		Pick_up_time:    order.Pickup_Time,
		Cooking_time: 	0,
		Cooking_details: []structs.CookingDetails{},
	})
	mutex.Unlock()
}

func sendFinishedOrders() {
		for i := 0; i < len(finishedOrders); i++ {
			if len(finishedOrders[i].Items) == len(finishedOrders[i].Cooking_details) {
				sendOrder(finishedOrders[i])
				finishedOrders = removeOrderFromFinishedOrders(finishedOrders, i)
				orderList = removeOrderFromOrderList(orderList, i)
			}
		}
}

func sendOrder(order structs.FinishedOrder) {
	data := order
	jsonData, errMarshall := json.Marshal(data)
	if errMarshall != nil {
		log.Fatal(errMarshall)
	}
	resp, err := http.Post("http://"+config.GetHallAddress()+"/distribution", "application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Order %d sent to hall. Status: %d\n", data.Order_id, resp.StatusCode)
}

func createApparatus() {
	for i := 0; i < config.NrOfStoves(); i++ {
		stoves = append(stoves, structs.Apparatus{0})
	}

	// Initialize ovens
	for i := 0; i < config.NrOfOvens(); i++ {
		ovens = append(ovens, structs.Apparatus{0})
	}
}

func removeOrderFromOrderList(orders []structs.Order, index int) []structs.Order {
	return append(orders[:index], orders[index+1:]...)
}

func removeOrderFromFinishedOrders(orders []structs.FinishedOrder, index int) []structs.FinishedOrder {
	return append(orders[:index], orders[index+1:]...)
}

func removeDish(dishes []int, index int) []int {
	return append(dishes[:index], dishes[index+1:]...)
}

//function to remove selected dish from orderList and add it to finishedOrders
//func updateOrderLists(orderId string, cookId int, dishId int) {
//
//	mutex.Lock()
//	//remove selected dish from order
//	for i := 0; i < len(orderList); i++ {
//		if orderList[i].Order_Id == orderId {
//			orderList[i].Items = removeDish(orderList[i].Items, dishId)
//		}
//	}
//
//	//add selected dish to finishedOrder
//	for i := 0; i < len(finishedOrders); i++ {
//		if finishedOrders[i].Order_id == orderId {
//			finishedOrders[i].Cooking_details = append(finishedOrders[i].Cooking_details, structs.CookingDetails{
//				Food_id: dishId,
//				Cook_id: cookId,
//			})
//		}
//	}
//	mutex.Unlock()
//
//}


func prepareDish(dishId int) {
	preparationTime := config.GetDish(dishId).Preparation_time
	time.Sleep(time.Duration(preparationTime)* time.Second)
}

func createCooks() {
	for i := 0; i < config.NrOfCooks(); i++ {
		go cook(i)
	}
}

func cook(cookId int) {
	my := config.GetCook(cookId)
	fmt.Printf("Cook: %s, %s\n", my.Name, my.Catch_phrase)
	//var currentOrder structs.Order

	for {
		time.Sleep(time.Duration(rand.Intn(500)+1000) * time.Millisecond)

		//highestPriorityIndex := 0
		mutex.Lock()
		if len(finishedOrders) > 0 {
			sendOrder(finishedOrders[0])
			finishedOrders = removeOrderFromFinishedOrders(finishedOrders, 0)
		}
		////check for finished orders and send them back to hall
		//sendFinishedOrders()
		//
		////find order with the highest priority
		//for i := 0; i < len(orderList); i++ {
		//	if orderList[i].Priority > highestPriorityIndex {
		//		highestPriorityIndex = orderList[i].Priority
		//		currentOrder = orderList[i]
		//	}
		//}
		//fmt.Printf("Cook %d took order: %+v\n", my.Cook_id ,currentOrder)
		////search dish to make from the highest priority order
		//orderId := currentOrder.Order_Id
		//dishes := currentOrder.Items
		//for i := 0; i < len(dishes); i++ {
		//	dish := config.GetDish(dishes[i])
		//	if dish.Complexity == my.Rank || dish.Complexity == my.Rank - 1{
		//		updateOrderLists(orderId, cookId, dish.Dish_id)
		//		time.Sleep(10 * time.Millisecond)
		//		//go prepareDish(dish.Dish_id)
		//	}
		//}
		mutex.Unlock()
	}
}

func printOrders() {
	for {
		time.Sleep(time.Duration(rand.Intn(1500)+500) * time.Millisecond)
		fmt.Printf("\nOrderList:\n")
		for i := 0; i < len(orderList); i++ {
			fmt.Printf("%+v\n", orderList[i])
		}

		fmt.Printf("\nFinishedOrders:\n")
		for i := 0; i < len(finishedOrders); i++ {
			fmt.Printf("%+v\n", finishedOrders[i])
		}
	}

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
