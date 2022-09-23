package config
import "github.com/Dunitrashuk/Kitchen/structs"

func GetCook(id int) structs.Cook {
	cooks := []structs.Cook{
		{1, 3, 4, "Gordon Ramsay", "Hey, panini head, are you listening to me?"},
		{2, 2, 3, "Jen Shelter", "Move, potato head!"},
		{3, 2, 2, "Michael Ray", "Wait for me!"},
		{4, 1, 2, "Kevin Heart", "I have to move faster!"},
	}
	return cooks[id]
}
