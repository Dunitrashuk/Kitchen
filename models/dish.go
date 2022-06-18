package models

type Dish struct {
	DishId int    `json:"dish_id"`
	Name   string `json:"name"`
	PreparationTime int    `json:"preparation_time"`
	Complexity      int    `json:"complexity"`
	CookingApparatus string `json:"cooking_apparatus"`
}
