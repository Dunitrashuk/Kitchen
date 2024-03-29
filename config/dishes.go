package config

import (
	"github.com/Dunitrashuk/Kitchen/structs"
)

var Dishes = []structs.Dish{
	{
		1,
		"pizza",
		20,
		2,
		"oven",
	},
	{
		2,
		"salad",
		10,
		1,
		"",
	},
	{
		3,
		"zeama",
		7,
		1,
		"stove",
	},
	{
		4,
		"Scallop Sashimi with Meyer Lemon Confit",
		32,
		3,
		"",
	},
	{
		5,
		"Island Duck with Mulberry Mustard",
		35,
		3,
		"oven",
	},
	{
		6,
		"Waffles",
		10,
		1,
		"stove",
	},
	{
		7,
		"Aubergine",
		20,
		2,
		"",
	},
	{
		8,
		"Lasagna",
		30,
		2,
		"oven",
	},
	{
		9,
		"Burger",
		15,
		1,
		"oven",
	},
	{
		10,
		"Gyros",
		15,
		1,
		"",
	},
}

func GetDish(id int) structs.Dish {
	return Dishes[id-1]
}

func GetDishLen() int {
	return len(Dishes)
}
