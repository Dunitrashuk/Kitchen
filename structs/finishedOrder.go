package structs

type FinishedOrder struct {
	Order_id        string              `json:"order_id"`
	Table_id        int              `json:"table_id"`
	Waiter_id       int              `json:"waiter_id"`
	Items           []int            `json:"items"`
	Priority        int              `json:"priority"`
	Max_wait        int              `json:"max_wait"`
	Pick_up_time    int              `json:"pick_up_time"`
	Cooking_time    int              `json:"cooking_time"`
	Cooking_details []CookingDetails `json:"cooking_details"`
}
